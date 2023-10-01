package contests

import (
	"context"
	"sort"
	"time"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	ent_contest_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contesttask"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskDetail struct {
	taskId              int
	score               int
	currentPenaltyCount int            // 今、順位表に最終結果として表示されるペナ(総合順位のペナに書かれる数字)
	nextPenaltyCount    int            // 次ACしたときに反映される分のペナ(総合順位のペナには足されない数字)
	acSubmitId          *int           // nil
	untilAc             *time.Duration // nil
}

type StandingsRecord struct {
	rank              int
	userName          string
	totalScore        int
	totalPenaltyCount int
	latestUntilAc     *time.Duration // nil
	taskDetailList    []TaskDetail
}

const STATUS_AC = "AC"
const STATUS_WA = "WA"

func (i *Interactor) GetStandings(ctx context.Context, req *backendv1.GetStandingsRequest) (*backendv1.GetStandingsResponse, error) {
	// get contest info
	contest, err := i.entClient.Contest.Query().
		Where(ent_contest.Slug(req.ContestSlug)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "the contest(slug: %s) is not found", req.ContestSlug)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get contest submits
	submits, err := i.entClient.Submit.Query().
		WithUser().
		WithContest().
		WithTask().
		Where(ent_submit.HasContestWith(ent_contest.ID(contest.ID))).
		All(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// sort submissions by submit_at (asc)
	sort.SliceStable(submits, func(i, j int) bool { return submits[i].SubmittedAt.Before(submits[j].SubmittedAt) })

	// get user info
	userInfo, err := separateSubmit(i, ctx, submits, contest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	standings_list := GetStandingsRecordSlice(userInfo)

	var standings_record []*backendv1.StandingsRecord
	for _, row := range standings_list {
		standings_record = append(standings_record, toStandingsRecord(row))
	}

	return &backendv1.GetStandingsResponse{
		StandingsList: standings_record,
	}, nil
}

func GetStandingsRecordSlice(userInfo map[int]StandingsRecord) []StandingsRecord {
	result := make([]StandingsRecord, 0)

	for _, value := range userInfo {
		result = append(result, value)
	}

	// sort by totalScore and total_penalty
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].totalScore > result[j].totalScore {
			return true
		} else if result[i].totalPenaltyCount < result[j].totalPenaltyCount {
			return true
		}
		return false
	})

	// allocate rank
	for index := range result {
		result[index].rank = index + 1
	}

	return result
}

/*
* separate submit by user_id
 */
func separateSubmit(i *Interactor, ctx context.Context, submissions []*ent.Submit, contest *ent.Contest) (map[int]StandingsRecord, error) {
	userInfo := make(map[int]StandingsRecord)

	for _, submission := range submissions {

		// exception handling
		if submission.SubmittedAt.Before(contest.StartAt) || contest.EndAt.After(submission.SubmittedAt) {
			continue
		}

		// initialize
		err := initializeContestTasksResult(i, ctx, userInfo, submission.Edges.User.ID, contest.ID)
		if err != nil {
			return nil, err
		}

		if !isHigherScore(userInfo, submission) {
			continue
		}

		index := getTaskDetailIndex(userInfo, submission.Edges.User.ID, submission.Edges.Task.ID)
		updateUserInfo := userInfo[submission.Edges.User.ID]
		if *submission.Status == STATUS_AC {
			untilAc := time.Until(contest.StartAt) * -1
			updateUserInfo.taskDetailList[index].acSubmitId = &submission.ID
			updateUserInfo.taskDetailList[index].untilAc = &untilAc
			updateUserInfo.taskDetailList[index].score = submission.Score
			updateUserInfo.latestUntilAc = &untilAc
			updateUserInfo.totalScore += updateUserInfo.taskDetailList[index].score

			// penalty
			updateUserInfo.taskDetailList[index].currentPenaltyCount += updateUserInfo.taskDetailList[index].nextPenaltyCount
			updateUserInfo.totalPenaltyCount += updateUserInfo.taskDetailList[index].nextPenaltyCount
			updateUserInfo.taskDetailList[index].nextPenaltyCount = 0
		} else {
			updateUserInfo.taskDetailList[index].nextPenaltyCount++
		}

		userInfo[submission.Edges.User.ID] = updateUserInfo
	}

	return userInfo, nil
}

func isHigherScore(userInfo map[int]StandingsRecord, submission *ent.Submit) bool {

	specific_user_taskDetailList := userInfo[submission.Edges.User.ID].taskDetailList

	// println("[isHigherScore: parent] ID:" + strconv.Itoa(submission.ID))
	for _, taskDetail := range specific_user_taskDetailList {
		if taskDetail.taskId == submission.Edges.Task.ID && taskDetail.score < submission.Score {
			// user get high score than prev submit
			return true
		}

	}

	return false
}

/*
* set initial values for userInfo taskDetail taskId
 */
func initializeContestTasksResult(i *Interactor, ctx context.Context, userInfo map[int]StandingsRecord, user_id int, contest_id int) error {
	_, initializedFlag := userInfo[user_id]

	if initializedFlag {
		// already initialized
		return nil
	}

	// get user info
	user, err := i.entClient.User.Query().Where(ent_user.ID(user_id)).Only(ctx)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	// get contestTasks
	contestTasks, err := i.entClient.ContestTask.Query().
		Where(ent_contest_task.HasContestWith(ent_contest.ID(contest_id))).
		All(ctx)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	taskDetailList := make([]TaskDetail, 0)
	for _, contest_task := range contestTasks {
		taskDetailList = append(taskDetailList, TaskDetail{contest_task.TaskID, 0, 0, 0, nil, nil})
	}

	// insert initialized variables
	userInfo[user_id] = StandingsRecord{1, user.Username, 0, 0, nil, taskDetailList}

	return nil
}

func getTaskDetailIndex(userInfo map[int]StandingsRecord, user_id int, target_taskId int) int {
	var targetIndex int = -1

	for index, taskDetail := range userInfo[user_id].taskDetailList {
		if taskDetail.taskId == target_taskId {
			targetIndex = index
			break
		}
	}

	return targetIndex
}

func toStandingsRecord(standings StandingsRecord) *backendv1.StandingsRecord {
	var taskDetailList []*backendv1.StandingsRecord_TaskDetail
	for _, row := range standings.taskDetailList {
		taskDetailList = append(taskDetailList, toStandingsRecordTaskDetail(row))
	}

	return &backendv1.StandingsRecord{
		Rank:              int32(standings.rank),
		Username:          standings.userName,
		TotalScore:        int32(standings.totalScore),
		TotalPenaltyCount: int32(standings.totalPenaltyCount),
		LatestAcAt:        toTimestamp(standings.latestUntilAc),
		TaskDetailList:    taskDetailList,
	}
}

func toTimestamp(d *time.Duration) *timestamppb.Timestamp {
	seconds := int64(d.Seconds())
	nanos := int64(d.Nanoseconds() % 1e9)
	return timestamppb.New(time.Unix(seconds, nanos))
}

func toStandingsRecordTaskDetail(td TaskDetail) *backendv1.StandingsRecord_TaskDetail {
	var acSubmitID int32
	if td.acSubmitId != nil {
		acSubmitID = int32(*td.acSubmitId)
	}

	var untilAc *durationpb.Duration
	untilAc = nil
	if td.untilAc != nil {
		untilAc = durationpb.New(*td.untilAc)
	}

	return &backendv1.StandingsRecord_TaskDetail{
		TaskId:       int32(td.taskId),
		Score:        int32(td.score),
		PenaltyCount: int32(td.currentPenaltyCount),
		AcSubmitId:   &acSubmitID,
		UntilAc:      untilAc,
	}
}
