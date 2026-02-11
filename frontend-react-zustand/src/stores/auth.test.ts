import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { useAuthStore } from "~/stores/auth";

vi.mock("~/gateway/auth", () => ({
  authService: {
    login: vi.fn(),
    logout: vi.fn(),
    getMe: vi.fn(),
  },
}));

async function importMockedAuthService() {
  const { authService } = await import("~/gateway/auth");
  return vi.mocked(authService);
}

describe("useAuthStore", () => {
  let mockAuthService: Awaited<ReturnType<typeof importMockedAuthService>>;

  beforeEach(async () => {
    mockAuthService = await importMockedAuthService();
    useAuthStore.setState({
      user: null,
      isLoading: false,
      error: null,
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("login", () => {
    it("should set user on successful login", async () => {
      // given
      const user = { userId: 1, loginId: "user1" };
      mockAuthService.login.mockResolvedValue({});
      mockAuthService.getMe.mockResolvedValue(user);

      // when
      await useAuthStore.getState().login("user1", "password1234");

      // then
      const state = useAuthStore.getState();
      expect(state.user).toEqual(user);
      expect(state.isLoading).toBe(false);
      expect(state.error).toBeNull();
    });

    it("should set error on login failure", async () => {
      // given
      const appError = new AppError("UNAUTHENTICATED", "Invalid credentials");
      mockAuthService.login.mockRejectedValue(appError);

      // when
      await useAuthStore.getState().login("user1", "wrongpass1");

      // then
      const state = useAuthStore.getState();
      expect(state.user).toBeNull();
      expect(state.isLoading).toBe(false);
      expect(state.error).toBe(appError);
    });

    it("should wrap non-AppError as UNKNOWN", async () => {
      // given
      mockAuthService.login.mockRejectedValue(new TypeError("unexpected"));

      // when
      await useAuthStore.getState().login("user1", "password1234");

      // then
      const state = useAuthStore.getState();
      expect(state.error).toBeInstanceOf(AppError);
      expect(state.error?.code).toBe("UNKNOWN");
      expect(state.error?.message).toBe("unexpected");
    });

    it("should set isLoading true during login", async () => {
      // given
      let capturedLoading = false;
      mockAuthService.login.mockImplementation(async () => {
        capturedLoading = useAuthStore.getState().isLoading;
        return {};
      });
      mockAuthService.getMe.mockResolvedValue({
        userId: 1,
        loginId: "user1",
      });

      // when
      await useAuthStore.getState().login("user1", "password1234");

      // then
      expect(capturedLoading).toBe(true);
      expect(useAuthStore.getState().isLoading).toBe(false);
    });
  });

  describe("logout", () => {
    it("should clear user on successful logout", async () => {
      // given
      useAuthStore.setState({ user: { userId: 1, loginId: "user1" } });
      mockAuthService.logout.mockResolvedValue(undefined);

      // when
      await useAuthStore.getState().logout();

      // then
      const state = useAuthStore.getState();
      expect(state.user).toBeNull();
      expect(state.isLoading).toBe(false);
      expect(state.error).toBeNull();
    });

    it("should set error on logout failure", async () => {
      // given
      const appError = new AppError("NETWORK_ERROR", "Network error occurred");
      mockAuthService.logout.mockRejectedValue(appError);

      // when
      await useAuthStore.getState().logout();

      // then
      const state = useAuthStore.getState();
      expect(state.isLoading).toBe(false);
      expect(state.error).toBe(appError);
    });
  });

  describe("fetchMe", () => {
    it("should set user on success", async () => {
      // given
      const user = { userId: 1, loginId: "user1" };
      mockAuthService.getMe.mockResolvedValue(user);

      // when
      await useAuthStore.getState().fetchMe();

      // then
      const state = useAuthStore.getState();
      expect(state.user).toEqual(user);
      expect(state.isLoading).toBe(false);
      expect(state.error).toBeNull();
    });

    it("should clear user and set error on failure", async () => {
      // given
      useAuthStore.setState({ user: { userId: 1, loginId: "user1" } });
      const appError = new AppError("UNAUTHENTICATED", "Not authenticated");
      mockAuthService.getMe.mockRejectedValue(appError);

      // when
      await useAuthStore.getState().fetchMe();

      // then
      const state = useAuthStore.getState();
      expect(state.user).toBeNull();
      expect(state.error).toBe(appError);
    });
  });

  describe("clearError", () => {
    it("should clear error state", () => {
      // given
      useAuthStore.setState({
        error: new AppError("API_ERROR", "Something failed"),
      });

      // when
      useAuthStore.getState().clearError();

      // then
      expect(useAuthStore.getState().error).toBeNull();
    });
  });
});
