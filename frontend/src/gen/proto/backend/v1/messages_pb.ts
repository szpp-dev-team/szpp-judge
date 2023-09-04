// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/messages.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { JudgeProgress, MutationTask, MutationTestcase, MutationTestcaseSet, Submission, Task, Testcase, TestcaseSet, User } from "./resources_pb";

/**
 * @generated from message backend.v1.GetUserRequest
 */
export class GetUserRequest extends Message<GetUserRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username = "";

  constructor(data?: PartialMessage<GetUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest {
    return new GetUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean {
    return proto3.util.equals(GetUserRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.GetUserResponse
 */
export class GetUserResponse extends Message<GetUserResponse> {
  /**
   * @generated from field: backend.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<GetUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserResponse {
    return new GetUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserResponse {
    return new GetUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserResponse {
    return new GetUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserResponse | PlainMessage<GetUserResponse> | undefined, b: GetUserResponse | PlainMessage<GetUserResponse> | undefined): boolean {
    return proto3.util.equals(GetUserResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.CreateUserRequest
 */
export class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username = "";

  /**
   * @generated from field: string password = 2;
   */
  password = "";

  /**
   * @generated from field: string email = 3;
   */
  email = "";

  constructor(data?: PartialMessage<CreateUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.CreateUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean {
    return proto3.util.equals(CreateUserRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.CreateUserResponse
 */
export class CreateUserResponse extends Message<CreateUserResponse> {
  /**
   * @generated from field: backend.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<CreateUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.CreateUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined, b: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined): boolean {
    return proto3.util.equals(CreateUserResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.LoginRequest
 */
export class LoginRequest extends Message<LoginRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username = "";

  /**
   * @generated from field: string password = 2;
   */
  password = "";

  constructor(data?: PartialMessage<LoginRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.LoginRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginRequest {
    return new LoginRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LoginRequest | PlainMessage<LoginRequest> | undefined, b: LoginRequest | PlainMessage<LoginRequest> | undefined): boolean {
    return proto3.util.equals(LoginRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.LoginResponse
 */
export class LoginResponse extends Message<LoginResponse> {
  /**
   * @generated from field: backend.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<LoginResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.LoginResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginResponse {
    return new LoginResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LoginResponse | PlainMessage<LoginResponse> | undefined, b: LoginResponse | PlainMessage<LoginResponse> | undefined): boolean {
    return proto3.util.equals(LoginResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.LogoutRequest
 */
export class LogoutRequest extends Message<LogoutRequest> {
  constructor(data?: PartialMessage<LogoutRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.LogoutRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutRequest {
    return new LogoutRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutRequest {
    return new LogoutRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutRequest {
    return new LogoutRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LogoutRequest | PlainMessage<LogoutRequest> | undefined, b: LogoutRequest | PlainMessage<LogoutRequest> | undefined): boolean {
    return proto3.util.equals(LogoutRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.LogoutResponse
 */
export class LogoutResponse extends Message<LogoutResponse> {
  constructor(data?: PartialMessage<LogoutResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.LogoutResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutResponse {
    return new LogoutResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutResponse {
    return new LogoutResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutResponse {
    return new LogoutResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LogoutResponse | PlainMessage<LogoutResponse> | undefined, b: LogoutResponse | PlainMessage<LogoutResponse> | undefined): boolean {
    return proto3.util.equals(LogoutResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.CreateTaskRequest
 */
export class CreateTaskRequest extends Message<CreateTaskRequest> {
  /**
   * @generated from field: backend.v1.MutationTask task = 1;
   */
  task?: MutationTask;

  /**
   * @generated from field: repeated backend.v1.MutationTestcaseSet testcase_sets = 2;
   */
  testcaseSets: MutationTestcaseSet[] = [];

  /**
   * @generated from field: repeated backend.v1.MutationTestcase testcases = 3;
   */
  testcases: MutationTestcase[] = [];

  constructor(data?: PartialMessage<CreateTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.CreateTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: MutationTask },
    { no: 2, name: "testcase_sets", kind: "message", T: MutationTestcaseSet, repeated: true },
    { no: 3, name: "testcases", kind: "message", T: MutationTestcase, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTaskRequest | PlainMessage<CreateTaskRequest> | undefined, b: CreateTaskRequest | PlainMessage<CreateTaskRequest> | undefined): boolean {
    return proto3.util.equals(CreateTaskRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.CreateTaskResponse
 */
export class CreateTaskResponse extends Message<CreateTaskResponse> {
  /**
   * @generated from field: backend.v1.Task task = 1;
   */
  task?: Task;

  /**
   * @generated from field: repeated backend.v1.TestcaseSet testcase_sets = 2;
   */
  testcaseSets: TestcaseSet[] = [];

  /**
   * @generated from field: repeated backend.v1.Testcase testcases = 3;
   */
  testcases: Testcase[] = [];

  constructor(data?: PartialMessage<CreateTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.CreateTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
    { no: 2, name: "testcase_sets", kind: "message", T: TestcaseSet, repeated: true },
    { no: 3, name: "testcases", kind: "message", T: Testcase, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined, b: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined): boolean {
    return proto3.util.equals(CreateTaskResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.GetTaskRequest
 */
export class GetTaskRequest extends Message<GetTaskRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<GetTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskRequest | PlainMessage<GetTaskRequest> | undefined, b: GetTaskRequest | PlainMessage<GetTaskRequest> | undefined): boolean {
    return proto3.util.equals(GetTaskRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.GetTaskResponse
 */
export class GetTaskResponse extends Message<GetTaskResponse> {
  /**
   * @generated from field: backend.v1.Task task = 1;
   */
  task?: Task;

  /**
   * @generated from field: repeated backend.v1.Testcase testcases = 2;
   */
  testcases: Testcase[] = [];

  constructor(data?: PartialMessage<GetTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
    { no: 2, name: "testcases", kind: "message", T: Testcase, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskResponse | PlainMessage<GetTaskResponse> | undefined, b: GetTaskResponse | PlainMessage<GetTaskResponse> | undefined): boolean {
    return proto3.util.equals(GetTaskResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.UpdateTaskRequest
 */
export class UpdateTaskRequest extends Message<UpdateTaskRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: backend.v1.MutationTask task = 2;
   */
  task?: MutationTask;

  /**
   * @generated from field: repeated backend.v1.MutationTestcaseSet testcase_sets = 3;
   */
  testcaseSets: MutationTestcaseSet[] = [];

  /**
   * @generated from field: repeated backend.v1.MutationTestcase testcases = 4;
   */
  testcases: MutationTestcase[] = [];

  constructor(data?: PartialMessage<UpdateTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.UpdateTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "task", kind: "message", T: MutationTask },
    { no: 3, name: "testcase_sets", kind: "message", T: MutationTestcaseSet, repeated: true },
    { no: 4, name: "testcases", kind: "message", T: MutationTestcase, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTaskRequest {
    return new UpdateTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTaskRequest {
    return new UpdateTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTaskRequest {
    return new UpdateTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTaskRequest | PlainMessage<UpdateTaskRequest> | undefined, b: UpdateTaskRequest | PlainMessage<UpdateTaskRequest> | undefined): boolean {
    return proto3.util.equals(UpdateTaskRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.UpdateTaskResponse
 */
export class UpdateTaskResponse extends Message<UpdateTaskResponse> {
  /**
   * @generated from field: backend.v1.Task task = 1;
   */
  task?: Task;

  /**
   * @generated from field: repeated backend.v1.TestcaseSet testcase_sets = 2;
   */
  testcaseSets: TestcaseSet[] = [];

  /**
   * @generated from field: repeated backend.v1.Testcase testcases = 3;
   */
  testcases: Testcase[] = [];

  constructor(data?: PartialMessage<UpdateTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.UpdateTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
    { no: 2, name: "testcase_sets", kind: "message", T: TestcaseSet, repeated: true },
    { no: 3, name: "testcases", kind: "message", T: Testcase, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTaskResponse {
    return new UpdateTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTaskResponse {
    return new UpdateTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTaskResponse {
    return new UpdateTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTaskResponse | PlainMessage<UpdateTaskResponse> | undefined, b: UpdateTaskResponse | PlainMessage<UpdateTaskResponse> | undefined): boolean {
    return proto3.util.equals(UpdateTaskResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.PingRequest
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
  static readonly typeName = "backend.v1.PingRequest";
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
 * @generated from message backend.v1.PingResponse
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
  static readonly typeName = "backend.v1.PingResponse";
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
   * @generated from field: string langage = 3;
   */
  langage = "";

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
    { no: 3, name: "langage", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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
 * @generated from message backend.v1.GetSubmissionRequest
 */
export class GetSubmissionRequest extends Message<GetSubmissionRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<GetSubmissionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetSubmissionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubmissionRequest {
    return new GetSubmissionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubmissionRequest {
    return new GetSubmissionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubmissionRequest {
    return new GetSubmissionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubmissionRequest | PlainMessage<GetSubmissionRequest> | undefined, b: GetSubmissionRequest | PlainMessage<GetSubmissionRequest> | undefined): boolean {
    return proto3.util.equals(GetSubmissionRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.GetSubmissionResponse
 */
export class GetSubmissionResponse extends Message<GetSubmissionResponse> {
  /**
   * @generated from field: backend.v1.Submission submission = 1;
   */
  submission?: Submission;

  constructor(data?: PartialMessage<GetSubmissionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.GetSubmissionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submission", kind: "message", T: Submission },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubmissionResponse {
    return new GetSubmissionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubmissionResponse {
    return new GetSubmissionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubmissionResponse {
    return new GetSubmissionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubmissionResponse | PlainMessage<GetSubmissionResponse> | undefined, b: GetSubmissionResponse | PlainMessage<GetSubmissionResponse> | undefined): boolean {
    return proto3.util.equals(GetSubmissionResponse, a, b);
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
 * @generated from message backend.v1.SubmissionSummary
 */
export class SubmissionSummary extends Message<SubmissionSummary> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: int32 user_id = 2;
   */
  userId = 0;

  /**
   * @generated from field: string user_name = 3;
   */
  userName = "";

  /**
   * @generated from field: optional int32 contest_id = 4;
   */
  contestId?: number;

  /**
   * @generated from field: int32 task_id = 5;
   */
  taskId = 0;

  /**
   * @generated from field: string task_title = 6;
   */
  taskTitle = "";

  /**
   * @generated from field: string lang_id = 7;
   */
  langId = "";

  /**
   * @generated from field: string status = 8;
   */
  status = "";

  /**
   * @generated from field: int32 score = 9;
   */
  score = 0;

  /**
   * @generated from field: optional uint32 exec_time = 10;
   */
  execTime?: number;

  /**
   * @generated from field: optional uint32 memory = 11;
   */
  memory?: number;

  /**
   * @generated from field: google.protobuf.Timestamp submited_at = 12;
   */
  submitedAt?: Timestamp;

  constructor(data?: PartialMessage<SubmissionSummary>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.SubmissionSummary";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "contest_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "task_title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "lang_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "exec_time", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 11, name: "memory", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 12, name: "submited_at", kind: "message", T: Timestamp },
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

