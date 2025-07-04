import 'package:flutter/material.dart';
import 'package:ui/models/tasks.dart';

class TaskLists extends StatelessWidget {
  const TaskLists({super.key, required this.taskRepository});

  final TaskRepository taskRepository;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Task Lists'),
      ),
      body: ListView.builder(
        itemBuilder: (context, index) {
          return ListTile(
            title: Text('Task ${index + 1}'),
          );
        },
      ),
    );
  }
}