import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/api/models/find_todo_response_todo.dart';
import 'package:todo/api/models/update_todo_request.dart';
import 'package:todo/api/todo/todo_client.dart';
import 'package:todo/core/network/dio_provider.dart';
import 'package:todo/features/todo/data/todo_exception.dart';

part 'todo_repository.g.dart';

@Riverpod(keepAlive: true)
TodoRepository todoRepository(Ref ref) {
  final dio = ref.watch(dioProvider);
  return TodoRepository(TodoClient(dio));
}

class TodoRepository {
  const TodoRepository(this._client);

  final TodoClient _client;

  Future<List<FindTodoResponseTodo>> fetchTodos() async {
    try {
      final response = await _client.getTodos();
      return response.todos;
    } on DioException catch (e) {
      debugPrint('TodoRepository.fetchTodos: DioException - ${e.type}: ${e.message}');
      _throwTodoException(e);
    } on Exception catch (e) {
      debugPrint('TodoRepository.fetchTodos: Unexpected exception - $e');
      throw TodoNetworkException(e);
    }
  }

  Future<void> updateTodo({required int id, required UpdateTodoRequest body}) async {
    try {
      await _client.updateTodo(id: id, body: body);
    } on DioException catch (e) {
      debugPrint('TodoRepository.updateTodo: DioException - ${e.type}: ${e.message}');
      _throwTodoException(e);
    } on Exception catch (e) {
      debugPrint('TodoRepository.updateTodo: Unexpected exception - $e');
      throw TodoNetworkException(e);
    }
  }

  Never _throwTodoException(DioException e) {
    if (e.response?.statusCode == 404) {
      throw TodoNotFoundException(e);
    }
    throw TodoNetworkException(e);
  }
}
