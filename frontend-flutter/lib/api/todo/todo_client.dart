// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';

import '../models/create_bulk_todos_request.dart';
import '../models/create_bulk_todos_response.dart';
import '../models/create_todo_request.dart';
import '../models/create_todo_response.dart';
import '../models/find_todo_response.dart';
import '../models/update_todo_request.dart';
import '../models/update_todo_response.dart';

part 'todo_client.g.dart';

@RestApi()
abstract class TodoClient {
  factory TodoClient(Dio dio, {String? baseUrl}) = _TodoClient;

  /// Get all todos.
  ///
  /// Get all todos for the authenticated user.
  @GET('/api/v1/todo')
  Future<FindTodoResponse> getTodos();

  /// Create a new todo.
  ///
  /// Create a new todo for the authenticated user.
  @POST('/api/v1/todo')
  Future<CreateTodoResponse> createTodo({
    @Body() required CreateTodoRequest body,
  });

  /// Create multiple todos.
  ///
  /// Create multiple todos for the authenticated user in a single request.
  @POST('/api/v1/todo/bulk')
  Future<CreateBulkTodosResponse> createBulkTodos({
    @Body() required CreateBulkTodosRequest body,
  });

  /// Update a todo.
  ///
  /// Update an existing todo for the authenticated user.
  ///
  /// [id] - Todo ID.
  @PUT('/api/v1/todo/{id}')
  Future<UpdateTodoResponse> updateTodo({
    @Path('id') required int id,
    @Body() required UpdateTodoRequest body,
  });

  /// Delete a todo.
  ///
  /// Delete an existing todo for the authenticated user.
  ///
  /// [id] - Todo ID.
  @DELETE('/api/v1/todo/{id}')
  Future<void> deleteTodo({
    @Path('id') required int id,
  });
}
