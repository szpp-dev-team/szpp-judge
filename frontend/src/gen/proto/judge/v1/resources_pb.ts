// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file judge/v1/resources.proto (package judge.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum judge.v1.JudgeStatus
 */
export enum JudgeStatus {
  /**
   * @generated from enum value: JUDGE_STATUS_UNSPECIFIED = 0;
   */
  JUDGE_STATUS_UNSPECIFIED = 0,

  /**
   * @generated from enum value: AC = 1;
   */
  AC = 1,

  /**
   * @generated from enum value: CE = 2;
   */
  CE = 2,

  /**
   * @generated from enum value: IE = 3;
   */
  IE = 3,

  /**
   * @generated from enum value: MLE = 4;
   */
  MLE = 4,

  /**
   * @generated from enum value: OLE = 5;
   */
  OLE = 5,

  /**
   * @generated from enum value: RE = 6;
   */
  RE = 6,

  /**
   * @generated from enum value: TLE = 7;
   */
  TLE = 7,

  /**
   * @generated from enum value: WA = 8;
   */
  WA = 8,
}
// Retrieve enum metadata with: proto3.getEnumType(JudgeStatus)
proto3.util.setEnumType(JudgeStatus, "judge.v1.JudgeStatus", [
  { no: 0, name: "JUDGE_STATUS_UNSPECIFIED" },
  { no: 1, name: "AC" },
  { no: 2, name: "CE" },
  { no: 3, name: "IE" },
  { no: 4, name: "MLE" },
  { no: 5, name: "OLE" },
  { no: 6, name: "RE" },
  { no: 7, name: "TLE" },
  { no: 8, name: "WA" },
]);

/**
 * @generated from message judge.v1.Testcase
 */
export class Testcase extends Message<Testcase> {
  /**
   * @generated from field: string input_path = 1;
   */
  inputPath = "";

  /**
   * @generated from field: string expected_path = 2;
   */
  expectedPath = "";

  constructor(data?: PartialMessage<Testcase>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.Testcase";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "input_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expected_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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
 * @generated from message judge.v1.ExecutionResultDetail
 */
export class ExecutionResultDetail extends Message<ExecutionResultDetail> {
  /**
   * TLE 等で kill した場合は exit_code はナシ
   *
   * @generated from field: optional uint32 exit_code = 1;
   */
  exitCode?: number;

  /**
   * @generated from field: string stdin = 2;
   */
  stdin = "";

  /**
   * @generated from field: string stdout = 3;
   */
  stdout = "";

  /**
   * @generated from field: string stderr = 4;
   */
  stderr = "";

  /**
   * stdout が巨大すぎて途中から省略した場合は true
   *
   * @generated from field: bool stdout_truncated = 5;
   */
  stdoutTruncated = false;

  /**
   * stderr が巨大すぎて途中から省略した場合は true
   *
   * @generated from field: bool stderr_truncated = 6;
   */
  stderrTruncated = false;

  constructor(data?: PartialMessage<ExecutionResultDetail>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.ExecutionResultDetail";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exit_code", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 2, name: "stdin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "stdout", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "stderr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "stdout_truncated", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "stderr_truncated", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExecutionResultDetail {
    return new ExecutionResultDetail().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExecutionResultDetail {
    return new ExecutionResultDetail().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExecutionResultDetail {
    return new ExecutionResultDetail().fromJsonString(jsonString, options);
  }

  static equals(a: ExecutionResultDetail | PlainMessage<ExecutionResultDetail> | undefined, b: ExecutionResultDetail | PlainMessage<ExecutionResultDetail> | undefined): boolean {
    return proto3.util.equals(ExecutionResultDetail, a, b);
  }
}

