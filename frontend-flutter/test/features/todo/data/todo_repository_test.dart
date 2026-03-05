import 'package:dio/dio.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/create_todo_request.dart';
import 'package:todo/api/models/create_todo_response.dart';
import 'package:todo/api/models/find_todo_response.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/api/models/update_todo_request.dart';
import 'package:todo/api/models/update_todo_response.dart';
import 'package:todo/api/todo/todo_client.dart';
import 'package:todo/features/todo/data/todo_exception.dart';
import 'package:todo/features/todo/data/todo_repository.dart';

class MockTodoClient extends Mock implements TodoClient {}

void main() {
  late MockTodoClient mockClient;
  late TodoRepository repository;

  setUpAll(() {
    registerFallbackValue(
      const UpdateTodoRequest(text: '', isComplete: false),
    );
    registerFallbackValue(
      const CreateTodoRequest(text: ''),
    );
  });

  setUp(() {
    mockClient = MockTodoClient();
    repository = TodoRepository(mockClient);
  });

  group('Test_TodoRepository_fetchTodos', () {
    test('shouldReturnTodoList_whenRequestSucceeds', () async {
      // given
      final now = DateTime.now();
      final todos = [
        FindTodoResponseTodo(id: 1, text: 'Buy milk', isComplete: false, createdAt: now, updatedAt: now),
        FindTodoResponseTodo(id: 2, text: 'Walk dog', isComplete: true, createdAt: now, updatedAt: now),
      ];
      when(() => mockClient.getTodos()).thenAnswer(
        (_) async => FindTodoResponse(todos: todos),
      );

      // when
      final result = await repository.fetchTodos();

      // then
      expect(result, todos);
    });

    test('shouldThrowTodoNetworkException_whenDioExceptionOccurs', () async {
      // given
      when(() => mockClient.getTodos()).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.fetchTodos(),
        throwsA(isA<TodoNetworkException>()),
      );
    });

    test('shouldThrowTodoNetworkException_whenNonDioExceptionOccurs', () async {
      // given
      when(() => mockClient.getTodos()).thenThrow(const FormatException('invalid json'));

      // when & then
      await expectLater(
        repository.fetchTodos(),
        throwsA(isA<TodoNetworkException>()),
      );
    });
  });

  group('Test_TodoRepository_createTodo', () {
    test('shouldComplete_whenRequestSucceeds', () async {
      // given
      final now = DateTime.now();
      when(
        () => mockClient.createTodo(body: any(named: 'body')),
      ).thenAnswer(
        (_) async => CreateTodoResponse(
          id: 3,
          text: 'New todo',
          isComplete: false,
          createdAt: now,
          updatedAt: now,
        ),
      );

      // when & then
      await expectLater(
        repository.createTodo(text: 'New todo'),
        completes,
      );
      verify(
        () => mockClient.createTodo(
          body: const CreateTodoRequest(text: 'New todo'),
        ),
      ).called(1);
    });

    test('shouldThrowTodoNetworkException_whenDioExceptionOccurs', () async {
      // given
      when(
        () => mockClient.createTodo(body: any(named: 'body')),
      ).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.createTodo(text: 'New todo'),
        throwsA(isA<TodoNetworkException>()),
      );
    });

    test('shouldThrowTodoNetworkException_whenNonDioExceptionOccurs', () async {
      // given
      when(
        () => mockClient.createTodo(body: any(named: 'body')),
      ).thenThrow(const FormatException('invalid json'));

      // when & then
      await expectLater(
        repository.createTodo(text: 'New todo'),
        throwsA(isA<TodoNetworkException>()),
      );
    });
  });

  group('Test_TodoRepository_updateTodo', () {
    test('shouldComplete_whenRequestSucceeds', () async {
      // given
      final now = DateTime.now();
      when(
        () => mockClient.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenAnswer(
        (_) async => UpdateTodoResponse(
          id: 1,
          text: 'Buy milk',
          isComplete: true,
          createdAt: now,
          updatedAt: now,
        ),
      );

      // when & then
      await expectLater(
        repository.updateTodo(
          id: 1,
          body: const UpdateTodoRequest(text: 'Buy milk', isComplete: true),
        ),
        completes,
      );
    });

    test('shouldThrowTodoNotFoundException_whenStatusIs404', () async {
      // given
      when(
        () => mockClient.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          response: Response(
            statusCode: 404,
            requestOptions: RequestOptions(),
          ),
          type: DioExceptionType.badResponse,
        ),
      );

      // when & then
      await expectLater(
        repository.updateTodo(
          id: 999,
          body: const UpdateTodoRequest(text: 'Buy milk', isComplete: true),
        ),
        throwsA(isA<TodoNotFoundException>()),
      );
    });

    test('shouldThrowTodoNetworkException_whenOtherDioExceptionOccurs', () async {
      // given
      when(
        () => mockClient.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenThrow(
        DioException(
          requestOptions: RequestOptions(),
          type: DioExceptionType.connectionTimeout,
        ),
      );

      // when & then
      await expectLater(
        repository.updateTodo(
          id: 1,
          body: const UpdateTodoRequest(text: 'Buy milk', isComplete: true),
        ),
        throwsA(isA<TodoNetworkException>()),
      );
    });
  });
}
