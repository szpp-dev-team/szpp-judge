// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/resources.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

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
 * @generated from message backend.v1.PersonalStandings
 */
export class PersonalStandings extends Message<PersonalStandings> {
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
   * @generated from field: repeated backend.v1.PersonalStandings.TaskDetail task_detail_list = 7;
   */
  taskDetailList: PersonalStandings_TaskDetail[] = [];

  constructor(data?: PartialMessage<PersonalStandings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.PersonalStandings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "rank", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "total_score", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "total_penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "latest_ac_at", kind: "message", T: Timestamp },
    { no: 7, name: "task_detail_list", kind: "message", T: PersonalStandings_TaskDetail, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonalStandings {
    return new PersonalStandings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonalStandings {
    return new PersonalStandings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonalStandings {
    return new PersonalStandings().fromJsonString(jsonString, options);
  }

  static equals(a: PersonalStandings | PlainMessage<PersonalStandings> | undefined, b: PersonalStandings | PlainMessage<PersonalStandings> | undefined): boolean {
    return proto3.util.equals(PersonalStandings, a, b);
  }
}

/**
 * @generated from message backend.v1.PersonalStandings.TaskDetail
 */
export class PersonalStandings_TaskDetail extends Message<PersonalStandings_TaskDetail> {
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
   * @generated from field: google.protobuf.Timestamp ac_required = 4;
   */
  acRequired?: Timestamp;

  constructor(data?: PartialMessage<PersonalStandings_TaskDetail>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "backend.v1.PersonalStandings.TaskDetail";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "penalty_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "ac_submit_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "ac_required", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonalStandings_TaskDetail {
    return new PersonalStandings_TaskDetail().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonalStandings_TaskDetail {
    return new PersonalStandings_TaskDetail().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonalStandings_TaskDetail {
    return new PersonalStandings_TaskDetail().fromJsonString(jsonString, options);
  }

  static equals(a: PersonalStandings_TaskDetail | PlainMessage<PersonalStandings_TaskDetail> | undefined, b: PersonalStandings_TaskDetail | PlainMessage<PersonalStandings_TaskDetail> | undefined): boolean {
    return proto3.util.equals(PersonalStandings_TaskDetail, a, b);
  }
}

