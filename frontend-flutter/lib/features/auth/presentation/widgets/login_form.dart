import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:todo/core/constants/app_constants.dart';
import 'package:todo/features/auth/presentation/auth_controller.dart';

class LoginForm extends ConsumerStatefulWidget {
  const LoginForm({required this.isLoading, super.key});

  final bool isLoading;

  @override
  ConsumerState<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends ConsumerState<LoginForm> {
  final _formKey = GlobalKey<FormState>();
  final _loginIdController = TextEditingController();
  final _passwordController = TextEditingController();
  bool _obscurePassword = true;

  @override
  void dispose() {
    _loginIdController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;

    await ref
        .read(authControllerProvider.notifier)
        .login(
          loginId: _loginIdController.text.trim(),
          password: _passwordController.text,
        );
  }

  String? _validateLoginId(String? value) {
    if (value == null || value.trim().isEmpty) {
      return 'Login ID is required';
    }
    if (value.trim().length > AppConstants.loginIdMaxLength) {
      return 'Login ID must be ${AppConstants.loginIdMaxLength} characters or less';
    }
    return null;
  }

  String? _validatePassword(String? value) {
    if (value == null || value.isEmpty) {
      return 'Password is required';
    }
    if (value.length < AppConstants.passwordMinLength) {
      return 'Password must be at least ${AppConstants.passwordMinLength} characters';
    }
    if (value.length > AppConstants.passwordMaxLength) {
      return 'Password must be ${AppConstants.passwordMaxLength} characters or less';
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: AutofillGroup(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextFormField(
              controller: _loginIdController,
              decoration: const InputDecoration(
                labelText: 'Login ID',
                border: OutlineInputBorder(),
              ),
              autofillHints: const [AutofillHints.username],
              textInputAction: TextInputAction.next,
              enabled: !widget.isLoading,
              validator: _validateLoginId,
            ),
            const SizedBox(height: 16),
            TextFormField(
              controller: _passwordController,
              decoration: InputDecoration(
                labelText: 'Password',
                border: const OutlineInputBorder(),
                suffixIcon: IconButton(
                  icon: Icon(
                    _obscurePassword ? Icons.visibility_off : Icons.visibility,
                  ),
                  onPressed: () {
                    setState(() => _obscurePassword = !_obscurePassword);
                  },
                ),
              ),
              obscureText: _obscurePassword,
              autofillHints: const [AutofillHints.password],
              textInputAction: TextInputAction.done,
              enabled: !widget.isLoading,
              validator: _validatePassword,
              onFieldSubmitted: (_) => _submit(),
            ),
            const SizedBox(height: 24),
            SizedBox(
              width: double.infinity,
              child: FilledButton(
                onPressed: widget.isLoading ? null : _submit,
                child: widget.isLoading
                    ? const SizedBox(
                        height: 20,
                        width: 20,
                        child: CircularProgressIndicator(strokeWidth: 2),
                      )
                    : const Text('Login'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
