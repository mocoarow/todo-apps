// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'find_todo_response.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$FindTodoResponse {

 List<FindTodoResponseTodo> get todos;
/// Create a copy of FindTodoResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$FindTodoResponseCopyWith<FindTodoResponse> get copyWith => _$FindTodoResponseCopyWithImpl<FindTodoResponse>(this as FindTodoResponse, _$identity);

  /// Serializes this FindTodoResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is FindTodoResponse&&const DeepCollectionEquality().equals(other.todos, todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(todos));

@override
String toString() {
  return 'FindTodoResponse(todos: $todos)';
}


}

/// @nodoc
abstract mixin class $FindTodoResponseCopyWith<$Res>  {
  factory $FindTodoResponseCopyWith(FindTodoResponse value, $Res Function(FindTodoResponse) _then) = _$FindTodoResponseCopyWithImpl;
@useResult
$Res call({
 List<FindTodoResponseTodo> todos
});




}
/// @nodoc
class _$FindTodoResponseCopyWithImpl<$Res>
    implements $FindTodoResponseCopyWith<$Res> {
  _$FindTodoResponseCopyWithImpl(this._self, this._then);

  final FindTodoResponse _self;
  final $Res Function(FindTodoResponse) _then;

/// Create a copy of FindTodoResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? todos = null,}) {
  return _then(_self.copyWith(
todos: null == todos ? _self.todos : todos // ignore: cast_nullable_to_non_nullable
as List<FindTodoResponseTodo>,
  ));
}

}


/// Adds pattern-matching-related methods to [FindTodoResponse].
extension FindTodoResponsePatterns on FindTodoResponse {
/// A variant of `map` that fallback to returning `orElse`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _FindTodoResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _FindTodoResponse() when $default != null:
return $default(_that);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// Callbacks receives the raw object, upcasted.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case final Subclass2 value:
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _FindTodoResponse value)  $default,){
final _that = this;
switch (_that) {
case _FindTodoResponse():
return $default(_that);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `map` that fallback to returning `null`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _FindTodoResponse value)?  $default,){
final _that = this;
switch (_that) {
case _FindTodoResponse() when $default != null:
return $default(_that);case _:
  return null;

}
}
/// A variant of `when` that fallback to an `orElse` callback.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<FindTodoResponseTodo> todos)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _FindTodoResponse() when $default != null:
return $default(_that.todos);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// As opposed to `map`, this offers destructuring.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case Subclass2(:final field2):
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<FindTodoResponseTodo> todos)  $default,) {final _that = this;
switch (_that) {
case _FindTodoResponse():
return $default(_that.todos);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `when` that fallback to returning `null`
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<FindTodoResponseTodo> todos)?  $default,) {final _that = this;
switch (_that) {
case _FindTodoResponse() when $default != null:
return $default(_that.todos);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _FindTodoResponse implements FindTodoResponse {
  const _FindTodoResponse({required final  List<FindTodoResponseTodo> todos}): _todos = todos;
  factory _FindTodoResponse.fromJson(Map<String, dynamic> json) => _$FindTodoResponseFromJson(json);

 final  List<FindTodoResponseTodo> _todos;
@override List<FindTodoResponseTodo> get todos {
  if (_todos is EqualUnmodifiableListView) return _todos;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_todos);
}


/// Create a copy of FindTodoResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$FindTodoResponseCopyWith<_FindTodoResponse> get copyWith => __$FindTodoResponseCopyWithImpl<_FindTodoResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$FindTodoResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _FindTodoResponse&&const DeepCollectionEquality().equals(other._todos, _todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_todos));

@override
String toString() {
  return 'FindTodoResponse(todos: $todos)';
}


}

/// @nodoc
abstract mixin class _$FindTodoResponseCopyWith<$Res> implements $FindTodoResponseCopyWith<$Res> {
  factory _$FindTodoResponseCopyWith(_FindTodoResponse value, $Res Function(_FindTodoResponse) _then) = __$FindTodoResponseCopyWithImpl;
@override @useResult
$Res call({
 List<FindTodoResponseTodo> todos
});




}
/// @nodoc
class __$FindTodoResponseCopyWithImpl<$Res>
    implements _$FindTodoResponseCopyWith<$Res> {
  __$FindTodoResponseCopyWithImpl(this._self, this._then);

  final _FindTodoResponse _self;
  final $Res Function(_FindTodoResponse) _then;

/// Create a copy of FindTodoResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? todos = null,}) {
  return _then(_FindTodoResponse(
todos: null == todos ? _self._todos : todos // ignore: cast_nullable_to_non_nullable
as List<FindTodoResponseTodo>,
  ));
}


}

// dart format on
