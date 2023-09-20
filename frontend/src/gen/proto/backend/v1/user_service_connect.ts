// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/user_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, ExistsEmailRequest, ExistsEmailResponse, ExistsUsernameRequest, ExistsUsernameResponse, GetUserRequest, GetUserResponse } from "./user_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service backend.v1.UserService
 */
export const UserService = {
  typeName: "backend.v1.UserService",
  methods: {
    /**
     * 指定された User を取得する
     *
     * @generated from rpc backend.v1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: GetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * User を新たに作成する
     *
     * @generated from rpc backend.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Username が存在するかを確認する
     *
     * @generated from rpc backend.v1.UserService.ExistsUsername
     */
    existsUsername: {
      name: "ExistsUsername",
      I: ExistsUsernameRequest,
      O: ExistsUsernameResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Email が存在するかを確認する
     *
     * @generated from rpc backend.v1.UserService.ExistsEmail
     */
    existsEmail: {
      name: "ExistsEmail",
      I: ExistsEmailRequest,
      O: ExistsEmailResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
