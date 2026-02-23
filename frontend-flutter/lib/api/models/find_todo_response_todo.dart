// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'find_todo_response_todo.freezed.dart';
part 'find_todo_response_todo.g.dart';

@Freezed()
abstract class FindTodoResponseTodo with _$FindTodoResponseTodo {
  const factory FindTodoResponseTodo({
    required int id,
    required String text,
    required bool isComplete,
    required DateTime createdAt,
    required DateTime updatedAt,
  }) = _FindTodoResponseTodo;

  factory FindTodoResponseTodo.fromJson(Map<String, Object?> json) => _$FindTodoResponseTodoFromJson(json);
}
