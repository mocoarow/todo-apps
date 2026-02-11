import { cleanup, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { useAuthStore } from "~/stores/auth";
import { TopPage } from "./TopPage";

vi.mock("~/gateway/auth", () => ({
  authService: {
    login: vi.fn(),
    logout: vi.fn(),
    getMe: vi.fn(),
  },
}));

describe("TopPage", () => {
  beforeEach(() => {
    useAuthStore.setState({
      user: { userId: 1, loginId: "testuser" },
      isLoading: false,
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
    render(<TopPage />);

    // then
    expect(screen.getByText("Welcome, testuser")).toBeDefined();
  });

  it("should call logout when logout button is clicked", async () => {
    // given
    const user = userEvent.setup();
    render(<TopPage />);

    // when
    await user.click(screen.getByRole("button", { name: "Logout" }));

    // then
    expect(useAuthStore.getState().user).toBeNull();
  });

  it("should disable logout button while loading", () => {
    // given
    useAuthStore.setState({ isLoading: true });

    // when
    render(<TopPage />);

    // then
    expect(screen.getByRole("button", { name: "Logout" })).toBeDisabled();
  });
});
