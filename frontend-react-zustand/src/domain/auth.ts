import type { AuthenticateRequest, AuthenticateResponse } from "~/api";

export interface AuthService {
  login(data: AuthenticateRequest): Promise<AuthenticateResponse>;
  logout(): void;
}
