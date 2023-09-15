// @generated by protoc-gen-connect-query v0.4.3 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/services.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, GetUserRequest, GetUserResponse } from "./messages_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService } from "@connectrpc/connect-query";

export const typeName = "backend.v1.UserService";

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
  }
} as const;

/**
 * 指定された User を取得する
 *
 * @generated from rpc backend.v1.UserService.GetUser
 */
export const getUser = createQueryService({
  service: UserService,
}).getUser;

/**
 * User を新たに作成する
 *
 * @generated from rpc backend.v1.UserService.CreateUser
 */
export const createUser = createQueryService({
  service: UserService,
}).createUser;
