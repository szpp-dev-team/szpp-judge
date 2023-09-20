// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/judge_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { JudgeProgress, SubmissionDetail, SubmissionSummary } from "./judge_resources_pb";

/**
 * @generated from message backend.v1.SubmitRequest
 */
export class SubmitRequest extends Message<SubmitRequest> {
  /**
   * @generated from field: optional int32 contest_id = 1;
   */
  contestId?: number;

  /**
   * @generated from field: int32 task_id = 2;
   */
  taskId = 0;

  /**
   * @generated from field: string lang_id = 3;
   */
  langId = "";

  /**
   * @generated from field: string source_code = 4;
   */
  sourceCode = "";

  constructor(data?: PartialMessage<SubmitRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmitRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "source_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitRequest {
    return new SubmitRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitRequest {
    return new SubmitRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitRequest {
    return new SubmitRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitRequest | PlainMessage<SubmitRequest> | undefined, b: SubmitRequest | PlainMessage<SubmitRequest> | undefined): boolean {
    return proto3.util.equals(SubmitRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.SubmitResponse
 */
export class SubmitResponse extends Message<SubmitResponse> {
  /**
   * @generated from field: int32 submission_id = 1;
   */
  submissionId = 0;

  constructor(data?: PartialMessage<SubmitResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmitResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submission_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitResponse {
    return new SubmitResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitResponse {
    return new SubmitResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitResponse {
    return new SubmitResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitResponse | PlainMessage<SubmitResponse> | undefined, b: SubmitResponse | PlainMessage<SubmitResponse> | undefined): boolean {
    return proto3.util.equals(SubmitResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.GetSubmissionDetailRequest
 */
export class GetSubmissionDetailRequest extends Message<GetSubmissionDetailRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<GetSubmissionDetailRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetSubmissionDetailRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubmissionDetailRequest {
    return new GetSubmissionDetailRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubmissionDetailRequest {
    return new GetSubmissionDetailRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubmissionDetailRequest {
    return new GetSubmissionDetailRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubmissionDetailRequest | PlainMessage<GetSubmissionDetailRequest> | undefined, b: GetSubmissionDetailRequest | PlainMessage<GetSubmissionDetailRequest> | undefined): boolean {
    return proto3.util.equals(GetSubmissionDetailRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.GetSubmissionDetailResponse
 */
export class GetSubmissionDetailResponse extends Message<GetSubmissionDetailResponse> {
  /**
   * @generated from field: backend.v1.SubmissionDetail submission_detail = 1;
   */
  submissionDetail?: SubmissionDetail;

  constructor(data?: PartialMessage<GetSubmissionDetailResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetSubmissionDetailResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submission_detail", kind: "message", T: SubmissionDetail },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubmissionDetailResponse {
    return new GetSubmissionDetailResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubmissionDetailResponse {
    return new GetSubmissionDetailResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubmissionDetailResponse {
    return new GetSubmissionDetailResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubmissionDetailResponse | PlainMessage<GetSubmissionDetailResponse> | undefined, b: GetSubmissionDetailResponse | PlainMessage<GetSubmissionDetailResponse> | undefined): boolean {
    return proto3.util.equals(GetSubmissionDetailResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.ListSubmissionsRequest
 */
export class ListSubmissionsRequest extends Message<ListSubmissionsRequest> {
  /**
   * @generated from field: optional int32 contest_id = 1;
   */
  contestId?: number;

  /**
   * @generated from field: optional int32 user_id = 2;
   */
  userId?: number;

  constructor(data?: PartialMessage<ListSubmissionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ListSubmissionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSubmissionsRequest {
    return new ListSubmissionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSubmissionsRequest {
    return new ListSubmissionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSubmissionsRequest {
    return new ListSubmissionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSubmissionsRequest | PlainMessage<ListSubmissionsRequest> | undefined, b: ListSubmissionsRequest | PlainMessage<ListSubmissionsRequest> | undefined): boolean {
    return proto3.util.equals(ListSubmissionsRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.ListSubmissionsResponse
 */
export class ListSubmissionsResponse extends Message<ListSubmissionsResponse> {
  /**
   * @generated from field: repeated backend.v1.SubmissionSummary submissions = 1;
   */
  submissions: SubmissionSummary[] = [];

  constructor(data?: PartialMessage<ListSubmissionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ListSubmissionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submissions", kind: "message", T: SubmissionSummary, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSubmissionsResponse {
    return new ListSubmissionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSubmissionsResponse {
    return new ListSubmissionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSubmissionsResponse {
    return new ListSubmissionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSubmissionsResponse | PlainMessage<ListSubmissionsResponse> | undefined, b: ListSubmissionsResponse | PlainMessage<ListSubmissionsResponse> | undefined): boolean {
    return proto3.util.equals(ListSubmissionsResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.GetJudgeProgressRequest
 */
export class GetJudgeProgressRequest extends Message<GetJudgeProgressRequest> {
  /**
   * @generated from field: int32 submission_id = 1;
   */
  submissionId = 0;

  constructor(data?: PartialMessage<GetJudgeProgressRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetJudgeProgressRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submission_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetJudgeProgressRequest {
    return new GetJudgeProgressRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetJudgeProgressRequest {
    return new GetJudgeProgressRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetJudgeProgressRequest {
    return new GetJudgeProgressRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetJudgeProgressRequest | PlainMessage<GetJudgeProgressRequest> | undefined, b: GetJudgeProgressRequest | PlainMessage<GetJudgeProgressRequest> | undefined): boolean {
    return proto3.util.equals(GetJudgeProgressRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.GetJudgeProgressResponse
 */
export class GetJudgeProgressResponse extends Message<GetJudgeProgressResponse> {
  /**
   * @generated from field: backend.v1.JudgeProgress judge_progress = 1;
   */
  judgeProgress?: JudgeProgress;

  constructor(data?: PartialMessage<GetJudgeProgressResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetJudgeProgressResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "judge_progress", kind: "message", T: JudgeProgress },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetJudgeProgressResponse {
    return new GetJudgeProgressResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetJudgeProgressResponse {
    return new GetJudgeProgressResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetJudgeProgressResponse {
    return new GetJudgeProgressResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetJudgeProgressResponse | PlainMessage<GetJudgeProgressResponse> | undefined, b: GetJudgeProgressResponse | PlainMessage<GetJudgeProgressResponse> | undefined): boolean {
    return proto3.util.equals(GetJudgeProgressResponse, a, b);
  }
}
