// @generated by protoc-gen-connect-query v0.5.1 with parameter "target=ts,import_extension=none"
// @generated from file backend/v1/healthcheck_service.proto (package backend.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PingRequest, PingResponse } from "./healthcheck_service_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks } from "@connectrpc/connect-query";

export const typeName = "backend.v1.HealthcheckService";

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

const $queryService = createQueryService({  service: HealthcheckService,});

/**
 * @generated from rpc backend.v1.HealthcheckService.Ping
 */
export const ping = {   ...$queryService.ping,  ...createUnaryHooks($queryService.ping)};
