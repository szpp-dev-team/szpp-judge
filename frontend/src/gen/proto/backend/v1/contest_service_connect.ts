// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/contest_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateContestRequest, CreateContestResponse, GetContestRequest, GetContestResponse, GetMySubmissionStatusesRequest, GetMySubmissionStatusesResponse, GetStandingsRequest, GetStandingsResponse, ListContestsRequest, ListContestsResponse, ListContestTasksRequest, ListContestTasksResponse, RegisterMeRequest, RegisterMeResponse, UpdateContestRequest, UpdateContestResponse } from "./contest_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

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
     * @generated from rpc backend.v1.ContestService.UpdateContest
     */
    updateContest: {
      name: "UpdateContest",
      I: UpdateContestRequest,
      O: UpdateContestResponse,
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
    /**
     * 参加登録
     *
     * @generated from rpc backend.v1.ContestService.RegisterMe
     */
    registerMe: {
      name: "RegisterMe",
      I: RegisterMeRequest,
      O: RegisterMeResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

