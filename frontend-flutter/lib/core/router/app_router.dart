import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/features/auth/presentation/auth_controller.dart';
import 'package:todo/features/auth/presentation/login_screen.dart';
import 'package:todo/features/todo/presentation/todo_screen.dart';

part 'app_router.g.dart';

@Riverpod(keepAlive: true)
GoRouter appRouter(Ref ref) {
  final authNotifier = ValueNotifier<AsyncValue<GetMeResponse?>>(const AsyncValue.loading());

  ref
    ..listen(authControllerProvider, (_, next) {
      authNotifier.value = next;
    })
    ..onDispose(authNotifier.dispose);

  return GoRouter(
    initialLocation: '/splash',
    refreshListenable: authNotifier,
    redirect: (context, state) {
      final authState = authNotifier.value;
      final isSplashRoute = state.matchedLocation == '/splash';

      if (authState.isLoading) {
        return isSplashRoute ? null : '/splash';
      }

      final isAuthenticated = authState.value != null;
      final isLoginRoute = state.matchedLocation == '/login';

      if (isSplashRoute) return isAuthenticated ? '/' : '/login';
      if (!isAuthenticated && !isLoginRoute) return '/login';
      if (isAuthenticated && isLoginRoute) return '/';
      return null;
    },
    routes: [
      GoRoute(
        path: '/splash',
        builder: (context, state) => const _SplashScreen(),
      ),
      GoRoute(
        path: '/login',
        builder: (context, state) => const LoginScreen(),
      ),
      GoRoute(
        path: '/',
        builder: (context, state) => const TodoScreen(),
      ),
    ],
  );
}

class _SplashScreen extends StatelessWidget {
  const _SplashScreen();

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(child: CircularProgressIndicator()),
    );
  }
}
