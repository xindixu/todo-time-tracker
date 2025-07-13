import 'package:ui/rpc/proto/context.pb.dart';
import 'package:ui/rpc/proto/model.pb.dart' as model;
import 'package:ui/rpc/proto/ttt_service.pbgrpc.dart';
import 'package:ui/utils/converter.dart';

// TODO: build context using user info
final context = Context(
  userName: 'xindixu',
  userUuid: '41b95446-dca2-400f-b839-7aa5e9a92e9f',
  token: '123',
);

class Task {
  final String uuid;
  final String name;
  final String description;
  final model.Task_Status status;
  final Duration? estimatedDuration;

  Task({
    required this.uuid,
    required this.name,
    required this.description,
    required this.status,
    required this.estimatedDuration,
  });

  Task copyWith({
    String? uuid,
    String? name,
    String? description,
    model.Task_Status? status,
    Duration? estimatedDuration,
  }) {
    return Task(
      uuid: uuid ?? this.uuid,
      name: name ?? this.name,
      description: description ?? this.description,
      status: status ?? this.status,
      estimatedDuration: estimatedDuration ?? this.estimatedDuration,
    );
  }
}



class TaskConverter {
  static Task fromProto(model.Task task) {
    return Task(
      uuid: task.uuid,
      name: task.name,
      description: task.description,
      status: task.status,
      estimatedDuration: DurationConverter.fromProtobuf(task.estimatedDuration),
    );
  }

  static model.Task toProto(Task task) {
    return model.Task(
      uuid: task.uuid,
      name: task.name,
      description: task.description,
      status: task.status,
      estimatedDuration: DurationConverter.toProtobuf(task.estimatedDuration ?? Duration.zero),
    );
  }
}


class TaskRepository {
  final TTTServiceClient tttServiceClient;

  TaskRepository({required this.tttServiceClient});

  Future<Task> getTask(String uuid) async {
    print('TaskRepository: Making gRPC call to getTask with uuid: $uuid');
    try {
      final resp = await tttServiceClient.getTask(GetTaskReq(context: context, uuid: uuid));
      print('TaskRepository: gRPC call successful, converting response');
      return TaskConverter.fromProto(resp.task);
    } catch (e) {
      print('TaskRepository: gRPC call failed: $e');
      rethrow;
    }
  }

  Future<Task> createTask(Task task) async {
    print('TaskRepository: Making gRPC call to createTask');
    try {
      final req = CreateTaskReq(
        context: context,
        name: task.name,
        description: task.description,
        status: task.status,
        estimatedDuration: DurationConverter.toProtobuf(task.estimatedDuration ?? Duration.zero),
      );
      final resp = await tttServiceClient.createTask(req);
      print('TaskRepository: createTask gRPC call successful');
      return TaskConverter.fromProto(resp.task);
    } catch (e) {
      print('TaskRepository: createTask gRPC call failed: $e');
      rethrow;
    }
  }
}