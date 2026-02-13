import { type FormEvent, useState } from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

interface TodoCreateFormProps {
  onCreateTodo: (text: string) => Promise<void>;
  isLoading?: boolean;
}

export function TodoCreateForm({
  onCreateTodo,
  isLoading = false,
}: TodoCreateFormProps) {
  const [text, setText] = useState("");

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const trimmed = text.trim();
    if (trimmed === "") return;
    await onCreateTodo(trimmed);
    setText("");
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2">
      <Input
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="What needs to be done?"
        maxLength={250}
        disabled={isLoading}
      />
      <Button type="submit" disabled={isLoading}>
        Add
      </Button>
    </form>
  );
}
