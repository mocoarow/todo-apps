import {
  type CreateTodoRequest,
  type CreateTodoResponse,
  CreateTodoResponseSchema,
  type GetTodosResponse,
  GetTodosResponseSchema,
  type UpdateTodoRequest,
  type UpdateTodoResponse,
  UpdateTodoResponseSchema,
} from "~/api";
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

  async createTodo(request: CreateTodoRequest): Promise<CreateTodoResponse> {
    const response = await this.client.fetchApi("/todo", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    });
    return this.client.parseJson(response, CreateTodoResponseSchema);
  }

  async updateTodo(
    id: number,
    request: UpdateTodoRequest,
  ): Promise<UpdateTodoResponse> {
    const response = await this.client.fetchApi(`/todo/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    });
    return this.client.parseJson(response, UpdateTodoResponseSchema);
  }
}

export const todoService = new HttpTodoService();
