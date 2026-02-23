abstract final class AppConstants {
  static const apiBaseUrl = String.fromEnvironment(
    'API_BASE_URL',
    defaultValue: 'http://localhost:8080',
  );

  static const appTitle = 'Todo App';

  // Validation limits (should match backend constraints)
  static const loginIdMaxLength = 100;
  static const passwordMinLength = 8;
  static const passwordMaxLength = 20;
}
