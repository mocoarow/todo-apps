import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { useTodoStore } from "~/stores/todo";

vi.mock("~/gateway/todo", () => ({
  todoService: {
    getTodos: vi.fn(),
    createTodo: vi.fn(),
    updateTodo: vi.fn(),
  },
}));

async function importMockedTodoService() {
  const { todoService } = await import("~/gateway/todo");
  return vi.mocked(todoService);
}

describe("useTodoStore", () => {
  let mockTodoService: Awaited<ReturnType<typeof importMockedTodoService>>;

  beforeEach(async () => {
    mockTodoService = await importMockedTodoService();
    useTodoStore.setState({
      todos: [],
      isLoading: false,
      isCreating: false,
      error: null,
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("fetchTodos", () => {
    it("should set todos on success", async () => {
      // given
      const todos = [
        {
          id: 1,
          text: "Buy milk",
          isComplete: false,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-01T00:00:00Z",
        },
        {
          id: 2,
          text: "Walk dog",
          isComplete: true,
          createdAt: "2025-01-02T00:00:00Z",
          updatedAt: "2025-01-02T00:00:00Z",
        },
      ];
      mockTodoService.getTodos.mockResolvedValue({ todos });

      // when
      await useTodoStore.getState().fetchTodos();

      // then
      const state = useTodoStore.getState();
      expect(state.todos).toEqual(todos);
      expect(state.isLoading).toBe(false);
      expect(state.error).toBeNull();
    });

    it("should set error on failure", async () => {
      // given
      const appError = new AppError("API_ERROR", "Server error");
      mockTodoService.getTodos.mockRejectedValue(appError);

      // when
      await useTodoStore.getState().fetchTodos();

      // then
      const state = useTodoStore.getState();
      expect(state.todos).toEqual([]);
      expect(state.isLoading).toBe(false);
      expect(state.error).toBe(appError);
    });

    it("should wrap non-AppError as UNKNOWN", async () => {
      // given
      mockTodoService.getTodos.mockRejectedValue(new TypeError("unexpected"));

      // when
      await useTodoStore.getState().fetchTodos();

      // then
      const state = useTodoStore.getState();
      expect(state.error).toBeInstanceOf(AppError);
      expect(state.error?.code).toBe("UNKNOWN");
      expect(state.error?.message).toBe("unexpected");
    });

    it("should set isLoading true during fetch", async () => {
      // given
      let capturedLoading = false;
      mockTodoService.getTodos.mockImplementation(async () => {
        capturedLoading = useTodoStore.getState().isLoading;
        return { todos: [] };
      });

      // when
      await useTodoStore.getState().fetchTodos();

      // then
      expect(capturedLoading).toBe(true);
      expect(useTodoStore.getState().isLoading).toBe(false);
    });
  });

  describe("createTodo", () => {
    it("should call createTodo and re-fetch todos on success", async () => {
      // given
      const todos = [
        {
          id: 1,
          text: "Buy milk",
          isComplete: false,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-01T00:00:00Z",
        },
      ];
      mockTodoService.createTodo.mockResolvedValue({
        id: 1,
        text: "Buy milk",
        isComplete: false,
        createdAt: "2025-01-01T00:00:00Z",
        updatedAt: "2025-01-01T00:00:00Z",
      });
      mockTodoService.getTodos.mockResolvedValue({ todos });

      // when
      await useTodoStore.getState().createTodo("Buy milk");

      // then
      expect(mockTodoService.createTodo).toHaveBeenCalledWith({
        text: "Buy milk",
      });
      expect(mockTodoService.getTodos).toHaveBeenCalled();
      expect(useTodoStore.getState().todos).toEqual(todos);
      expect(useTodoStore.getState().isCreating).toBe(false);
    });

    it("should set error on failure", async () => {
      // given
      const appError = new AppError("API_ERROR", "Create failed");
      mockTodoService.createTodo.mockRejectedValue(appError);

      // when
      await useTodoStore.getState().createTodo("Buy milk");

      // then
      expect(useTodoStore.getState().error).toBe(appError);
      expect(useTodoStore.getState().isCreating).toBe(false);
    });

    it("should set isCreating true during create", async () => {
      // given
      let capturedLoading = false;
      mockTodoService.createTodo.mockImplementation(async () => {
        capturedLoading = useTodoStore.getState().isCreating;
        return {
          id: 1,
          text: "Buy milk",
          isComplete: false,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-01T00:00:00Z",
        };
      });
      mockTodoService.getTodos.mockResolvedValue({ todos: [] });

      // when
      await useTodoStore.getState().createTodo("Buy milk");

      // then
      expect(capturedLoading).toBe(true);
      expect(useTodoStore.getState().isCreating).toBe(false);
    });
  });

  describe("updateTodo", () => {
    it("should call updateTodo and re-fetch todos on success", async () => {
      // given
      const todos = [
        {
          id: 1,
          text: "Buy milk",
          isComplete: true,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-02T00:00:00Z",
        },
      ];
      mockTodoService.updateTodo.mockResolvedValue(todos[0]!);
      mockTodoService.getTodos.mockResolvedValue({ todos });

      // when
      await useTodoStore
        .getState()
        .updateTodo(1, { text: "Buy milk", isComplete: true });

      // then
      expect(mockTodoService.updateTodo).toHaveBeenCalledWith(1, {
        text: "Buy milk",
        isComplete: true,
      });
      expect(mockTodoService.getTodos).toHaveBeenCalled();
      expect(useTodoStore.getState().todos).toEqual(todos);
    });

    it("should set error on failure", async () => {
      // given
      const appError = new AppError("API_ERROR", "Update failed");
      mockTodoService.updateTodo.mockRejectedValue(appError);

      // when
      await useTodoStore
        .getState()
        .updateTodo(1, { text: "Buy milk", isComplete: true });

      // then
      expect(useTodoStore.getState().error).toBe(appError);
    });

    it("should not set isLoading during update", async () => {
      // given
      let capturedLoading = false;
      mockTodoService.updateTodo.mockImplementation(async () => {
        capturedLoading = useTodoStore.getState().isLoading;
        return {
          id: 1,
          text: "Buy milk",
          isComplete: true,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-02T00:00:00Z",
        };
      });
      mockTodoService.getTodos.mockResolvedValue({ todos: [] });

      // when
      await useTodoStore
        .getState()
        .updateTodo(1, { text: "Buy milk", isComplete: true });

      // then
      expect(capturedLoading).toBe(false);
    });
  });

  describe("clearError", () => {
    it("should clear error state", () => {
      // given
      useTodoStore.setState({
        error: new AppError("API_ERROR", "Something failed"),
      });

      // when
      useTodoStore.getState().clearError();

      // then
      expect(useTodoStore.getState().error).toBeNull();
    });
  });
});
