// This is a generated file - do not edit.
//
// Generated from proto/ttt_service.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'ttt_service.pb.dart' as $0;

export 'ttt_service.pb.dart';

/// TTTService defines the main gRPC service for Todo Time Tracker
@$pb.GrpcServiceName('ttt.TTTService')
class TTTServiceClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  TTTServiceClient(super.channel, {super.options, super.interceptors});

  /// User operations
  $grpc.ResponseFuture<$0.CreateUserResp> createUser($0.CreateUserReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$createUser, request, options: options);
  }

  /// Tag operations
  $grpc.ResponseFuture<$0.GetTagResp> getTag($0.GetTagReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$getTag, request, options: options);
  }

  $grpc.ResponseFuture<$0.ListTagsResp> listTags($0.ListTagsReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$listTags, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateTagResp> createTag($0.CreateTagReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$createTag, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateTagResp> updateTag($0.UpdateTagReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$updateTag, request, options: options);
  }

  $grpc.ResponseFuture<$0.DeleteTagResp> deleteTag($0.DeleteTagReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$deleteTag, request, options: options);
  }

  /// Task operations
  $grpc.ResponseFuture<$0.CreateTaskResp> createTask($0.CreateTaskReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$createTask, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetTaskResp> getTask($0.GetTaskReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$getTask, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateTaskLinksResp> createTaskLinks($0.CreateTaskLinksReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$createTaskLinks, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetTaskLinksResp> getTaskLinks($0.GetTaskLinksReq request, {$grpc.CallOptions? options,}) {
    return $createUnaryCall(_$getTaskLinks, request, options: options);
  }

    // method descriptors

  static final _$createUser = $grpc.ClientMethod<$0.CreateUserReq, $0.CreateUserResp>(
      '/ttt.TTTService/CreateUser',
      ($0.CreateUserReq value) => value.writeToBuffer(),
      $0.CreateUserResp.fromBuffer);
  static final _$getTag = $grpc.ClientMethod<$0.GetTagReq, $0.GetTagResp>(
      '/ttt.TTTService/GetTag',
      ($0.GetTagReq value) => value.writeToBuffer(),
      $0.GetTagResp.fromBuffer);
  static final _$listTags = $grpc.ClientMethod<$0.ListTagsReq, $0.ListTagsResp>(
      '/ttt.TTTService/ListTags',
      ($0.ListTagsReq value) => value.writeToBuffer(),
      $0.ListTagsResp.fromBuffer);
  static final _$createTag = $grpc.ClientMethod<$0.CreateTagReq, $0.CreateTagResp>(
      '/ttt.TTTService/CreateTag',
      ($0.CreateTagReq value) => value.writeToBuffer(),
      $0.CreateTagResp.fromBuffer);
  static final _$updateTag = $grpc.ClientMethod<$0.UpdateTagReq, $0.UpdateTagResp>(
      '/ttt.TTTService/UpdateTag',
      ($0.UpdateTagReq value) => value.writeToBuffer(),
      $0.UpdateTagResp.fromBuffer);
  static final _$deleteTag = $grpc.ClientMethod<$0.DeleteTagReq, $0.DeleteTagResp>(
      '/ttt.TTTService/DeleteTag',
      ($0.DeleteTagReq value) => value.writeToBuffer(),
      $0.DeleteTagResp.fromBuffer);
  static final _$createTask = $grpc.ClientMethod<$0.CreateTaskReq, $0.CreateTaskResp>(
      '/ttt.TTTService/CreateTask',
      ($0.CreateTaskReq value) => value.writeToBuffer(),
      $0.CreateTaskResp.fromBuffer);
  static final _$getTask = $grpc.ClientMethod<$0.GetTaskReq, $0.GetTaskResp>(
      '/ttt.TTTService/GetTask',
      ($0.GetTaskReq value) => value.writeToBuffer(),
      $0.GetTaskResp.fromBuffer);
  static final _$createTaskLinks = $grpc.ClientMethod<$0.CreateTaskLinksReq, $0.CreateTaskLinksResp>(
      '/ttt.TTTService/CreateTaskLinks',
      ($0.CreateTaskLinksReq value) => value.writeToBuffer(),
      $0.CreateTaskLinksResp.fromBuffer);
  static final _$getTaskLinks = $grpc.ClientMethod<$0.GetTaskLinksReq, $0.GetTaskLinksResp>(
      '/ttt.TTTService/GetTaskLinks',
      ($0.GetTaskLinksReq value) => value.writeToBuffer(),
      $0.GetTaskLinksResp.fromBuffer);
}

@$pb.GrpcServiceName('ttt.TTTService')
abstract class TTTServiceBase extends $grpc.Service {
  $core.String get $name => 'ttt.TTTService';

  TTTServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CreateUserReq, $0.CreateUserResp>(
        'CreateUser',
        createUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateUserReq.fromBuffer(value),
        ($0.CreateUserResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetTagReq, $0.GetTagResp>(
        'GetTag',
        getTag_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetTagReq.fromBuffer(value),
        ($0.GetTagResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ListTagsReq, $0.ListTagsResp>(
        'ListTags',
        listTags_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ListTagsReq.fromBuffer(value),
        ($0.ListTagsResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateTagReq, $0.CreateTagResp>(
        'CreateTag',
        createTag_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateTagReq.fromBuffer(value),
        ($0.CreateTagResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateTagReq, $0.UpdateTagResp>(
        'UpdateTag',
        updateTag_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateTagReq.fromBuffer(value),
        ($0.UpdateTagResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteTagReq, $0.DeleteTagResp>(
        'DeleteTag',
        deleteTag_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DeleteTagReq.fromBuffer(value),
        ($0.DeleteTagResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateTaskReq, $0.CreateTaskResp>(
        'CreateTask',
        createTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateTaskReq.fromBuffer(value),
        ($0.CreateTaskResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetTaskReq, $0.GetTaskResp>(
        'GetTask',
        getTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetTaskReq.fromBuffer(value),
        ($0.GetTaskResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateTaskLinksReq, $0.CreateTaskLinksResp>(
        'CreateTaskLinks',
        createTaskLinks_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateTaskLinksReq.fromBuffer(value),
        ($0.CreateTaskLinksResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetTaskLinksReq, $0.GetTaskLinksResp>(
        'GetTaskLinks',
        getTaskLinks_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetTaskLinksReq.fromBuffer(value),
        ($0.GetTaskLinksResp value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateUserResp> createUser_Pre($grpc.ServiceCall $call, $async.Future<$0.CreateUserReq> $request) async {
    return createUser($call, await $request);
  }

  $async.Future<$0.CreateUserResp> createUser($grpc.ServiceCall call, $0.CreateUserReq request);

  $async.Future<$0.GetTagResp> getTag_Pre($grpc.ServiceCall $call, $async.Future<$0.GetTagReq> $request) async {
    return getTag($call, await $request);
  }

  $async.Future<$0.GetTagResp> getTag($grpc.ServiceCall call, $0.GetTagReq request);

  $async.Future<$0.ListTagsResp> listTags_Pre($grpc.ServiceCall $call, $async.Future<$0.ListTagsReq> $request) async {
    return listTags($call, await $request);
  }

  $async.Future<$0.ListTagsResp> listTags($grpc.ServiceCall call, $0.ListTagsReq request);

  $async.Future<$0.CreateTagResp> createTag_Pre($grpc.ServiceCall $call, $async.Future<$0.CreateTagReq> $request) async {
    return createTag($call, await $request);
  }

  $async.Future<$0.CreateTagResp> createTag($grpc.ServiceCall call, $0.CreateTagReq request);

  $async.Future<$0.UpdateTagResp> updateTag_Pre($grpc.ServiceCall $call, $async.Future<$0.UpdateTagReq> $request) async {
    return updateTag($call, await $request);
  }

  $async.Future<$0.UpdateTagResp> updateTag($grpc.ServiceCall call, $0.UpdateTagReq request);

  $async.Future<$0.DeleteTagResp> deleteTag_Pre($grpc.ServiceCall $call, $async.Future<$0.DeleteTagReq> $request) async {
    return deleteTag($call, await $request);
  }

  $async.Future<$0.DeleteTagResp> deleteTag($grpc.ServiceCall call, $0.DeleteTagReq request);

  $async.Future<$0.CreateTaskResp> createTask_Pre($grpc.ServiceCall $call, $async.Future<$0.CreateTaskReq> $request) async {
    return createTask($call, await $request);
  }

  $async.Future<$0.CreateTaskResp> createTask($grpc.ServiceCall call, $0.CreateTaskReq request);

  $async.Future<$0.GetTaskResp> getTask_Pre($grpc.ServiceCall $call, $async.Future<$0.GetTaskReq> $request) async {
    return getTask($call, await $request);
  }

  $async.Future<$0.GetTaskResp> getTask($grpc.ServiceCall call, $0.GetTaskReq request);

  $async.Future<$0.CreateTaskLinksResp> createTaskLinks_Pre($grpc.ServiceCall $call, $async.Future<$0.CreateTaskLinksReq> $request) async {
    return createTaskLinks($call, await $request);
  }

  $async.Future<$0.CreateTaskLinksResp> createTaskLinks($grpc.ServiceCall call, $0.CreateTaskLinksReq request);

  $async.Future<$0.GetTaskLinksResp> getTaskLinks_Pre($grpc.ServiceCall $call, $async.Future<$0.GetTaskLinksReq> $request) async {
    return getTaskLinks($call, await $request);
  }

  $async.Future<$0.GetTaskLinksResp> getTaskLinks($grpc.ServiceCall call, $0.GetTaskLinksReq request);

}
