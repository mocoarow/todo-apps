import 'package:flutter/material.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/features/todo/presentation/widgets/todo_item_tile.dart';

class TodoListView extends StatelessWidget {
  const TodoListView({required this.todos, super.key});

  final List<FindTodoResponseTodo> todos;

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: todos.length,
      itemBuilder: (context, index) => TodoItemTile(todo: todos[index]),
    );
  }
}
