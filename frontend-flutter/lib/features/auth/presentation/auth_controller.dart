import 'package:flutter/foundation.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';

part 'auth_controller.g.dart';

@Riverpod(keepAlive: true)
class AuthController extends _$AuthController {
  @override
  Future<GetMeResponse?> build() async {
    try {
      return await ref.read(authRepositoryProvider).getMe();
    } on UnauthenticatedException {
      return null;
    }
  }

  Future<void> login({
    required String loginId,
    required String password,
  }) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final repository = ref.read(authRepositoryProvider);
      await repository.authenticate(loginId: loginId, password: password);
      return repository.getMe();
    });
  }

  Future<void> logout() async {
    try {
      await ref.read(authRepositoryProvider).logout();
    } on Exception catch (e) {
      debugPrint('Logout failed: $e');
    }
    state = const AsyncData(null);
  }
}
