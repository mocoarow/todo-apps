import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:todo/core/constants/app_constants.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/presentation/auth_controller.dart';
import 'package:todo/features/auth/presentation/widgets/login_form.dart';

class LoginScreen extends ConsumerWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    ref.listen(authControllerProvider, (previous, next) {
      if (next.hasError && !next.isLoading) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(_resolveErrorMessage(next.error)),
            backgroundColor: Theme.of(context).colorScheme.error,
          ),
        );
      }
    });

    final isLoading = ref.watch(authControllerProvider).isLoading;

    return Scaffold(
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 400),
          child: Padding(
            padding: const EdgeInsets.all(24),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  AppConstants.appTitle,
                  style: Theme.of(context).textTheme.headlineMedium,
                ),
                const SizedBox(height: 32),
                LoginForm(isLoading: isLoading),
              ],
            ),
          ),
        ),
      ),
    );
  }

  String _resolveErrorMessage(Object? error) {
    return switch (error) {
      InvalidCredentialsException() => 'Invalid login ID or password.',
      AuthNetworkException() => 'Connection failed. Please check your network.',
      _ => 'An unexpected error occurred. Please try again.',
    };
  }
}
