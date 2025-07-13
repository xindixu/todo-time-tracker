import 'package:flutter/material.dart';
import 'package:ui/components/tasks/todo.dart';
import 'package:ui/models/tasks.dart';
import 'package:ui/rpc/client.dart';

class ListTasks extends StatefulWidget {
  const ListTasks({super.key});

  @override
  State<ListTasks> createState() => _ListTasksState();
}

class _ListTasksState extends State<ListTasks> {
  late Future<Task> _futureTask;
  late final TaskRepository _taskRepository;

  @override
  void initState() {
    super.initState();
    _taskRepository = TaskRepository(
      tttServiceClient: TTTService().tttServiceClient,
    );
    _futureTask = _taskRepository.getTask(
      '902350bc-3246-4b0d-9768-01abf6e77fa8',
    );
  }

  void _retryLoad() {
    setState(() {
      _futureTask = _taskRepository.getTask(
        '902350bc-3246-4b0d-9768-01abf6e77fa8',
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: FutureBuilder<Task>(
        future: _futureTask,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text('Error: ${snapshot.error}'),
                  const SizedBox(height: 16),
                  ElevatedButton(
                    onPressed: _retryLoad,
                    child: const Text('Retry'),
                  ),
                ],
              ),
            );
          }
          if (!snapshot.hasData) {
            return const Center(child: Text('No task found'));
          }
          final task = snapshot.data!;
          return TaskTodo(task: task, onUpdate: (task) {
            print('ListTasks: onUpdate: $task');
          });
        },
      ),
    );
  }
}
