// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'authenticate_request.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$AuthenticateRequest {

 String get loginId; String get password;
/// Create a copy of AuthenticateRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$AuthenticateRequestCopyWith<AuthenticateRequest> get copyWith => _$AuthenticateRequestCopyWithImpl<AuthenticateRequest>(this as AuthenticateRequest, _$identity);

  /// Serializes this AuthenticateRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is AuthenticateRequest&&(identical(other.loginId, loginId) || other.loginId == loginId)&&(identical(other.password, password) || other.password == password));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,loginId,password);

@override
String toString() {
  return 'AuthenticateRequest(loginId: $loginId, password: $password)';
}


}

/// @nodoc
abstract mixin class $AuthenticateRequestCopyWith<$Res>  {
  factory $AuthenticateRequestCopyWith(AuthenticateRequest value, $Res Function(AuthenticateRequest) _then) = _$AuthenticateRequestCopyWithImpl;
@useResult
$Res call({
 String loginId, String password
});




}
/// @nodoc
class _$AuthenticateRequestCopyWithImpl<$Res>
    implements $AuthenticateRequestCopyWith<$Res> {
  _$AuthenticateRequestCopyWithImpl(this._self, this._then);

  final AuthenticateRequest _self;
  final $Res Function(AuthenticateRequest) _then;

/// Create a copy of AuthenticateRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? loginId = null,Object? password = null,}) {
  return _then(_self.copyWith(
loginId: null == loginId ? _self.loginId : loginId // ignore: cast_nullable_to_non_nullable
as String,password: null == password ? _self.password : password // ignore: cast_nullable_to_non_nullable
as String,
  ));
}

}


/// Adds pattern-matching-related methods to [AuthenticateRequest].
extension AuthenticateRequestPatterns on AuthenticateRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _AuthenticateRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _AuthenticateRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _AuthenticateRequest value)  $default,){
final _that = this;
switch (_that) {
case _AuthenticateRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _AuthenticateRequest value)?  $default,){
final _that = this;
switch (_that) {
case _AuthenticateRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String loginId,  String password)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _AuthenticateRequest() when $default != null:
return $default(_that.loginId,_that.password);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String loginId,  String password)  $default,) {final _that = this;
switch (_that) {
case _AuthenticateRequest():
return $default(_that.loginId,_that.password);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String loginId,  String password)?  $default,) {final _that = this;
switch (_that) {
case _AuthenticateRequest() when $default != null:
return $default(_that.loginId,_that.password);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _AuthenticateRequest implements AuthenticateRequest {
  const _AuthenticateRequest({required this.loginId, required this.password});
  factory _AuthenticateRequest.fromJson(Map<String, dynamic> json) => _$AuthenticateRequestFromJson(json);

@override final  String loginId;
@override final  String password;

/// Create a copy of AuthenticateRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$AuthenticateRequestCopyWith<_AuthenticateRequest> get copyWith => __$AuthenticateRequestCopyWithImpl<_AuthenticateRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$AuthenticateRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _AuthenticateRequest&&(identical(other.loginId, loginId) || other.loginId == loginId)&&(identical(other.password, password) || other.password == password));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,loginId,password);

@override
String toString() {
  return 'AuthenticateRequest(loginId: $loginId, password: $password)';
}


}

/// @nodoc
abstract mixin class _$AuthenticateRequestCopyWith<$Res> implements $AuthenticateRequestCopyWith<$Res> {
  factory _$AuthenticateRequestCopyWith(_AuthenticateRequest value, $Res Function(_AuthenticateRequest) _then) = __$AuthenticateRequestCopyWithImpl;
@override @useResult
$Res call({
 String loginId, String password
});




}
/// @nodoc
class __$AuthenticateRequestCopyWithImpl<$Res>
    implements _$AuthenticateRequestCopyWith<$Res> {
  __$AuthenticateRequestCopyWithImpl(this._self, this._then);

  final _AuthenticateRequest _self;
  final $Res Function(_AuthenticateRequest) _then;

/// Create a copy of AuthenticateRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? loginId = null,Object? password = null,}) {
  return _then(_AuthenticateRequest(
loginId: null == loginId ? _self.loginId : loginId // ignore: cast_nullable_to_non_nullable
as String,password: null == password ? _self.password : password // ignore: cast_nullable_to_non_nullable
as String,
  ));
}


}

// dart format on
