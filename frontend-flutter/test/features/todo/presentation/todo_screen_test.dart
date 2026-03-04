import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/features/auth/data/auth_exception.dart';
import 'package:todo/features/auth/data/auth_repository.dart';
import 'package:todo/features/todo/data/todo_exception.dart';
import 'package:todo/features/todo/data/todo_repository.dart';
import 'package:todo/features/todo/presentation/todo_screen.dart';

import '../../../helpers/fixtures.dart';
import '../../../helpers/test_helpers.dart';

class MockTodoRepository extends Mock implements TodoRepository {}

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockTodoRepository mockTodoRepository;
  late MockAuthRepository mockAuthRepository;

  setUp(() {
    mockTodoRepository = MockTodoRepository();
    mockAuthRepository = MockAuthRepository();
    when(() => mockAuthRepository.getMe()).thenThrow(const UnauthenticatedException());
  });

  final todos = sampleTodos();

  group('Test_TodoScreen', () {
    testWidgets('shouldShowProgressIndicator_whenLoading', (tester) async {
      // given
      final completer = Completer<List<FindTodoResponseTodo>>();
      when(() => mockTodoRepository.fetchTodos()).thenAnswer((_) => completer.future);

      // when
      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pump();

      // then
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });

    testWidgets('shouldDisplayTodos_whenDataIsLoaded', (tester) async {
      // given
      when(() => mockTodoRepository.fetchTodos()).thenAnswer((_) async => todos);

      // when
      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.text('Buy milk'), findsOneWidget);
      expect(find.text('Walk dog'), findsOneWidget);
    });

    testWidgets('shouldDisplayEmptyMessage_whenTodoListIsEmpty', (tester) async {
      // given
      when(() => mockTodoRepository.fetchTodos()).thenAnswer((_) async => []);

      // when
      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.text('No todos yet.'), findsOneWidget);
    });

    testWidgets('shouldDisplayRetryButton_whenErrorOccurs', (tester) async {
      // given
      when(() => mockTodoRepository.fetchTodos()).thenThrow(const TodoNetworkException());

      // when
      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.text('Connection failed. Please check your network.'), findsOneWidget);
      expect(find.text('Retry'), findsOneWidget);
    });

    testWidgets('shouldDisplayLogoutButton', (tester) async {
      // given
      when(() => mockTodoRepository.fetchTodos()).thenAnswer((_) async => todos);

      // when
      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // then
      expect(find.byIcon(Icons.logout), findsOneWidget);
    });
  });
}
