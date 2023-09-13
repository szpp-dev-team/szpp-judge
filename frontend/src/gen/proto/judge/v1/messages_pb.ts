// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file judge/v1/messages.proto (package judge.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { JudgeType } from "./judge_type_pb";
import { ExecutionResultDetail, JudgeStatus, Testcase } from "./resources_pb";

/**
 * @generated from message judge.v1.PingRequest
 */
export class PingRequest extends Message<PingRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<PingRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.PingRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingRequest {
    return new PingRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingRequest {
    return new PingRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingRequest {
    return new PingRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PingRequest | PlainMessage<PingRequest> | undefined, b: PingRequest | PlainMessage<PingRequest> | undefined): boolean {
    return proto3.util.equals(PingRequest, a, b);
  }
}

/**
 * @generated from message judge.v1.PingResponse
 */
export class PingResponse extends Message<PingResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<PingResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.PingResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingResponse {
    return new PingResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingResponse {
    return new PingResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingResponse {
    return new PingResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PingResponse | PlainMessage<PingResponse> | undefined, b: PingResponse | PlainMessage<PingResponse> | undefined): boolean {
    return proto3.util.equals(PingResponse, a, b);
  }
}

/**
 * @generated from message judge.v1.JudgeRequest
 */
export class JudgeRequest extends Message<JudgeRequest> {
  /**
   * GCS上のパス
   *
   * @generated from field: string source_code_path = 1;
   */
  sourceCodePath = "";

  /**
   * @generated from field: string lang_id = 2;
   */
  langId = "";

  /**
   * @generated from field: judge.v1.JudgeType judge_type = 3;
   */
  judgeType?: JudgeType;

  /**
   * 実行時間制限 [ms]
   *
   * @generated from field: uint32 exec_time_limit_ms = 4;
   */
  execTimeLimitMs = 0;

  /**
   * 実行時メモリ制限 [MiB]
   *
   * @generated from field: uint32 exec_memory_limit_mib = 5;
   */
  execMemoryLimitMib = 0;

  /**
   * @generated from field: repeated judge.v1.Testcase testcases = 7;
   */
  testcases: Testcase[] = [];

  /**
   * @generated from field: optional bool want_result_detail = 8;
   */
  wantResultDetail?: boolean;

  /**
   * @generated from field: optional uint32 stdout_limit_kib = 9;
   */
  stdoutLimitKib?: number;

  /**
   * @generated from field: optional uint32 stderr_limit_kib = 10;
   */
  stderrLimitKib?: number;

  /**
   * @generated from field: int32 submit_id = 11;
   */
  submitId = 0;

  constructor(data?: PartialMessage<JudgeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source_code_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "judge_type", kind: "message", T: JudgeType },
    { no: 4, name: "exec_time_limit_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "exec_memory_limit_mib", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 7, name: "testcases", kind: "message", T: Testcase, repeated: true },
    { no: 8, name: "want_result_detail", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 9, name: "stdout_limit_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 10, name: "stderr_limit_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 11, name: "submit_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeRequest {
    return new JudgeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeRequest {
    return new JudgeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeRequest {
    return new JudgeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeRequest | PlainMessage<JudgeRequest> | undefined, b: JudgeRequest | PlainMessage<JudgeRequest> | undefined): boolean {
    return proto3.util.equals(JudgeRequest, a, b);
  }
}

/**
 * @generated from message judge.v1.JudgeResponse
 */
export class JudgeResponse extends Message<JudgeResponse> {
  /**
   * テストケースの ID(DB に保存するときに使う)
   *
   * @generated from field: uint32 testcase_id = 1;
   */
  testcaseId = 0;

  /**
   * @generated from field: judge.v1.JudgeStatus status = 2;
   */
  status = JudgeStatus.JUDGE_STATUS_UNSPECIFIED;

  /**
   * @generated from field: uint32 exec_time_ms = 3;
   */
  execTimeMs = 0;

  /**
   * @generated from field: uint32 exec_memory_kib = 4;
   */
  execMemoryKib = 0;

  /**
   * @generated from field: string compiler_message = 5;
   */
  compilerMessage = "";

  /**
   * want_result_detail=true のときに値がセットされる
   *
   * @generated from field: optional judge.v1.ExecutionResultDetail detail = 6;
   */
  detail?: ExecutionResultDetail;

  constructor(data?: PartialMessage<JudgeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "testcase_id", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(JudgeStatus) },
    { no: 3, name: "exec_time_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "exec_memory_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "compiler_message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "detail", kind: "message", T: ExecutionResultDetail, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeResponse {
    return new JudgeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeResponse {
    return new JudgeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeResponse {
    return new JudgeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeResponse | PlainMessage<JudgeResponse> | undefined, b: JudgeResponse | PlainMessage<JudgeResponse> | undefined): boolean {
    return proto3.util.equals(JudgeResponse, a, b);
  }
}

/**
 * AtCoder でいうコードテスト向けのリクエスト
 *
 * @generated from message judge.v1.RunRequest
 */
export class RunRequest extends Message<RunRequest> {
  /**
   * @generated from field: string source_code = 1;
   */
  sourceCode = "";

  /**
   * @generated from field: string lang_id = 2;
   */
  langId = "";

  /**
   * @generated from field: string stdin = 3;
   */
  stdin = "";

  /**
   * 実行時間制限 [ms]
   *
   * @generated from field: uint32 exec_time_limit_ms = 4;
   */
  execTimeLimitMs = 0;

  /**
   * 実行時メモリ制限 [MiB]
   *
   * @generated from field: uint32 exec_memory_limit_mib = 5;
   */
  execMemoryLimitMib = 0;

  /**
   * @generated from field: optional uint32 stdout_limit_kib = 6;
   */
  stdoutLimitKib?: number;

  /**
   * @generated from field: optional uint32 stderr_limit_kib = 7;
   */
  stderrLimitKib?: number;

  constructor(data?: PartialMessage<RunRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.RunRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "stdin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "exec_time_limit_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "exec_memory_limit_mib", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 6, name: "stdout_limit_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 7, name: "stderr_limit_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RunRequest {
    return new RunRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RunRequest {
    return new RunRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RunRequest {
    return new RunRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RunRequest | PlainMessage<RunRequest> | undefined, b: RunRequest | PlainMessage<RunRequest> | undefined): boolean {
    return proto3.util.equals(RunRequest, a, b);
  }
}

/**
 * @generated from message judge.v1.RunResponse
 */
export class RunResponse extends Message<RunResponse> {
  /**
   * @generated from field: uint32 exec_time_ms = 1;
   */
  execTimeMs = 0;

  /**
   * @generated from field: uint32 exec_memory_kib = 2;
   */
  execMemoryKib = 0;

  /**
   * @generated from field: string compiler_message = 3;
   */
  compilerMessage = "";

  /**
   * @generated from field: judge.v1.ExecutionResultDetail detail = 4;
   */
  detail?: ExecutionResultDetail;

  constructor(data?: PartialMessage<RunResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.RunResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exec_time_ms", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "exec_memory_kib", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "compiler_message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "detail", kind: "message", T: ExecutionResultDetail },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RunResponse {
    return new RunResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RunResponse {
    return new RunResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RunResponse {
    return new RunResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RunResponse | PlainMessage<RunResponse> | undefined, b: RunResponse | PlainMessage<RunResponse> | undefined): boolean {
    return proto3.util.equals(RunResponse, a, b);
  }
}

