import type { RequestHandler } from "msw";
import { authHandlers } from "./auth";
import { contestHandlers } from "./contest";
import { healthcheckHandlers } from "./healthcheck";
import { taskHandlers } from "./task";
import { userHandlers } from "./user";
import { judgeHandlers } from "./judge";

export const handlers: readonly RequestHandler[] = [
  ...healthcheckHandlers,
  ...authHandlers,
  ...userHandlers,
  ...taskHandlers,
  ...contestHandlers,
  ...judgeHandlers,
] as const;
