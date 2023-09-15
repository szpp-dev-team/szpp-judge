// @generated by protoc-gen-connect-query v0.4.3 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/services.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { LoginRequest, LoginResponse, LogoutRequest, LogoutResponse, RefreshAccessTokenRequest, RefreshAccessTokenResponse } from "./messages_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService } from "@connectrpc/connect-query";

export const typeName = "backend.v1.AuthService";

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
    /**
     * refresh token を使って access token を更新する
     *
     * @generated from rpc backend.v1.AuthService.RefreshAccessToken
     */
    refreshAccessToken: {
      name: "RefreshAccessToken",
      I: RefreshAccessTokenRequest,
      O: RefreshAccessTokenResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * ログイン
 *
 * @generated from rpc backend.v1.AuthService.Login
 */
export const login = createQueryService({
  service: AuthService,
}).login;

/**
 * ログアウト
 *
 * @generated from rpc backend.v1.AuthService.Logout
 */
export const logout = createQueryService({
  service: AuthService,
}).logout;

/**
 * refresh token を使って access token を更新する
 *
 * @generated from rpc backend.v1.AuthService.RefreshAccessToken
 */
export const refreshAccessToken = createQueryService({
  service: AuthService,
}).refreshAccessToken;
