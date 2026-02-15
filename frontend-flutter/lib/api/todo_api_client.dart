// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint, unused_import, invalid_annotation_target, unnecessary_import

import 'package:dio/dio.dart';

import 'auth/auth_client.dart';
import 'todo/todo_client.dart';

/// Todo App API `v1.0`.
///
/// A Todo application API with authentication.
class TodoApiClient {
  TodoApiClient(
    Dio dio, {
    String? baseUrl,
  })  : _dio = dio,
        _baseUrl = baseUrl;

  final Dio _dio;
  final String? _baseUrl;

  static String get version => '1.0';

  AuthClient? _auth;
  TodoClient? _todo;

  AuthClient get auth => _auth ??= AuthClient(_dio, baseUrl: _baseUrl);

  TodoClient get todo => _todo ??= TodoClient(_dio, baseUrl: _baseUrl);
}
