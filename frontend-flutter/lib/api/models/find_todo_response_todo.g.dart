// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'find_todo_response_todo.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_FindTodoResponseTodo _$FindTodoResponseTodoFromJson(
  Map<String, dynamic> json,
) => _FindTodoResponseTodo(
  id: (json['id'] as num).toInt(),
  text: json['text'] as String,
  isComplete: json['isComplete'] as bool,
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
);

Map<String, dynamic> _$FindTodoResponseTodoToJson(
  _FindTodoResponseTodo instance,
) => <String, dynamic>{
  'id': instance.id,
  'text': instance.text,
  'isComplete': instance.isComplete,
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
};
