import { Button } from "~/components/ui/button";
import { useAuthStore } from "~/stores/auth";

export function TopPage() {
  const { user, logout, isLoading } = useAuthStore();

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-4">
      <h1 className="text-2xl font-bold">Todo App</h1>
      <p className="text-muted-foreground">Welcome, {user?.loginId}</p>
      <Button variant="outline" onClick={logout} disabled={isLoading}>
        Logout
      </Button>
    </div>
  );
}
