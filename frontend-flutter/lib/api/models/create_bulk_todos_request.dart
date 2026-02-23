// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

import 'create_todo_request.dart';

part 'create_bulk_todos_request.freezed.dart';
part 'create_bulk_todos_request.g.dart';

@Freezed()
abstract class CreateBulkTodosRequest with _$CreateBulkTodosRequest {
  const factory CreateBulkTodosRequest({
    required List<CreateTodoRequest> todos,
  }) = _CreateBulkTodosRequest;

  factory CreateBulkTodosRequest.fromJson(Map<String, Object?> json) => _$CreateBulkTodosRequestFromJson(json);
}
