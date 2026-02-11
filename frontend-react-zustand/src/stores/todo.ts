import { create } from "zustand";
import type { GetTodosResponseTodo } from "~/api";
import { type AppError, toAppError } from "~/domain/error";
import { todoService } from "~/gateway/todo";

interface TodoState {
  todos: GetTodosResponseTodo[];
  isLoading: boolean;
  error: AppError | null;
}

interface TodoActions {
  fetchTodos(): Promise<void>;
  clearError(): void;
}

export const useTodoStore = create<TodoState & TodoActions>()((set) => ({
  todos: [],
  isLoading: false,
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

  clearError: () => {
    set({ error: null });
  },
}));
