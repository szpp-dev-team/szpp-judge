package contests

import (
	"context"
	"sort"
	"time"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskDetail struct {
	task_id               int
	score                 int
	current_penalty_count int            // 今、順位表に最終結果として表示されるペナ(総合順位のペナに書かれる数字)
	next_penalty_count    int            // 次ACしたときに反映される分のペナ(総合順位のペナには足されない数字)
	ac_submit_id          *int           // nil
	until_ac              *time.Duration // nil
}

type StandingsRecord struct {
	rank                int
	user_name           string
	total_score         int
	total_penalty_count int
	latest_until_ac     *time.Duration // nil
	task_detail_list    []TaskDetail
}

const STATUS_AC = "AC"
const STATUS_WA = "WA"

func (i *Interactor) GetStandings(ctx context.Context, req *backendv1.GetStandingsRequest) (*backendv1.GetStandingsResponse, error) {
	// get contest info
	contest, err := i.entClient.Contest.Query().
		Where(ent_contest.ID(int(req.ContestId))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "the contest(id: %d) is not found", req.ContestId)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get contest submits
	submits, err := i.entClient.Submit.Query().
		Where(ent_submit.ID(int(req.ContestId))).
		All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// sort submissions by submit_at (asc)
	sort.SliceStable(submits, func(i, j int) bool { return submits[i].SubmittedAt.Before(submits[j].SubmittedAt) })

	// get user info
	user_info, err := separateSubmit(i, ctx, submits, contest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	standings_list := GetStandingsRecordSlice(user_info)

	var standings_record []*backendv1.StandingsRecord
	for _, row := range standings_list {
		standings_record = append(standings_record, toStandingsRecord(row))
	}

	return &backendv1.GetStandingsResponse{
		StandingsList: standings_record,
	}, nil
}

func GetStandingsRecordSlice(user_info map[int]StandingsRecord) []StandingsRecord {
	result := make([]StandingsRecord, 0)

	for _, value := range user_info {
		result = append(result, value)
	}

	// sort by total_score and total_penalty
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].total_score > result[j].total_score {
			return true
		} else if result[i].total_penalty_count < result[j].total_penalty_count {
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
	user_info := make(map[int]StandingsRecord)

	for _, submission := range submissions {

		// exception handling
		if submission.SubmittedAt.Before(contest.StartAt) || contest.EndAt.After(submission.SubmittedAt) {
			continue
		}

		if !isHigherScore(user_info, submission) {
			continue
		}

		// initialize
		err := initializeContestTasksResult(i, ctx, user_info, submission.Edges.User.ID, contest.ID)
		if err != nil {
			return nil, err
		}

		index := getTaskDetailIndex(user_info, submission.Edges.User.ID, submission.Edges.Task.ID)
		update_user_info := user_info[submission.Edges.User.ID]
		if *submission.Status == STATUS_AC {
			until_ac := time.Until(contest.StartAt) * -1
			update_user_info.task_detail_list[index].ac_submit_id = &submission.ID
			update_user_info.task_detail_list[index].until_ac = &until_ac
			update_user_info.task_detail_list[index].score = submission.Score
			update_user_info.latest_until_ac = &until_ac
			update_user_info.total_score += update_user_info.task_detail_list[index].score

			// penalty
			update_user_info.task_detail_list[index].current_penalty_count += update_user_info.task_detail_list[index].next_penalty_count
			update_user_info.total_penalty_count += update_user_info.task_detail_list[index].next_penalty_count
			update_user_info.task_detail_list[index].next_penalty_count = 0
		} else {
			update_user_info.task_detail_list[index].next_penalty_count++
		}

		user_info[submission.Edges.User.ID] = update_user_info
	}

	return user_info, nil
}

func isHigherScore(user_info map[int]StandingsRecord, submission *ent.Submit) bool {

	specific_user_task_detail_list := user_info[submission.Edges.User.ID].task_detail_list

	for _, task_detail := range specific_user_task_detail_list {

		if task_detail.task_id == submission.Edges.Task.ID && task_detail.score < submission.Score {
			// user get high score than prev submit
			return true
		}

	}

	return false
}

/*
* set initial values for user_info task_detail task_id
 */
func initializeContestTasksResult(i *Interactor, ctx context.Context, user_info map[int]StandingsRecord, user_id int, contest_id int) error {

	_, initialized_flag := user_info[user_id]

	if initialized_flag {
		// already initialized
		return nil
	}

	// get user info
	user, err := i.entClient.User.Query().Where(ent_user.ID(user_id)).Only(ctx)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	// get contest_tasks
	contest_tasks, err := i.entClient.ContestTask.Query().Where(predicate.ContestTask(ent_contest.ID(contest_id))).All(ctx)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	task_detail_list := make([]TaskDetail, 0)
	for _, contest_task := range contest_tasks {
		task_detail_list = append(task_detail_list, TaskDetail{contest_task.ID, 0, 0, 0, nil, nil})
	}

	// insert initialized variables
	user_info[user_id] = StandingsRecord{1, user.Username, 0, 0, nil, task_detail_list}

	return nil
}

func getTaskDetailIndex(user_info map[int]StandingsRecord, user_id int, target_task_id int) int {
	var target_index int = -1

	for index, task_detail := range user_info[user_id].task_detail_list {
		if task_detail.task_id == target_task_id {
			target_index = index
			break
		}
	}

	return target_index
}

func toStandingsRecord(standings StandingsRecord) *backendv1.StandingsRecord {
	var taskDetailList []*backendv1.StandingsRecord_TaskDetail
	for _, row := range standings.task_detail_list {
		taskDetailList = append(taskDetailList, toStandingsRecordTaskDetail(row))
	}

	return &backendv1.StandingsRecord{
		Rank:              int32(standings.rank),
		Username:          standings.user_name,
		TotalScore:        int32(standings.total_score),
		TotalPenaltyCount: int32(standings.total_penalty_count),
		LatestAcAt:        toTimestamp(standings.latest_until_ac),
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
	if td.ac_submit_id != nil {
		acSubmitID = int32(*td.ac_submit_id)
	}

	return &backendv1.StandingsRecord_TaskDetail{
		TaskId:       int32(td.task_id),
		Score:        int32(td.score),
		PenaltyCount: int32(td.current_penalty_count),
		AcSubmitId:   &acSubmitID,
		UntilAc:      durationpb.New(*td.until_ac),
	}
}
