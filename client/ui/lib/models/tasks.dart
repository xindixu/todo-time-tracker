import 'package:ui/rpc/proto/model.pb.dart' as model;
import 'package:ui/rpc/proto/ttt_service.pbgrpc.dart';
import 'package:ui/utils/converter.dart';


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
    final resp = await tttServiceClient.getTask(GetTaskReq(uuid: uuid));
    return TaskConverter.fromProto(resp.task);
  }

  Future<Task> createTask(Task task) async {
    final req = CreateTaskReq(
      name: task.name,
      description: task.description,
      status: task.status,
      estimatedDuration: DurationConverter.toProtobuf(task.estimatedDuration ?? Duration.zero),
    );
    final resp = await tttServiceClient.createTask(req);
    return TaskConverter.fromProto(resp.task);
  }
}