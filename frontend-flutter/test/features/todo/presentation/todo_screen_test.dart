import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/api/models/update_todo_request.dart';
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

  setUpAll(() {
    registerFallbackValue(
      const UpdateTodoRequest(text: '', isComplete: false),
    );
  });

  setUp(() {
    mockTodoRepository = MockTodoRepository();
    mockAuthRepository = MockAuthRepository();
    when(
      () => mockAuthRepository.getMe(),
    ).thenThrow(const UnauthenticatedException());
  });

  final todos = sampleTodos();

  group('Test_TodoScreen', () {
    testWidgets('shouldShowProgressIndicator_whenLoading', (tester) async {
      // given
      final completer = Completer<List<FindTodoResponseTodo>>();
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) => completer.future);

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
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

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

    testWidgets('shouldDisplayEmptyMessage_whenTodoListIsEmpty', (
      tester,
    ) async {
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
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenThrow(const TodoNetworkException());

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
      expect(
        find.text('Connection failed. Please check your network.'),
        findsOneWidget,
      );
      expect(find.text('Retry'), findsOneWidget);
    });

    testWidgets('shouldDisplayLogoutButton', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

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

    testWidgets('shouldDisplayAddButton', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

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
      expect(find.byType(FloatingActionButton), findsOneWidget);
      expect(find.byIcon(Icons.add), findsOneWidget);
    });

    testWidgets('shouldShowDialog_whenAddButtonTapped', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.byType(FloatingActionButton));
      await tester.pumpAndSettle();

      // then
      expect(find.text('Add Todo'), findsOneWidget);
      expect(find.byType(TextField), findsOneWidget);
      expect(find.text('Cancel'), findsOneWidget);
      expect(find.text('Add'), findsOneWidget);
    });

    testWidgets('shouldCallAddTodo_whenDialogSubmitted', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);
      when(
        () => mockTodoRepository.createTodo(text: any(named: 'text')),
      ).thenAnswer((_) async {});

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.byType(FloatingActionButton));
      await tester.pumpAndSettle();
      await tester.enterText(find.byType(TextField), 'New todo');
      await tester.pump();
      await tester.tap(find.text('Add'));
      await tester.pumpAndSettle();

      // then
      verify(() => mockTodoRepository.createTodo(text: 'New todo')).called(1);
    });

    testWidgets('shouldNotCallAddTodo_whenDialogCancelled', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.byType(FloatingActionButton));
      await tester.pumpAndSettle();
      await tester.enterText(find.byType(TextField), 'New todo');
      await tester.tap(find.text('Cancel'));
      await tester.pumpAndSettle();

      // then
      verifyNever(
        () => mockTodoRepository.createTodo(text: any(named: 'text')),
      );
    });

    testWidgets('shouldDisableAddButton_whenTextIsEmpty', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.byType(FloatingActionButton));
      await tester.pumpAndSettle();

      // then — Add button should be disabled when text is empty
      final addButton = tester.widget<TextButton>(
        find.widgetWithText(TextButton, 'Add'),
      );
      expect(addButton.onPressed, isNull);
    });

    testWidgets('shouldShowEditDialog_whenTodoTitleTapped', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.text('Buy milk'));
      await tester.pumpAndSettle();

      // then
      expect(find.text('Edit Todo'), findsOneWidget);
      expect(find.text('Save'), findsOneWidget);
      expect(find.text('Cancel'), findsOneWidget);
    });

    testWidgets('shouldCallUpdateTodo_whenEditDialogSaved', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);
      when(
        () => mockTodoRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenAnswer((_) async {});

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.text('Buy milk'));
      await tester.pumpAndSettle();
      // Clear existing text and enter new text
      await tester.enterText(find.byType(TextField), 'Buy eggs');
      await tester.pump();
      await tester.tap(find.text('Save'));
      await tester.pumpAndSettle();

      // then
      verify(
        () => mockTodoRepository.updateTodo(
          id: 1,
          body: const UpdateTodoRequest(text: 'Buy eggs', isComplete: false),
        ),
      ).called(1);
    });

    testWidgets('shouldNotCallUpdateTodo_whenEditDialogCancelled', (
      tester,
    ) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.text('Buy milk'));
      await tester.pumpAndSettle();
      await tester.enterText(find.byType(TextField), 'Buy eggs');
      await tester.tap(find.text('Cancel'));
      await tester.pumpAndSettle();

      // then
      verifyNever(
        () => mockTodoRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      );
    });

    testWidgets('shouldDisableSaveButton_whenTextIsUnchanged', (tester) async {
      // given
      when(
        () => mockTodoRepository.fetchTodos(),
      ).thenAnswer((_) async => todos);

      await tester.pumpApp(
        const TodoScreen(),
        overrides: [
          todoRepositoryProvider.overrideWithValue(mockTodoRepository),
          authRepositoryProvider.overrideWithValue(mockAuthRepository),
        ],
      );
      await tester.pumpAndSettle();

      // when
      await tester.tap(find.text('Buy milk'));
      await tester.pumpAndSettle();

      // then — Save button should be disabled when text is unchanged
      final saveButton = tester.widget<TextButton>(
        find.widgetWithText(TextButton, 'Save'),
      );
      expect(saveButton.onPressed, isNull);
    });
  });
}
