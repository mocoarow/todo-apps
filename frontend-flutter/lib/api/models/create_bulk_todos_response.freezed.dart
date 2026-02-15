// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'create_bulk_todos_response.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$CreateBulkTodosResponse {

 List<CreateTodoResponse> get todos;
/// Create a copy of CreateBulkTodosResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateBulkTodosResponseCopyWith<CreateBulkTodosResponse> get copyWith => _$CreateBulkTodosResponseCopyWithImpl<CreateBulkTodosResponse>(this as CreateBulkTodosResponse, _$identity);

  /// Serializes this CreateBulkTodosResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateBulkTodosResponse&&const DeepCollectionEquality().equals(other.todos, todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(todos));

@override
String toString() {
  return 'CreateBulkTodosResponse(todos: $todos)';
}


}

/// @nodoc
abstract mixin class $CreateBulkTodosResponseCopyWith<$Res>  {
  factory $CreateBulkTodosResponseCopyWith(CreateBulkTodosResponse value, $Res Function(CreateBulkTodosResponse) _then) = _$CreateBulkTodosResponseCopyWithImpl;
@useResult
$Res call({
 List<CreateTodoResponse> todos
});




}
/// @nodoc
class _$CreateBulkTodosResponseCopyWithImpl<$Res>
    implements $CreateBulkTodosResponseCopyWith<$Res> {
  _$CreateBulkTodosResponseCopyWithImpl(this._self, this._then);

  final CreateBulkTodosResponse _self;
  final $Res Function(CreateBulkTodosResponse) _then;

/// Create a copy of CreateBulkTodosResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? todos = null,}) {
  return _then(_self.copyWith(
todos: null == todos ? _self.todos : todos // ignore: cast_nullable_to_non_nullable
as List<CreateTodoResponse>,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateBulkTodosResponse].
extension CreateBulkTodosResponsePatterns on CreateBulkTodosResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateBulkTodosResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateBulkTodosResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateBulkTodosResponse value)  $default,){
final _that = this;
switch (_that) {
case _CreateBulkTodosResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateBulkTodosResponse value)?  $default,){
final _that = this;
switch (_that) {
case _CreateBulkTodosResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<CreateTodoResponse> todos)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateBulkTodosResponse() when $default != null:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<CreateTodoResponse> todos)  $default,) {final _that = this;
switch (_that) {
case _CreateBulkTodosResponse():
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<CreateTodoResponse> todos)?  $default,) {final _that = this;
switch (_that) {
case _CreateBulkTodosResponse() when $default != null:
return $default(_that.todos);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateBulkTodosResponse implements CreateBulkTodosResponse {
  const _CreateBulkTodosResponse({required final  List<CreateTodoResponse> todos}): _todos = todos;
  factory _CreateBulkTodosResponse.fromJson(Map<String, dynamic> json) => _$CreateBulkTodosResponseFromJson(json);

 final  List<CreateTodoResponse> _todos;
@override List<CreateTodoResponse> get todos {
  if (_todos is EqualUnmodifiableListView) return _todos;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_todos);
}


/// Create a copy of CreateBulkTodosResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateBulkTodosResponseCopyWith<_CreateBulkTodosResponse> get copyWith => __$CreateBulkTodosResponseCopyWithImpl<_CreateBulkTodosResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateBulkTodosResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateBulkTodosResponse&&const DeepCollectionEquality().equals(other._todos, _todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_todos));

@override
String toString() {
  return 'CreateBulkTodosResponse(todos: $todos)';
}


}

/// @nodoc
abstract mixin class _$CreateBulkTodosResponseCopyWith<$Res> implements $CreateBulkTodosResponseCopyWith<$Res> {
  factory _$CreateBulkTodosResponseCopyWith(_CreateBulkTodosResponse value, $Res Function(_CreateBulkTodosResponse) _then) = __$CreateBulkTodosResponseCopyWithImpl;
@override @useResult
$Res call({
 List<CreateTodoResponse> todos
});




}
/// @nodoc
class __$CreateBulkTodosResponseCopyWithImpl<$Res>
    implements _$CreateBulkTodosResponseCopyWith<$Res> {
  __$CreateBulkTodosResponseCopyWithImpl(this._self, this._then);

  final _CreateBulkTodosResponse _self;
  final $Res Function(_CreateBulkTodosResponse) _then;

/// Create a copy of CreateBulkTodosResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? todos = null,}) {
  return _then(_CreateBulkTodosResponse(
todos: null == todos ? _self._todos : todos // ignore: cast_nullable_to_non_nullable
as List<CreateTodoResponse>,
  ));
}


}

// dart format on
