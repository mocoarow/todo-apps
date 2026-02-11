import { cleanup, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { AppError } from "~/domain/error";
import { useAuthStore } from "~/stores/auth";
import { LoginPage } from "./LoginPage";

vi.mock("~/gateway/auth", () => ({
  authService: {
    login: vi.fn(),
    logout: vi.fn(),
    getMe: vi.fn(),
  },
}));

describe("LoginPage", () => {
  beforeEach(() => {
    useAuthStore.setState({
      user: null,
      isLoading: false,
      error: null,
    });
  });

  afterEach(() => {
    cleanup();
    vi.clearAllMocks();
  });

  it("should render login form with loginId and password fields", () => {
    // given
    // (default state)

    // when
    render(<LoginPage />);

    // then
    expect(screen.getByLabelText("Login ID")).toBeDefined();
    expect(screen.getByLabelText("Password")).toBeDefined();
    expect(screen.getByRole("button", { name: "Login" })).toBeDefined();
  });

  it("should display error message when error exists", () => {
    // given
    useAuthStore.setState({
      error: new AppError("UNAUTHENTICATED", "Invalid credentials"),
    });

    // when
    render(<LoginPage />);

    // then
    expect(screen.getByText("Invalid credentials")).toBeDefined();
  });

  it("should disable button and show loading text while loading", () => {
    // given
    useAuthStore.setState({ isLoading: true });

    // when
    render(<LoginPage />);

    // then
    const button = screen.getByRole("button", { name: "Logging in..." });
    expect(button).toBeDisabled();
  });

  it("should call login with input values on submit", async () => {
    // given
    const user = userEvent.setup();
    const mockLogin = vi.fn();
    useAuthStore.setState({ login: mockLogin });
    render(<LoginPage />);

    // when
    await user.type(screen.getByLabelText("Login ID"), "testuser");
    await user.type(screen.getByLabelText("Password"), "password1234");
    await user.click(screen.getByRole("button", { name: "Login" }));

    // then
    expect(mockLogin).toHaveBeenCalledWith("testuser", "password1234");
  });
});
