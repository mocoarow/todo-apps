// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'update_todo_request.freezed.dart';
part 'update_todo_request.g.dart';

@Freezed()
abstract class UpdateTodoRequest with _$UpdateTodoRequest {
  const factory UpdateTodoRequest({
    required String text,
    required bool isComplete,
  }) = _UpdateTodoRequest;

  factory UpdateTodoRequest.fromJson(Map<String, Object?> json) => _$UpdateTodoRequestFromJson(json);
}
