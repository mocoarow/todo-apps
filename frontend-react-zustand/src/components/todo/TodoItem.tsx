import { type KeyboardEvent, useRef, useState } from "react";
import type { GetTodosResponseTodo } from "~/api";
import { Checkbox } from "~/components/ui/checkbox";
import { Input } from "~/components/ui/input";

interface TodoItemProps {
  todo: GetTodosResponseTodo;
  onToggleComplete: (id: number, isComplete: boolean) => void;
  onUpdateText: (id: number, text: string) => void;
}

export function TodoItem({
  todo,
  onToggleComplete,
  onUpdateText,
}: TodoItemProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [editText, setEditText] = useState(todo.text);
  const cancelledRef = useRef(false);

  const handleDoubleClick = () => {
    setEditText(todo.text);
    cancelledRef.current = false;
    setIsEditing(true);
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      const trimmed = editText.trim();
      if (trimmed !== "" && trimmed !== todo.text) {
        onUpdateText(todo.id, trimmed);
      }
      setIsEditing(false);
    } else if (e.key === "Escape") {
      cancelledRef.current = true;
      setEditText(todo.text);
      setIsEditing(false);
    }
  };

  const handleBlur = () => {
    if (cancelledRef.current) return;
    const trimmed = editText.trim();
    if (trimmed !== "" && trimmed !== todo.text) {
      onUpdateText(todo.id, trimmed);
    }
    setIsEditing(false);
  };

  return (
    <li className="flex items-center gap-3 rounded-md border px-4 py-3">
      <Checkbox
        checked={todo.isComplete}
        onCheckedChange={() => onToggleComplete(todo.id, !todo.isComplete)}
        aria-label={todo.text}
      />
      {isEditing ? (
        <Input
          value={editText}
          onChange={(e) => setEditText(e.target.value)}
          onKeyDown={handleKeyDown}
          onBlur={handleBlur}
          maxLength={250}
          autoFocus
        />
      ) : (
        <button
          type="button"
          className={
            todo.isComplete ? "text-muted-foreground line-through" : undefined
          }
          onDoubleClick={handleDoubleClick}
        >
          {todo.text}
        </button>
      )}
    </li>
  );
}
