import { cleanup, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { useAuthStore } from "~/stores/auth";
import { useTodoStore } from "~/stores/todo";
import { TodoPage } from "./TodoPage";

vi.mock("~/gateway/auth", () => ({
  authService: {
    login: vi.fn(),
    logout: vi.fn(),
    getMe: vi.fn(),
  },
}));

vi.mock("~/gateway/todo", () => ({
  todoService: {
    getTodos: vi.fn(),
    createTodo: vi.fn(),
    updateTodo: vi.fn(),
  },
}));

describe("TodoPage", () => {
  beforeEach(() => {
    useAuthStore.setState({
      user: { userId: 1, loginId: "testuser" },
      isLoading: false,
      error: null,
    });
    useTodoStore.setState({
      todos: [],
      isLoading: false,
      isCreating: false,
      error: null,
    });
  });

  afterEach(() => {
    cleanup();
    vi.clearAllMocks();
  });

  it("should display the user's loginId", () => {
    // given
    // (state set in beforeEach)

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByText("testuser")).toBeDefined();
  });

  it("should display loading state", () => {
    // given
    useTodoStore.setState({ isLoading: true });

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByText("Loading...")).toBeDefined();
  });

  it("should display error message with retry button", () => {
    // given
    useTodoStore.setState({
      error: new AppError("API_ERROR", "Failed to fetch todos"),
      fetchTodos: vi.fn(),
    });

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByText("Failed to fetch todos")).toBeDefined();
    expect(screen.getByRole("button", { name: "Retry" })).toBeDefined();
  });

  it("should call fetchTodos on retry button click", async () => {
    // given
    const user = userEvent.setup();
    const mockFetchTodos = vi.fn();
    useTodoStore.setState({
      error: new AppError("API_ERROR", "Failed to fetch todos"),
      fetchTodos: mockFetchTodos,
    });
    render(<TodoPage />);

    // when
    await user.click(screen.getByRole("button", { name: "Retry" }));

    // then
    expect(mockFetchTodos).toHaveBeenCalled();
  });

  it("should display todos when loaded", () => {
    // given
    useTodoStore.setState({
      todos: [
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
      ],
    });

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByText("Buy milk")).toBeDefined();
    expect(screen.getByText("Walk dog")).toBeDefined();
  });

  it("should display empty message when no todos", () => {
    // given
    useTodoStore.setState({ todos: [] });

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByText("No todos yet.")).toBeDefined();
  });

  it("should call logout when logout button is clicked", async () => {
    // given
    const user = userEvent.setup();
    const mockLogout = vi.fn();
    useAuthStore.setState({ logout: mockLogout });
    render(<TodoPage />);

    // when
    await user.click(screen.getByRole("button", { name: "Logout" }));

    // then
    expect(mockLogout).toHaveBeenCalled();
  });

  it("should disable logout button while auth is loading", () => {
    // given
    useAuthStore.setState({ isLoading: true });

    // when
    render(<TodoPage />);

    // then
    expect(screen.getByRole("button", { name: "Logout" })).toBeDisabled();
  });

  it("should render the create todo form", () => {
    // given / when
    render(<TodoPage />);

    // then
    expect(screen.getByPlaceholderText("What needs to be done?")).toBeDefined();
    expect(screen.getByRole("button", { name: "Add" })).toBeDefined();
  });

  it("should call createTodo when form is submitted", async () => {
    // given
    const user = userEvent.setup();
    const mockCreateTodo = vi.fn();
    useTodoStore.setState({ createTodo: mockCreateTodo });
    render(<TodoPage />);

    // when
    await user.type(
      screen.getByPlaceholderText("What needs to be done?"),
      "New task{Enter}",
    );

    // then
    expect(mockCreateTodo).toHaveBeenCalledWith("New task");
  });

  it("should call updateTodo when checkbox is toggled", async () => {
    // given
    const user = userEvent.setup();
    const mockUpdateTodo = vi.fn();
    useTodoStore.setState({
      todos: [
        {
          id: 1,
          text: "Buy milk",
          isComplete: false,
          createdAt: "2025-01-01T00:00:00Z",
          updatedAt: "2025-01-01T00:00:00Z",
        },
      ],
      updateTodo: mockUpdateTodo,
    });
    render(<TodoPage />);

    // when
    await user.click(screen.getByRole("checkbox", { name: "Buy milk" }));

    // then
    expect(mockUpdateTodo).toHaveBeenCalledWith(1, {
      text: "Buy milk",
      isComplete: true,
    });
  });
});
