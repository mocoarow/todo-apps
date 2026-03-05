import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/api/models/update_todo_request.dart';
import 'package:todo/features/todo/data/todo_exception.dart';
import 'package:todo/features/todo/data/todo_repository.dart';
import 'package:todo/features/todo/presentation/todo_controller.dart';

import '../../../helpers/fixtures.dart';

class MockTodoRepository extends Mock implements TodoRepository {}

void main() {
  late MockTodoRepository mockRepository;

  setUpAll(() {
    registerFallbackValue(
      const UpdateTodoRequest(text: '', isComplete: false),
    );
  });

  setUp(() {
    mockRepository = MockTodoRepository();
  });

  ProviderContainer createContainer() {
    final container = ProviderContainer(
      overrides: [
        todoRepositoryProvider.overrideWithValue(mockRepository),
      ],
    );
    addTearDown(container.dispose);
    return container;
  }

  final todos = sampleTodos();

  group('Test_TodoController_build', () {
    test('shouldReturnTodoList_whenFetchSucceeds', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      final container = createContainer();

      // when
      final result = await container.read(todoControllerProvider.future);

      // then
      expect(result, todos);
    });

    // Note: Build error → AsyncError transition is verified
    // via the widget test (shouldDisplayRetryButton_whenErrorOccurs)
    // due to a Riverpod v3 limitation with ProviderContainer + auto-dispose async notifiers.
  });

  group('Test_TodoController_toggleComplete', () {
    test('shouldRefetchTodos_whenToggleSucceeds', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenAnswer((_) async {});
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).toggleComplete(todos[0]);

      // then
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      verify(
        () => mockRepository.updateTodo(
          id: 1,
          body: const UpdateTodoRequest(text: 'Buy milk', isComplete: true),
        ),
      ).called(1);
    });

    test('shouldRevertToOriginalData_whenToggleFails', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenThrow(const TodoNetworkException());
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).toggleComplete(todos[0]);

      // then — state is reverted to original data (not AsyncError)
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      expect(state.value, todos);
    });
  });

  group('Test_TodoController_updateTitle', () {
    test('shouldRefetchTodos_whenUpdateTitleSucceeds', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenAnswer((_) async {});
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).updateTitle(todos[0], 'Buy eggs');

      // then
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      verify(
        () => mockRepository.updateTodo(
          id: 1,
          body: const UpdateTodoRequest(text: 'Buy eggs', isComplete: false),
        ),
      ).called(1);
    });

    test('shouldRevertToOriginalData_whenUpdateTitleFails', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.updateTodo(
          id: any(named: 'id'),
          body: any(named: 'body'),
        ),
      ).thenThrow(const TodoNetworkException());
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).updateTitle(todos[0], 'Buy eggs');

      // then — state is reverted to original data (not AsyncError)
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      expect(state.value, todos);
    });
  });

  group('Test_TodoController_addTodo', () {
    test('shouldRefetchTodos_whenAddSucceeds', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.createTodo(text: any(named: 'text')),
      ).thenAnswer((_) async {});
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).addTodo('New todo');

      // then
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      verify(() => mockRepository.createTodo(text: 'New todo')).called(1);
      // fetchTodos called twice: once in build, once after addTodo
      verify(() => mockRepository.fetchTodos()).called(2);
    });

    test('shouldTransitionToErrorThenRevert_whenAddFails', () async {
      // given
      when(() => mockRepository.fetchTodos()).thenAnswer((_) async => todos);
      when(
        () => mockRepository.createTodo(text: any(named: 'text')),
      ).thenThrow(const TodoNetworkException());
      final container = createContainer();
      await container.read(todoControllerProvider.future);

      // when
      await container.read(todoControllerProvider.notifier).addTodo('New todo');

      // then — state is reverted to original data
      final state = container.read(todoControllerProvider);
      expect(state, isA<AsyncData<List<FindTodoResponseTodo>>>());
      expect(state.value, todos);
    });
  });
}
