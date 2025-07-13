import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class NavbarItem {
  const NavbarItem({
    required this.title,
    required this.icon,
    required this.route,
  });

  final String title;
  final IconData icon;
  final String route;
}

const navbarItems = [
  NavbarItem(title: 'Tasks', icon: Icons.list, route: '/tasks'),
  NavbarItem(title: 'Schedule', icon: Icons.calendar_today, route: '/schedule'),
  NavbarItem(title: 'Track', icon: Icons.work, route: '/track'),
];

const settingsItem = NavbarItem(title: 'Settings', icon: Icons.settings, route: '/settings');

class Navbar extends StatelessWidget {
  const Navbar({super.key});

  @override
  Widget build(BuildContext context) {
    final currentPath = GoRouterState.of(context).uri.path;

    return Drawer(
      child: Column(
        children: [
          // Top section with main navigation items
          Expanded(
            child: ListView(
              children: navbarItems
                  .map(
                    (item) => ListTile(
                      title: Text(item.title),
                      leading: Icon(item.icon),
                      selected: currentPath.startsWith(item.route),
                      onTap: () => context.go(item.route),
                    ),
                  )
                  .toList(),
            ),
          ),
          ListTile(
            title: Text(settingsItem.title),
            leading: Icon(settingsItem.icon),
            selected: currentPath.startsWith(settingsItem.route),
            onTap: () => context.go(settingsItem.route),
          ),
        ],
      ),
    );
  }
}
