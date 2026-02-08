export type AppErrorCode =
  | "UNAUTHENTICATED"
  | "API_ERROR"
  | "NETWORK_ERROR"
  | "UNKNOWN";

export class AppError extends Error {
  readonly code: AppErrorCode;

  constructor(code: AppErrorCode, message: string) {
    super(message);
    this.name = "AppError";
    this.code = code;
  }
}
