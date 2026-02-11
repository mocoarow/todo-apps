import { defineConfig } from "orval";

export default defineConfig({
  todoApiZod: {
    input: {
      target: "../openapi/openapi.yaml",
    },
    output: {
      client: "zod",
      target: "./src/api/types.gen.ts",
      mode: "single",
    },
  },
});
