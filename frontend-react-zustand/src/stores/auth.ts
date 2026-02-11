import { create } from "zustand";
import type { GetMeResponse } from "~/api";
import { type AppError, toAppError } from "~/domain/error";
import { authService } from "~/gateway/auth";

interface AuthState {
  user: GetMeResponse | null;
  isLoading: boolean;
  error: AppError | null;
}

interface AuthActions {
  login(loginId: string, password: string): Promise<void>;
  logout(): Promise<void>;
  fetchMe(): Promise<void>;
  clearError(): void;
}

export const useAuthStore = create<AuthState & AuthActions>()((set) => ({
  user: null,
  isLoading: false,
  error: null,

  login: async (loginId, password) => {
    set({ isLoading: true, error: null });
    try {
      await authService.login({ loginId, password });
      const user = await authService.getMe();
      set({ user, isLoading: false });
    } catch (error) {
      set({ error: toAppError(error), isLoading: false });
    }
  },

  logout: async () => {
    set({ isLoading: true, error: null });
    try {
      await authService.logout();
      set({ user: null, isLoading: false });
    } catch (error) {
      set({ error: toAppError(error), isLoading: false });
    }
  },

  fetchMe: async () => {
    set({ isLoading: true, error: null });
    try {
      const user = await authService.getMe();
      set({ user, isLoading: false });
    } catch (error) {
      set({ user: null, error: toAppError(error), isLoading: false });
    }
  },

  clearError: () => {
    set({ error: null });
  },
}));
