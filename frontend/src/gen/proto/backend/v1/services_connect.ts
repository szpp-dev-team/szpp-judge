// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/services.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAnswerRequest, CreateAnswerResponse, CreateClarificationRequest, CreateClarificationResponse, CreateContestRequest, CreateContestResponse, CreateTaskRequest, CreateTaskResponse, CreateUserRequest, CreateUserResponse, DeleteAnswerRequest, DeleteAnswerResponse, DeleteClarificationRequest, DeleteClarificationResponse, GetAnswerRequest, GetAnswerResponse, GetClarificationRequest, GetClarificationResponse, GetContestRequest, GetContestResponse, GetJudgeProgressRequest, GetJudgeProgressResponse, GetMySubmissionStatusesRequest, GetMySubmissionStatusesResponse, GetStandingsRequest, GetStandingsResponse, GetSubmissionDetailRequest, GetSubmissionDetailResponse, GetTaskRequest, GetTaskResponse, GetTestcaseSetsRequest, GetTestcaseSetsResponse, GetUserRequest, GetUserResponse, ListClarificationsRequest, ListClarificationsResponse, ListContestsRequest, ListContestsResponse, ListContestTasksRequest, ListContestTasksResponse, ListSubmissionsRequest, ListSubmissionsResponse, LoginRequest, LoginResponse, LogoutRequest, LogoutResponse, PingRequest, PingResponse, SubmitRequest, SubmitResponse, SyncTestcaseSetsRequest, SyncTestcaseSetsResponse, UpdateAnswerRequest, UpdateAnswerResponse, UpdateClarificationRequest, UpdateClarificationResponse, UpdateTaskRequest, UpdateTaskResponse } from "./messages_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service backend.v1.UserService
 */
export const UserService = {
  typeName: "backend.v1.UserService",
  methods: {
    /**
     * 指定された User を取得する
     *
     * @generated from rpc backend.v1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: GetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * User を新たに作成する
     *
     * @generated from rpc backend.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.AuthService
 */
export const AuthService = {
  typeName: "backend.v1.AuthService",
  methods: {
    /**
     * ログイン
     *
     * @generated from rpc backend.v1.AuthService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ログアウト
     *
     * @generated from rpc backend.v1.AuthService.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutRequest,
      O: LogoutResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.TaskService
 */
export const TaskService = {
  typeName: "backend.v1.TaskService",
  methods: {
    /**
     * Task を作成する
     *
     * @generated from rpc backend.v1.TaskService.CreateTask
     */
    createTask: {
      name: "CreateTask",
      I: CreateTaskRequest,
      O: CreateTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 指定された Task を取得する
     *
     * @generated from rpc backend.v1.TaskService.GetTask
     */
    getTask: {
      name: "GetTask",
      I: GetTaskRequest,
      O: GetTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Task を更新する
     *
     * @generated from rpc backend.v1.TaskService.UpdateTask
     */
    updateTask: {
      name: "UpdateTask",
      I: UpdateTaskRequest,
      O: UpdateTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
     * contestant によるリクエストの場合は sample のみ取得する。
     *
     * @generated from rpc backend.v1.TaskService.GetTestcaseSets
     */
    getTestcaseSets: {
      name: "GetTestcaseSets",
      I: GetTestcaseSetsRequest,
      O: GetTestcaseSetsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
     *
     * @generated from rpc backend.v1.TaskService.SyncTestcaseSets
     */
    syncTestcaseSets: {
      name: "SyncTestcaseSets",
      I: SyncTestcaseSetsRequest,
      O: SyncTestcaseSetsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 提出する
     *
     * @generated from rpc backend.v1.TaskService.Submit
     */
    submit: {
      name: "Submit",
      I: SubmitRequest,
      O: SubmitResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 提出の詳細を取得
     *
     * @generated from rpc backend.v1.TaskService.GetSubmissionDetail
     */
    getSubmissionDetail: {
      name: "GetSubmissionDetail",
      I: GetSubmissionDetailRequest,
      O: GetSubmissionDetailResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 提出一覧を取得
     *
     * @generated from rpc backend.v1.TaskService.ListSubmissions
     */
    listSubmissions: {
      name: "ListSubmissions",
      I: ListSubmissionsRequest,
      O: ListSubmissionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ジャッジの進捗を取得
     *
     * @generated from rpc backend.v1.TaskService.GetJudgeProgress
     */
    getJudgeProgress: {
      name: "GetJudgeProgress",
      I: GetJudgeProgressRequest,
      O: GetJudgeProgressResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Clarification を作成する
     *
     * @generated from rpc backend.v1.TaskService.CreateClarification
     */
    createClarification: {
      name: "CreateClarification",
      I: CreateClarificationRequest,
      O: CreateClarificationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 指定された Clarification を取得する
     *
     * @generated from rpc backend.v1.TaskService.GetClarification
     */
    getClarification: {
      name: "GetClarification",
      I: GetClarificationRequest,
      O: GetClarificationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ClarificationListを取得する
     *
     * @generated from rpc backend.v1.TaskService.ListClarifications
     */
    listClarifications: {
      name: "ListClarifications",
      I: ListClarificationsRequest,
      O: ListClarificationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Clarification を更新する
     *
     * @generated from rpc backend.v1.TaskService.UpdateClarification
     */
    updateClarification: {
      name: "UpdateClarification",
      I: UpdateClarificationRequest,
      O: UpdateClarificationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Clarification を削除する
     *
     * @generated from rpc backend.v1.TaskService.DeleteClarification
     */
    deleteClarification: {
      name: "DeleteClarification",
      I: DeleteClarificationRequest,
      O: DeleteClarificationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Answerを追加する
     *
     * @generated from rpc backend.v1.TaskService.CreateAnswer
     */
    createAnswer: {
      name: "CreateAnswer",
      I: CreateAnswerRequest,
      O: CreateAnswerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Answerを取得する
     *
     * @generated from rpc backend.v1.TaskService.GetAnswer
     */
    getAnswer: {
      name: "GetAnswer",
      I: GetAnswerRequest,
      O: GetAnswerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Answerを更新する
     *
     * @generated from rpc backend.v1.TaskService.UpdateAnswer
     */
    updateAnswer: {
      name: "UpdateAnswer",
      I: UpdateAnswerRequest,
      O: UpdateAnswerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Answerを削除する
     *
     * @generated from rpc backend.v1.TaskService.DeleteAnswer
     */
    deleteAnswer: {
      name: "DeleteAnswer",
      I: DeleteAnswerRequest,
      O: DeleteAnswerResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.HealthcheckService
 */
export const HealthcheckService = {
  typeName: "backend.v1.HealthcheckService",
  methods: {
    /**
     * @generated from rpc backend.v1.HealthcheckService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.ContestService
 */
export const ContestService = {
  typeName: "backend.v1.ContestService",
  methods: {
    /**
     * @generated from rpc backend.v1.ContestService.CreateContest
     */
    createContest: {
      name: "CreateContest",
      I: CreateContestRequest,
      O: CreateContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc backend.v1.ContestService.GetContest
     */
    getContest: {
      name: "GetContest",
      I: GetContestRequest,
      O: GetContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc backend.v1.ContestService.ListContests
     */
    listContests: {
      name: "ListContests",
      I: ListContestsRequest,
      O: ListContestsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc backend.v1.ContestService.ListContestTasks
     */
    listContestTasks: {
      name: "ListContestTasks",
      I: ListContestTasksRequest,
      O: ListContestTasksResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 自分の問題ごとの結果情報を返す
     *
     * @generated from rpc backend.v1.ContestService.GetMySubmissionStatuses
     */
    getMySubmissionStatuses: {
      name: "GetMySubmissionStatuses",
      I: GetMySubmissionStatusesRequest,
      O: GetMySubmissionStatusesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 順位表取得
     *
     * @generated from rpc backend.v1.ContestService.GetStandings
     */
    getStandings: {
      name: "GetStandings",
      I: GetStandingsRequest,
      O: GetStandingsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

