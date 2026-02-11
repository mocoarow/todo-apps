import { type GetTodosResponse, GetTodosResponseSchema } from "~/api";
import type { TodoService } from "~/domain/todo";
import { HttpClient } from "~/gateway/http-client";

export class HttpTodoService implements TodoService {
  private readonly client: HttpClient;
  constructor(client: HttpClient = new HttpClient()) {
    this.client = client;
  }

  async getTodos(): Promise<GetTodosResponse> {
    const response = await this.client.fetchApi("/todo", { method: "GET" });
    return this.client.parseJson(response, GetTodosResponseSchema);
  }
}

export const todoService = new HttpTodoService();
