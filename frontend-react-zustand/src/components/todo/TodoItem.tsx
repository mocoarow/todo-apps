import type { GetTodosResponseTodo } from "~/api";
import { Checkbox } from "~/components/ui/checkbox";

interface TodoItemProps {
  todo: GetTodosResponseTodo;
}

export function TodoItem({ todo }: TodoItemProps) {
  return (
    <li className="flex items-center gap-3 rounded-md border px-4 py-3">
      <Checkbox checked={todo.isComplete} disabled aria-label={todo.text} />
      <span
        className={
          todo.isComplete ? "text-muted-foreground line-through" : undefined
        }
      >
        {todo.text}
      </span>
    </li>
  );
}
