import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/components/app_layout.dart';
import 'package:ui/views/landing/main.dart';
import 'package:ui/views/tasks/main.dart';

// GoRouter configuration
final router = GoRouter(
  routes: [
    GoRoute(
      path: '/',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const LandingMain(),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
    GoRoute(
      path: '/tasks',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const TasksMain(),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
    GoRoute(
      path: '/schedule',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const AppLayout(child: Center(child: Text('Schedule'))),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
    GoRoute(
      path: '/track',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const AppLayout(child: Center(child: Text('Track'))),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
    GoRoute(
      path: '/analytics',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const AppLayout(child: Center(child: Text('Analytics'))),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
    GoRoute(
      path: '/settings',
      pageBuilder: (BuildContext context, GoRouterState state) {
        return CustomTransitionPage(
          key: state.pageKey,
          child: const AppLayout(child: Center(child: Text('Settings'))),
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            return FadeTransition(opacity: animation, child: child);
          },
        );
      },
    ),
  ],
);



