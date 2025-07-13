import 'package:flutter/material.dart';
import 'package:ui/models/tasks.dart';
import 'package:ui/rpc/proto/model.pbenum.dart';
import 'package:ui/components/tasks/status_dropdown.dart';

class TaskTodo extends StatelessWidget {
  const TaskTodo({super.key, required this.task, required this.onUpdate});
  final Task task;
  final Function(Task) onUpdate;

  IconData statusIcon(Task_Status status) {
    switch (status) {
      case Task_Status.TODO:
        return Icons.check_box;
      case Task_Status.IN_PROGRESS:
        return Icons.indeterminate_check_box;
      case Task_Status.BLOCKED:
        return Icons.block;
      case Task_Status.DONE:
        return Icons.check_box_outline_blank;
      default:
        return Icons.check_box_outline_blank;
    }
  }

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(task.name),
      subtitle: task.description.isNotEmpty ? Text(task.description) : null,
      // trailing: Flexible(
      //   child: Row(
      //     children: [
      //       StatusDropdown(
      //         status: task.status,
      //         onChange: (status) {
      //           onUpdate(task.copyWith(status: status));
      //         },
      //       ),
      //     ],
      //   ),
      // ),
    );
  }
}
