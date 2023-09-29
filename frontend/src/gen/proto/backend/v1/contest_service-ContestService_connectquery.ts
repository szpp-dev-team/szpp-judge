// @generated by protoc-gen-connect-query v0.5.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/contest_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAnswerRequest, CreateAnswerResponse, CreateClarificationRequest, CreateClarificationResponse, CreateContestRequest, CreateContestResponse, DeleteAnswerRequest, DeleteAnswerResponse, DeleteClarificationRequest, DeleteClarificationResponse, GetContestRequest, GetContestResponse, GetContestTaskRequest, GetContestTaskResponse, GetMySubmissionStatusesRequest, GetMySubmissionStatusesResponse, GetStandingsRequest, GetStandingsResponse, ListClarificationsRequest, ListClarificationsResponse, ListContestsRequest, ListContestsResponse, ListContestTasksRequest, ListContestTasksResponse, RegisterMeRequest, RegisterMeResponse, SyncContestTasksRequest, SyncContestTasksResponse, UpdateAnswerRequest, UpdateAnswerResponse, UpdateContestRequest, UpdateContestResponse } from "./contest_service_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks } from "@connectrpc/connect-query";

export const typeName = "backend.v1.ContestService";

/**
 * @generated from service backend.v1.ContestService
 */
export const ContestService = {
  typeName: "backend.v1.ContestService",
  methods: {
    /**
     * コンテストを作成する
     *
     * @generated from rpc backend.v1.ContestService.CreateContest
     */
    createContest: {
      name: "CreateContest",
      I: CreateContestRequest,
      O: CreateContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * コンテスト情報を取得する
     *
     * @generated from rpc backend.v1.ContestService.GetContest
     */
    getContest: {
      name: "GetContest",
      I: GetContestRequest,
      O: GetContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * コンテスト情報を更新する
     *
     * @generated from rpc backend.v1.ContestService.UpdateContest
     */
    updateContest: {
      name: "UpdateContest",
      I: UpdateContestRequest,
      O: UpdateContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * コンテストの一覧を取得する
     *
     * @generated from rpc backend.v1.ContestService.ListContests
     */
    listContests: {
      name: "ListContests",
      I: ListContestsRequest,
      O: ListContestsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * コンテストに紐づく問題の一覧を取得する
     *
     * @generated from rpc backend.v1.ContestService.ListContestTasks
     */
    listContestTasks: {
      name: "ListContestTasks",
      I: ListContestTasksRequest,
      O: ListContestTasksResponse,
      kind: MethodKind.Unary,
    },
    /**
     * コンテストに紐づく問題を取得する
     *
     * @generated from rpc backend.v1.ContestService.GetContestTask
     */
    getContestTask: {
      name: "GetContestTask",
      I: GetContestTaskRequest,
      O: GetContestTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 問題をコンテストに紐づかせる
     *
     * @generated from rpc backend.v1.ContestService.SyncContestTasks
     */
    syncContestTasks: {
      name: "SyncContestTasks",
      I: SyncContestTasksRequest,
      O: SyncContestTasksResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 自分の問題ごとの結果情報を取得する
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
     * 順位表を取得する
     *
     * @generated from rpc backend.v1.ContestService.GetStandings
     */
    getStandings: {
      name: "GetStandings",
      I: GetStandingsRequest,
      O: GetStandingsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 参加登録をする
     *
     * @generated from rpc backend.v1.ContestService.RegisterMe
     */
    registerMe: {
      name: "RegisterMe",
      I: RegisterMeRequest,
      O: RegisterMeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Clarification を作成する
     *
     * @generated from rpc backend.v1.ContestService.CreateClarification
     */
    createClarification: {
      name: "CreateClarification",
      I: CreateClarificationRequest,
      O: CreateClarificationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ClarificationListを取得する
     *
     * @generated from rpc backend.v1.ContestService.ListClarifications
     */
    listClarifications: {
      name: "ListClarifications",
      I: ListClarificationsRequest,
      O: ListClarificationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Clarification を削除する
     *
     * @generated from rpc backend.v1.ContestService.DeleteClarification
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
     * @generated from rpc backend.v1.ContestService.CreateAnswer
     */
    createAnswer: {
      name: "CreateAnswer",
      I: CreateAnswerRequest,
      O: CreateAnswerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Answerを更新する
     *
     * @generated from rpc backend.v1.ContestService.UpdateAnswer
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
     * @generated from rpc backend.v1.ContestService.DeleteAnswer
     */
    deleteAnswer: {
      name: "DeleteAnswer",
      I: DeleteAnswerRequest,
      O: DeleteAnswerResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

const $queryService = createQueryService({  service: ContestService,});

/**
 * コンテストを作成する
 *
 * @generated from rpc backend.v1.ContestService.CreateContest
 */
export const createContest = {   ...$queryService.createContest,  ...createUnaryHooks($queryService.createContest)};

/**
 * コンテスト情報を取得する
 *
 * @generated from rpc backend.v1.ContestService.GetContest
 */
export const getContest = {   ...$queryService.getContest,  ...createUnaryHooks($queryService.getContest)};

/**
 * コンテスト情報を更新する
 *
 * @generated from rpc backend.v1.ContestService.UpdateContest
 */
export const updateContest = {   ...$queryService.updateContest,  ...createUnaryHooks($queryService.updateContest)};

/**
 * コンテストの一覧を取得する
 *
 * @generated from rpc backend.v1.ContestService.ListContests
 */
export const listContests = {   ...$queryService.listContests,  ...createUnaryHooks($queryService.listContests)};

/**
 * コンテストに紐づく問題の一覧を取得する
 *
 * @generated from rpc backend.v1.ContestService.ListContestTasks
 */
export const listContestTasks = {   ...$queryService.listContestTasks,  ...createUnaryHooks($queryService.listContestTasks)};

/**
 * コンテストに紐づく問題を取得する
 *
 * @generated from rpc backend.v1.ContestService.GetContestTask
 */
export const getContestTask = {   ...$queryService.getContestTask,  ...createUnaryHooks($queryService.getContestTask)};

/**
 * 問題をコンテストに紐づかせる
 *
 * @generated from rpc backend.v1.ContestService.SyncContestTasks
 */
export const syncContestTasks = {   ...$queryService.syncContestTasks,  ...createUnaryHooks($queryService.syncContestTasks)};

/**
 * 自分の問題ごとの結果情報を取得する
 *
 * @generated from rpc backend.v1.ContestService.GetMySubmissionStatuses
 */
export const getMySubmissionStatuses = {   ...$queryService.getMySubmissionStatuses,  ...createUnaryHooks($queryService.getMySubmissionStatuses)};

/**
 * 順位表を取得する
 *
 * @generated from rpc backend.v1.ContestService.GetStandings
 */
export const getStandings = {   ...$queryService.getStandings,  ...createUnaryHooks($queryService.getStandings)};

/**
 * 参加登録をする
 *
 * @generated from rpc backend.v1.ContestService.RegisterMe
 */
export const registerMe = {   ...$queryService.registerMe,  ...createUnaryHooks($queryService.registerMe)};

/**
 * Clarification を作成する
 *
 * @generated from rpc backend.v1.ContestService.CreateClarification
 */
export const createClarification = {   ...$queryService.createClarification,  ...createUnaryHooks($queryService.createClarification)};

/**
 * ClarificationListを取得する
 *
 * @generated from rpc backend.v1.ContestService.ListClarifications
 */
export const listClarifications = {   ...$queryService.listClarifications,  ...createUnaryHooks($queryService.listClarifications)};

/**
 * Clarification を削除する
 *
 * @generated from rpc backend.v1.ContestService.DeleteClarification
 */
export const deleteClarification = {   ...$queryService.deleteClarification,  ...createUnaryHooks($queryService.deleteClarification)};

/**
 * Answerを追加する
 *
 * @generated from rpc backend.v1.ContestService.CreateAnswer
 */
export const createAnswer = {   ...$queryService.createAnswer,  ...createUnaryHooks($queryService.createAnswer)};

/**
 * Answerを更新する
 *
 * @generated from rpc backend.v1.ContestService.UpdateAnswer
 */
export const updateAnswer = {   ...$queryService.updateAnswer,  ...createUnaryHooks($queryService.updateAnswer)};

/**
 * Answerを削除する
 *
 * @generated from rpc backend.v1.ContestService.DeleteAnswer
 */
export const deleteAnswer = {   ...$queryService.deleteAnswer,  ...createUnaryHooks($queryService.deleteAnswer)};
