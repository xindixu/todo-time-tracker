import 'package:flutter/material.dart';
import 'package:ui/rpc/client.dart';
import 'package:ui/utils/routes.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await TTTService().init();
  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {

    return MaterialApp.router(
      routerConfig: router,
      title: 'Todo time tracker',
      theme: ThemeData(
        brightness: Brightness.light,
        primaryColor: Colors.purple,
      ),
    );
  }
}