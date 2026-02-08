import { z } from "zod";

const DEFAULT_API_BASE_URL = "http://localhost:8080/api/v1" as const;

const envSchema = z.object({
  VITE_API_BASE_URL: z
    .string()
    .url()
    .default(DEFAULT_API_BASE_URL)
    .describe("Backend API base URL"),
});

function validateEnv() {
  return envSchema.parse({
    VITE_API_BASE_URL: import.meta.env.VITE_API_BASE_URL,
  });
}

const validatedEnv = validateEnv();

export const config = Object.freeze({
  apiBaseUrl: validatedEnv.VITE_API_BASE_URL,
} as const);
