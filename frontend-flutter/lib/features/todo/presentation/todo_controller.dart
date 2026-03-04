import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/api/models/update_todo_request.dart';
import 'package:todo/features/todo/data/todo_repository.dart';

part 'todo_controller.g.dart';

@riverpod
class TodoController extends _$TodoController {
  @override
  Future<List<FindTodoResponseTodo>> build() async {
    return ref.watch(todoRepositoryProvider).fetchTodos();
  }

  Future<void> toggleComplete(FindTodoResponseTodo todo) async {
    final previousData = state.value;
    if (previousData != null) {
      state = AsyncData(
        previousData.map((t) {
          if (t.id == todo.id) {
            return t.copyWith(isComplete: !todo.isComplete);
          }
          return t;
        }).toList(),
      );
    }

    state = await AsyncValue.guard(() async {
      final repository = ref.read(todoRepositoryProvider);
      await repository.updateTodo(
        id: todo.id,
        body: UpdateTodoRequest(
          text: todo.text,
          isComplete: !todo.isComplete,
        ),
      );
      return repository.fetchTodos();
    });

    // Revert optimistic update on error so the list stays visible.
    // The error SnackBar is triggered by ref.listen in TodoScreen
    // during the intermediate AsyncError state.
    if (state.hasError && previousData != null) {
      state = AsyncData(previousData);
    }
  }
}
