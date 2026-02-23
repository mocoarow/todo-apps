import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';
import 'package:todo/features/auth/presentation/widgets/login_form.dart';

import '../../../../helpers/test_helpers.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository mockRepository;

  setUp(() {
    mockRepository = MockAuthRepository();
    when(() => mockRepository.getMe())
        .thenThrow(const UnauthenticatedException());
  });

  group('Test_LoginForm_validation', () {
    testWidgets('shouldShowError_whenLoginIdIsEmpty', (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      expect(find.text('Login ID is required'), findsOneWidget);
    });

    testWidgets('shouldShowError_whenLoginIdExceeds100Characters',
        (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(
        find.byType(TextFormField).first,
        'a' * 101,
      );
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      expect(
        find.text('Login ID must be 100 characters or less'),
        findsOneWidget,
      );
    });

    testWidgets('shouldShowError_whenPasswordIsEmpty', (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(find.byType(TextFormField).first, 'testuser');
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      expect(find.text('Password is required'), findsOneWidget);
    });

    testWidgets('shouldShowError_whenPasswordIsTooShort', (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(find.byType(TextFormField).last, '1234567');
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      expect(
        find.text('Password must be at least 8 characters'),
        findsOneWidget,
      );
    });

    testWidgets('shouldShowError_whenPasswordIsTooLong', (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.enterText(
        find.byType(TextFormField).last,
        'a' * 21,
      );
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      expect(
        find.text('Password must be 20 characters or less'),
        findsOneWidget,
      );
    });
  });

  group('Test_LoginForm_submit', () {
    testWidgets('shouldCallLogin_whenInputIsValid', (tester) async {
      // given
      when(
        () => mockRepository.authenticate(
          loginId: any(named: 'loginId'),
          password: any(named: 'password'),
        ),
      ).thenAnswer((_) async {});
      when(() => mockRepository.getMe()).thenAnswer(
        (_) async => const GetMeResponse(userId: 1, loginId: 'testuser'),
      );

      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // Re-stub getMe for the login flow (build already consumed the throw)
      when(() => mockRepository.getMe()).thenAnswer(
        (_) async => const GetMeResponse(userId: 1, loginId: 'testuser'),
      );

      // when
      await tester.enterText(find.byType(TextFormField).first, 'testuser');
      await tester.enterText(
        find.byType(TextFormField).last,
        'password123',
      );
      await tester.tap(find.text('Login'));
      await tester.pumpAndSettle();

      // then
      verify(
        () => mockRepository.authenticate(
          loginId: 'testuser',
          password: 'password123',
        ),
      ).called(1);
    });
  });

  group('Test_LoginForm_loading', () {
    testWidgets('shouldDisableButton_whenIsLoading', (tester) async {
      // given & when
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: true)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pump();

      // then
      final button = tester.widget<FilledButton>(find.byType(FilledButton));
      expect(button.onPressed, isNull);
    });

    testWidgets('shouldShowProgressIndicator_whenIsLoading', (tester) async {
      // given & when
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: true)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pump();

      // then
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
      expect(find.text('Login'), findsNothing);
    });
  });

  group('Test_LoginForm_passwordToggle', () {
    testWidgets('shouldTogglePasswordVisibility', (tester) async {
      // given
      await tester.pumpApp(
        Scaffold(body: const LoginForm(isLoading: false)),
        overrides: [
          authRepositoryProvider.overrideWithValue(mockRepository),
        ],
      );
      await tester.pumpAndSettle();

      // Initially password is obscured
      expect(find.byIcon(Icons.visibility_off), findsOneWidget);

      // when
      await tester.tap(find.byIcon(Icons.visibility_off));
      await tester.pump();

      // then - password is now visible
      expect(find.byIcon(Icons.visibility), findsOneWidget);
      expect(find.byIcon(Icons.visibility_off), findsNothing);
    });
  });
}
