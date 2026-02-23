// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'get_me_response.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$GetMeResponse {

 int get userId; String get loginId;
/// Create a copy of GetMeResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$GetMeResponseCopyWith<GetMeResponse> get copyWith => _$GetMeResponseCopyWithImpl<GetMeResponse>(this as GetMeResponse, _$identity);

  /// Serializes this GetMeResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is GetMeResponse&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.loginId, loginId) || other.loginId == loginId));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,userId,loginId);

@override
String toString() {
  return 'GetMeResponse(userId: $userId, loginId: $loginId)';
}


}

/// @nodoc
abstract mixin class $GetMeResponseCopyWith<$Res>  {
  factory $GetMeResponseCopyWith(GetMeResponse value, $Res Function(GetMeResponse) _then) = _$GetMeResponseCopyWithImpl;
@useResult
$Res call({
 int userId, String loginId
});




}
/// @nodoc
class _$GetMeResponseCopyWithImpl<$Res>
    implements $GetMeResponseCopyWith<$Res> {
  _$GetMeResponseCopyWithImpl(this._self, this._then);

  final GetMeResponse _self;
  final $Res Function(GetMeResponse) _then;

/// Create a copy of GetMeResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? userId = null,Object? loginId = null,}) {
  return _then(_self.copyWith(
userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as int,loginId: null == loginId ? _self.loginId : loginId // ignore: cast_nullable_to_non_nullable
as String,
  ));
}

}


/// Adds pattern-matching-related methods to [GetMeResponse].
extension GetMeResponsePatterns on GetMeResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _GetMeResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _GetMeResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _GetMeResponse value)  $default,){
final _that = this;
switch (_that) {
case _GetMeResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _GetMeResponse value)?  $default,){
final _that = this;
switch (_that) {
case _GetMeResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( int userId,  String loginId)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _GetMeResponse() when $default != null:
return $default(_that.userId,_that.loginId);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( int userId,  String loginId)  $default,) {final _that = this;
switch (_that) {
case _GetMeResponse():
return $default(_that.userId,_that.loginId);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( int userId,  String loginId)?  $default,) {final _that = this;
switch (_that) {
case _GetMeResponse() when $default != null:
return $default(_that.userId,_that.loginId);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _GetMeResponse implements GetMeResponse {
  const _GetMeResponse({required this.userId, required this.loginId});
  factory _GetMeResponse.fromJson(Map<String, dynamic> json) => _$GetMeResponseFromJson(json);

@override final  int userId;
@override final  String loginId;

/// Create a copy of GetMeResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$GetMeResponseCopyWith<_GetMeResponse> get copyWith => __$GetMeResponseCopyWithImpl<_GetMeResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$GetMeResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _GetMeResponse&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.loginId, loginId) || other.loginId == loginId));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,userId,loginId);

@override
String toString() {
  return 'GetMeResponse(userId: $userId, loginId: $loginId)';
}


}

/// @nodoc
abstract mixin class _$GetMeResponseCopyWith<$Res> implements $GetMeResponseCopyWith<$Res> {
  factory _$GetMeResponseCopyWith(_GetMeResponse value, $Res Function(_GetMeResponse) _then) = __$GetMeResponseCopyWithImpl;
@override @useResult
$Res call({
 int userId, String loginId
});




}
/// @nodoc
class __$GetMeResponseCopyWithImpl<$Res>
    implements _$GetMeResponseCopyWith<$Res> {
  __$GetMeResponseCopyWithImpl(this._self, this._then);

  final _GetMeResponse _self;
  final $Res Function(_GetMeResponse) _then;

/// Create a copy of GetMeResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? userId = null,Object? loginId = null,}) {
  return _then(_GetMeResponse(
userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as int,loginId: null == loginId ? _self.loginId : loginId // ignore: cast_nullable_to_non_nullable
as String,
  ));
}


}

// dart format on
