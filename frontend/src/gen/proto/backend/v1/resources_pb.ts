// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/resources.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { JudgeType } from "../../judge/v1/judge_type_pb";
import { JudgeStatus } from "../../judge/v1/resources_pb";

/**
 * @generated from enum backend.v1.Visibility
 */
export enum Visibility {
  /**
   * @generated from enum value: VISIBILITY_UNSPECIFIED = 0;
   */
  VISIBILITY_UNSPECIFIED = 0,

  /**
   * @generated from enum value: PUBLIC = 1;
   */
  PUBLIC = 1,

  /**
   * @generated from enum value: PRIVATE = 2;
   */
  PRIVATE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Visibility)
proto3.util.setEnumType(Visibility, "backend.v1.Visibility", [
  { no: 0, name: "VISIBILITY_UNSPECIFIED" },
  { no: 1, name: "PUBLIC" },
  { no: 2, name: "PRIVATE" },
]);

/**
 * @generated from enum backend.v1.Difficulty
 */
export enum Difficulty {
  /**
   * @generated from enum value: DIFFICULTY_UNSPECIFIED = 0;
   */
  DIFFICULTY_UNSPECIFIED = 0,

  /**
   * @generated from enum value: BEGINNER = 1;
   */
  BEGINNER = 1,

  /**
   * @generated from enum value: EASY = 2;
   */
  EASY = 2,

  /**
   * @generated from enum value: MEDIUM = 3;
   */
  MEDIUM = 3,

  /**
   * @generated from enum value: HARD = 4;
   */
  HARD = 4,

  /**
   * @generated from enum value: IMPOSSIBLE = 5;
   */
  IMPOSSIBLE = 5,
}
// Retrieve enum metadata with: proto3.getEnumType(Difficulty)
proto3.util.setEnumType(Difficulty, "backend.v1.Difficulty", [
  { no: 0, name: "DIFFICULTY_UNSPECIFIED" },
  { no: 1, name: "BEGINNER" },
  { no: 2, name: "EASY" },
  { no: 3, name: "MEDIUM" },
  { no: 4, name: "HARD" },
  { no: 5, name: "IMPOSSIBLE" },
]);

/**
 * @generated from message backend.v1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string username = 2;
   */
  username = "";

  /**
   * @generated from field: bool is_admin = 4;
   */
  isAdmin = false;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 6;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "is_admin", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "updated_at", kind: "message", T: Timestamp, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
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
 * @generated from message backend.v1.StandingsElement
 */
export class StandingsElement extends Message<StandingsElement> {
  /**
   * @generated from field: int32 rank = 1;
   */
  rank = 0;

  /**
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string user_name = 3;
   */
  userName = "";

  /**
   * @generated from field: int32 total_score = 4;
   */
  totalScore = 0;

  /**
   * @generated from field: int32 total_penalty_count = 5;
   */
  totalPenaltyCount = 0;

  /**
   * @generated from field: google.protobuf.Timestamp latest_ac_at = 6;
   */
  latestAcAt?: Timestamp;

  /**
   * @generated from field: repeated backend.v1.StandingsElement.TaskDetail task_detail_list = 7;
   */
  taskDetailList: StandingsElement_TaskDetail[] = [];

  constructor(data?: PartialMessage<StandingsElement>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.StandingsElement";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rank", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "total_score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "total_penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "latest_ac_at", kind: "message", T: Timestamp },
    { no: 7, name: "task_detail_list", kind: "message", T: StandingsElement_TaskDetail, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StandingsElement {
    return new StandingsElement().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StandingsElement {
    return new StandingsElement().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StandingsElement {
    return new StandingsElement().fromJsonString(jsonString, options);
  }

  static equals(a: StandingsElement | PlainMessage<StandingsElement> | undefined, b: StandingsElement | PlainMessage<StandingsElement> | undefined): boolean {
    return proto3.util.equals(StandingsElement, a, b);
  }
}

/**
 * @generated from message backend.v1.StandingsElement.TaskDetail
 */
export class StandingsElement_TaskDetail extends Message<StandingsElement_TaskDetail> {
  /**
   * @generated from field: int32 task_id = 1;
   */
  taskId = 0;

  /**
   * @generated from field: int32 penalty_count = 2;
   */
  penaltyCount = 0;

  /**
   * @generated from field: int32 ac_submit_id = 3;
   */
  acSubmitId = 0;

  /**
   * @generated from field: google.protobuf.Timestamp ac_elapsed = 4;
   */
  acElapsed?: Timestamp;

  constructor(data?: PartialMessage<StandingsElement_TaskDetail>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.StandingsElement.TaskDetail";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "ac_submit_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "ac_elapsed", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StandingsElement_TaskDetail {
    return new StandingsElement_TaskDetail().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StandingsElement_TaskDetail {
    return new StandingsElement_TaskDetail().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StandingsElement_TaskDetail {
    return new StandingsElement_TaskDetail().fromJsonString(jsonString, options);
  }

  static equals(a: StandingsElement_TaskDetail | PlainMessage<StandingsElement_TaskDetail> | undefined, b: StandingsElement_TaskDetail | PlainMessage<StandingsElement_TaskDetail> | undefined): boolean {
    return proto3.util.equals(StandingsElement_TaskDetail, a, b);
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
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string content = 3;
   */
  content = "";

  /**
   * @generated from field: int32 contest_id = 4;
   */
  contestId = 0;

  /**
   * @generated from field: bool is_public = 5;
   */
  isPublic = false;

  /**
   * @generated from field: backend.v1.Clarification.Answer answer = 6;
   */
  answer?: Clarification_Answer;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 7;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 8;
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
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "is_public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "answer", kind: "message", T: Clarification_Answer },
    { no: 7, name: "created_at", kind: "message", T: Timestamp },
    { no: 8, name: "updated_at", kind: "message", T: Timestamp },
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
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string content = 3;
   */
  content = "";

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
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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

/**
 * @generated from message backend.v1.Task
 */
export class Task extends Message<Task> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * 問題名
   *
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * 問題文
   *
   * @generated from field: string statement = 3;
   */
  statement = "";

  /**
   * 実行時間制限[ms]
   *
   * @generated from field: int32 exec_time_limit = 4;
   */
  execTimeLimit = 0;

  /**
   * 実行メモリ制限[MB]
   *
   * @generated from field: int32 exec_memory_limit = 5;
   */
  execMemoryLimit = 0;

  /**
   * 属するコンテストの ID
   *
   * @generated from field: repeated int32 contest_ids = 6;
   */
  contestIds: number[] = [];

  /**
   * Judge の type(完全一致、誤差など)
   *
   * @generated from field: judge.v1.JudgeType judge_type = 7;
   */
  judgeType?: JudgeType;

  /**
   * @generated from field: backend.v1.Difficulty difficulty = 8;
   */
  difficulty = Difficulty.DIFFICULTY_UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 9;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 10;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Task>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.Task";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "statement", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "exec_time_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "exec_memory_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "contest_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 7, name: "judge_type", kind: "message", T: JudgeType },
    { no: 8, name: "difficulty", kind: "enum", T: proto3.getEnumType(Difficulty) },
    { no: 9, name: "created_at", kind: "message", T: Timestamp },
    { no: 10, name: "updated_at", kind: "message", T: Timestamp, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Task {
    return new Task().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Task {
    return new Task().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Task {
    return new Task().fromJsonString(jsonString, options);
  }

  static equals(a: Task | PlainMessage<Task> | undefined, b: Task | PlainMessage<Task> | undefined): boolean {
    return proto3.util.equals(Task, a, b);
  }
}

/**
 * @generated from message backend.v1.MutationTask
 */
export class MutationTask extends Message<MutationTask> {
  /**
   * 問題名
   *
   * @generated from field: string title = 1;
   */
  title = "";

  /**
   * 問題文
   *
   * @generated from field: string statement = 2;
   */
  statement = "";

  /**
   * 実行時間制限[ms]
   *
   * @generated from field: int32 exec_time_limit = 3;
   */
  execTimeLimit = 0;

  /**
   * 実行メモリ制限[MB]
   *
   * @generated from field: int32 exec_memory_limit = 4;
   */
  execMemoryLimit = 0;

  /**
   * Judge の type(完全一致、誤差など)
   *
   * @generated from field: judge.v1.JudgeType judge_type = 5;
   */
  judgeType?: JudgeType;

  /**
   * @generated from field: backend.v1.Difficulty difficulty = 6;
   */
  difficulty = Difficulty.DIFFICULTY_UNSPECIFIED;

  constructor(data?: PartialMessage<MutationTask>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.MutationTask";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "statement", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "exec_time_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "exec_memory_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "judge_type", kind: "message", T: JudgeType },
    { no: 6, name: "difficulty", kind: "enum", T: proto3.getEnumType(Difficulty) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MutationTask {
    return new MutationTask().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MutationTask {
    return new MutationTask().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MutationTask {
    return new MutationTask().fromJsonString(jsonString, options);
  }

  static equals(a: MutationTask | PlainMessage<MutationTask> | undefined, b: MutationTask | PlainMessage<MutationTask> | undefined): boolean {
    return proto3.util.equals(MutationTask, a, b);
  }
}

/**
 * (slug, task_id) は複合 unique
 *
 * @generated from message backend.v1.TestcaseSet
 */
export class TestcaseSet extends Message<TestcaseSet> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * means name
   *
   * @generated from field: string slug = 2;
   */
  slug = "";

  /**
   * @generated from field: int32 score = 3;
   */
  score = 0;

  /**
   * @generated from field: repeated string testcase_slugs = 4;
   */
  testcaseSlugs: string[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 6;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: int32 task_id = 7;
   */
  taskId = 0;

  constructor(data?: PartialMessage<TestcaseSet>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.TestcaseSet";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "testcase_slugs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "updated_at", kind: "message", T: Timestamp, opt: true },
    { no: 7, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestcaseSet {
    return new TestcaseSet().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestcaseSet {
    return new TestcaseSet().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestcaseSet {
    return new TestcaseSet().fromJsonString(jsonString, options);
  }

  static equals(a: TestcaseSet | PlainMessage<TestcaseSet> | undefined, b: TestcaseSet | PlainMessage<TestcaseSet> | undefined): boolean {
    return proto3.util.equals(TestcaseSet, a, b);
  }
}

/**
 * @generated from message backend.v1.MutationTestcaseSet
 */
export class MutationTestcaseSet extends Message<MutationTestcaseSet> {
  /**
   * means name
   *
   * @generated from field: string slug = 1;
   */
  slug = "";

  /**
   * @generated from field: int32 score = 2;
   */
  score = 0;

  /**
   * @generated from field: repeated string testcase_slugs = 3;
   */
  testcaseSlugs: string[] = [];

  constructor(data?: PartialMessage<MutationTestcaseSet>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.MutationTestcaseSet";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "testcase_slugs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MutationTestcaseSet {
    return new MutationTestcaseSet().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MutationTestcaseSet {
    return new MutationTestcaseSet().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MutationTestcaseSet {
    return new MutationTestcaseSet().fromJsonString(jsonString, options);
  }

  static equals(a: MutationTestcaseSet | PlainMessage<MutationTestcaseSet> | undefined, b: MutationTestcaseSet | PlainMessage<MutationTestcaseSet> | undefined): boolean {
    return proto3.util.equals(MutationTestcaseSet, a, b);
  }
}

/**
 * (slug, task_id) は複合 unique
 *
 * @generated from message backend.v1.Testcase
 */
export class Testcase extends Message<Testcase> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string slug = 2;
   */
  slug = "";

  /**
   * @generated from field: optional string description = 3;
   */
  description?: string;

  /**
   * @generated from field: string input = 4;
   */
  input = "";

  /**
   * @generated from field: string output = 5;
   */
  output = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: int32 task_id = 8;
   */
  taskId = 0;

  constructor(data?: PartialMessage<Testcase>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.Testcase";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "input", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "output", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp, opt: true },
    { no: 8, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Testcase {
    return new Testcase().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Testcase {
    return new Testcase().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Testcase {
    return new Testcase().fromJsonString(jsonString, options);
  }

  static equals(a: Testcase | PlainMessage<Testcase> | undefined, b: Testcase | PlainMessage<Testcase> | undefined): boolean {
    return proto3.util.equals(Testcase, a, b);
  }
}

/**
 * @generated from message backend.v1.MutationTestcase
 */
export class MutationTestcase extends Message<MutationTestcase> {
  /**
   * @generated from field: string slug = 1;
   */
  slug = "";

  /**
   * @generated from field: optional string description = 2;
   */
  description?: string;

  /**
   * @generated from field: string input = 3;
   */
  input = "";

  /**
   * @generated from field: string output = 4;
   */
  output = "";

  constructor(data?: PartialMessage<MutationTestcase>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.MutationTestcase";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "input", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "output", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MutationTestcase {
    return new MutationTestcase().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MutationTestcase {
    return new MutationTestcase().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MutationTestcase {
    return new MutationTestcase().fromJsonString(jsonString, options);
  }

  static equals(a: MutationTestcase | PlainMessage<MutationTestcase> | undefined, b: MutationTestcase | PlainMessage<MutationTestcase> | undefined): boolean {
    return proto3.util.equals(MutationTestcase, a, b);
  }
}

/**
 * 提出の詳細
 *
 * @generated from message backend.v1.SubmissionDetail
 */
export class SubmissionDetail extends Message<SubmissionDetail> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * user
   *
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string username = 3;
   */
  username = "";

  /**
   * contest
   *
   * コンテスト中の提出でなければ null
   *
   * @generated from field: optional int32 contest_id = 4;
   */
  contestId?: number;

  /**
   * task
   *
   * @generated from field: int32 task_id = 5;
   */
  taskId = 0;

  /**
   * @generated from field: string task_title = 6;
   */
  taskTitle = "";

  /**
   * @generated from field: int32 score = 7;
   */
  score = 0;

  /**
   * submission
   *
   * @generated from field: string lang_id = 8;
   */
  langId = "";

  /**
   * @generated from field: string source_code = 9;
   */
  sourceCode = "";

  /**
   * judge
   *
   * ジャッジ中はnull
   *
   * @generated from field: optional judge.v1.JudgeStatus status = 10;
   */
  status?: JudgeStatus;

  /**
   * ジャッジ中はnull
   *
   * @generated from field: optional uint32 exec_time_ms = 11;
   */
  execTimeMs?: number;

  /**
   * ジャッジ中はnull
   *
   * @generated from field: optional uint32 exec_memory_kib = 12;
   */
  execMemoryKib?: number;

  /**
   * ジャッジ完了したテストケースの実行情報
   *
   * @generated from field: repeated backend.v1.TestcaseResult testcase_results = 13;
   */
  testcaseResults: TestcaseResult[] = [];

  /**
   * timestamp
   *
   * @generated from field: google.protobuf.Timestamp submitted_at = 14;
   */
  submittedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 15;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 16;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<SubmissionDetail>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmissionDetail";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "task_title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "source_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "status", kind: "enum", T: proto3.getEnumType(JudgeStatus), opt: true },
    { no: 11, name: "exec_time_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 12, name: "exec_memory_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 13, name: "testcase_results", kind: "message", T: TestcaseResult, repeated: true },
    { no: 14, name: "submitted_at", kind: "message", T: Timestamp },
    { no: 15, name: "created_at", kind: "message", T: Timestamp },
    { no: 16, name: "updated_at", kind: "message", T: Timestamp, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmissionDetail {
    return new SubmissionDetail().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmissionDetail {
    return new SubmissionDetail().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmissionDetail {
    return new SubmissionDetail().fromJsonString(jsonString, options);
  }

  static equals(a: SubmissionDetail | PlainMessage<SubmissionDetail> | undefined, b: SubmissionDetail | PlainMessage<SubmissionDetail> | undefined): boolean {
    return proto3.util.equals(SubmissionDetail, a, b);
  }
}

/**
 * 提出一覧用
 *
 * @generated from message backend.v1.SubmissionSummary
 */
export class SubmissionSummary extends Message<SubmissionSummary> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * user
   *
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string username = 3;
   */
  username = "";

  /**
   * contest
   *
   * @generated from field: optional int32 contest_id = 4;
   */
  contestId?: number;

  /**
   * task
   *
   * @generated from field: int32 task_id = 5;
   */
  taskId = 0;

  /**
   * @generated from field: string task_title = 6;
   */
  taskTitle = "";

  /**
   * @generated from field: int32 score = 7;
   */
  score = 0;

  /**
   * submission
   *
   * @generated from field: string lang_id = 8;
   */
  langId = "";

  /**
   * judge
   *
   * ジャッジ中はnull
   *
   * @generated from field: optional judge.v1.JudgeStatus judge_status = 9;
   */
  judgeStatus?: JudgeStatus;

  /**
   * ジャッジ中はnull
   *
   * @generated from field: optional uint32 exec_time_ms = 10;
   */
  execTimeMs?: number;

  /**
   * ジャッジ中はnull
   *
   * @generated from field: optional uint32 exec_memory_kib = 11;
   */
  execMemoryKib?: number;

  /**
   * timestamp
   *
   * @generated from field: google.protobuf.Timestamp submitted_at = 12;
   */
  submittedAt?: Timestamp;

  constructor(data?: PartialMessage<SubmissionSummary>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmissionSummary";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "task_title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "judge_status", kind: "enum", T: proto3.getEnumType(JudgeStatus), opt: true },
    { no: 10, name: "exec_time_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 11, name: "exec_memory_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 12, name: "submitted_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmissionSummary {
    return new SubmissionSummary().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmissionSummary {
    return new SubmissionSummary().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmissionSummary {
    return new SubmissionSummary().fromJsonString(jsonString, options);
  }

  static equals(a: SubmissionSummary | PlainMessage<SubmissionSummary> | undefined, b: SubmissionSummary | PlainMessage<SubmissionSummary> | undefined): boolean {
    return proto3.util.equals(SubmissionSummary, a, b);
  }
}

/**
 * テストケース一つあたりの結果
 *
 * @generated from message backend.v1.TestcaseResult
 */
export class TestcaseResult extends Message<TestcaseResult> {
  /**
   * @generated from field: string testcase_name = 1;
   */
  testcaseName = "";

  /**
   * @generated from field: judge.v1.JudgeStatus judge_status = 2;
   */
  judgeStatus = JudgeStatus.JUDGE_STATUS_UNSPECIFIED;

  /**
   * @generated from field: uint32 exec_time_ms = 3;
   */
  execTimeMs = 0;

  /**
   * @generated from field: uint32 exec_memory_kib = 4;
   */
  execMemoryKib = 0;

  constructor(data?: PartialMessage<TestcaseResult>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.TestcaseResult";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "testcase_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "judge_status", kind: "enum", T: proto3.getEnumType(JudgeStatus) },
    { no: 3, name: "exec_time_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "exec_memory_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestcaseResult {
    return new TestcaseResult().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestcaseResult {
    return new TestcaseResult().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestcaseResult {
    return new TestcaseResult().fromJsonString(jsonString, options);
  }

  static equals(a: TestcaseResult | PlainMessage<TestcaseResult> | undefined, b: TestcaseResult | PlainMessage<TestcaseResult> | undefined): boolean {
    return proto3.util.equals(TestcaseResult, a, b);
  }
}

/**
 * ジャッジの途中経過　WA(2/10) など。
 *
 * @generated from message backend.v1.JudgeProgress
 */
export class JudgeProgress extends Message<JudgeProgress> {
  /**
   * ジャッジ済みのケースが全てACであれば UNSPECIFIED そうでなければ最初に出た非ACの結果
   *
   * @generated from field: judge.v1.JudgeStatus status = 1;
   */
  status = JudgeStatus.JUDGE_STATUS_UNSPECIFIED;

  /**
   * テストケースの総数
   *
   * @generated from field: int32 total_testcases = 2;
   */
  totalTestcases = 0;

  /**
   * ジャッジ済みのテストケース数
   *
   * @generated from field: int32 completed_testcases = 3;
   */
  completedTestcases = 0;

  constructor(data?: PartialMessage<JudgeProgress>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.JudgeProgress";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(JudgeStatus) },
    { no: 2, name: "total_testcases", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "completed_testcases", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeProgress {
    return new JudgeProgress().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeProgress {
    return new JudgeProgress().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeProgress {
    return new JudgeProgress().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeProgress | PlainMessage<JudgeProgress> | undefined, b: JudgeProgress | PlainMessage<JudgeProgress> | undefined): boolean {
    return proto3.util.equals(JudgeProgress, a, b);
  }
}

