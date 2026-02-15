// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'create_bulk_todos_request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_CreateBulkTodosRequest _$CreateBulkTodosRequestFromJson(
  Map<String, dynamic> json,
) => _CreateBulkTodosRequest(
  todos: (json['todos'] as List<dynamic>)
      .map((e) => CreateTodoRequest.fromJson(e as Map<String, dynamic>))
      .toList(),
);

Map<String, dynamic> _$CreateBulkTodosRequestToJson(
  _CreateBulkTodosRequest instance,
) => <String, dynamic>{'todos': instance.todos};
