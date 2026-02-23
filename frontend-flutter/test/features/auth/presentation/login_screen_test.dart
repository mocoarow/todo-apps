import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';
import 'package:todo/features/auth/presentation/login_screen.dart';
import 'package:todo/features/auth/presentation/widgets/login_form.dart';

import '../../../helpers/test_helpers.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository mockRepository;

  setUp(() {
    mockRepository = MockAuthRepository();
    when(() => mockRepository.getMe()).thenThrow(const UnauthenticatedException());
  });

  group('Test_LoginScreen', () {
    testWidgets('shouldDisplayAppTitle', (tester) async {
      // given & when
      await tester.pumpApp(
        const LoginScreen(),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.text('Todo App'), findsOneWidget);
    });

    testWidgets('shouldDisplayLoginForm', (tester) async {
      // given & when
      await tester.pumpApp(
        const LoginScreen(),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.byType(LoginForm), findsOneWidget);
    });

    testWidgets('shouldShowSnackBar_whenInvalidCredentialsExceptionOccurs', (tester) async {
      // given
      when(
        () => mockRepository.authenticate(
          loginId: any(named: 'loginId'),
          password: any(named: 'password'),
        ),
      ).thenThrow(const InvalidCredentialsException());

      await tester.pumpApp(
        const LoginScreen(),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(
        find.byType(TextFormField).first,
        'testuser',
      );
      await tester.enterText(
        find.byType(TextFormField).last,
        'password123',
      );
      await tester.tap(find.text('Login'));
      await tester.pump();
      await tester.pump();

      // then
      expect(find.byType(SnackBar), findsOneWidget);
      expect(find.text('Invalid login ID or password.'), findsOneWidget);
    });

    testWidgets('shouldShowSnackBar_whenAuthNetworkExceptionOccurs', (tester) async {
      // given
      when(
        () => mockRepository.authenticate(
          loginId: any(named: 'loginId'),
          password: any(named: 'password'),
        ),
      ).thenThrow(const AuthNetworkException());

      await tester.pumpApp(
        const LoginScreen(),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(
        find.byType(TextFormField).first,
        'testuser',
      );
      await tester.enterText(
        find.byType(TextFormField).last,
        'password123',
      );
      await tester.tap(find.text('Login'));
      await tester.pump();
      await tester.pump();

      // then
      expect(find.byType(SnackBar), findsOneWidget);
      expect(
        find.text('Connection failed. Please check your network.'),
        findsOneWidget,
      );
    });
  });
}
