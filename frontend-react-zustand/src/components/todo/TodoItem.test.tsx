import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it } from "vitest";
import type { GetTodosResponseTodo } from "~/api";
import { TodoItem } from "./TodoItem";

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

describe("TodoItem", () => {
  afterEach(() => {
    cleanup();
  });

  it("should display todo text", () => {
    // given
    const todo = createTodo();

    // when
    render(<TodoItem todo={todo} />);

    // then
    expect(screen.getByText("Buy milk")).toBeDefined();
  });

  it("should show line-through for completed todo", () => {
    // given
    const todo = createTodo({ text: "Walk dog", isComplete: true });

    // when
    render(<TodoItem todo={todo} />);

    // then
    const span = screen.getByText("Walk dog");
    expect(span.className).toContain("line-through");
  });

  it("should not show line-through for incomplete todo", () => {
    // given
    const todo = createTodo();

    // when
    render(<TodoItem todo={todo} />);

    // then
    const span = screen.getByText("Buy milk");
    expect(span.className ?? "").not.toContain("line-through");
  });

  it("should render checkbox with aria-label matching todo text", () => {
    // given
    const todo = createTodo();

    // when
    render(<TodoItem todo={todo} />);

    // then
    expect(screen.getByRole("checkbox", { name: "Buy milk" })).toBeDefined();
  });

  it("should render disabled checkbox", () => {
    // given
    const todo = createTodo();

    // when
    render(<TodoItem todo={todo} />);

    // then
    const checkbox = screen.getByRole("checkbox", { name: "Buy milk" });
    expect(checkbox).toBeDisabled();
  });
});
