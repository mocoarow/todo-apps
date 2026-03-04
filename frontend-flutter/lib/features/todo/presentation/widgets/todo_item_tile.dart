import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/features/todo/presentation/todo_controller.dart';

class TodoItemTile extends ConsumerWidget {
  const TodoItemTile({required this.todo, super.key});

  final FindTodoResponseTodo todo;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return CheckboxListTile(
      value: todo.isComplete,
      onChanged: (_) => ref.read(todoControllerProvider.notifier).toggleComplete(todo),
      title: Text(
        todo.text,
        style: todo.isComplete ? const TextStyle(decoration: TextDecoration.lineThrough) : null,
      ),
    );
  }
}
