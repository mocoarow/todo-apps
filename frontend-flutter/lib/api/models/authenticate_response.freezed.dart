// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'authenticate_response.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$AuthenticateResponse {

/// JWT access token (omitted when delivered via cookie)
 String? get accessToken;
/// Create a copy of AuthenticateResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$AuthenticateResponseCopyWith<AuthenticateResponse> get copyWith => _$AuthenticateResponseCopyWithImpl<AuthenticateResponse>(this as AuthenticateResponse, _$identity);

  /// Serializes this AuthenticateResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is AuthenticateResponse&&(identical(other.accessToken, accessToken) || other.accessToken == accessToken));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,accessToken);

@override
String toString() {
  return 'AuthenticateResponse(accessToken: $accessToken)';
}


}

/// @nodoc
abstract mixin class $AuthenticateResponseCopyWith<$Res>  {
  factory $AuthenticateResponseCopyWith(AuthenticateResponse value, $Res Function(AuthenticateResponse) _then) = _$AuthenticateResponseCopyWithImpl;
@useResult
$Res call({
 String? accessToken
});




}
/// @nodoc
class _$AuthenticateResponseCopyWithImpl<$Res>
    implements $AuthenticateResponseCopyWith<$Res> {
  _$AuthenticateResponseCopyWithImpl(this._self, this._then);

  final AuthenticateResponse _self;
  final $Res Function(AuthenticateResponse) _then;

/// Create a copy of AuthenticateResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? accessToken = freezed,}) {
  return _then(_self.copyWith(
accessToken: freezed == accessToken ? _self.accessToken : accessToken // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}

}


/// Adds pattern-matching-related methods to [AuthenticateResponse].
extension AuthenticateResponsePatterns on AuthenticateResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _AuthenticateResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _AuthenticateResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _AuthenticateResponse value)  $default,){
final _that = this;
switch (_that) {
case _AuthenticateResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _AuthenticateResponse value)?  $default,){
final _that = this;
switch (_that) {
case _AuthenticateResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String? accessToken)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _AuthenticateResponse() when $default != null:
return $default(_that.accessToken);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String? accessToken)  $default,) {final _that = this;
switch (_that) {
case _AuthenticateResponse():
return $default(_that.accessToken);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String? accessToken)?  $default,) {final _that = this;
switch (_that) {
case _AuthenticateResponse() when $default != null:
return $default(_that.accessToken);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _AuthenticateResponse implements AuthenticateResponse {
  const _AuthenticateResponse({this.accessToken});
  factory _AuthenticateResponse.fromJson(Map<String, dynamic> json) => _$AuthenticateResponseFromJson(json);

/// JWT access token (omitted when delivered via cookie)
@override final  String? accessToken;

/// Create a copy of AuthenticateResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$AuthenticateResponseCopyWith<_AuthenticateResponse> get copyWith => __$AuthenticateResponseCopyWithImpl<_AuthenticateResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$AuthenticateResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _AuthenticateResponse&&(identical(other.accessToken, accessToken) || other.accessToken == accessToken));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,accessToken);

@override
String toString() {
  return 'AuthenticateResponse(accessToken: $accessToken)';
}


}

/// @nodoc
abstract mixin class _$AuthenticateResponseCopyWith<$Res> implements $AuthenticateResponseCopyWith<$Res> {
  factory _$AuthenticateResponseCopyWith(_AuthenticateResponse value, $Res Function(_AuthenticateResponse) _then) = __$AuthenticateResponseCopyWithImpl;
@override @useResult
$Res call({
 String? accessToken
});




}
/// @nodoc
class __$AuthenticateResponseCopyWithImpl<$Res>
    implements _$AuthenticateResponseCopyWith<$Res> {
  __$AuthenticateResponseCopyWithImpl(this._self, this._then);

  final _AuthenticateResponse _self;
  final $Res Function(_AuthenticateResponse) _then;

/// Create a copy of AuthenticateResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? accessToken = freezed,}) {
  return _then(_AuthenticateResponse(
accessToken: freezed == accessToken ? _self.accessToken : accessToken // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}


}

// dart format on
