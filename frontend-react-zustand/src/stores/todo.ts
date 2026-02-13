import { create } from "zustand";
import type { GetTodosResponseTodo, UpdateTodoRequest } from "~/api";
import { type AppError, toAppError } from "~/domain/error";
import { todoService } from "~/gateway/todo";

interface TodoState {
  todos: GetTodosResponseTodo[];
  isLoading: boolean;
  isCreating: boolean;
  error: AppError | null;
}

interface TodoActions {
  fetchTodos(): Promise<void>;
  createTodo(text: string): Promise<void>;
  updateTodo(id: number, request: UpdateTodoRequest): Promise<void>;
  clearError(): void;
}

export const useTodoStore = create<TodoState & TodoActions>()((set) => ({
  todos: [],
  isLoading: false,
  isCreating: false,
  error: null,

  fetchTodos: async () => {
    set({ isLoading: true, error: null });
    try {
      const response = await todoService.getTodos();
      set({ todos: response.todos, isLoading: false });
    } catch (error) {
      set({ error: toAppError(error), isLoading: false });
    }
  },

  createTodo: async (text: string) => {
    set({ isCreating: true, error: null });
    try {
      await todoService.createTodo({ text });
      const response = await todoService.getTodos();
      set({ todos: response.todos, isCreating: false });
    } catch (error) {
      set({ error: toAppError(error), isCreating: false });
    }
  },

  updateTodo: async (id: number, request: UpdateTodoRequest) => {
    set({ error: null });
    try {
      await todoService.updateTodo(id, request);
      const response = await todoService.getTodos();
      set({ todos: response.todos });
    } catch (error) {
      set({ error: toAppError(error) });
    }
  },

  clearError: () => {
    set({ error: null });
  },
}));
