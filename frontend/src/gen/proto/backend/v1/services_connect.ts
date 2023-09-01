// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/services.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateContestRequest, CreateContestResponse, CreateUserRequest, CreateUserResponse, GetUserRequest, GetUserResponse, LoginRequest, LoginResponse, LogoutRequest, LogoutResponse, PingRequest, PingResponse, SearchContestRequest, SearchContestResponse } from "./messages_pb";
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
  }
} as const;

/**
 * @generated from service backend.v1.AuthService
 */
export const AuthService = {
  typeName: "backend.v1.AuthService",
  methods: {
    /**
     * ログイン
     *
     * @generated from rpc backend.v1.AuthService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ログアウト
     *
     * @generated from rpc backend.v1.AuthService.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutRequest,
      O: LogoutResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.HealthcheckService
 */
export const HealthcheckService = {
  typeName: "backend.v1.HealthcheckService",
  methods: {
    /**
     * @generated from rpc backend.v1.HealthcheckService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service backend.v1.ContestService
 */
export const ContestService = {
  typeName: "backend.v1.ContestService",
  methods: {
    /**
     * @generated from rpc backend.v1.ContestService.CreateContest
     */
    createContest: {
      name: "CreateContest",
      I: CreateContestRequest,
      O: CreateContestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc backend.v1.ContestService.SearchContest
     */
    searchContest: {
      name: "SearchContest",
      I: SearchContestRequest,
      O: SearchContestResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

