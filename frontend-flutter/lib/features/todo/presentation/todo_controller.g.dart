// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'todo_controller.dart';

// **************************************************************************
// RiverpodGenerator
// **************************************************************************

// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, type=warning

@ProviderFor(TodoController)
final todoControllerProvider = TodoControllerProvider._();

final class TodoControllerProvider extends $AsyncNotifierProvider<TodoController, List<FindTodoResponseTodo>> {
  TodoControllerProvider._()
    : super(
        from: null,
        argument: null,
        retry: null,
        name: r'todoControllerProvider',
        isAutoDispose: true,
        dependencies: null,
        $allTransitiveDependencies: null,
      );

  @override
  String debugGetCreateSourceHash() => _$todoControllerHash();

  @$internal
  @override
  TodoController create() => TodoController();
}

String _$todoControllerHash() => r'e7667d8a1121a31763e3e2c92521bff3b7908df0';

abstract class _$TodoController extends $AsyncNotifier<List<FindTodoResponseTodo>> {
  FutureOr<List<FindTodoResponseTodo>> build();
  @$mustCallSuper
  @override
  void runBuild() {
    final ref = this.ref as $Ref<AsyncValue<List<FindTodoResponseTodo>>, List<FindTodoResponseTodo>>;
    final element =
        ref.element
            as $ClassProviderElement<
              AnyNotifier<AsyncValue<List<FindTodoResponseTodo>>, List<FindTodoResponseTodo>>,
              AsyncValue<List<FindTodoResponseTodo>>,
              Object?,
              Object?
            >;
    element.handleCreate(ref, build);
  }
}
