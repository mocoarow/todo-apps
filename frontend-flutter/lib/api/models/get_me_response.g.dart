// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'get_me_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_GetMeResponse _$GetMeResponseFromJson(Map<String, dynamic> json) =>
    _GetMeResponse(
      userId: (json['userId'] as num).toInt(),
      loginId: json['loginId'] as String,
    );

Map<String, dynamic> _$GetMeResponseToJson(_GetMeResponse instance) =>
    <String, dynamic>{'userId': instance.userId, 'loginId': instance.loginId};
