import type {
  CreateTodoRequest,
  CreateTodoResponse,
  GetTodosResponse,
  UpdateTodoRequest,
  UpdateTodoResponse,
} from "~/api";

export interface TodoService {
  getTodos(): Promise<GetTodosResponse>;
  createTodo(request: CreateTodoRequest): Promise<CreateTodoResponse>;
  updateTodo(
    id: number,
    request: UpdateTodoRequest,
  ): Promise<UpdateTodoResponse>;
}
