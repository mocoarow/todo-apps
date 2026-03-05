import 'package:flutter/material.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/features/todo/presentation/widgets/todo_item_tile.dart';

class TodoListView extends StatelessWidget {
  const TodoListView({required this.todos, this.onEditTodo, super.key});

  final List<FindTodoResponseTodo> todos;
  final void Function(FindTodoResponseTodo todo)? onEditTodo;

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: todos.length,
      itemBuilder: (context, index) {
        final todo = todos[index];
        return TodoItemTile(
          todo: todo,
          onTitleTap: onEditTodo != null ? () => onEditTodo!(todo) : null,
        );
      },
    );
  }
}
