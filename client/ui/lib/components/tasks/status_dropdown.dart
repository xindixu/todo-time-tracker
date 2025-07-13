import 'package:flutter/material.dart';
import 'package:ui/rpc/proto/model.pbenum.dart';

class StatusDropdown extends StatelessWidget {
  const StatusDropdown({
    super.key,
    required this.status,
    required this.onChange,
  });
  final Task_Status status;
  final Function(Task_Status) onChange;

  @override
  Widget build(BuildContext context) {
    return DropdownButton<Task_Status>(
      value: status,
      iconSize: 24,
      isDense: true,
      icon: const Icon(Icons.arrow_drop_down),
      onChanged: (Task_Status? value) {
        onChange(value!);
      },
      items: Task_Status.values
          .where((e) => e != Task_Status.STATUS_NONE)
          .map<DropdownMenuItem<Task_Status>>((Task_Status value) {
            return DropdownMenuItem<Task_Status>(
              value: value,
              child: Text(value.name, style: const TextStyle(fontSize: 12)),
            );
          })
          .toList(),
    );
  }
}
