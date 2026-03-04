import 'package:todo/api/models/find_todo_response_todo.dart';

final _now = DateTime(2024);

List<FindTodoResponseTodo> sampleTodos() => [
  FindTodoResponseTodo(id: 1, text: 'Buy milk', isComplete: false, createdAt: _now, updatedAt: _now),
  FindTodoResponseTodo(id: 2, text: 'Walk dog', isComplete: true, createdAt: _now, updatedAt: _now),
];
