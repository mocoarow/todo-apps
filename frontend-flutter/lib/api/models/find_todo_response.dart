// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

import 'find_todo_response_todo.dart';

part 'find_todo_response.freezed.dart';
part 'find_todo_response.g.dart';

@Freezed()
abstract class FindTodoResponse with _$FindTodoResponse {
  const factory FindTodoResponse({
    required List<FindTodoResponseTodo> todos,
  }) = _FindTodoResponse;

  factory FindTodoResponse.fromJson(Map<String, Object?> json) => _$FindTodoResponseFromJson(json);
}
