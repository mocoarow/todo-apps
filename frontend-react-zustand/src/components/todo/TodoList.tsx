import type { GetTodosResponseTodo } from "~/api";
import { TodoItem } from "~/components/todo/TodoItem";

interface TodoListProps {
  todos: GetTodosResponseTodo[];
}

export function TodoList({ todos }: TodoListProps) {
  if (todos.length === 0) {
    return (
      <p className="text-muted-foreground py-8 text-center">No todos yet.</p>
    );
  }

  return (
    <ul className="space-y-2">
      {todos.map((todo) => (
        <TodoItem key={todo.id} todo={todo} />
      ))}
    </ul>
  );
}
