// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

import 'create_todo_response.dart';

part 'create_bulk_todos_response.freezed.dart';
part 'create_bulk_todos_response.g.dart';

@Freezed()
abstract class CreateBulkTodosResponse with _$CreateBulkTodosResponse {
  const factory CreateBulkTodosResponse({
    required List<CreateTodoResponse> todos,
  }) = _CreateBulkTodosResponse;
  
  factory CreateBulkTodosResponse.fromJson(Map<String, Object?> json) => _$CreateBulkTodosResponseFromJson(json);
}
