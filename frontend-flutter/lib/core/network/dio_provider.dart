import 'package:cookie_jar/cookie_jar.dart';
import 'package:dio/dio.dart';
import 'package:dio_cookie_manager/dio_cookie_manager.dart';
import 'package:flutter/foundation.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:todo/core/constants/app_constants.dart';

part 'dio_provider.g.dart';

@Riverpod(keepAlive: true)
Dio dio(Ref ref) {
  final dio = Dio(
    BaseOptions(
      baseUrl: AppConstants.apiBaseUrl,
      connectTimeout: const Duration(seconds: 10),
      receiveTimeout: const Duration(seconds: 10),
      extra: {
        if (kIsWeb) 'withCredentials': true,
      },
    ),
  );

  if (!kIsWeb) {
    // In-memory cookie store. Sessions are lost on app restart.
    // Upgrade to PersistCookieJar when persistent sessions are needed.
    dio.interceptors.add(CookieManager(CookieJar()));
  }

  return dio;
}
