import 'package:dio/dio.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/auth/auth_client.dart';
import 'package:todo/api/models/authenticate_request.dart';
import 'package:todo/api/models/authenticate_response.dart';
import 'package:todo/api/models/get_me_response.dart';
import 'package:todo/api/models/x_token_delivery.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';

class MockAuthClient extends Mock implements AuthClient {}

void main() {
  late MockAuthClient mockClient;
  late AuthRepository repository;

  setUpAll(() {
    registerFallbackValue(
      const AuthenticateRequest(loginId: '', password: ''),
    );
    registerFallbackValue(XTokenDelivery.cookie);
  });

  setUp(() {
    mockClient = MockAuthClient();
    repository = AuthRepository(mockClient);
  });

  group('Test_AuthRepository_authenticate', () {
    test('shouldComplete_whenRequestSucceeds', () async {
      // given
      when(
        () => mockClient.authenticate(
          body: any(named: 'body'),
          xTokenDelivery: any(named: 'xTokenDelivery'),
        ),
      ).thenAnswer((_) async => const AuthenticateResponse());

      // when & then
      await expectLater(
        repository.authenticate(loginId: 'user', password: 'password'),
        completes,
      );
    });

    test('shouldThrowInvalidCredentialsException_whenStatusIs401', () async {
      // given
      when(
        () => mockClient.authenticate(
          body: any(named: 'body'),
          xTokenDelivery: any(named: 'xTokenDelivery'),
        ),
      ).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          response: Response(
            statusCode: 401,
            requestOptions: RequestOptions(),
          ),
          type: DioExceptionType.badResponse,
        ),
      );

      // when & then
      await expectLater(
        repository.authenticate(loginId: 'user', password: 'password'),
        throwsA(isA<InvalidCredentialsException>()),
      );
    });

    test('shouldThrowAuthNetworkException_whenOtherDioExceptionOccurs', () async {
      // given
      when(
        () => mockClient.authenticate(
          body: any(named: 'body'),
          xTokenDelivery: any(named: 'xTokenDelivery'),
        ),
      ).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.authenticate(loginId: 'user', password: 'password'),
        throwsA(isA<AuthNetworkException>()),
      );
    });

    test('shouldThrowAuthNetworkException_whenNonDioExceptionOccurs', () async {
      // given
      when(
        () => mockClient.authenticate(
          body: any(named: 'body'),
          xTokenDelivery: any(named: 'xTokenDelivery'),
        ),
      ).thenThrow(const FormatException('invalid json'));

      // when & then
      await expectLater(
        repository.authenticate(loginId: 'user', password: 'password'),
        throwsA(isA<AuthNetworkException>()),
      );
    });
  });

  group('Test_AuthRepository_getMe', () {
    test('shouldReturnGetMeResponse_whenRequestSucceeds', () async {
      // given
      const expected = GetMeResponse(userId: 1, loginId: 'user');
      when(() => mockClient.getMe()).thenAnswer((_) async => expected);

      // when
      final result = await repository.getMe();

      // then
      expect(result, expected);
    });

    test('shouldThrowUnauthenticatedException_whenStatusIs401', () async {
      // given
      when(() => mockClient.getMe()).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          response: Response(
            statusCode: 401,
            requestOptions: RequestOptions(),
          ),
          type: DioExceptionType.badResponse,
        ),
      );

      // when & then
      await expectLater(
        repository.getMe(),
        throwsA(isA<UnauthenticatedException>()),
      );
    });

    test('shouldThrowAuthNetworkException_whenOtherDioExceptionOccurs', () async {
      // given
      when(() => mockClient.getMe()).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.getMe(),
        throwsA(isA<AuthNetworkException>()),
      );
    });

    test('shouldThrowAuthNetworkException_whenNonDioExceptionOccurs', () async {
      // given
      when(() => mockClient.getMe()).thenThrow(const FormatException('invalid json'));

      // when & then
      await expectLater(
        repository.getMe(),
        throwsA(isA<AuthNetworkException>()),
      );
    });
  });

  group('Test_AuthRepository_logout', () {
    test('shouldComplete_whenRequestSucceeds', () async {
      // given
      when(() => mockClient.logout()).thenAnswer((_) async {});

      // when & then
      await expectLater(repository.logout(), completes);
    });

    test('shouldThrowAuthNetworkException_whenDioExceptionOccurs', () async {
      // given
      when(() => mockClient.logout()).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.logout(),
        throwsA(isA<AuthNetworkException>()),
      );
    });

    test('shouldThrowAuthNetworkException_whenNonDioExceptionOccurs', () async {
      // given
      when(() => mockClient.logout()).thenThrow(const FormatException('invalid json'));

      // when & then
      await expectLater(
        repository.logout(),
        throwsA(isA<AuthNetworkException>()),
      );
    });
  });
}
