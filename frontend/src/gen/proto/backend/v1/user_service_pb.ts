// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/user_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { User } from "./user_resources_pb";

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
 * @generated from message backend.v1.ExistsUsernameRequest
 */
export class ExistsUsernameRequest extends Message<ExistsUsernameRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username = "";

  constructor(data?: PartialMessage<ExistsUsernameRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ExistsUsernameRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExistsUsernameRequest {
    return new ExistsUsernameRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExistsUsernameRequest {
    return new ExistsUsernameRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExistsUsernameRequest {
    return new ExistsUsernameRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ExistsUsernameRequest | PlainMessage<ExistsUsernameRequest> | undefined, b: ExistsUsernameRequest | PlainMessage<ExistsUsernameRequest> | undefined): boolean {
    return proto3.util.equals(ExistsUsernameRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.ExistsUsernameResponse
 */
export class ExistsUsernameResponse extends Message<ExistsUsernameResponse> {
  /**
   * @generated from field: bool exists = 1;
   */
  exists = false;

  constructor(data?: PartialMessage<ExistsUsernameResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ExistsUsernameResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exists", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExistsUsernameResponse {
    return new ExistsUsernameResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExistsUsernameResponse {
    return new ExistsUsernameResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExistsUsernameResponse {
    return new ExistsUsernameResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ExistsUsernameResponse | PlainMessage<ExistsUsernameResponse> | undefined, b: ExistsUsernameResponse | PlainMessage<ExistsUsernameResponse> | undefined): boolean {
    return proto3.util.equals(ExistsUsernameResponse, a, b);
  }
}

/**
 * @generated from message backend.v1.ExistsEmailRequest
 */
export class ExistsEmailRequest extends Message<ExistsEmailRequest> {
  /**
   * @generated from field: string email = 1;
   */
  email = "";

  constructor(data?: PartialMessage<ExistsEmailRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ExistsEmailRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExistsEmailRequest {
    return new ExistsEmailRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExistsEmailRequest {
    return new ExistsEmailRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExistsEmailRequest {
    return new ExistsEmailRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ExistsEmailRequest | PlainMessage<ExistsEmailRequest> | undefined, b: ExistsEmailRequest | PlainMessage<ExistsEmailRequest> | undefined): boolean {
    return proto3.util.equals(ExistsEmailRequest, a, b);
  }
}

/**
 * @generated from message backend.v1.ExistsEmailResponse
 */
export class ExistsEmailResponse extends Message<ExistsEmailResponse> {
  /**
   * @generated from field: bool exists = 1;
   */
  exists = false;

  constructor(data?: PartialMessage<ExistsEmailResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.ExistsEmailResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exists", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExistsEmailResponse {
    return new ExistsEmailResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExistsEmailResponse {
    return new ExistsEmailResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExistsEmailResponse {
    return new ExistsEmailResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ExistsEmailResponse | PlainMessage<ExistsEmailResponse> | undefined, b: ExistsEmailResponse | PlainMessage<ExistsEmailResponse> | undefined): boolean {
    return proto3.util.equals(ExistsEmailResponse, a, b);
  }
}

