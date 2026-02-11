import {
  type AuthenticateRequest,
  type AuthenticateResponse,
  AuthenticateResponseSchema,
  type GetMeResponse,
  GetMeResponseSchema,
} from "~/api";
import type { AuthService } from "~/domain/auth";
import { HttpClient } from "~/gateway/http-client";

export class HttpAuthService implements AuthService {
  private readonly client: HttpClient;
  constructor(client: HttpClient = new HttpClient()) {
    this.client = client;
  }

  async login(data: AuthenticateRequest): Promise<AuthenticateResponse> {
    const response = await this.client.fetchApi("/auth/authenticate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-token-delivery": "cookie",
      },
      body: JSON.stringify(data),
    });
    return this.client.parseJson(response, AuthenticateResponseSchema);
  }

  async logout(): Promise<void> {
    await this.client.fetchApi("/auth/logout", { method: "POST" });
  }

  async getMe(): Promise<GetMeResponse> {
    const response = await this.client.fetchApi("/auth/me", { method: "GET" });
    return this.client.parseJson(response, GetMeResponseSchema);
  }
}

export const authService = new HttpAuthService();
