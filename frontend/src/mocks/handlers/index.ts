import type { RequestHandler } from "msw";
import { authHandlers } from "./auth";
import { contestHandlers } from "./contest";
import { healthcheckHandlers } from "./healthcheck";
import { userHandlers } from "./user";

export const handlers: RequestHandler[] = [
  ...healthcheckHandlers,
  ...authHandlers,
  ...userHandlers,
  ...contestHandlers,
];
