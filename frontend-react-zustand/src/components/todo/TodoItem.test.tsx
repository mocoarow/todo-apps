import { cleanup, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, describe, expect, it, vi } from "vitest";
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
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // then
    expect(screen.getByText("Buy milk")).toBeDefined();
  });

  it("should show line-through for completed todo", () => {
    // given
    const todo = createTodo({ text: "Walk dog", isComplete: true });

    // when
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // then
    const span = screen.getByText("Walk dog");
    expect(span.className).toContain("line-through");
  });

  it("should not show line-through for incomplete todo", () => {
    // given
    const todo = createTodo();

    // when
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // then
    const span = screen.getByText("Buy milk");
    expect(span.className ?? "").not.toContain("line-through");
  });

  it("should render checkbox with aria-label matching todo text", () => {
    // given
    const todo = createTodo();

    // when
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // then
    expect(screen.getByRole("checkbox", { name: "Buy milk" })).toBeDefined();
  });

  it("should render enabled checkbox", () => {
    // given
    const todo = createTodo();

    // when
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // then
    const checkbox = screen.getByRole("checkbox", { name: "Buy milk" });
    expect(checkbox).not.toBeDisabled();
  });

  it("should call onToggleComplete when checkbox is clicked", async () => {
    // given
    const user = userEvent.setup();
    const onToggleComplete = vi.fn();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={onToggleComplete}
        onUpdateText={vi.fn()}
      />,
    );

    // when
    await user.click(screen.getByRole("checkbox", { name: "Buy milk" }));

    // then
    expect(onToggleComplete).toHaveBeenCalledWith(1, true);
  });

  it("should call onToggleComplete with false for completed todo", async () => {
    // given
    const user = userEvent.setup();
    const onToggleComplete = vi.fn();
    const todo = createTodo({ isComplete: true });
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={onToggleComplete}
        onUpdateText={vi.fn()}
      />,
    );

    // when
    await user.click(screen.getByRole("checkbox", { name: "Buy milk" }));

    // then
    expect(onToggleComplete).toHaveBeenCalledWith(1, false);
  });

  it("should enter edit mode on double-click", async () => {
    // given
    const user = userEvent.setup();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={vi.fn()}
      />,
    );

    // when
    await user.dblClick(screen.getByText("Buy milk"));

    // then
    const input = screen.getByDisplayValue("Buy milk");
    expect(input).toBeDefined();
  });

  it("should call onUpdateText on Enter in edit mode", async () => {
    // given
    const user = userEvent.setup();
    const onUpdateText = vi.fn();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={onUpdateText}
      />,
    );

    // when
    await user.dblClick(screen.getByText("Buy milk"));
    const input = screen.getByDisplayValue("Buy milk");
    await user.clear(input);
    await user.type(input, "Buy eggs{Enter}");

    // then
    expect(onUpdateText).toHaveBeenCalledWith(1, "Buy eggs");
  });

  it("should cancel edit on Escape", async () => {
    // given
    const user = userEvent.setup();
    const onUpdateText = vi.fn();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={onUpdateText}
      />,
    );

    // when
    await user.dblClick(screen.getByText("Buy milk"));
    const input = screen.getByDisplayValue("Buy milk");
    await user.clear(input);
    await user.type(input, "Buy eggs{Escape}");

    // then
    expect(onUpdateText).not.toHaveBeenCalled();
    expect(screen.getByText("Buy milk")).toBeDefined();
  });

  it("should call onUpdateText on blur in edit mode", async () => {
    // given
    const user = userEvent.setup();
    const onUpdateText = vi.fn();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={onUpdateText}
      />,
    );

    // when
    await user.dblClick(screen.getByText("Buy milk"));
    const input = screen.getByDisplayValue("Buy milk");
    await user.clear(input);
    await user.type(input, "Buy eggs");
    await user.tab();

    // then
    expect(onUpdateText).toHaveBeenCalledWith(1, "Buy eggs");
  });

  it("should not call onUpdateText if text is unchanged on blur", async () => {
    // given
    const user = userEvent.setup();
    const onUpdateText = vi.fn();
    const todo = createTodo();
    render(
      <TodoItem
        todo={todo}
        onToggleComplete={vi.fn()}
        onUpdateText={onUpdateText}
      />,
    );

    // when
    await user.dblClick(screen.getByText("Buy milk"));
    await user.tab();

    // then
    expect(onUpdateText).not.toHaveBeenCalled();
  });
});
