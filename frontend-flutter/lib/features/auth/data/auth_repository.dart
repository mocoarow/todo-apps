import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/api/auth/auth_client.dart';
import 'package:todo/api/models/authenticate_request.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/core/network/dio_provider.dart';
import 'package:todo/features/auth/data/auth_exception.dart';

part 'auth_repository.g.dart';

@Riverpod(keepAlive: true)
AuthRepository authRepository(Ref ref) {
  final dio = ref.watch(dioProvider);
  return AuthRepository(AuthClient(dio));
}

class AuthRepository {
  const AuthRepository(this._client);

  final AuthClient _client;

  Future<void> authenticate({
    required String loginId,
    required String password,
  }) async {
    try {
      await _client.authenticate(
        body: AuthenticateRequest(loginId: loginId, password: password),
      );
    } on DioException catch (e) {
      _throwAuthException(e, onUnauthorized: const InvalidCredentialsException());
    } on Exception {
      throw const AuthNetworkException();
    }
  }

  Future<GetMeResponse> getMe() async {
    try {
      return await _client.getMe();
    } on DioException catch (e) {
      _throwAuthException(e, onUnauthorized: const UnauthenticatedException());
    } on Exception {
      throw const AuthNetworkException();
    }
  }

  Future<void> logout() async {
    try {
      await _client.logout();
    } on DioException catch (e) {
      _throwAuthException(e);
    } on Exception {
      throw const AuthNetworkException();
    }
  }

  Never _throwAuthException(DioException e, {AuthException? onUnauthorized}) {
    if (e.response?.statusCode == 401 && onUnauthorized != null) {
      throw onUnauthorized;
    }
    throw const AuthNetworkException();
  }
}
