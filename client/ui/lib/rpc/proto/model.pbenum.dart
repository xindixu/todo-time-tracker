// This is a generated file - do not edit.
//
// Generated from proto/model.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Account_Type extends $pb.ProtobufEnum {
  static const Account_Type NONE = Account_Type._(0, _omitEnumNames ? '' : 'NONE');
  static const Account_Type USER = Account_Type._(1, _omitEnumNames ? '' : 'USER');
  static const Account_Type ORGANIZATION = Account_Type._(2, _omitEnumNames ? '' : 'ORGANIZATION');

  static const $core.List<Account_Type> values = <Account_Type> [
    NONE,
    USER,
    ORGANIZATION,
  ];

  static final $core.List<Account_Type?> _byValue = $pb.ProtobufEnum.$_initByValueList(values, 2);
  static Account_Type? valueOf($core.int value) =>  value < 0 || value >= _byValue.length ? null : _byValue[value];

  const Account_Type._(super.value, super.name);
}

class Task_Status extends $pb.ProtobufEnum {
  static const Task_Status STATUS_NONE = Task_Status._(0, _omitEnumNames ? '' : 'STATUS_NONE');
  static const Task_Status TODO = Task_Status._(1, _omitEnumNames ? '' : 'TODO');
  static const Task_Status IN_PROGRESS = Task_Status._(2, _omitEnumNames ? '' : 'IN_PROGRESS');
  static const Task_Status DONE = Task_Status._(3, _omitEnumNames ? '' : 'DONE');
  static const Task_Status BLOCKED = Task_Status._(4, _omitEnumNames ? '' : 'BLOCKED');

  static const $core.List<Task_Status> values = <Task_Status> [
    STATUS_NONE,
    TODO,
    IN_PROGRESS,
    DONE,
    BLOCKED,
  ];

  static final $core.List<Task_Status?> _byValue = $pb.ProtobufEnum.$_initByValueList(values, 4);
  static Task_Status? valueOf($core.int value) =>  value < 0 || value >= _byValue.length ? null : _byValue[value];

  const Task_Status._(super.value, super.name);
}

class Task_Link extends $pb.ProtobufEnum {
  static const Task_Link LINK_NONE = Task_Link._(0, _omitEnumNames ? '' : 'LINK_NONE');
  static const Task_Link PARENT_OF = Task_Link._(1, _omitEnumNames ? '' : 'PARENT_OF');
  static const Task_Link BLOCKS = Task_Link._(2, _omitEnumNames ? '' : 'BLOCKS');
  static const Task_Link RELATES_TO = Task_Link._(3, _omitEnumNames ? '' : 'RELATES_TO');
  static const Task_Link DUPLICATE_OF = Task_Link._(4, _omitEnumNames ? '' : 'DUPLICATE_OF');

  static const $core.List<Task_Link> values = <Task_Link> [
    LINK_NONE,
    PARENT_OF,
    BLOCKS,
    RELATES_TO,
    DUPLICATE_OF,
  ];

  static final $core.List<Task_Link?> _byValue = $pb.ProtobufEnum.$_initByValueList(values, 4);
  static Task_Link? valueOf($core.int value) =>  value < 0 || value >= _byValue.length ? null : _byValue[value];

  const Task_Link._(super.value, super.name);
}


const $core.bool _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
