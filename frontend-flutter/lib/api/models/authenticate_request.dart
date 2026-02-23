// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'authenticate_request.freezed.dart';
part 'authenticate_request.g.dart';

@Freezed()
abstract class AuthenticateRequest with _$AuthenticateRequest {
  const factory AuthenticateRequest({
    required String loginId,
    required String password,
  }) = _AuthenticateRequest;

  factory AuthenticateRequest.fromJson(Map<String, Object?> json) => _$AuthenticateRequestFromJson(json);
}
