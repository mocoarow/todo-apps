import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:todo/features/auth/presentation/auth_controller.dart';
import 'package:todo/features/todo/data/todo_exception.dart';
import 'package:todo/features/todo/presentation/todo_controller.dart';
import 'package:todo/features/todo/presentation/widgets/todo_list_view.dart';

class TodoScreen extends ConsumerWidget {
  const TodoScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    // Only show SnackBar for errors during operations (toggle, etc.),
    // not for initial load failures which are shown inline via error view.
    ref.listen(todoControllerProvider, (previous, next) {
      if (previous != null && previous.hasValue && next.hasError && !next.isLoading) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(_resolveErrorMessage(next.error)),
            backgroundColor: Theme.of(context).colorScheme.error,
          ),
        );
      }
    });

    final todosAsync = ref.watch(todoControllerProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Todos'),
        actions: [
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: () => ref.read(authControllerProvider.notifier).logout(),
          ),
        ],
      ),
      body: todosAsync.when(
        data: (todos) {
          if (todos.isEmpty) {
            return const Center(child: Text('No todos yet.'));
          }
          return TodoListView(todos: todos);
        },
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (error, _) => Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                _resolveErrorMessage(error),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: () => ref.invalidate(todoControllerProvider),
                child: const Text('Retry'),
              ),
            ],
          ),
        ),
      ),
    );
  }

  static String _resolveErrorMessage(Object? error) {
    return switch (error) {
      TodoNotFoundException() => 'Todo not found.',
      TodoNetworkException() => 'Connection failed. Please check your network.',
      _ => 'An unexpected error occurred. Please try again.',
    };
  }
}
