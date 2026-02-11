import {
  type AuthenticateRequest,
  type AuthenticateResponse,
  AuthenticateResponseSchema,
  type ErrorResponse,
  type GetMeResponse,
  GetMeResponseSchema,
} from "~/api";
import { config } from "~/config/config";
import type { AuthService } from "~/domain/auth";
import { AppError, type AppErrorCode } from "~/domain/error";

export class HttpAuthService implements AuthService {
  private readonly baseUrl: string;

  constructor(baseUrl: string = config.apiBaseUrl) {
    this.baseUrl = baseUrl;
  }

  async login(data: AuthenticateRequest): Promise<AuthenticateResponse> {
    const response = await this.fetchApi("/auth/authenticate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    return this.parseJson(response, AuthenticateResponseSchema);
  }

  async logout(): Promise<void> {
    await this.fetchApi("/auth/logout", { method: "POST" });
  }

  async getMe(): Promise<GetMeResponse> {
    const response = await this.fetchApi("/auth/me", { method: "GET" });
    return this.parseJson(response, GetMeResponseSchema);
  }

  private async fetchApi(path: string, init?: RequestInit): Promise<Response> {
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

  private async parseJson<T>(
    response: Response,
    schema: { parse: (data: unknown) => T },
  ): Promise<T> {
    try {
      return schema.parse(await response.json());
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

export const authService = new HttpAuthService();
