import {
  type AuthenticateRequest,
  type AuthenticateResponse,
  AuthenticateResponseSchema,
  type ErrorResponse,
} from "~/api";
import { config } from "~/config/config";
import type { AuthService } from "../domain/auth";
import { AppError } from "../domain/error";

export class HttpAuthService implements AuthService {
  private readonly baseUrl: string;

  constructor(baseUrl: string = config.apiBaseUrl) {
    this.baseUrl = baseUrl;
  }

  async login(data: AuthenticateRequest): Promise<AuthenticateResponse> {
    let response: Response;
    try {
      response = await fetch(`${this.baseUrl}/auth/authenticate`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(data),
      });
    } catch {
      throw new AppError("NETWORK_ERROR", "Network error occurred");
    }

    if (!response.ok) {
      const code = response.status === 401 ? "UNAUTHENTICATED" : "API_ERROR";
      const message = await this.extractErrorMessage(response);
      throw new AppError(code, message);
    }

    try {
      return AuthenticateResponseSchema.parse(await response.json());
    } catch {
      throw new AppError("API_ERROR", "Invalid response format");
    }
  }

  logout(): void {
    // httpOnly Cookieのクリアはサーバー側で行う
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
