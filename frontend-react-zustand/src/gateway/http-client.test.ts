import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { HttpClient } from "~/gateway/http-client";

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

describe("HttpClient", () => {
  let client: HttpClient;

  beforeEach(() => {
    client = new HttpClient(BASE_URL);
    vi.stubGlobal("fetch", vi.fn());
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe("fetchApi", () => {
    it("should return response on success", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(jsonResponse({ ok: true }));

      // when
      const response = await client.fetchApi("/test", { method: "GET" });

      // then
      expect(response.ok).toBe(true);
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/test`, {
        method: "GET",
        credentials: "include",
      });
    });

    it("should include credentials in every request", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(jsonResponse({}));

      // when
      await client.fetchApi("/test", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });

      // then
      expect(fetch).toHaveBeenCalledWith(`${BASE_URL}/test`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    });

    it("should throw NETWORK_ERROR on fetch failure", async () => {
      // given
      vi.mocked(fetch).mockRejectedValue(new TypeError("Failed to fetch"));

      // when & then
      await expect(client.fetchApi("/test")).rejects.toMatchObject({
        code: "NETWORK_ERROR",
        message: "Network error occurred",
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
      await expect(client.fetchApi("/test")).rejects.toMatchObject({
        code: "UNAUTHENTICATED",
      });
    });

    it("should throw API_ERROR on 500 with JSON error body", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        jsonResponse({ code: "INTERNAL", message: "Server error" }, 500),
      );

      // when & then
      await expect(client.fetchApi("/test")).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Server error",
      });
    });

    it("should throw API_ERROR on 500 with text body", async () => {
      // given
      vi.mocked(fetch).mockResolvedValue(
        textResponse("Internal Server Error", 500),
      );

      // when & then
      await expect(client.fetchApi("/test")).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Internal Server Error",
      });
    });

    it("should use statusText when error body parsing fails", async () => {
      // given
      const response = new Response(null, {
        status: 502,
        statusText: "Bad Gateway",
      });
      // Make .text() throw
      vi.spyOn(response, "text").mockRejectedValue(new Error("read failed"));
      vi.spyOn(response.headers, "get").mockReturnValue(null);
      vi.mocked(fetch).mockResolvedValue(response);

      // when & then
      await expect(client.fetchApi("/test")).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Bad Gateway",
      });
    });
  });

  describe("parseJson", () => {
    it("should parse response with schema", async () => {
      // given
      const response = jsonResponse({ name: "test" });
      const schema = { parse: (data: unknown) => data as { name: string } };

      // when
      const result = await client.parseJson(response, schema);

      // then
      expect(result).toEqual({ name: "test" });
    });

    it("should throw API_ERROR when schema validation fails", async () => {
      // given
      const response = jsonResponse({ bad: "data" });
      const schema = {
        parse: () => {
          throw new Error("validation failed");
        },
      };

      // when & then
      await expect(client.parseJson(response, schema)).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Invalid response format",
      });
    });

    it("should throw API_ERROR when response is not valid JSON", async () => {
      // given
      const response = new Response("not json", { status: 200 });
      const schema = { parse: (data: unknown) => data };

      // when & then
      await expect(client.parseJson(response, schema)).rejects.toMatchObject({
        code: "API_ERROR",
        message: "Invalid response format",
      });
    });
  });
});
