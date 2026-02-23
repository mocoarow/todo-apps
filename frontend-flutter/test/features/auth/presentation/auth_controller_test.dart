import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';
import 'package:todo/features/auth/presentation/auth_controller.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository mockRepository;

  setUp(() {
    mockRepository = MockAuthRepository();
  });

  ProviderContainer createContainer() {
    final container = ProviderContainer(
      overrides: [
        authRepositoryProvider.overrideWithValue(mockRepository),
      ],
    );
    addTearDown(container.dispose);
    return container;
  }

  group('Test_AuthController_build', () {
    test('shouldReturnUser_whenGetMeSucceeds', () async {
      // given
      const user = GetMeResponse(userId: 1, loginId: 'testuser');
      when(() => mockRepository.getMe()).thenAnswer((_) async => user);
      final container = createContainer();

      // when
      final result = await container.read(authControllerProvider.future);

      // then
      expect(result, user);
    });

    test('shouldReturnNull_whenUnauthenticatedExceptionIsThrown', () async {
      // given
      when(() => mockRepository.getMe())
          .thenThrow(const UnauthenticatedException());
      final container = createContainer();

      // when
      final result = await container.read(authControllerProvider.future);

      // then
      expect(result, isNull);
    });
  });

  group('Test_AuthController_login', () {
    test('shouldSetAsyncDataWithUser_whenLoginSucceeds', () async {
      // given
      const user = GetMeResponse(userId: 1, loginId: 'testuser');
      when(() => mockRepository.getMe())
          .thenThrow(const UnauthenticatedException());
      when(
        () => mockRepository.authenticate(
          loginId: any(named: 'loginId'),
          password: any(named: 'password'),
        ),
      ).thenAnswer((_) async {});
      final container = createContainer();
      await container.read(authControllerProvider.future);

      // After build completes, update getMe to return user
      when(() => mockRepository.getMe()).thenAnswer((_) async => user);

      // when
      await container.read(authControllerProvider.notifier).login(
            loginId: 'testuser',
            password: 'password123',
          );

      // then
      final state = container.read(authControllerProvider);
      expect(state, isA<AsyncData<GetMeResponse?>>());
      expect(state.value, user);
    });

    test('shouldSetAsyncError_whenInvalidCredentialsExceptionIsThrown',
        () async {
      // given
      when(() => mockRepository.getMe())
          .thenThrow(const UnauthenticatedException());
      when(
        () => mockRepository.authenticate(
          loginId: any(named: 'loginId'),
          password: any(named: 'password'),
        ),
      ).thenThrow(const InvalidCredentialsException());
      final container = createContainer();
      await container.read(authControllerProvider.future);

      // when
      await container.read(authControllerProvider.notifier).login(
            loginId: 'testuser',
            password: 'wrongpass',
          );

      // then
      final state = container.read(authControllerProvider);
      expect(state, isA<AsyncError<GetMeResponse?>>());
      expect(state.error, isA<InvalidCredentialsException>());
    });
  });

  group('Test_AuthController_logout', () {
    test('shouldSetAsyncDataNull_whenLogoutSucceeds', () async {
      // given
      const user = GetMeResponse(userId: 1, loginId: 'testuser');
      when(() => mockRepository.getMe()).thenAnswer((_) async => user);
      when(() => mockRepository.logout()).thenAnswer((_) async {});
      final container = createContainer();
      await container.read(authControllerProvider.future);

      // when
      await container.read(authControllerProvider.notifier).logout();

      // then
      final state = container.read(authControllerProvider);
      expect(state, isA<AsyncData<GetMeResponse?>>());
      expect(state.value, isNull);
    });

    test('shouldSetAsyncDataNull_whenLogoutFails', () async {
      // given
      const user = GetMeResponse(userId: 1, loginId: 'testuser');
      when(() => mockRepository.getMe()).thenAnswer((_) async => user);
      when(() => mockRepository.logout())
          .thenThrow(const AuthNetworkException());
      final container = createContainer();
      await container.read(authControllerProvider.future);

      // when
      await container.read(authControllerProvider.notifier).logout();

      // then
      final state = container.read(authControllerProvider);
      expect(state, isA<AsyncData<GetMeResponse?>>());
      expect(state.value, isNull);
    });
  });
}
