import type { GetTodosResponse } from "~/api";

export interface TodoService {
  getTodos(): Promise<GetTodosResponse>;
}
