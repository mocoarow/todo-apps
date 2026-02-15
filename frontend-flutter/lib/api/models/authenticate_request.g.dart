// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'authenticate_request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_AuthenticateRequest _$AuthenticateRequestFromJson(Map<String, dynamic> json) =>
    _AuthenticateRequest(
      loginId: json['loginId'] as String,
      password: json['password'] as String,
    );

Map<String, dynamic> _$AuthenticateRequestToJson(
  _AuthenticateRequest instance,
) => <String, dynamic>{
  'loginId': instance.loginId,
  'password': instance.password,
};
