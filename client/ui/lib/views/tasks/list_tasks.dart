import 'package:flutter/material.dart';
import 'package:ui/models/tasks.dart';

class ListTasks extends StatelessWidget {
  // const ListTasks({super.key, required this.taskRepository});
  // final TaskRepository taskRepository;

  const ListTasks({super.key});


  @override
  Widget build(BuildContext context) {
    return Scaffold(
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