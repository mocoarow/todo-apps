import { cleanup, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, describe, expect, it, vi } from "vitest";
import { TodoCreateForm } from "./TodoCreateForm";

describe("TodoCreateForm", () => {
  afterEach(() => {
    cleanup();
  });

  it("should render input and add button", () => {
    // given
    const onCreateTodo = vi.fn();

    // when
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);

    // then
    expect(screen.getByPlaceholderText("What needs to be done?")).toBeDefined();
    expect(screen.getByRole("button", { name: "Add" })).toBeDefined();
  });

  it("should call onCreateTodo with text on submit", async () => {
    // given
    const user = userEvent.setup();
    const onCreateTodo = vi.fn();
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);

    // when
    await user.type(
      screen.getByPlaceholderText("What needs to be done?"),
      "Buy milk",
    );
    await user.click(screen.getByRole("button", { name: "Add" }));

    // then
    expect(onCreateTodo).toHaveBeenCalledWith("Buy milk");
  });

  it("should clear input after successful submit", async () => {
    // given
    const user = userEvent.setup();
    const onCreateTodo = vi.fn().mockResolvedValue(undefined);
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);
    const input = screen.getByPlaceholderText("What needs to be done?");

    // when
    await user.type(input, "Buy milk");
    await user.click(screen.getByRole("button", { name: "Add" }));

    // then
    expect(input).toHaveValue("");
  });

  it("should not call onCreateTodo with empty text", async () => {
    // given
    const user = userEvent.setup();
    const onCreateTodo = vi.fn();
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);

    // when
    await user.click(screen.getByRole("button", { name: "Add" }));

    // then
    expect(onCreateTodo).not.toHaveBeenCalled();
  });

  it("should not call onCreateTodo with whitespace-only text", async () => {
    // given
    const user = userEvent.setup();
    const onCreateTodo = vi.fn();
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);

    // when
    await user.type(
      screen.getByPlaceholderText("What needs to be done?"),
      "   ",
    );
    await user.click(screen.getByRole("button", { name: "Add" }));

    // then
    expect(onCreateTodo).not.toHaveBeenCalled();
  });

  it("should submit on Enter key press", async () => {
    // given
    const user = userEvent.setup();
    const onCreateTodo = vi.fn().mockResolvedValue(undefined);
    render(<TodoCreateForm onCreateTodo={onCreateTodo} />);

    // when
    await user.type(
      screen.getByPlaceholderText("What needs to be done?"),
      "Buy milk{Enter}",
    );

    // then
    expect(onCreateTodo).toHaveBeenCalledWith("Buy milk");
  });

  it("should disable button while isLoading is true", () => {
    // given
    const onCreateTodo = vi.fn();

    // when
    render(<TodoCreateForm onCreateTodo={onCreateTodo} isLoading />);

    // then
    expect(screen.getByRole("button", { name: "Add" })).toBeDisabled();
  });
});
