import { useEffect } from "react";
import { TodoList } from "~/components/todo/TodoList";
import { Button } from "~/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "~/components/ui/card";
import { useAuthStore } from "~/stores/auth";
import { useTodoStore } from "~/stores/todo";

export function TodoPage() {
  const { user, logout, isLoading: isAuthLoading } = useAuthStore();
  const { todos, isLoading, error, fetchTodos } = useTodoStore();

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  return (
    <div className="mx-auto flex min-h-svh max-w-2xl flex-col gap-6 p-6">
      <header className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">Todo App</h1>
          <p className="text-muted-foreground text-sm">{user?.loginId}</p>
        </div>
        <Button
          variant="outline"
          size="sm"
          onClick={logout}
          disabled={isAuthLoading}
        >
          Logout
        </Button>
      </header>

      <Card>
        <CardHeader>
          <CardTitle>Todos</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <p className="text-muted-foreground py-8 text-center">Loading...</p>
          ) : error ? (
            <div className="flex flex-col items-center gap-2 py-8">
              <p className="text-destructive text-sm">{error.message}</p>
              <Button variant="outline" size="sm" onClick={fetchTodos}>
                Retry
              </Button>
            </div>
          ) : (
            <TodoList todos={todos} />
          )}
        </CardContent>
      </Card>
    </div>
  );
}
