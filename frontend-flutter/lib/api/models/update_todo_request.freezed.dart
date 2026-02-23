// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'update_todo_request.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$UpdateTodoRequest {

 String get text; bool get isComplete;
/// Create a copy of UpdateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UpdateTodoRequestCopyWith<UpdateTodoRequest> get copyWith => _$UpdateTodoRequestCopyWithImpl<UpdateTodoRequest>(this as UpdateTodoRequest, _$identity);

  /// Serializes this UpdateTodoRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is UpdateTodoRequest&&(identical(other.text, text) || other.text == text)&&(identical(other.isComplete, isComplete) || other.isComplete == isComplete));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,text,isComplete);

@override
String toString() {
  return 'UpdateTodoRequest(text: $text, isComplete: $isComplete)';
}


}

/// @nodoc
abstract mixin class $UpdateTodoRequestCopyWith<$Res>  {
  factory $UpdateTodoRequestCopyWith(UpdateTodoRequest value, $Res Function(UpdateTodoRequest) _then) = _$UpdateTodoRequestCopyWithImpl;
@useResult
$Res call({
 String text, bool isComplete
});




}
/// @nodoc
class _$UpdateTodoRequestCopyWithImpl<$Res>
    implements $UpdateTodoRequestCopyWith<$Res> {
  _$UpdateTodoRequestCopyWithImpl(this._self, this._then);

  final UpdateTodoRequest _self;
  final $Res Function(UpdateTodoRequest) _then;

/// Create a copy of UpdateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? text = null,Object? isComplete = null,}) {
  return _then(_self.copyWith(
text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,isComplete: null == isComplete ? _self.isComplete : isComplete // ignore: cast_nullable_to_non_nullable
as bool,
  ));
}

}


/// Adds pattern-matching-related methods to [UpdateTodoRequest].
extension UpdateTodoRequestPatterns on UpdateTodoRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _UpdateTodoRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _UpdateTodoRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _UpdateTodoRequest value)  $default,){
final _that = this;
switch (_that) {
case _UpdateTodoRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _UpdateTodoRequest value)?  $default,){
final _that = this;
switch (_that) {
case _UpdateTodoRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String text,  bool isComplete)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _UpdateTodoRequest() when $default != null:
return $default(_that.text,_that.isComplete);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String text,  bool isComplete)  $default,) {final _that = this;
switch (_that) {
case _UpdateTodoRequest():
return $default(_that.text,_that.isComplete);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String text,  bool isComplete)?  $default,) {final _that = this;
switch (_that) {
case _UpdateTodoRequest() when $default != null:
return $default(_that.text,_that.isComplete);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _UpdateTodoRequest implements UpdateTodoRequest {
  const _UpdateTodoRequest({required this.text, required this.isComplete});
  factory _UpdateTodoRequest.fromJson(Map<String, dynamic> json) => _$UpdateTodoRequestFromJson(json);

@override final  String text;
@override final  bool isComplete;

/// Create a copy of UpdateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$UpdateTodoRequestCopyWith<_UpdateTodoRequest> get copyWith => __$UpdateTodoRequestCopyWithImpl<_UpdateTodoRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$UpdateTodoRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _UpdateTodoRequest&&(identical(other.text, text) || other.text == text)&&(identical(other.isComplete, isComplete) || other.isComplete == isComplete));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,text,isComplete);

@override
String toString() {
  return 'UpdateTodoRequest(text: $text, isComplete: $isComplete)';
}


}

/// @nodoc
abstract mixin class _$UpdateTodoRequestCopyWith<$Res> implements $UpdateTodoRequestCopyWith<$Res> {
  factory _$UpdateTodoRequestCopyWith(_UpdateTodoRequest value, $Res Function(_UpdateTodoRequest) _then) = __$UpdateTodoRequestCopyWithImpl;
@override @useResult
$Res call({
 String text, bool isComplete
});




}
/// @nodoc
class __$UpdateTodoRequestCopyWithImpl<$Res>
    implements _$UpdateTodoRequestCopyWith<$Res> {
  __$UpdateTodoRequestCopyWithImpl(this._self, this._then);

  final _UpdateTodoRequest _self;
  final $Res Function(_UpdateTodoRequest) _then;

/// Create a copy of UpdateTodoRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? text = null,Object? isComplete = null,}) {
  return _then(_UpdateTodoRequest(
text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,isComplete: null == isComplete ? _self.isComplete : isComplete // ignore: cast_nullable_to_non_nullable
as bool,
  ));
}


}

// dart format on
