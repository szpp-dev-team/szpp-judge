// @generated by protoc-gen-connect-query v0.4.3 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/services.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAnswerRequest, CreateAnswerResponse, CreateClarificationRequest, CreateClarificationResponse, CreateTaskRequest, CreateTaskResponse, DeleteAnswerRequest, DeleteAnswerResponse, DeleteClarificationRequest, DeleteClarificationResponse, GetAnswerRequest, GetAnswerResponse, GetClarificationRequest, GetClarificationResponse, GetJudgeProgressRequest, GetJudgeProgressResponse, GetSubmissionDetailRequest, GetSubmissionDetailResponse, GetTaskRequest, GetTaskResponse, GetTestcaseSetsRequest, GetTestcaseSetsResponse, ListClarificationsRequest, ListClarificationsResponse, ListSubmissionsRequest, ListSubmissionsResponse, SubmitRequest, SubmitResponse, SyncTestcaseSetsRequest, SyncTestcaseSetsResponse, UpdateAnswerRequest, UpdateAnswerResponse, UpdateClarificationRequest, UpdateClarificationResponse, UpdateTaskRequest, UpdateTaskResponse } from "./messages_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService } from "@connectrpc/connect-query";

export const typeName = "backend.v1.TaskService";

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
 * Task を作成する
 *
 * @generated from rpc backend.v1.TaskService.CreateTask
 */
export const createTask = createQueryService({
  service: TaskService,
}).createTask;

/**
 * 指定された Task を取得する
 *
 * @generated from rpc backend.v1.TaskService.GetTask
 */
export const getTask = createQueryService({
  service: TaskService,
}).getTask;

/**
 * Task を更新する
 *
 * @generated from rpc backend.v1.TaskService.UpdateTask
 */
export const updateTask = createQueryService({
  service: TaskService,
}).updateTask;

/**
 * TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
 * contestant によるリクエストの場合は sample のみ取得する。
 *
 * @generated from rpc backend.v1.TaskService.GetTestcaseSets
 */
export const getTestcaseSets = createQueryService({
  service: TaskService,
}).getTestcaseSets;

/**
 * TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
 *
 * @generated from rpc backend.v1.TaskService.SyncTestcaseSets
 */
export const syncTestcaseSets = createQueryService({
  service: TaskService,
}).syncTestcaseSets;

/**
 * 提出する
 *
 * @generated from rpc backend.v1.TaskService.Submit
 */
export const submit = createQueryService({
  service: TaskService,
}).submit;

/**
 * 提出の詳細を取得
 *
 * @generated from rpc backend.v1.TaskService.GetSubmissionDetail
 */
export const getSubmissionDetail = createQueryService({
  service: TaskService,
}).getSubmissionDetail;

/**
 * 提出一覧を取得
 *
 * @generated from rpc backend.v1.TaskService.ListSubmissions
 */
export const listSubmissions = createQueryService({
  service: TaskService,
}).listSubmissions;

/**
 * ジャッジの進捗を取得
 *
 * @generated from rpc backend.v1.TaskService.GetJudgeProgress
 */
export const getJudgeProgress = createQueryService({
  service: TaskService,
}).getJudgeProgress;

/**
 * Clarification を作成する
 *
 * @generated from rpc backend.v1.TaskService.CreateClarification
 */
export const createClarification = createQueryService({
  service: TaskService,
}).createClarification;

/**
 * 指定された Clarification を取得する
 *
 * @generated from rpc backend.v1.TaskService.GetClarification
 */
export const getClarification = createQueryService({
  service: TaskService,
}).getClarification;

/**
 * ClarificationListを取得する
 *
 * @generated from rpc backend.v1.TaskService.ListClarifications
 */
export const listClarifications = createQueryService({
  service: TaskService,
}).listClarifications;

/**
 * Clarification を更新する
 *
 * @generated from rpc backend.v1.TaskService.UpdateClarification
 */
export const updateClarification = createQueryService({
  service: TaskService,
}).updateClarification;

/**
 * Clarification を削除する
 *
 * @generated from rpc backend.v1.TaskService.DeleteClarification
 */
export const deleteClarification = createQueryService({
  service: TaskService,
}).deleteClarification;

/**
 * Answerを追加する
 *
 * @generated from rpc backend.v1.TaskService.CreateAnswer
 */
export const createAnswer = createQueryService({
  service: TaskService,
}).createAnswer;

/**
 * Answerを取得する
 *
 * @generated from rpc backend.v1.TaskService.GetAnswer
 */
export const getAnswer = createQueryService({
  service: TaskService,
}).getAnswer;

/**
 * Answerを更新する
 *
 * @generated from rpc backend.v1.TaskService.UpdateAnswer
 */
export const updateAnswer = createQueryService({
  service: TaskService,
}).updateAnswer;

/**
 * Answerを削除する
 *
 * @generated from rpc backend.v1.TaskService.DeleteAnswer
 */
export const deleteAnswer = createQueryService({
  service: TaskService,
}).deleteAnswer;
