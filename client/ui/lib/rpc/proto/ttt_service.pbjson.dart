// This is a generated file - do not edit.
//
// Generated from proto/ttt_service.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use createUserReqDescriptor instead')
const CreateUserReq$json = {
  '1': 'CreateUserReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'email', '3': 3, '4': 1, '5': 9, '10': 'email'},
    {'1': 'password', '3': 4, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `CreateUserReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createUserReqDescriptor = $convert.base64Decode(
    'Cg1DcmVhdGVVc2VyUmVxEioKB2NvbnRleHQYASABKAsyEC5jb250ZXh0LkNvbnRleHRSB2Nvbn'
    'RleHQSEgoEbmFtZRgCIAEoCVIEbmFtZRIUCgVlbWFpbBgDIAEoCVIFZW1haWwSGgoIcGFzc3dv'
    'cmQYBCABKAlSCHBhc3N3b3Jk');

@$core.Deprecated('Use createUserRespDescriptor instead')
const CreateUserResp$json = {
  '1': 'CreateUserResp',
  '2': [
    {'1': 'user', '3': 1, '4': 1, '5': 11, '6': '.model.User', '10': 'user'},
  ],
};

/// Descriptor for `CreateUserResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createUserRespDescriptor = $convert.base64Decode(
    'Cg5DcmVhdGVVc2VyUmVzcBIfCgR1c2VyGAEgASgLMgsubW9kZWwuVXNlclIEdXNlcg==');

@$core.Deprecated('Use getTagReqDescriptor instead')
const GetTagReq$json = {
  '1': 'GetTagReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
  ],
};

/// Descriptor for `GetTagReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTagReqDescriptor = $convert.base64Decode(
    'CglHZXRUYWdSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dFIHY29udGV4dB'
    'ISCgR1dWlkGAIgASgJUgR1dWlk');

@$core.Deprecated('Use getTagRespDescriptor instead')
const GetTagResp$json = {
  '1': 'GetTagResp',
  '2': [
    {'1': 'tag', '3': 1, '4': 1, '5': 11, '6': '.model.Tag', '10': 'tag'},
  ],
};

/// Descriptor for `GetTagResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTagRespDescriptor = $convert.base64Decode(
    'CgpHZXRUYWdSZXNwEhwKA3RhZxgBIAEoCzIKLm1vZGVsLlRhZ1IDdGFn');

@$core.Deprecated('Use listTagsReqDescriptor instead')
const ListTagsReq$json = {
  '1': 'ListTagsReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
  ],
};

/// Descriptor for `ListTagsReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTagsReqDescriptor = $convert.base64Decode(
    'CgtMaXN0VGFnc1JlcRIqCgdjb250ZXh0GAEgASgLMhAuY29udGV4dC5Db250ZXh0Ugdjb250ZX'
    'h0');

@$core.Deprecated('Use listTagsRespDescriptor instead')
const ListTagsResp$json = {
  '1': 'ListTagsResp',
  '2': [
    {'1': 'tags', '3': 1, '4': 3, '5': 11, '6': '.model.Tag', '10': 'tags'},
  ],
};

/// Descriptor for `ListTagsResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTagsRespDescriptor = $convert.base64Decode(
    'CgxMaXN0VGFnc1Jlc3ASHgoEdGFncxgBIAMoCzIKLm1vZGVsLlRhZ1IEdGFncw==');

@$core.Deprecated('Use createTagReqDescriptor instead')
const CreateTagReq$json = {
  '1': 'CreateTagReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `CreateTagReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTagReqDescriptor = $convert.base64Decode(
    'CgxDcmVhdGVUYWdSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dFIHY29udG'
    'V4dBISCgRuYW1lGAIgASgJUgRuYW1l');

@$core.Deprecated('Use createTagRespDescriptor instead')
const CreateTagResp$json = {
  '1': 'CreateTagResp',
  '2': [
    {'1': 'tag', '3': 1, '4': 1, '5': 11, '6': '.model.Tag', '10': 'tag'},
  ],
};

/// Descriptor for `CreateTagResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTagRespDescriptor = $convert.base64Decode(
    'Cg1DcmVhdGVUYWdSZXNwEhwKA3RhZxgBIAEoCzIKLm1vZGVsLlRhZ1IDdGFn');

@$core.Deprecated('Use updateTagReqDescriptor instead')
const UpdateTagReq$json = {
  '1': 'UpdateTagReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
    {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `UpdateTagReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTagReqDescriptor = $convert.base64Decode(
    'CgxVcGRhdGVUYWdSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dFIHY29udG'
    'V4dBISCgR1dWlkGAIgASgJUgR1dWlkEhIKBG5hbWUYAyABKAlSBG5hbWU=');

@$core.Deprecated('Use updateTagRespDescriptor instead')
const UpdateTagResp$json = {
  '1': 'UpdateTagResp',
  '2': [
    {'1': 'tag', '3': 1, '4': 1, '5': 11, '6': '.model.Tag', '10': 'tag'},
  ],
};

/// Descriptor for `UpdateTagResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTagRespDescriptor = $convert.base64Decode(
    'Cg1VcGRhdGVUYWdSZXNwEhwKA3RhZxgBIAEoCzIKLm1vZGVsLlRhZ1IDdGFn');

@$core.Deprecated('Use deleteTagReqDescriptor instead')
const DeleteTagReq$json = {
  '1': 'DeleteTagReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
  ],
};

/// Descriptor for `DeleteTagReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTagReqDescriptor = $convert.base64Decode(
    'CgxEZWxldGVUYWdSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dFIHY29udG'
    'V4dBISCgR1dWlkGAIgASgJUgR1dWlk');

@$core.Deprecated('Use deleteTagRespDescriptor instead')
const DeleteTagResp$json = {
  '1': 'DeleteTagResp',
};

/// Descriptor for `DeleteTagResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTagRespDescriptor = $convert.base64Decode(
    'Cg1EZWxldGVUYWdSZXNw');

@$core.Deprecated('Use createTaskReqDescriptor instead')
const CreateTaskReq$json = {
  '1': 'CreateTaskReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'status', '3': 4, '4': 1, '5': 14, '6': '.model.Task.Status', '10': 'status'},
    {'1': 'estimated_duration', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'estimatedDuration'},
  ],
};

/// Descriptor for `CreateTaskReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskReqDescriptor = $convert.base64Decode(
    'Cg1DcmVhdGVUYXNrUmVxEioKB2NvbnRleHQYASABKAsyEC5jb250ZXh0LkNvbnRleHRSB2Nvbn'
    'RleHQSEgoEbmFtZRgCIAEoCVIEbmFtZRIgCgtkZXNjcmlwdGlvbhgDIAEoCVILZGVzY3JpcHRp'
    'b24SKgoGc3RhdHVzGAQgASgOMhIubW9kZWwuVGFzay5TdGF0dXNSBnN0YXR1cxJIChJlc3RpbW'
    'F0ZWRfZHVyYXRpb24YBSABKAsyGS5nb29nbGUucHJvdG9idWYuRHVyYXRpb25SEWVzdGltYXRl'
    'ZER1cmF0aW9u');

@$core.Deprecated('Use createTaskRespDescriptor instead')
const CreateTaskResp$json = {
  '1': 'CreateTaskResp',
  '2': [
    {'1': 'task', '3': 1, '4': 1, '5': 11, '6': '.model.Task', '10': 'task'},
  ],
};

/// Descriptor for `CreateTaskResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskRespDescriptor = $convert.base64Decode(
    'Cg5DcmVhdGVUYXNrUmVzcBIfCgR0YXNrGAEgASgLMgsubW9kZWwuVGFza1IEdGFzaw==');

@$core.Deprecated('Use getTaskReqDescriptor instead')
const GetTaskReq$json = {
  '1': 'GetTaskReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'uuid', '3': 2, '4': 1, '5': 9, '10': 'uuid'},
  ],
};

/// Descriptor for `GetTaskReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskReqDescriptor = $convert.base64Decode(
    'CgpHZXRUYXNrUmVxEioKB2NvbnRleHQYASABKAsyEC5jb250ZXh0LkNvbnRleHRSB2NvbnRleH'
    'QSEgoEdXVpZBgCIAEoCVIEdXVpZA==');

@$core.Deprecated('Use getTaskRespDescriptor instead')
const GetTaskResp$json = {
  '1': 'GetTaskResp',
  '2': [
    {'1': 'task', '3': 1, '4': 1, '5': 11, '6': '.model.Task', '10': 'task'},
  ],
};

/// Descriptor for `GetTaskResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskRespDescriptor = $convert.base64Decode(
    'CgtHZXRUYXNrUmVzcBIfCgR0YXNrGAEgASgLMgsubW9kZWwuVGFza1IEdGFzaw==');

@$core.Deprecated('Use createTaskLinksReqDescriptor instead')
const CreateTaskLinksReq$json = {
  '1': 'CreateTaskLinksReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'from_task_uuid', '3': 2, '4': 1, '5': 9, '10': 'fromTaskUuid'},
    {'1': 'to_task_uuid', '3': 3, '4': 1, '5': 9, '10': 'toTaskUuid'},
    {'1': 'link', '3': 4, '4': 1, '5': 14, '6': '.model.Task.Link', '10': 'link'},
  ],
};

/// Descriptor for `CreateTaskLinksReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskLinksReqDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVUYXNrTGlua3NSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dF'
    'IHY29udGV4dBIkCg5mcm9tX3Rhc2tfdXVpZBgCIAEoCVIMZnJvbVRhc2tVdWlkEiAKDHRvX3Rh'
    'c2tfdXVpZBgDIAEoCVIKdG9UYXNrVXVpZBIkCgRsaW5rGAQgASgOMhAubW9kZWwuVGFzay5MaW'
    '5rUgRsaW5r');

@$core.Deprecated('Use createTaskLinksRespDescriptor instead')
const CreateTaskLinksResp$json = {
  '1': 'CreateTaskLinksResp',
};

/// Descriptor for `CreateTaskLinksResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskLinksRespDescriptor = $convert.base64Decode(
    'ChNDcmVhdGVUYXNrTGlua3NSZXNw');

@$core.Deprecated('Use getTaskLinksReqDescriptor instead')
const GetTaskLinksReq$json = {
  '1': 'GetTaskLinksReq',
  '2': [
    {'1': 'context', '3': 1, '4': 1, '5': 11, '6': '.context.Context', '10': 'context'},
    {'1': 'from_task_uuid', '3': 2, '4': 1, '5': 9, '10': 'fromTaskUuid'},
    {'1': 'to_task_uuid', '3': 3, '4': 1, '5': 9, '10': 'toTaskUuid'},
  ],
};

/// Descriptor for `GetTaskLinksReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskLinksReqDescriptor = $convert.base64Decode(
    'Cg9HZXRUYXNrTGlua3NSZXESKgoHY29udGV4dBgBIAEoCzIQLmNvbnRleHQuQ29udGV4dFIHY2'
    '9udGV4dBIkCg5mcm9tX3Rhc2tfdXVpZBgCIAEoCVIMZnJvbVRhc2tVdWlkEiAKDHRvX3Rhc2tf'
    'dXVpZBgDIAEoCVIKdG9UYXNrVXVpZA==');

@$core.Deprecated('Use getTaskLinksRespDescriptor instead')
const GetTaskLinksResp$json = {
  '1': 'GetTaskLinksResp',
  '2': [
    {'1': 'links', '3': 1, '4': 3, '5': 14, '6': '.model.Task.Link', '10': 'links'},
  ],
};

/// Descriptor for `GetTaskLinksResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskLinksRespDescriptor = $convert.base64Decode(
    'ChBHZXRUYXNrTGlua3NSZXNwEiYKBWxpbmtzGAEgAygOMhAubW9kZWwuVGFzay5MaW5rUgVsaW'
    '5rcw==');

