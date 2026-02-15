// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'error_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_ErrorResponse _$ErrorResponseFromJson(Map<String, dynamic> json) =>
    _ErrorResponse(
      code: json['code'] as String,
      message: json['message'] as String,
    );

Map<String, dynamic> _$ErrorResponseToJson(_ErrorResponse instance) =>
    <String, dynamic>{'code': instance.code, 'message': instance.message};
