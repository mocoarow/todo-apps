// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:freezed_annotation/freezed_annotation.dart';

part 'get_me_response.freezed.dart';
part 'get_me_response.g.dart';

@Freezed()
abstract class GetMeResponse with _$GetMeResponse {
  const factory GetMeResponse({
    required int userId,
    required String loginId,
  }) = _GetMeResponse;
  
  factory GetMeResponse.fromJson(Map<String, Object?> json) => _$GetMeResponseFromJson(json);
}
