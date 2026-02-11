import type {
  AuthenticateRequest,
  AuthenticateResponse,
  GetMeResponse,
} from "~/api";

export interface AuthService {
  login(data: AuthenticateRequest): Promise<AuthenticateResponse>;
  logout(): Promise<void>;
  getMe(): Promise<GetMeResponse>;
}
