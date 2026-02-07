export type {
  components,
  operations,
  paths,
} from "./types.gen";

type Schemas = components["schemas"];

export type AuthenticateRequest = Schemas["AuthenticateRequest"];
export type AuthenticateResponse = Schemas["AuthenticateResponse"];
export type CreateTodoRequest = Schemas["CreateTodoRequest"];
export type CreateTodoResponse = Schemas["CreateTodoResponse"];
export type CreateBulkTodosRequest = Schemas["CreateBulkTodosRequest"];
export type CreateBulkTodosResponse = Schemas["CreateBulkTodosResponse"];
export type UpdateTodoRequest = Schemas["UpdateTodoRequest"];
export type UpdateTodoResponse = Schemas["UpdateTodoResponse"];
export type FindTodoResponse = Schemas["FindTodoResponse"];
export type FindTodoResponseTodo = Schemas["FindTodoResponseTodo"];
export type ErrorResponse = Schemas["ErrorResponse"];
