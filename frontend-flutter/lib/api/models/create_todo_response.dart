// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'create_todo_response.freezed.dart';
part 'create_todo_response.g.dart';

@Freezed()
abstract class CreateTodoResponse with _$CreateTodoResponse {
  const factory CreateTodoResponse({
    required int id,
    required String text,
    required bool isComplete,
    required DateTime createdAt,
    required DateTime updatedAt,
  }) = _CreateTodoResponse;
  
  factory CreateTodoResponse.fromJson(Map<String, Object?> json) => _$CreateTodoResponseFromJson(json);
}
