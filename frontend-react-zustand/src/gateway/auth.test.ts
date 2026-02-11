import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { HttpAuthService } from "~/gateway/auth";

const BASE_URL = "http://localhost:8080/api/v1";

function jsonResponse(body: unknown, status = 200): Response {
  return new Response(JSON.stringify(body), {
    status,
    headers: { "content-type": "application/json" },
  });
}

function textResponse(text: string, status: number): Response {
  return new Response(text, { status });
}

describe("HttpAuthService", () => {
  let service: HttpAuthService;

  beforeEach(() => {
    service = new HttpAuthService(BASE_URL);
    vi.stubGlobal("fetch", vi.fn());
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe("login", () => {
    it("should return parsed response on success", async () => {
      // given
      const body = { accessToken: "token-123" };
      vi.mocked(fetch).mockResolvedValue(jsonResponse(body));

      // when
      const result = await service.login({
        loginId: "user1",
        password: "password1234",
      });

      // then
      expect(result).toEqual({ accessToken: "token-123" });
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/auth/authenticate`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ loginId: "user1", password: "password1234" }),
      });
    });

    it("should throw UNAUTHENTICATED on 401", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        jsonResponse(
          { code: "UNAUTHENTICATED", message: "Invalid credentials" },
          401,
        ),
      );

      // when & then
      await expect(
        service.login({ loginId: "user1", password: "wrongpass1" }),
      ).rejects.toThrow(AppError);

      await expect(
        service.login({ loginId: "user1", password: "wrongpass1" }),
      ).rejects.toMatchObject({ code: "UNAUTHENTICATED" });
    });

    it("should throw NETWORK_ERROR on fetch failure", async () => {
      // given
      vi.mocked(fetch).mockRejectedValue(new TypeError("Failed to fetch"));

      // when & then
      await expect(
        service.login({ loginId: "user1", password: "password1234" }),
      ).rejects.toMatchObject({
        code: "NETWORK_ERROR",
        message: "Network error occurred",
      });
    });

    it("should throw API_ERROR on invalid response format", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        new Response("not json", { status: 200 }),
      );

      // when & then
      await expect(
        service.login({ loginId: "user1", password: "password1234" }),
      ).rejects.toMatchObject({ code: "API_ERROR" });
    });
  });

  describe("logout", () => {
    it("should call logout endpoint", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(new Response(null, { status: 204 }));

      // when
      await service.logout();

      // then
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/auth/logout`, {
        method: "POST",
        credentials: "include",
      });
    });

    it("should throw API_ERROR on server error", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        textResponse("Internal Server Error", 500),
      );

      // when & then
      await expect(service.logout()).rejects.toMatchObject({
        code: "API_ERROR",
      });
    });

    it("should throw NETWORK_ERROR on fetch failure", async () => {
      // given
      vi.mocked(fetch).mockRejectedValue(new TypeError("Failed to fetch"));

      // when & then
      await expect(service.logout()).rejects.toMatchObject({
        code: "NETWORK_ERROR",
      });
    });
  });

  describe("getMe", () => {
    it("should return parsed user on success", async () => {
      // given
      const body = { userId: 1, loginId: "user1" };
      vi.mocked(fetch).mockResolvedValue(jsonResponse(body));

      // when
      const result = await service.getMe();

      // then
      expect(result).toEqual({ userId: 1, loginId: "user1" });
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/auth/me`, {
        method: "GET",
        credentials: "include",
      });
    });

    it("should throw UNAUTHENTICATED on 401", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        jsonResponse(
          { code: "UNAUTHENTICATED", message: "Not authenticated" },
          401,
        ),
      );

      // when & then
      await expect(service.getMe()).rejects.toMatchObject({
        code: "UNAUTHENTICATED",
      });
    });

    it("should throw NETWORK_ERROR on fetch failure", async () => {
      // given
      vi.mocked(fetch).mockRejectedValue(new TypeError("Failed to fetch"));

      // when & then
      await expect(service.getMe()).rejects.toMatchObject({
        code: "NETWORK_ERROR",
      });
    });

    it("should throw API_ERROR on invalid response format", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(jsonResponse({ bad: "data" }));

      // when & then
      await expect(service.getMe()).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Invalid response format",
      });
    });
  });
});
