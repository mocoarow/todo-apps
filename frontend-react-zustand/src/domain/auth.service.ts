import type { AuthenticateRequest, AuthenticateResponse } from "~/api";

export interface IAuthService {
  login(data: AuthenticateRequest): Promise<AuthenticateResponse>;
  logout(): void;
  getToken(): string | null;
  getUserId(): number | null;
  getLoginId(): string | null;
  saveAuth(token: string): void;
  isAuthenticated(): boolean;
}
