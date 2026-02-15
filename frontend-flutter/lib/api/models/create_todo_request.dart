// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'create_todo_request.freezed.dart';
part 'create_todo_request.g.dart';

@Freezed()
abstract class CreateTodoRequest with _$CreateTodoRequest {
  const factory CreateTodoRequest({
    required String text,
  }) = _CreateTodoRequest;
  
  factory CreateTodoRequest.fromJson(Map<String, Object?> json) => _$CreateTodoRequestFromJson(json);
}
