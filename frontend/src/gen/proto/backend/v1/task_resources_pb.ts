// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/task_resources.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { JudgeType } from "../../judge/v1/judge_type_pb";

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
   * Judge の type(完全一致、誤差など)
   *
   * @generated from field: judge.v1.JudgeType judge_type = 6;
   */
  judgeType?: JudgeType;

  /**
   * @generated from field: backend.v1.Difficulty difficulty = 7;
   */
  difficulty = Difficulty.DIFFICULTY_UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 8;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp updated_at = 9;
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
    { no: 6, name: "judge_type", kind: "message", T: JudgeType },
    { no: 7, name: "difficulty", kind: "enum", T: proto3.getEnumType(Difficulty) },
    { no: 8, name: "created_at", kind: "message", T: Timestamp },
    { no: 9, name: "updated_at", kind: "message", T: Timestamp, opt: true },
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
   * update の時のみ
   *
   * @generated from field: optional int32 id = 1;
   */
  id?: number;

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
   * Judge の type(完全一致、誤差など)
   *
   * @generated from field: judge.v1.JudgeType judge_type = 6;
   */
  judgeType?: JudgeType;

  /**
   * @generated from field: backend.v1.Difficulty difficulty = 7;
   */
  difficulty = Difficulty.DIFFICULTY_UNSPECIFIED;

  /**
   * @generated from field: bool is_public = 8;
   */
  isPublic = false;

  constructor(data?: PartialMessage<MutationTask>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.MutationTask";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "statement", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "exec_time_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "exec_memory_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "judge_type", kind: "message", T: JudgeType },
    { no: 7, name: "difficulty", kind: "enum", T: proto3.getEnumType(Difficulty) },
    { no: 8, name: "is_public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
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
   * 得点比率(%)。TestcaseSet.score_ratio の総和が100になるように設定する。
   *
   * @generated from field: int32 score_ratio = 3;
   */
  scoreRatio = 0;

  /**
   * @generated from field: bool is_sample = 4;
   */
  isSample = false;

  /**
   * @generated from field: repeated string testcase_slugs = 5;
   */
  testcaseSlugs: string[] = [];

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

  constructor(data?: PartialMessage<TestcaseSet>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.TestcaseSet";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "score_ratio", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "is_sample", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "testcase_slugs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp, opt: true },
    { no: 8, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
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
   * @generated from field: int32 score_ratio = 2;
   */
  scoreRatio = 0;

  /**
   * @generated from field: bool is_sample = 3;
   */
  isSample = false;

  /**
   * @generated from field: repeated string testcase_slugs = 4;
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
    { no: 2, name: "score_ratio", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "is_sample", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "testcase_slugs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
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

