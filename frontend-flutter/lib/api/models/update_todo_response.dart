// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'update_todo_response.freezed.dart';
part 'update_todo_response.g.dart';

@Freezed()
abstract class UpdateTodoResponse with _$UpdateTodoResponse {
  const factory UpdateTodoResponse({
    required int id,
    required String text,
    required bool isComplete,
    required DateTime createdAt,
    required DateTime updatedAt,
  }) = _UpdateTodoResponse;

  factory UpdateTodoResponse.fromJson(Map<String, Object?> json) => _$UpdateTodoResponseFromJson(json);
}
