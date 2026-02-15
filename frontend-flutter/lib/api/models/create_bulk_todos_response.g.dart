// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'create_bulk_todos_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_CreateBulkTodosResponse _$CreateBulkTodosResponseFromJson(
  Map<String, dynamic> json,
) => _CreateBulkTodosResponse(
  todos: (json['todos'] as List<dynamic>)
      .map((e) => CreateTodoResponse.fromJson(e as Map<String, dynamic>))
      .toList(),
);

Map<String, dynamic> _$CreateBulkTodosResponseToJson(
  _CreateBulkTodosResponse instance,
) => <String, dynamic>{'todos': instance.todos};
