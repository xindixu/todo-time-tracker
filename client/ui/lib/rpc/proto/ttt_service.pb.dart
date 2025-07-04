// This is a generated file - do not edit.
//
// Generated from proto/ttt_service.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/duration.pb.dart' as $3;
import 'context.pb.dart' as $1;
import 'model.pb.dart' as $2;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

/// -- User operations --
class CreateUserReq extends $pb.GeneratedMessage {
  factory CreateUserReq({
    $1.Context? context,
    $core.String? name,
    $core.String? email,
    $core.String? password,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (name != null) result.name = name;
    if (email != null) result.email = email;
    if (password != null) result.password = password;
    return result;
  }

  CreateUserReq._();

  factory CreateUserReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateUserReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateUserReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'name')
    ..aOS(3, _omitFieldNames ? '' : 'email')
    ..aOS(4, _omitFieldNames ? '' : 'password')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateUserReq clone() => CreateUserReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateUserReq copyWith(void Function(CreateUserReq) updates) => super.copyWith((message) => updates(message as CreateUserReq)) as CreateUserReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateUserReq create() => CreateUserReq._();
  @$core.override
  CreateUserReq createEmptyInstance() => create();
  static $pb.PbList<CreateUserReq> createRepeated() => $pb.PbList<CreateUserReq>();
  @$core.pragma('dart2js:noInline')
  static CreateUserReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateUserReq>(create);
  static CreateUserReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get email => $_getSZ(2);
  @$pb.TagNumber(3)
  set email($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasEmail() => $_has(2);
  @$pb.TagNumber(3)
  void clearEmail() => $_clearField(3);

  @$pb.TagNumber(4)
  $core.String get password => $_getSZ(3);
  @$pb.TagNumber(4)
  set password($core.String value) => $_setString(3, value);
  @$pb.TagNumber(4)
  $core.bool hasPassword() => $_has(3);
  @$pb.TagNumber(4)
  void clearPassword() => $_clearField(4);
}

class CreateUserResp extends $pb.GeneratedMessage {
  factory CreateUserResp({
    $2.User? user,
  }) {
    final result = create();
    if (user != null) result.user = user;
    return result;
  }

  CreateUserResp._();

  factory CreateUserResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateUserResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateUserResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.User>(1, _omitFieldNames ? '' : 'user', subBuilder: $2.User.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateUserResp clone() => CreateUserResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateUserResp copyWith(void Function(CreateUserResp) updates) => super.copyWith((message) => updates(message as CreateUserResp)) as CreateUserResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateUserResp create() => CreateUserResp._();
  @$core.override
  CreateUserResp createEmptyInstance() => create();
  static $pb.PbList<CreateUserResp> createRepeated() => $pb.PbList<CreateUserResp>();
  @$core.pragma('dart2js:noInline')
  static CreateUserResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateUserResp>(create);
  static CreateUserResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.User get user => $_getN(0);
  @$pb.TagNumber(1)
  set user($2.User value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearUser() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.User ensureUser() => $_ensure(0);
}

class GetTagReq extends $pb.GeneratedMessage {
  factory GetTagReq({
    $1.Context? context,
    $core.String? uuid,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (uuid != null) result.uuid = uuid;
    return result;
  }

  GetTagReq._();

  factory GetTagReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTagReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTagReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'uuid')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTagReq clone() => GetTagReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTagReq copyWith(void Function(GetTagReq) updates) => super.copyWith((message) => updates(message as GetTagReq)) as GetTagReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTagReq create() => GetTagReq._();
  @$core.override
  GetTagReq createEmptyInstance() => create();
  static $pb.PbList<GetTagReq> createRepeated() => $pb.PbList<GetTagReq>();
  @$core.pragma('dart2js:noInline')
  static GetTagReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTagReq>(create);
  static GetTagReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get uuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUuid() => $_clearField(2);
}

class GetTagResp extends $pb.GeneratedMessage {
  factory GetTagResp({
    $2.Tag? tag,
  }) {
    final result = create();
    if (tag != null) result.tag = tag;
    return result;
  }

  GetTagResp._();

  factory GetTagResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTagResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTagResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.Tag>(1, _omitFieldNames ? '' : 'tag', subBuilder: $2.Tag.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTagResp clone() => GetTagResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTagResp copyWith(void Function(GetTagResp) updates) => super.copyWith((message) => updates(message as GetTagResp)) as GetTagResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTagResp create() => GetTagResp._();
  @$core.override
  GetTagResp createEmptyInstance() => create();
  static $pb.PbList<GetTagResp> createRepeated() => $pb.PbList<GetTagResp>();
  @$core.pragma('dart2js:noInline')
  static GetTagResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTagResp>(create);
  static GetTagResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Tag get tag => $_getN(0);
  @$pb.TagNumber(1)
  set tag($2.Tag value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasTag() => $_has(0);
  @$pb.TagNumber(1)
  void clearTag() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.Tag ensureTag() => $_ensure(0);
}

class ListTagsReq extends $pb.GeneratedMessage {
  factory ListTagsReq({
    $1.Context? context,
  }) {
    final result = create();
    if (context != null) result.context = context;
    return result;
  }

  ListTagsReq._();

  factory ListTagsReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory ListTagsReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTagsReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTagsReq clone() => ListTagsReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTagsReq copyWith(void Function(ListTagsReq) updates) => super.copyWith((message) => updates(message as ListTagsReq)) as ListTagsReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTagsReq create() => ListTagsReq._();
  @$core.override
  ListTagsReq createEmptyInstance() => create();
  static $pb.PbList<ListTagsReq> createRepeated() => $pb.PbList<ListTagsReq>();
  @$core.pragma('dart2js:noInline')
  static ListTagsReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTagsReq>(create);
  static ListTagsReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);
}

class ListTagsResp extends $pb.GeneratedMessage {
  factory ListTagsResp({
    $core.Iterable<$2.Tag>? tags,
  }) {
    final result = create();
    if (tags != null) result.tags.addAll(tags);
    return result;
  }

  ListTagsResp._();

  factory ListTagsResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory ListTagsResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTagsResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..pc<$2.Tag>(1, _omitFieldNames ? '' : 'tags', $pb.PbFieldType.PM, subBuilder: $2.Tag.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTagsResp clone() => ListTagsResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTagsResp copyWith(void Function(ListTagsResp) updates) => super.copyWith((message) => updates(message as ListTagsResp)) as ListTagsResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTagsResp create() => ListTagsResp._();
  @$core.override
  ListTagsResp createEmptyInstance() => create();
  static $pb.PbList<ListTagsResp> createRepeated() => $pb.PbList<ListTagsResp>();
  @$core.pragma('dart2js:noInline')
  static ListTagsResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTagsResp>(create);
  static ListTagsResp? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$2.Tag> get tags => $_getList(0);
}

class CreateTagReq extends $pb.GeneratedMessage {
  factory CreateTagReq({
    $1.Context? context,
    $core.String? name,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (name != null) result.name = name;
    return result;
  }

  CreateTagReq._();

  factory CreateTagReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTagReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTagReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTagReq clone() => CreateTagReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTagReq copyWith(void Function(CreateTagReq) updates) => super.copyWith((message) => updates(message as CreateTagReq)) as CreateTagReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTagReq create() => CreateTagReq._();
  @$core.override
  CreateTagReq createEmptyInstance() => create();
  static $pb.PbList<CreateTagReq> createRepeated() => $pb.PbList<CreateTagReq>();
  @$core.pragma('dart2js:noInline')
  static CreateTagReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTagReq>(create);
  static CreateTagReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => $_clearField(2);
}

class CreateTagResp extends $pb.GeneratedMessage {
  factory CreateTagResp({
    $2.Tag? tag,
  }) {
    final result = create();
    if (tag != null) result.tag = tag;
    return result;
  }

  CreateTagResp._();

  factory CreateTagResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTagResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTagResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.Tag>(1, _omitFieldNames ? '' : 'tag', subBuilder: $2.Tag.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTagResp clone() => CreateTagResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTagResp copyWith(void Function(CreateTagResp) updates) => super.copyWith((message) => updates(message as CreateTagResp)) as CreateTagResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTagResp create() => CreateTagResp._();
  @$core.override
  CreateTagResp createEmptyInstance() => create();
  static $pb.PbList<CreateTagResp> createRepeated() => $pb.PbList<CreateTagResp>();
  @$core.pragma('dart2js:noInline')
  static CreateTagResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTagResp>(create);
  static CreateTagResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Tag get tag => $_getN(0);
  @$pb.TagNumber(1)
  set tag($2.Tag value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasTag() => $_has(0);
  @$pb.TagNumber(1)
  void clearTag() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.Tag ensureTag() => $_ensure(0);
}

class UpdateTagReq extends $pb.GeneratedMessage {
  factory UpdateTagReq({
    $1.Context? context,
    $core.String? uuid,
    $core.String? name,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (uuid != null) result.uuid = uuid;
    if (name != null) result.name = name;
    return result;
  }

  UpdateTagReq._();

  factory UpdateTagReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory UpdateTagReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UpdateTagReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'uuid')
    ..aOS(3, _omitFieldNames ? '' : 'name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UpdateTagReq clone() => UpdateTagReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UpdateTagReq copyWith(void Function(UpdateTagReq) updates) => super.copyWith((message) => updates(message as UpdateTagReq)) as UpdateTagReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UpdateTagReq create() => UpdateTagReq._();
  @$core.override
  UpdateTagReq createEmptyInstance() => create();
  static $pb.PbList<UpdateTagReq> createRepeated() => $pb.PbList<UpdateTagReq>();
  @$core.pragma('dart2js:noInline')
  static UpdateTagReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateTagReq>(create);
  static UpdateTagReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get uuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUuid() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => $_clearField(3);
}

class UpdateTagResp extends $pb.GeneratedMessage {
  factory UpdateTagResp({
    $2.Tag? tag,
  }) {
    final result = create();
    if (tag != null) result.tag = tag;
    return result;
  }

  UpdateTagResp._();

  factory UpdateTagResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory UpdateTagResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UpdateTagResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.Tag>(1, _omitFieldNames ? '' : 'tag', subBuilder: $2.Tag.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UpdateTagResp clone() => UpdateTagResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UpdateTagResp copyWith(void Function(UpdateTagResp) updates) => super.copyWith((message) => updates(message as UpdateTagResp)) as UpdateTagResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UpdateTagResp create() => UpdateTagResp._();
  @$core.override
  UpdateTagResp createEmptyInstance() => create();
  static $pb.PbList<UpdateTagResp> createRepeated() => $pb.PbList<UpdateTagResp>();
  @$core.pragma('dart2js:noInline')
  static UpdateTagResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateTagResp>(create);
  static UpdateTagResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Tag get tag => $_getN(0);
  @$pb.TagNumber(1)
  set tag($2.Tag value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasTag() => $_has(0);
  @$pb.TagNumber(1)
  void clearTag() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.Tag ensureTag() => $_ensure(0);
}

class DeleteTagReq extends $pb.GeneratedMessage {
  factory DeleteTagReq({
    $1.Context? context,
    $core.String? uuid,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (uuid != null) result.uuid = uuid;
    return result;
  }

  DeleteTagReq._();

  factory DeleteTagReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory DeleteTagReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeleteTagReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'uuid')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeleteTagReq clone() => DeleteTagReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeleteTagReq copyWith(void Function(DeleteTagReq) updates) => super.copyWith((message) => updates(message as DeleteTagReq)) as DeleteTagReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeleteTagReq create() => DeleteTagReq._();
  @$core.override
  DeleteTagReq createEmptyInstance() => create();
  static $pb.PbList<DeleteTagReq> createRepeated() => $pb.PbList<DeleteTagReq>();
  @$core.pragma('dart2js:noInline')
  static DeleteTagReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteTagReq>(create);
  static DeleteTagReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get uuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUuid() => $_clearField(2);
}

class DeleteTagResp extends $pb.GeneratedMessage {
  factory DeleteTagResp() => create();

  DeleteTagResp._();

  factory DeleteTagResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory DeleteTagResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeleteTagResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeleteTagResp clone() => DeleteTagResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeleteTagResp copyWith(void Function(DeleteTagResp) updates) => super.copyWith((message) => updates(message as DeleteTagResp)) as DeleteTagResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeleteTagResp create() => DeleteTagResp._();
  @$core.override
  DeleteTagResp createEmptyInstance() => create();
  static $pb.PbList<DeleteTagResp> createRepeated() => $pb.PbList<DeleteTagResp>();
  @$core.pragma('dart2js:noInline')
  static DeleteTagResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteTagResp>(create);
  static DeleteTagResp? _defaultInstance;
}

class CreateTaskReq extends $pb.GeneratedMessage {
  factory CreateTaskReq({
    $1.Context? context,
    $core.String? name,
    $core.String? description,
    $2.Task_Status? status,
    $3.Duration? estimatedDuration,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (name != null) result.name = name;
    if (description != null) result.description = description;
    if (status != null) result.status = status;
    if (estimatedDuration != null) result.estimatedDuration = estimatedDuration;
    return result;
  }

  CreateTaskReq._();

  factory CreateTaskReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTaskReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTaskReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'name')
    ..aOS(3, _omitFieldNames ? '' : 'description')
    ..e<$2.Task_Status>(4, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: $2.Task_Status.STATUS_NONE, valueOf: $2.Task_Status.valueOf, enumValues: $2.Task_Status.values)
    ..aOM<$3.Duration>(5, _omitFieldNames ? '' : 'estimatedDuration', subBuilder: $3.Duration.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskReq clone() => CreateTaskReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskReq copyWith(void Function(CreateTaskReq) updates) => super.copyWith((message) => updates(message as CreateTaskReq)) as CreateTaskReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskReq create() => CreateTaskReq._();
  @$core.override
  CreateTaskReq createEmptyInstance() => create();
  static $pb.PbList<CreateTaskReq> createRepeated() => $pb.PbList<CreateTaskReq>();
  @$core.pragma('dart2js:noInline')
  static CreateTaskReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTaskReq>(create);
  static CreateTaskReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get description => $_getSZ(2);
  @$pb.TagNumber(3)
  set description($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasDescription() => $_has(2);
  @$pb.TagNumber(3)
  void clearDescription() => $_clearField(3);

  @$pb.TagNumber(4)
  $2.Task_Status get status => $_getN(3);
  @$pb.TagNumber(4)
  set status($2.Task_Status value) => $_setField(4, value);
  @$pb.TagNumber(4)
  $core.bool hasStatus() => $_has(3);
  @$pb.TagNumber(4)
  void clearStatus() => $_clearField(4);

  @$pb.TagNumber(5)
  $3.Duration get estimatedDuration => $_getN(4);
  @$pb.TagNumber(5)
  set estimatedDuration($3.Duration value) => $_setField(5, value);
  @$pb.TagNumber(5)
  $core.bool hasEstimatedDuration() => $_has(4);
  @$pb.TagNumber(5)
  void clearEstimatedDuration() => $_clearField(5);
  @$pb.TagNumber(5)
  $3.Duration ensureEstimatedDuration() => $_ensure(4);
}

class CreateTaskResp extends $pb.GeneratedMessage {
  factory CreateTaskResp({
    $2.Task? task,
  }) {
    final result = create();
    if (task != null) result.task = task;
    return result;
  }

  CreateTaskResp._();

  factory CreateTaskResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTaskResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTaskResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.Task>(1, _omitFieldNames ? '' : 'task', subBuilder: $2.Task.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskResp clone() => CreateTaskResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskResp copyWith(void Function(CreateTaskResp) updates) => super.copyWith((message) => updates(message as CreateTaskResp)) as CreateTaskResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskResp create() => CreateTaskResp._();
  @$core.override
  CreateTaskResp createEmptyInstance() => create();
  static $pb.PbList<CreateTaskResp> createRepeated() => $pb.PbList<CreateTaskResp>();
  @$core.pragma('dart2js:noInline')
  static CreateTaskResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTaskResp>(create);
  static CreateTaskResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Task get task => $_getN(0);
  @$pb.TagNumber(1)
  set task($2.Task value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasTask() => $_has(0);
  @$pb.TagNumber(1)
  void clearTask() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.Task ensureTask() => $_ensure(0);
}

class GetTaskReq extends $pb.GeneratedMessage {
  factory GetTaskReq({
    $1.Context? context,
    $core.String? uuid,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (uuid != null) result.uuid = uuid;
    return result;
  }

  GetTaskReq._();

  factory GetTaskReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTaskReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTaskReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'uuid')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskReq clone() => GetTaskReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskReq copyWith(void Function(GetTaskReq) updates) => super.copyWith((message) => updates(message as GetTaskReq)) as GetTaskReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTaskReq create() => GetTaskReq._();
  @$core.override
  GetTaskReq createEmptyInstance() => create();
  static $pb.PbList<GetTaskReq> createRepeated() => $pb.PbList<GetTaskReq>();
  @$core.pragma('dart2js:noInline')
  static GetTaskReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTaskReq>(create);
  static GetTaskReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get uuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUuid() => $_clearField(2);
}

class GetTaskResp extends $pb.GeneratedMessage {
  factory GetTaskResp({
    $2.Task? task,
  }) {
    final result = create();
    if (task != null) result.task = task;
    return result;
  }

  GetTaskResp._();

  factory GetTaskResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTaskResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTaskResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$2.Task>(1, _omitFieldNames ? '' : 'task', subBuilder: $2.Task.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskResp clone() => GetTaskResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskResp copyWith(void Function(GetTaskResp) updates) => super.copyWith((message) => updates(message as GetTaskResp)) as GetTaskResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTaskResp create() => GetTaskResp._();
  @$core.override
  GetTaskResp createEmptyInstance() => create();
  static $pb.PbList<GetTaskResp> createRepeated() => $pb.PbList<GetTaskResp>();
  @$core.pragma('dart2js:noInline')
  static GetTaskResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTaskResp>(create);
  static GetTaskResp? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Task get task => $_getN(0);
  @$pb.TagNumber(1)
  set task($2.Task value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasTask() => $_has(0);
  @$pb.TagNumber(1)
  void clearTask() => $_clearField(1);
  @$pb.TagNumber(1)
  $2.Task ensureTask() => $_ensure(0);
}

class CreateTaskLinksReq extends $pb.GeneratedMessage {
  factory CreateTaskLinksReq({
    $1.Context? context,
    $core.String? fromTaskUuid,
    $core.String? toTaskUuid,
    $2.Task_Link? link,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (fromTaskUuid != null) result.fromTaskUuid = fromTaskUuid;
    if (toTaskUuid != null) result.toTaskUuid = toTaskUuid;
    if (link != null) result.link = link;
    return result;
  }

  CreateTaskLinksReq._();

  factory CreateTaskLinksReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTaskLinksReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTaskLinksReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'fromTaskUuid')
    ..aOS(3, _omitFieldNames ? '' : 'toTaskUuid')
    ..e<$2.Task_Link>(4, _omitFieldNames ? '' : 'link', $pb.PbFieldType.OE, defaultOrMaker: $2.Task_Link.LINK_NONE, valueOf: $2.Task_Link.valueOf, enumValues: $2.Task_Link.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskLinksReq clone() => CreateTaskLinksReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskLinksReq copyWith(void Function(CreateTaskLinksReq) updates) => super.copyWith((message) => updates(message as CreateTaskLinksReq)) as CreateTaskLinksReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskLinksReq create() => CreateTaskLinksReq._();
  @$core.override
  CreateTaskLinksReq createEmptyInstance() => create();
  static $pb.PbList<CreateTaskLinksReq> createRepeated() => $pb.PbList<CreateTaskLinksReq>();
  @$core.pragma('dart2js:noInline')
  static CreateTaskLinksReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTaskLinksReq>(create);
  static CreateTaskLinksReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get fromTaskUuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set fromTaskUuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasFromTaskUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearFromTaskUuid() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get toTaskUuid => $_getSZ(2);
  @$pb.TagNumber(3)
  set toTaskUuid($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasToTaskUuid() => $_has(2);
  @$pb.TagNumber(3)
  void clearToTaskUuid() => $_clearField(3);

  @$pb.TagNumber(4)
  $2.Task_Link get link => $_getN(3);
  @$pb.TagNumber(4)
  set link($2.Task_Link value) => $_setField(4, value);
  @$pb.TagNumber(4)
  $core.bool hasLink() => $_has(3);
  @$pb.TagNumber(4)
  void clearLink() => $_clearField(4);
}

class CreateTaskLinksResp extends $pb.GeneratedMessage {
  factory CreateTaskLinksResp() => create();

  CreateTaskLinksResp._();

  factory CreateTaskLinksResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory CreateTaskLinksResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTaskLinksResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskLinksResp clone() => CreateTaskLinksResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskLinksResp copyWith(void Function(CreateTaskLinksResp) updates) => super.copyWith((message) => updates(message as CreateTaskLinksResp)) as CreateTaskLinksResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskLinksResp create() => CreateTaskLinksResp._();
  @$core.override
  CreateTaskLinksResp createEmptyInstance() => create();
  static $pb.PbList<CreateTaskLinksResp> createRepeated() => $pb.PbList<CreateTaskLinksResp>();
  @$core.pragma('dart2js:noInline')
  static CreateTaskLinksResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTaskLinksResp>(create);
  static CreateTaskLinksResp? _defaultInstance;
}

class GetTaskLinksReq extends $pb.GeneratedMessage {
  factory GetTaskLinksReq({
    $1.Context? context,
    $core.String? fromTaskUuid,
    $core.String? toTaskUuid,
  }) {
    final result = create();
    if (context != null) result.context = context;
    if (fromTaskUuid != null) result.fromTaskUuid = fromTaskUuid;
    if (toTaskUuid != null) result.toTaskUuid = toTaskUuid;
    return result;
  }

  GetTaskLinksReq._();

  factory GetTaskLinksReq.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTaskLinksReq.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTaskLinksReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..aOM<$1.Context>(1, _omitFieldNames ? '' : 'context', subBuilder: $1.Context.create)
    ..aOS(2, _omitFieldNames ? '' : 'fromTaskUuid')
    ..aOS(3, _omitFieldNames ? '' : 'toTaskUuid')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskLinksReq clone() => GetTaskLinksReq()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskLinksReq copyWith(void Function(GetTaskLinksReq) updates) => super.copyWith((message) => updates(message as GetTaskLinksReq)) as GetTaskLinksReq;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTaskLinksReq create() => GetTaskLinksReq._();
  @$core.override
  GetTaskLinksReq createEmptyInstance() => create();
  static $pb.PbList<GetTaskLinksReq> createRepeated() => $pb.PbList<GetTaskLinksReq>();
  @$core.pragma('dart2js:noInline')
  static GetTaskLinksReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTaskLinksReq>(create);
  static GetTaskLinksReq? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Context get context => $_getN(0);
  @$pb.TagNumber(1)
  set context($1.Context value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasContext() => $_has(0);
  @$pb.TagNumber(1)
  void clearContext() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.Context ensureContext() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get fromTaskUuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set fromTaskUuid($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasFromTaskUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearFromTaskUuid() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get toTaskUuid => $_getSZ(2);
  @$pb.TagNumber(3)
  set toTaskUuid($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasToTaskUuid() => $_has(2);
  @$pb.TagNumber(3)
  void clearToTaskUuid() => $_clearField(3);
}

class GetTaskLinksResp extends $pb.GeneratedMessage {
  factory GetTaskLinksResp({
    $core.Iterable<$2.Task_Link>? links,
  }) {
    final result = create();
    if (links != null) result.links.addAll(links);
    return result;
  }

  GetTaskLinksResp._();

  factory GetTaskLinksResp.fromBuffer($core.List<$core.int> data, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(data, registry);
  factory GetTaskLinksResp.fromJson($core.String json, [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTaskLinksResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'ttt'), createEmptyInstance: create)
    ..pc<$2.Task_Link>(1, _omitFieldNames ? '' : 'links', $pb.PbFieldType.KE, valueOf: $2.Task_Link.valueOf, enumValues: $2.Task_Link.values, defaultEnumValue: $2.Task_Link.LINK_NONE)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskLinksResp clone() => GetTaskLinksResp()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetTaskLinksResp copyWith(void Function(GetTaskLinksResp) updates) => super.copyWith((message) => updates(message as GetTaskLinksResp)) as GetTaskLinksResp;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTaskLinksResp create() => GetTaskLinksResp._();
  @$core.override
  GetTaskLinksResp createEmptyInstance() => create();
  static $pb.PbList<GetTaskLinksResp> createRepeated() => $pb.PbList<GetTaskLinksResp>();
  @$core.pragma('dart2js:noInline')
  static GetTaskLinksResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTaskLinksResp>(create);
  static GetTaskLinksResp? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$2.Task_Link> get links => $_getList(0);
}


const $core.bool _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
