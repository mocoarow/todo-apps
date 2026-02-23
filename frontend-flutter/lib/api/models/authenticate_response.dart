// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'authenticate_response.freezed.dart';
part 'authenticate_response.g.dart';

/// Authentication response. When X-Token-Delivery is 'json' (default), accessToken is returned in the body. When 'cookie', the token is delivered via Set-Cookie header and accessToken is omitted.
@Freezed()
abstract class AuthenticateResponse with _$AuthenticateResponse {
  const factory AuthenticateResponse({
    /// JWT access token (omitted when delivered via cookie)
    String? accessToken,
  }) = _AuthenticateResponse;

  factory AuthenticateResponse.fromJson(Map<String, Object?> json) => _$AuthenticateResponseFromJson(json);
}
