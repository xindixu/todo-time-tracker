// This is a generated file - do not edit.
//
// Generated from proto/model.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use accountDescriptor instead')
const Account$json = {
  '1': 'Account',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'type', '3': 5, '4': 1, '5': 14, '6': '.model.Account.Type', '10': 'type'},
  ],
  '4': [Account_Type$json],
};

@$core.Deprecated('Use accountDescriptor instead')
const Account_Type$json = {
  '1': 'Type',
  '2': [
    {'1': 'NONE', '2': 0},
    {'1': 'USER', '2': 1},
    {'1': 'ORGANIZATION', '2': 2},
  ],
};

/// Descriptor for `Account`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List accountDescriptor = $convert.base64Decode(
    'CgdBY2NvdW50Eg4KAmlkGAEgASgDUgJpZBISCgR1dWlkGAIgASgJUgR1dWlkEjkKCmNyZWF0ZW'
    'RfYXQYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdGVkQXQSOQoKdXBk'
    'YXRlZF9hdBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXVwZGF0ZWRBdBInCg'
    'R0eXBlGAUgASgOMhMubW9kZWwuQWNjb3VudC5UeXBlUgR0eXBlIiwKBFR5cGUSCAoETk9ORRAA'
    'EggKBFVTRVIQARIQCgxPUkdBTklaQVRJT04QAg==');

@$core.Deprecated('Use organizationDescriptor instead')
const Organization$json = {
  '1': 'Organization',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'account_uuid', '3': 5, '4': 1, '5': 3, '10': 'accountUuid'},
    {'1': 'name', '3': 6, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `Organization`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List organizationDescriptor = $convert.base64Decode(
    'CgxPcmdhbml6YXRpb24SDgoCaWQYASABKANSAmlkEhIKBHV1aWQYAiABKAlSBHV1aWQSOQoKY3'
    'JlYXRlZF9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCWNyZWF0ZWRBdBI5'
    'Cgp1cGRhdGVkX2F0GAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJdXBkYXRlZE'
    'F0EiEKDGFjY291bnRfdXVpZBgFIAEoA1ILYWNjb3VudFV1aWQSEgoEbmFtZRgGIAEoCVIEbmFt'
    'ZQ==');

@$core.Deprecated('Use userDescriptor instead')
const User$json = {
  '1': 'User',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'account_uuid', '3': 5, '4': 1, '5': 3, '10': 'accountUuid'},
    {'1': 'name', '3': 6, '4': 1, '5': 9, '10': 'name'},
    {'1': 'email', '3': 7, '4': 1, '5': 9, '10': 'email'},
  ],
};

/// Descriptor for `User`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDescriptor = $convert.base64Decode(
    'CgRVc2VyEg4KAmlkGAEgASgDUgJpZBISCgR1dWlkGAIgASgJUgR1dWlkEjkKCmNyZWF0ZWRfYX'
    'QYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdGVkQXQSOQoKdXBkYXRl'
    'ZF9hdBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXVwZGF0ZWRBdBIhCgxhY2'
    'NvdW50X3V1aWQYBSABKANSC2FjY291bnRVdWlkEhIKBG5hbWUYBiABKAlSBG5hbWUSFAoFZW1h'
    'aWwYByABKAlSBWVtYWls');

@$core.Deprecated('Use tagDescriptor instead')
const Tag$json = {
  '1': 'Tag',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'account_uuid', '3': 5, '4': 1, '5': 3, '10': 'accountUuid'},
    {'1': 'name', '3': 6, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 7, '4': 1, '5': 9, '10': 'description'},
    {'1': 'color', '3': 8, '4': 1, '5': 9, '10': 'color'},
  ],
};

/// Descriptor for `Tag`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tagDescriptor = $convert.base64Decode(
    'CgNUYWcSDgoCaWQYASABKANSAmlkEhIKBHV1aWQYAiABKAlSBHV1aWQSOQoKY3JlYXRlZF9hdB'
    'gDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCWNyZWF0ZWRBdBI5Cgp1cGRhdGVk'
    'X2F0GAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJdXBkYXRlZEF0EiEKDGFjY2'
    '91bnRfdXVpZBgFIAEoA1ILYWNjb3VudFV1aWQSEgoEbmFtZRgGIAEoCVIEbmFtZRIgCgtkZXNj'
    'cmlwdGlvbhgHIAEoCVILZGVzY3JpcHRpb24SFAoFY29sb3IYCCABKAlSBWNvbG9y');

@$core.Deprecated('Use taskDescriptor instead')
const Task$json = {
  '1': 'Task',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'account_uuid', '3': 5, '4': 1, '5': 3, '10': 'accountUuid'},
    {'1': 'name', '3': 6, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 7, '4': 1, '5': 9, '10': 'description'},
    {'1': 'status', '3': 8, '4': 1, '5': 14, '6': '.model.Task.Status', '10': 'status'},
    {'1': 'estimated_duration', '3': 9, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'estimatedDuration'},
  ],
  '4': [Task_Status$json, Task_Link$json],
};

@$core.Deprecated('Use taskDescriptor instead')
const Task_Status$json = {
  '1': 'Status',
  '2': [
    {'1': 'STATUS_NONE', '2': 0},
    {'1': 'TODO', '2': 1},
    {'1': 'IN_PROGRESS', '2': 2},
    {'1': 'DONE', '2': 3},
    {'1': 'BLOCKED', '2': 4},
  ],
};

@$core.Deprecated('Use taskDescriptor instead')
const Task_Link$json = {
  '1': 'Link',
  '2': [
    {'1': 'LINK_NONE', '2': 0},
    {'1': 'PARENT_OF', '2': 1},
    {'1': 'BLOCKS', '2': 2},
    {'1': 'RELATES_TO', '2': 3},
    {'1': 'DUPLICATE_OF', '2': 4},
  ],
};

/// Descriptor for `Task`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taskDescriptor = $convert.base64Decode(
    'CgRUYXNrEg4KAmlkGAEgASgDUgJpZBISCgR1dWlkGAIgASgJUgR1dWlkEjkKCmNyZWF0ZWRfYX'
    'QYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdGVkQXQSOQoKdXBkYXRl'
    'ZF9hdBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXVwZGF0ZWRBdBIhCgxhY2'
    'NvdW50X3V1aWQYBSABKANSC2FjY291bnRVdWlkEhIKBG5hbWUYBiABKAlSBG5hbWUSIAoLZGVz'
    'Y3JpcHRpb24YByABKAlSC2Rlc2NyaXB0aW9uEioKBnN0YXR1cxgIIAEoDjISLm1vZGVsLlRhc2'
    'suU3RhdHVzUgZzdGF0dXMSSAoSZXN0aW1hdGVkX2R1cmF0aW9uGAkgASgLMhkuZ29vZ2xlLnBy'
    'b3RvYnVmLkR1cmF0aW9uUhFlc3RpbWF0ZWREdXJhdGlvbiJLCgZTdGF0dXMSDwoLU1RBVFVTX0'
    '5PTkUQABIICgRUT0RPEAESDwoLSU5fUFJPR1JFU1MQAhIICgRET05FEAMSCwoHQkxPQ0tFRBAE'
    'IlIKBExpbmsSDQoJTElOS19OT05FEAASDQoJUEFSRU5UX09GEAESCgoGQkxPQ0tTEAISDgoKUk'
    'VMQVRFU19UTxADEhAKDERVUExJQ0FURV9PRhAE');

@$core.Deprecated('Use sessionDescriptor instead')
const Session$json = {
  '1': 'Session',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'created_at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'updated_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    {'1': 'task_uuid', '3': 4, '4': 1, '5': 9, '10': 'taskUuid'},
    {'1': 'user_uuid', '3': 5, '4': 1, '5': 9, '10': 'userUuid'},
    {'1': 'starts_at', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'startsAt'},
    {'1': 'ends_at', '3': 7, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'endsAt'},
  ],
};

/// Descriptor for `Session`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sessionDescriptor = $convert.base64Decode(
    'CgdTZXNzaW9uEg4KAmlkGAEgASgDUgJpZBI5CgpjcmVhdGVkX2F0GAIgASgLMhouZ29vZ2xlLn'
    'Byb3RvYnVmLlRpbWVzdGFtcFIJY3JlYXRlZEF0EjkKCnVwZGF0ZWRfYXQYAyABKAsyGi5nb29n'
    'bGUucHJvdG9idWYuVGltZXN0YW1wUgl1cGRhdGVkQXQSGwoJdGFza191dWlkGAQgASgJUgh0YX'
    'NrVXVpZBIbCgl1c2VyX3V1aWQYBSABKAlSCHVzZXJVdWlkEjcKCXN0YXJ0c19hdBgGIAEoCzIa'
    'Lmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCHN0YXJ0c0F0EjMKB2VuZHNfYXQYByABKAsyGi'
    '5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgZlbmRzQXQ=');

