import { RequestHandler } from "msw";
import { authHandlers } from "./auth";
import { healthcheckHandlers } from "./healthcheck";

export const handlers: RequestHandler[] = [
  ...healthcheckHandlers,
  ...authHandlers,
];
