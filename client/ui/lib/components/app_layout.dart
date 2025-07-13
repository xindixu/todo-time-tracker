import 'package:flutter/material.dart';
import 'package:ui/components/navbar.dart';

class AppLayout extends StatelessWidget {
  const AppLayout({super.key, required this.child});
  final Widget child;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: Row(
          children: [
            const Navbar(),
            Expanded(child: child),
          ],
        ),
      ),
    );
  }
}