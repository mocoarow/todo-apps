// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'find_todo_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_FindTodoResponse _$FindTodoResponseFromJson(Map<String, dynamic> json) =>
    _FindTodoResponse(
      todos: (json['todos'] as List<dynamic>)
          .map((e) => FindTodoResponseTodo.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$FindTodoResponseToJson(_FindTodoResponse instance) =>
    <String, dynamic>{'todos': instance.todos};
