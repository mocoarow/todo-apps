sealed class AuthException implements Exception {
  const AuthException(this.message);
  final String message;

  @override
  String toString() => message;
}

class InvalidCredentialsException extends AuthException {
  const InvalidCredentialsException() : super('Invalid credentials');
}

class UnauthenticatedException extends AuthException {
  const UnauthenticatedException() : super('Not authenticated');
}

class AuthNetworkException extends AuthException {
  const AuthNetworkException() : super('Network error');
}
