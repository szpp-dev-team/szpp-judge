package contests

import (
	"testing"
	"time"
)

func createTaskDetail(id int, score int, currentPenaltyCount int, nextPenaltyCount int, acSubmitId *int, untilAc *time.Duration) TaskDetail {
	return TaskDetail{
		taskId:              id,
		score:               score,
		currentPenaltyCount: currentPenaltyCount,
		nextPenaltyCount:    nextPenaltyCount,
		acSubmitId:          acSubmitId,
		untilAc:             untilAc,
	}
}

func createStandingsRecord(userName string, totalScore int, totalPenaltyCount int, latestUntilAc *time.Duration, taskDetailList []TaskDetail) StandingsRecord {
	return StandingsRecord{
		rank:              1,
		userName:          userName,
		totalScore:        totalScore,
		totalPenaltyCount: totalPenaltyCount,
		latestUntilAc:     latestUntilAc,
		taskDetailList:    taskDetailList,
	}
}

func Test_GetStandings(t *testing.T) {
	// entClient := utils.NewTestClient(t)
	// defer entClient.Close()
	// interactor := NewInteractor(entClient, nil)

	// ctx := context.Background()
	// resp, err := interactor.GetStandings(ctx, &backendv1.GetStandingsRequest{
	// 	ContestSlug: "test001",
	// })

	// var emptyInt *int
	// var emptyTime *time.Duration
	// emptyInt = nil
	// emptyTime = nil

	// userInfo := make(map[int]StandingsRecord)

	// navle_task1 := createTaskDetail(4, 0, 0, 0, emptyInt, emptyTime)
	// navle_task2 := createTaskDetail(5, 0, 0, 0, emptyInt, emptyTime)
	// var navle_list []TaskDetail = []TaskDetail{navle_task1, navle_task2}
	// userInfo[0] = createStandingsRecord("navle", 0, 0, emptyTime, navle_list)

	// admin_time, _ := time.ParseDuration("68593026s")
	// admin_task1 := createTaskDetail(4, 100, 0, 0, emptyInt, &admin_time)
	// admin_task2 := createTaskDetail(5, 200, 0, 0, emptyInt, &admin_time)
	// var admin_list []TaskDetail = []TaskDetail{admin_task1, admin_task2}
	// userInfo[1] = createStandingsRecord("admin", 300, 0, &admin_time, admin_list)

	// aya_time, _ := time.ParseDuration("68603810s")
	// aya_task1 := createTaskDetail(4, 100, 0, 0, emptyInt, &aya_time)
	// aya_task2 := createTaskDetail(5, 200, 0, 0, emptyInt, &aya_time)
	// var aya_list []TaskDetail = []TaskDetail{aya_task1, aya_task2}
	// userInfo[2] = createStandingsRecord("aya", 300, 0, &aya_time, aya_list)

	// piddy_time, _ := time.ParseDuration("68593103s")
	// piddy_task1 := createTaskDetail(4, 100, 0, 0, emptyInt, &piddy_time)
	// piddy_task2 := createTaskDetail(5, 200, 0, 0, emptyInt, &piddy_time)
	// var piddy_list []TaskDetail = []TaskDetail{piddy_task1, piddy_task2}
	// userInfo[3] = createStandingsRecord("piddy", 300, 0, &piddy_time, piddy_list)

	// test_time, _ := time.ParseDuration("68594201s")
	// test_task1 := createTaskDetail(4, 100, 0, 0, emptyInt, &test_time)
	// test_task2 := createTaskDetail(5, 200, 0, 0, emptyInt, &test_time)
	// var test_list []TaskDetail = []TaskDetail{test_task1, test_task2}
	// userInfo[4] = createStandingsRecord("test", 300, 0, &test_time, test_list)

	// res := GetStandingsRecordSlice(userInfo)

	// require.NoError(t, err)

	test_time, _ := time.ParseDuration("68594201s")
	t.Log(test_time)

	var sec = 300
	var pena = 1
	test_time += time.Duration(sec * pena * int(time.Second))

	t.Log(test_time)
}
