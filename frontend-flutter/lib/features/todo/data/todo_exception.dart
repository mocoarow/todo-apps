sealed class TodoException implements Exception {
  const TodoException(this.message, [this.cause]);
  final String message;
  final Object? cause;

  @override
  String toString() => cause != null ? '$message (cause: $cause)' : message;
}

class TodoNotFoundException extends TodoException {
  const TodoNotFoundException([Object? cause]) : super('Todo not found', cause);
}

class TodoNetworkException extends TodoException {
  const TodoNetworkException([Object? cause]) : super('Network error', cause);
}
