// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'update_todo_request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_UpdateTodoRequest _$UpdateTodoRequestFromJson(Map<String, dynamic> json) => _UpdateTodoRequest(
  text: json['text'] as String,
  isComplete: json['isComplete'] as bool,
);

Map<String, dynamic> _$UpdateTodoRequestToJson(_UpdateTodoRequest instance) => <String, dynamic>{
  'text': instance.text,
  'isComplete': instance.isComplete,
};
