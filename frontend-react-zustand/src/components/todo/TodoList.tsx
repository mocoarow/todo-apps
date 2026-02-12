import type { GetTodosResponseTodo } from "~/api";
import { TodoItem } from "~/components/todo/TodoItem";

interface TodoListProps {
  todos: GetTodosResponseTodo[];
  onToggleComplete: (id: number, isComplete: boolean) => void;
  onUpdateText: (id: number, text: string) => void;
}

export function TodoList({
  todos,
  onToggleComplete,
  onUpdateText,
}: TodoListProps) {
  if (todos.length === 0) {
    return (
      <p className="text-muted-foreground py-8 text-center">No todos yet.</p>
    );
  }

  return (
    <ul className="space-y-2">
      {todos.map((todo) => (
        <TodoItem
          key={todo.id}
          todo={todo}
          onToggleComplete={onToggleComplete}
          onUpdateText={onUpdateText}
        />
      ))}
    </ul>
  );
}
