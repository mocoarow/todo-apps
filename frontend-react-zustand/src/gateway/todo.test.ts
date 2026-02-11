import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { HttpClient } from "~/gateway/http-client";
import { HttpTodoService } from "~/gateway/todo";

const BASE_URL = "http://localhost:8080/api/v1";

function jsonResponse(body: unknown, status = 200): Response {
  return new Response(JSON.stringify(body), {
    status,
    headers: { "content-type": "application/json" },
  });
}

describe("HttpTodoService", () => {
  let service: HttpTodoService;

  beforeEach(() => {
    service = new HttpTodoService(new HttpClient(BASE_URL));
    vi.stubGlobal("fetch", vi.fn());
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe("getTodos", () => {
    it("should return parsed todos on success", async () => {
      // given
      const body = {
        todos: [
          {
            id: 1,
            text: "Buy milk",
            isComplete: false,
            createdAt: "2025-01-01T00:00:00Z",
            updatedAt: "2025-01-01T00:00:00Z",
          },
          {
            id: 2,
            text: "Walk dog",
            isComplete: true,
            createdAt: "2025-01-02T00:00:00Z",
            updatedAt: "2025-01-02T00:00:00Z",
          },
        ],
      };
      vi.mocked(fetch).mockResolvedValue(jsonResponse(body));

      // when
      const result = await service.getTodos();

      // then
      expect(result).toEqual(body);
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/todo`, {
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
      await expect(service.getTodos()).rejects.toMatchObject({
        code: "UNAUTHENTICATED",
      });
    });

    it("should throw NETWORK_ERROR on fetch failure", async () => {
      // given
      vi.mocked(fetch).mockRejectedValue(new TypeError("Failed to fetch"));

      // when & then
      await expect(service.getTodos()).rejects.toMatchObject({
        code: "NETWORK_ERROR",
        message: "Network error occurred",
      });
    });

    it("should throw API_ERROR on invalid response format", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(jsonResponse({ bad: "data" }));

      // when & then
      await expect(service.getTodos()).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Invalid response format",
      });
    });

    it("should throw API_ERROR on server error", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        jsonResponse({ code: "INTERNAL", message: "Server error" }, 500),
      );

      // when & then
      await expect(service.getTodos()).rejects.toThrow(AppError);
      await expect(service.getTodos()).rejects.toMatchObject({
        code: "API_ERROR",
      });
    });
  });
});
