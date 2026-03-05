import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
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
          return TodoListView(
            todos: todos,
            onEditTodo: (todo) => _showEditTodoDialog(context, ref, todo),
          );
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
      floatingActionButton: FloatingActionButton(
        onPressed: () => _showAddTodoDialog(context, ref),
        child: const Icon(Icons.add),
      ),
    );
  }

  Future<void> _showAddTodoDialog(BuildContext context, WidgetRef ref) async {
    final text = await showDialog<String>(
      context: context,
      builder: (context) => const _AddTodoDialog(),
    );
    if (text != null && text.isNotEmpty) {
      unawaited(ref.read(todoControllerProvider.notifier).addTodo(text));
    }
  }

  Future<void> _showEditTodoDialog(
    BuildContext context,
    WidgetRef ref,
    FindTodoResponseTodo todo,
  ) async {
    final newText = await showDialog<String>(
      context: context,
      builder: (context) => _EditTodoDialog(initialText: todo.text),
    );
    if (newText != null && newText != todo.text) {
      unawaited(
        ref.read(todoControllerProvider.notifier).updateTitle(todo, newText),
      );
    }
  }

  static String _resolveErrorMessage(Object? error) {
    return switch (error) {
      TodoNotFoundException() => 'Todo not found.',
      TodoNetworkException() => 'Connection failed. Please check your network.',
      _ => 'An unexpected error occurred. Please try again.',
    };
  }
}

class _AddTodoDialog extends StatefulWidget {
  const _AddTodoDialog();

  @override
  State<_AddTodoDialog> createState() => _AddTodoDialogState();
}

class _AddTodoDialogState extends State<_AddTodoDialog> {
  final _controller = TextEditingController();

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Add Todo'),
      content: TextField(
        controller: _controller,
        autofocus: true,
        decoration: const InputDecoration(hintText: 'What needs to be done?'),
        onChanged: (_) => setState(() {}),
        onSubmitted: (text) {
          if (text.trim().isNotEmpty) {
            Navigator.of(context).pop(text.trim());
          }
        },
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        TextButton(
          onPressed: _controller.text.trim().isEmpty ? null : () => Navigator.of(context).pop(_controller.text.trim()),
          child: const Text('Add'),
        ),
      ],
    );
  }
}

class _EditTodoDialog extends StatefulWidget {
  const _EditTodoDialog({required this.initialText});

  final String initialText;

  @override
  State<_EditTodoDialog> createState() => _EditTodoDialogState();
}

class _EditTodoDialogState extends State<_EditTodoDialog> {
  late final TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController(text: widget.initialText);
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  bool get _canSave {
    final trimmed = _controller.text.trim();
    return trimmed.isNotEmpty && trimmed != widget.initialText;
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Edit Todo'),
      content: TextField(
        controller: _controller,
        autofocus: true,
        maxLength: 250,
        decoration: const InputDecoration(hintText: 'Update your todo'),
        onChanged: (_) => setState(() {}),
        onSubmitted: (text) {
          if (_canSave) {
            Navigator.of(context).pop(text.trim());
          }
        },
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        TextButton(
          onPressed: _canSave ? () => Navigator.of(context).pop(_controller.text.trim()) : null,
          child: const Text('Save'),
        ),
      ],
    );
  }
}
