import type { ErrorResponse } from "~/api";
import { config } from "~/config/config";
import { AppError, type AppErrorCode } from "~/domain/error";

export class HttpClient {
  private readonly baseUrl: string;

  constructor(baseUrl: string = config.apiBaseUrl) {
    this.baseUrl = baseUrl;
  }

  async fetchApi(path: string, init?: RequestInit): Promise<Response> {
    let response: Response;
    try {
      response = await fetch(`${this.baseUrl}${path}`, {
        ...init,
        credentials: "include",
      });
    } catch {
      throw new AppError("NETWORK_ERROR", "Network error occurred");
    }

    if (!response.ok) {
      const code: AppErrorCode =
        response.status === 401 ? "UNAUTHENTICATED" : "API_ERROR";
      const message = await this.extractErrorMessage(response);
      throw new AppError(code, message);
    }

    return response;
  }

  async parseJson<T>(
    response: Response,
    schema: { parse: (data: unknown) => T },
  ): Promise<T> {
    let data: unknown;
    try {
      data = await response.json();
    } catch {
      throw new AppError("API_ERROR", "Invalid response format");
    }
    try {
      return schema.parse(data);
    } catch {
      throw new AppError("API_ERROR", "Invalid response format");
    }
  }

  private async extractErrorMessage(response: Response): Promise<string> {
    try {
      const contentType = response.headers.get("content-type");
      if (contentType?.includes("application/json") === true) {
        const errorResponse: ErrorResponse = await response.json();
        return errorResponse.message ?? response.statusText;
      }
      return await response.text();
    } catch {
      return response.statusText;
    }
  }
}
