import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it } from "vitest";
import type { GetTodosResponseTodo } from "~/api";
import { TodoList } from "./TodoList";

function createTodo(
  overrides: Partial<GetTodosResponseTodo> = {},
): GetTodosResponseTodo {
  return {
    id: 1,
    text: "Buy milk",
    isComplete: false,
    createdAt: "2025-01-01T00:00:00Z",
    updatedAt: "2025-01-01T00:00:00Z",
    ...overrides,
  };
}

describe("TodoList", () => {
  afterEach(() => {
    cleanup();
  });

  it("should display empty message when no todos", () => {
    // given
    const todos: GetTodosResponseTodo[] = [];

    // when
    render(<TodoList todos={todos} />);

    // then
    expect(screen.getByText("No todos yet.")).toBeDefined();
  });

  it("should render all todos", () => {
    // given
    const todos = [
      createTodo({ id: 1, text: "Buy milk", isComplete: false }),
      createTodo({ id: 2, text: "Walk dog", isComplete: true }),
    ];

    // when
    render(<TodoList todos={todos} />);

    // then
    expect(screen.getByText("Buy milk")).toBeDefined();
    expect(screen.getByText("Walk dog")).toBeDefined();
  });

  it("should render todos as list items", () => {
    // given
    const todos = [
      createTodo({ id: 1, text: "Buy milk", isComplete: false }),
      createTodo({ id: 2, text: "Walk dog", isComplete: true }),
    ];

    // when
    render(<TodoList todos={todos} />);

    // then
    const listItems = screen.getAllByRole("listitem");
    expect(listItems).toHaveLength(2);
  });
});
