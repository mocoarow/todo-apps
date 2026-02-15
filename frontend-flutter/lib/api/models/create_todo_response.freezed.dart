// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'create_todo_response.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$CreateTodoResponse {

 int get id; String get text; bool get isComplete; DateTime get createdAt; DateTime get updatedAt;
/// Create a copy of CreateTodoResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateTodoResponseCopyWith<CreateTodoResponse> get copyWith => _$CreateTodoResponseCopyWithImpl<CreateTodoResponse>(this as CreateTodoResponse, _$identity);

  /// Serializes this CreateTodoResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateTodoResponse&&(identical(other.id, id) || other.id == id)&&(identical(other.text, text) || other.text == text)&&(identical(other.isComplete, isComplete) || other.isComplete == isComplete)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,text,isComplete,createdAt,updatedAt);

@override
String toString() {
  return 'CreateTodoResponse(id: $id, text: $text, isComplete: $isComplete, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $CreateTodoResponseCopyWith<$Res>  {
  factory $CreateTodoResponseCopyWith(CreateTodoResponse value, $Res Function(CreateTodoResponse) _then) = _$CreateTodoResponseCopyWithImpl;
@useResult
$Res call({
 int id, String text, bool isComplete, DateTime createdAt, DateTime updatedAt
});




}
/// @nodoc
class _$CreateTodoResponseCopyWithImpl<$Res>
    implements $CreateTodoResponseCopyWith<$Res> {
  _$CreateTodoResponseCopyWithImpl(this._self, this._then);

  final CreateTodoResponse _self;
  final $Res Function(CreateTodoResponse) _then;

/// Create a copy of CreateTodoResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? text = null,Object? isComplete = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as int,text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,isComplete: null == isComplete ? _self.isComplete : isComplete // ignore: cast_nullable_to_non_nullable
as bool,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateTodoResponse].
extension CreateTodoResponsePatterns on CreateTodoResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateTodoResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateTodoResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateTodoResponse value)  $default,){
final _that = this;
switch (_that) {
case _CreateTodoResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateTodoResponse value)?  $default,){
final _that = this;
switch (_that) {
case _CreateTodoResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( int id,  String text,  bool isComplete,  DateTime createdAt,  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateTodoResponse() when $default != null:
return $default(_that.id,_that.text,_that.isComplete,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( int id,  String text,  bool isComplete,  DateTime createdAt,  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _CreateTodoResponse():
return $default(_that.id,_that.text,_that.isComplete,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( int id,  String text,  bool isComplete,  DateTime createdAt,  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _CreateTodoResponse() when $default != null:
return $default(_that.id,_that.text,_that.isComplete,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateTodoResponse implements CreateTodoResponse {
  const _CreateTodoResponse({required this.id, required this.text, required this.isComplete, required this.createdAt, required this.updatedAt});
  factory _CreateTodoResponse.fromJson(Map<String, dynamic> json) => _$CreateTodoResponseFromJson(json);

@override final  int id;
@override final  String text;
@override final  bool isComplete;
@override final  DateTime createdAt;
@override final  DateTime updatedAt;

/// Create a copy of CreateTodoResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateTodoResponseCopyWith<_CreateTodoResponse> get copyWith => __$CreateTodoResponseCopyWithImpl<_CreateTodoResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateTodoResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateTodoResponse&&(identical(other.id, id) || other.id == id)&&(identical(other.text, text) || other.text == text)&&(identical(other.isComplete, isComplete) || other.isComplete == isComplete)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,text,isComplete,createdAt,updatedAt);

@override
String toString() {
  return 'CreateTodoResponse(id: $id, text: $text, isComplete: $isComplete, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$CreateTodoResponseCopyWith<$Res> implements $CreateTodoResponseCopyWith<$Res> {
  factory _$CreateTodoResponseCopyWith(_CreateTodoResponse value, $Res Function(_CreateTodoResponse) _then) = __$CreateTodoResponseCopyWithImpl;
@override @useResult
$Res call({
 int id, String text, bool isComplete, DateTime createdAt, DateTime updatedAt
});




}
/// @nodoc
class __$CreateTodoResponseCopyWithImpl<$Res>
    implements _$CreateTodoResponseCopyWith<$Res> {
  __$CreateTodoResponseCopyWithImpl(this._self, this._then);

  final _CreateTodoResponse _self;
  final $Res Function(_CreateTodoResponse) _then;

/// Create a copy of CreateTodoResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? text = null,Object? isComplete = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_CreateTodoResponse(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as int,text: null == text ? _self.text : text // ignore: cast_nullable_to_non_nullable
as String,isComplete: null == isComplete ? _self.isComplete : isComplete // ignore: cast_nullable_to_non_nullable
as bool,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}

// dart format on
