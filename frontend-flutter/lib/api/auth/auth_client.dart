// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';

import '../models/authenticate_request.dart';
import '../models/authenticate_response.dart';
import '../models/get_me_response.dart';
import '../models/x_token_delivery.dart';

part 'auth_client.g.dart';

@RestApi()
abstract class AuthClient {
  factory AuthClient(Dio dio, {String? baseUrl}) = _AuthClient;

  /// User authentication.
  ///
  /// Authenticate user with login ID and password.
  ///
  /// [xTokenDelivery] - Token delivery method (json or cookie).
  @POST('/api/v1/auth/authenticate')
  Future<AuthenticateResponse> authenticate({
    @Body() required AuthenticateRequest body,
    @Header('X-Token-Delivery') XTokenDelivery? xTokenDelivery = XTokenDelivery.cookie,
  });

  /// User logout.
  ///
  /// Clear the access-token cookie to log the user out.
  @POST('/api/v1/auth/logout')
  Future<void> logout();

  /// Get current user.
  ///
  /// Get the authenticated user's information.
  @GET('/api/v1/auth/me')
  Future<GetMeResponse> getMe();
}
