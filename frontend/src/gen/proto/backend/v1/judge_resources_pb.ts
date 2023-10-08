// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/judge_resources.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { JudgeStatus } from "../../judge/v1/resources_pb";

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
   * ジャッジ中は judge.v1.JudgeStatus.WJ
   *
   * @generated from field: judge.v1.JudgeStatus status = 10;
   */
  status = JudgeStatus.JUDGE_STATUS_UNSPECIFIED;

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
    { no: 10, name: "status", kind: "enum", T: proto3.getEnumType(JudgeStatus) },
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

