// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/contest_resources.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Difficulty } from "./task_resources_pb";

/**
 * @generated from message backend.v1.ContestTask
 */
export class ContestTask extends Message<ContestTask> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * @generated from field: int32 exec_time_limit = 3;
   */
  execTimeLimit = 0;

  /**
   * @generated from field: int32 exec_memory_limit = 4;
   */
  execMemoryLimit = 0;

  /**
   * @generated from field: backend.v1.Difficulty difficulty = 5;
   */
  difficulty = Difficulty.DIFFICULTY_UNSPECIFIED;

  /**
   * @generated from field: int32 score = 6;
   */
  score = 0;

  constructor(data?: PartialMessage<ContestTask>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ContestTask";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "exec_time_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "exec_memory_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "difficulty", kind: "enum", T: proto3.getEnumType(Difficulty) },
    { no: 6, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContestTask {
    return new ContestTask().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContestTask {
    return new ContestTask().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContestTask {
    return new ContestTask().fromJsonString(jsonString, options);
  }

  static equals(a: ContestTask | PlainMessage<ContestTask> | undefined, b: ContestTask | PlainMessage<ContestTask> | undefined): boolean {
    return proto3.util.equals(ContestTask, a, b);
  }
}

/**
 * @generated from message backend.v1.Contest
 */
export class Contest extends Message<Contest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string slug = 3;
   */
  slug = "";

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: repeated int32 task_ids = 5;
   */
  taskIds: number[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp start_at = 6;
   */
  startAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_at = 7;
   */
  endAt?: Timestamp;

  constructor(data?: PartialMessage<Contest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.Contest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "task_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 6, name: "start_at", kind: "message", T: Timestamp },
    { no: 7, name: "end_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Contest {
    return new Contest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Contest {
    return new Contest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Contest {
    return new Contest().fromJsonString(jsonString, options);
  }

  static equals(a: Contest | PlainMessage<Contest> | undefined, b: Contest | PlainMessage<Contest> | undefined): boolean {
    return proto3.util.equals(Contest, a, b);
  }
}

/**
 * @generated from message backend.v1.MutationContest
 */
export class MutationContest extends Message<MutationContest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string slug = 2;
   */
  slug = "";

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * @generated from field: google.protobuf.Timestamp start_at = 4;
   */
  startAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_at = 5;
   */
  endAt?: Timestamp;

  /**
   * @generated from field: bool is_public = 6;
   */
  isPublic = false;

  constructor(data?: PartialMessage<MutationContest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.MutationContest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "start_at", kind: "message", T: Timestamp },
    { no: 5, name: "end_at", kind: "message", T: Timestamp },
    { no: 6, name: "is_public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MutationContest {
    return new MutationContest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MutationContest {
    return new MutationContest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MutationContest {
    return new MutationContest().fromJsonString(jsonString, options);
  }

  static equals(a: MutationContest | PlainMessage<MutationContest> | undefined, b: MutationContest | PlainMessage<MutationContest> | undefined): boolean {
    return proto3.util.equals(MutationContest, a, b);
  }
}

/**
 * @generated from message backend.v1.StandingsRecord
 */
export class StandingsRecord extends Message<StandingsRecord> {
  /**
   * @generated from field: int32 rank = 1;
   */
  rank = 0;

  /**
   * @generated from field: string username = 2;
   */
  username = "";

  /**
   * @generated from field: int32 total_score = 3;
   */
  totalScore = 0;

  /**
   * @generated from field: int32 total_penalty_count = 4;
   */
  totalPenaltyCount = 0;

  /**
   * @generated from field: optional google.protobuf.Timestamp latest_ac_at = 5;
   */
  latestAcAt?: Timestamp;

  /**
   * @generated from field: repeated backend.v1.StandingsRecord.TaskDetail task_detail_list = 6;
   */
  taskDetailList: StandingsRecord_TaskDetail[] = [];

  constructor(data?: PartialMessage<StandingsRecord>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.StandingsRecord";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rank", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "total_score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "total_penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "latest_ac_at", kind: "message", T: Timestamp, opt: true },
    { no: 6, name: "task_detail_list", kind: "message", T: StandingsRecord_TaskDetail, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StandingsRecord {
    return new StandingsRecord().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StandingsRecord {
    return new StandingsRecord().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StandingsRecord {
    return new StandingsRecord().fromJsonString(jsonString, options);
  }

  static equals(a: StandingsRecord | PlainMessage<StandingsRecord> | undefined, b: StandingsRecord | PlainMessage<StandingsRecord> | undefined): boolean {
    return proto3.util.equals(StandingsRecord, a, b);
  }
}

/**
 * @generated from message backend.v1.StandingsRecord.TaskDetail
 */
export class StandingsRecord_TaskDetail extends Message<StandingsRecord_TaskDetail> {
  /**
   * @generated from field: int32 task_id = 1;
   */
  taskId = 0;

  /**
   * @generated from field: int32 score = 2;
   */
  score = 0;

  /**
   * @generated from field: int32 penalty_count = 3;
   */
  penaltyCount = 0;

  /**
   * @generated from field: optional int32 ac_submit_id = 4;
   */
  acSubmitId?: number;

  /**
   * コンテスト開始からの経過時間
   *
   * @generated from field: optional google.protobuf.Duration until_ac = 5;
   */
  untilAc?: Duration;

  constructor(data?: PartialMessage<StandingsRecord_TaskDetail>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.StandingsRecord.TaskDetail";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "ac_submit_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "until_ac", kind: "message", T: Duration, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StandingsRecord_TaskDetail {
    return new StandingsRecord_TaskDetail().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StandingsRecord_TaskDetail {
    return new StandingsRecord_TaskDetail().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StandingsRecord_TaskDetail {
    return new StandingsRecord_TaskDetail().fromJsonString(jsonString, options);
  }

  static equals(a: StandingsRecord_TaskDetail | PlainMessage<StandingsRecord_TaskDetail> | undefined, b: StandingsRecord_TaskDetail | PlainMessage<StandingsRecord_TaskDetail> | undefined): boolean {
    return proto3.util.equals(StandingsRecord_TaskDetail, a, b);
  }
}

/**
 * @generated from message backend.v1.SubmissionStatus
 */
export class SubmissionStatus extends Message<SubmissionStatus> {
  /**
   * @generated from field: int32 task_id = 1;
   */
  taskId = 0;

  /**
   * @generated from field: optional int32 score = 2;
   */
  score?: number;

  constructor(data?: PartialMessage<SubmissionStatus>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmissionStatus";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmissionStatus {
    return new SubmissionStatus().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmissionStatus {
    return new SubmissionStatus().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmissionStatus {
    return new SubmissionStatus().fromJsonString(jsonString, options);
  }

  static equals(a: SubmissionStatus | PlainMessage<SubmissionStatus> | undefined, b: SubmissionStatus | PlainMessage<SubmissionStatus> | undefined): boolean {
    return proto3.util.equals(SubmissionStatus, a, b);
  }
}

/**
 * @generated from message backend.v1.Clarification
 */
export class Clarification extends Message<Clarification> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string content = 2;
   */
  content = "";

  /**
   * @generated from field: bool is_public = 3;
   */
  isPublic = false;

  /**
   * 自分が質問したものかどうか
   *
   * @generated from field: bool is_mine = 4;
   */
  isMine = false;

  /**
   * @generated from field: optional backend.v1.Clarification.Answer answer = 5;
   */
  answer?: Clarification_Answer;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Clarification>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.Clarification";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "is_mine", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "answer", kind: "message", T: Clarification_Answer, opt: true },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Clarification {
    return new Clarification().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Clarification {
    return new Clarification().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Clarification {
    return new Clarification().fromJsonString(jsonString, options);
  }

  static equals(a: Clarification | PlainMessage<Clarification> | undefined, b: Clarification | PlainMessage<Clarification> | undefined): boolean {
    return proto3.util.equals(Clarification, a, b);
  }
}

/**
 * @generated from message backend.v1.Clarification.Answer
 */
export class Clarification_Answer extends Message<Clarification_Answer> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string content = 2;
   */
  content = "";

  /**
   * 自分が回答したものかどうか 
   *
   * @generated from field: bool is_mine = 3;
   */
  isMine = false;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 4;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 5;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Clarification_Answer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.Clarification.Answer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_mine", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "created_at", kind: "message", T: Timestamp },
    { no: 5, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Clarification_Answer {
    return new Clarification_Answer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Clarification_Answer {
    return new Clarification_Answer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Clarification_Answer {
    return new Clarification_Answer().fromJsonString(jsonString, options);
  }

  static equals(a: Clarification_Answer | PlainMessage<Clarification_Answer> | undefined, b: Clarification_Answer | PlainMessage<Clarification_Answer> | undefined): boolean {
    return proto3.util.equals(Clarification_Answer, a, b);
  }
}

