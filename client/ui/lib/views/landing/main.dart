import 'package:flutter/material.dart';
import 'package:ui/components/app_layout.dart';

class LandingMain extends StatelessWidget {
  const LandingMain({super.key});

  @override
  Widget build(BuildContext context) {
    return const AppLayout(
      child: Center(
        child: Text('Hello World'),
      ),
    );
  }
}