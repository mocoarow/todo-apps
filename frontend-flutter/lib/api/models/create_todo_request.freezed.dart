// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'create_todo_request.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$CreateTodoRequest {

 String get text;
/// Create a copy of CreateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateTodoRequestCopyWith<CreateTodoRequest> get copyWith => _$CreateTodoRequestCopyWithImpl<CreateTodoRequest>(this as CreateTodoRequest, _$identity);

  /// Serializes this CreateTodoRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateTodoRequest&&(identical(other.text, text) || other.text == text));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,text);

@override
String toString() {
  return 'CreateTodoRequest(text: $text)';
}


}

/// @nodoc
abstract mixin class $CreateTodoRequestCopyWith<$Res>  {
  factory $CreateTodoRequestCopyWith(CreateTodoRequest value, $Res Function(CreateTodoRequest) _then) = _$CreateTodoRequestCopyWithImpl;
@useResult
$Res call({
 String text
});




}
/// @nodoc
class _$CreateTodoRequestCopyWithImpl<$Res>
    implements $CreateTodoRequestCopyWith<$Res> {
  _$CreateTodoRequestCopyWithImpl(this._self, this._then);

  final CreateTodoRequest _self;
  final $Res Function(CreateTodoRequest) _then;

/// Create a copy of CreateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? text = null,}) {
  return _then(_self.copyWith(
text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateTodoRequest].
extension CreateTodoRequestPatterns on CreateTodoRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateTodoRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateTodoRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateTodoRequest value)  $default,){
final _that = this;
switch (_that) {
case _CreateTodoRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateTodoRequest value)?  $default,){
final _that = this;
switch (_that) {
case _CreateTodoRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String text)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateTodoRequest() when $default != null:
return $default(_that.text);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String text)  $default,) {final _that = this;
switch (_that) {
case _CreateTodoRequest():
return $default(_that.text);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String text)?  $default,) {final _that = this;
switch (_that) {
case _CreateTodoRequest() when $default != null:
return $default(_that.text);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateTodoRequest implements CreateTodoRequest {
  const _CreateTodoRequest({required this.text});
  factory _CreateTodoRequest.fromJson(Map<String, dynamic> json) => _$CreateTodoRequestFromJson(json);

@override final  String text;

/// Create a copy of CreateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateTodoRequestCopyWith<_CreateTodoRequest> get copyWith => __$CreateTodoRequestCopyWithImpl<_CreateTodoRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateTodoRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateTodoRequest&&(identical(other.text, text) || other.text == text));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,text);

@override
String toString() {
  return 'CreateTodoRequest(text: $text)';
}


}

/// @nodoc
abstract mixin class _$CreateTodoRequestCopyWith<$Res> implements $CreateTodoRequestCopyWith<$Res> {
  factory _$CreateTodoRequestCopyWith(_CreateTodoRequest value, $Res Function(_CreateTodoRequest) _then) = __$CreateTodoRequestCopyWithImpl;
@override @useResult
$Res call({
 String text
});




}
/// @nodoc
class __$CreateTodoRequestCopyWithImpl<$Res>
    implements _$CreateTodoRequestCopyWith<$Res> {
  __$CreateTodoRequestCopyWithImpl(this._self, this._then);

  final _CreateTodoRequest _self;
  final $Res Function(_CreateTodoRequest) _then;

/// Create a copy of CreateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? text = null,}) {
  return _then(_CreateTodoRequest(
text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,
  ));
}


}

// dart format on
