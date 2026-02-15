// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'create_bulk_todos_request.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$CreateBulkTodosRequest {

 List<CreateTodoRequest> get todos;
/// Create a copy of CreateBulkTodosRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateBulkTodosRequestCopyWith<CreateBulkTodosRequest> get copyWith => _$CreateBulkTodosRequestCopyWithImpl<CreateBulkTodosRequest>(this as CreateBulkTodosRequest, _$identity);

  /// Serializes this CreateBulkTodosRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateBulkTodosRequest&&const DeepCollectionEquality().equals(other.todos, todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(todos));

@override
String toString() {
  return 'CreateBulkTodosRequest(todos: $todos)';
}


}

/// @nodoc
abstract mixin class $CreateBulkTodosRequestCopyWith<$Res>  {
  factory $CreateBulkTodosRequestCopyWith(CreateBulkTodosRequest value, $Res Function(CreateBulkTodosRequest) _then) = _$CreateBulkTodosRequestCopyWithImpl;
@useResult
$Res call({
 List<CreateTodoRequest> todos
});




}
/// @nodoc
class _$CreateBulkTodosRequestCopyWithImpl<$Res>
    implements $CreateBulkTodosRequestCopyWith<$Res> {
  _$CreateBulkTodosRequestCopyWithImpl(this._self, this._then);

  final CreateBulkTodosRequest _self;
  final $Res Function(CreateBulkTodosRequest) _then;

/// Create a copy of CreateBulkTodosRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? todos = null,}) {
  return _then(_self.copyWith(
todos: null == todos ? _self.todos : todos // ignore: cast_nullable_to_non_nullable
as List<CreateTodoRequest>,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateBulkTodosRequest].
extension CreateBulkTodosRequestPatterns on CreateBulkTodosRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateBulkTodosRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateBulkTodosRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateBulkTodosRequest value)  $default,){
final _that = this;
switch (_that) {
case _CreateBulkTodosRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateBulkTodosRequest value)?  $default,){
final _that = this;
switch (_that) {
case _CreateBulkTodosRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<CreateTodoRequest> todos)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateBulkTodosRequest() when $default != null:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<CreateTodoRequest> todos)  $default,) {final _that = this;
switch (_that) {
case _CreateBulkTodosRequest():
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<CreateTodoRequest> todos)?  $default,) {final _that = this;
switch (_that) {
case _CreateBulkTodosRequest() when $default != null:
return $default(_that.todos);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateBulkTodosRequest implements CreateBulkTodosRequest {
  const _CreateBulkTodosRequest({required final  List<CreateTodoRequest> todos}): _todos = todos;
  factory _CreateBulkTodosRequest.fromJson(Map<String, dynamic> json) => _$CreateBulkTodosRequestFromJson(json);

 final  List<CreateTodoRequest> _todos;
@override List<CreateTodoRequest> get todos {
  if (_todos is EqualUnmodifiableListView) return _todos;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_todos);
}


/// Create a copy of CreateBulkTodosRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateBulkTodosRequestCopyWith<_CreateBulkTodosRequest> get copyWith => __$CreateBulkTodosRequestCopyWithImpl<_CreateBulkTodosRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateBulkTodosRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateBulkTodosRequest&&const DeepCollectionEquality().equals(other._todos, _todos));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_todos));

@override
String toString() {
  return 'CreateBulkTodosRequest(todos: $todos)';
}


}

/// @nodoc
abstract mixin class _$CreateBulkTodosRequestCopyWith<$Res> implements $CreateBulkTodosRequestCopyWith<$Res> {
  factory _$CreateBulkTodosRequestCopyWith(_CreateBulkTodosRequest value, $Res Function(_CreateBulkTodosRequest) _then) = __$CreateBulkTodosRequestCopyWithImpl;
@override @useResult
$Res call({
 List<CreateTodoRequest> todos
});




}
/// @nodoc
class __$CreateBulkTodosRequestCopyWithImpl<$Res>
    implements _$CreateBulkTodosRequestCopyWith<$Res> {
  __$CreateBulkTodosRequestCopyWithImpl(this._self, this._then);

  final _CreateBulkTodosRequest _self;
  final $Res Function(_CreateBulkTodosRequest) _then;

/// Create a copy of CreateBulkTodosRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? todos = null,}) {
  return _then(_CreateBulkTodosRequest(
todos: null == todos ? _self._todos : todos // ignore: cast_nullable_to_non_nullable
as List<CreateTodoRequest>,
  ));
}


}

// dart format on
