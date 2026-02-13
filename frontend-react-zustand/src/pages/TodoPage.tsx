import { useCallback, useEffect } from "react";
import { TodoCreateForm } from "~/components/todo/TodoCreateForm";
import { TodoList } from "~/components/todo/TodoList";
import { Button } from "~/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "~/components/ui/card";
import { useAuthStore } from "~/stores/auth";
import { useTodoStore } from "~/stores/todo";

export function TodoPage() {
  const { user, logout, isLoading: isAuthLoading } = useAuthStore();
  const {
    todos,
    isLoading,
    isCreating,
    error,
    fetchTodos,
    createTodo,
    updateTodo,
  } = useTodoStore();

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  const handleToggleComplete = useCallback(
    (id: number, isComplete: boolean) => {
      const todo = todos.find((t) => t.id === id);
      if (todo == null) return;
      updateTodo(id, { text: todo.text, isComplete });
    },
    [todos, updateTodo],
  );

  const handleUpdateText = useCallback(
    (id: number, text: string) => {
      const todo = todos.find((t) => t.id === id);
      if (todo == null) return;
      updateTodo(id, { text, isComplete: todo.isComplete });
    },
    [todos, updateTodo],
  );

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
        <CardContent className="space-y-4">
          <TodoCreateForm onCreateTodo={createTodo} isLoading={isCreating} />
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
            <TodoList
              todos={todos}
              onToggleComplete={handleToggleComplete}
              onUpdateText={handleUpdateText}
            />
          )}
        </CardContent>
      </Card>
    </div>
  );
}
