import type { z } from "zod";
import {
  type AuthenticateBody,
  AuthenticateResponse as AuthenticateResponseSchema,
  type CreateBulkTodosBody,
  type CreateTodoBody,
  type GetTodosResponse as FindTodoResponseSchema,
  GetMeResponse as GetMeResponseSchema,
  type UpdateTodoBody,
  type UpdateTodoResponse as UpdateTodoResponseSchema,
} from "./types.gen";

// Zod schemas for runtime validation
export { AuthenticateResponseSchema, GetMeResponseSchema };

// TypeScript types derived from Zod schemas
export type AuthenticateRequest = z.infer<typeof AuthenticateBody>;
export type AuthenticateResponse = z.infer<typeof AuthenticateResponseSchema>;
export type CreateTodoRequest = z.infer<typeof CreateTodoBody>;
export type CreateBulkTodosRequest = z.infer<typeof CreateBulkTodosBody>;
export type UpdateTodoRequest = z.infer<typeof UpdateTodoBody>;
export type UpdateTodoResponse = z.infer<typeof UpdateTodoResponseSchema>;
export type FindTodoResponse = z.infer<typeof FindTodoResponseSchema>;
export type FindTodoResponseTodo = FindTodoResponse["todos"][number];
export type CreateTodoResponse = UpdateTodoResponse;
export type CreateBulkTodosResponse = { todos: CreateTodoResponse[] };
export type GetMeResponse = z.infer<typeof GetMeResponseSchema>;
export type ErrorResponse = { code: string; message: string };
