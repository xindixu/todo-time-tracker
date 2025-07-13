import 'package:flutter/material.dart';
import 'package:ui/components/app_layout.dart';
import 'package:ui/views/tasks/list_tasks.dart';

class TabItem {
  const TabItem({required this.tab, required this.child});
  final Tab tab;
  final Widget child;
}

class TasksMain extends StatelessWidget {
  const TasksMain({super.key});

  static const List<TabItem> tabs = <TabItem>[
    TabItem(
      tab: Tab(text: 'List tasks'),
      child: ListTasks(),
    ),
    TabItem(
      tab: Tab(text: 'Link tasks'),
      child: Placeholder(),
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return AppLayout(
      child: DefaultTabController(
        length: tabs.length,
        child: Scaffold(
          appBar: TabBar(
            tabs: tabs.map((TabItem tabItem) => tabItem.tab).toList(),
          ),
          body: TabBarView(
            children: tabs.map((TabItem tabItem) => tabItem.child).toList(),
          ),
        ),
      ),
    );
  }
}
