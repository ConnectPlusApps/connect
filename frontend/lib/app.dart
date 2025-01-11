import 'package:flutter/material.dart';
import 'screens/login_screen.dart';

class ConnectApp extends StatelessWidget {
  const ConnectApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Connect+',
      theme: ThemeData(
        primarySwatch: Colors.pink,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: const LoginScreen(),
      routes: {
        '/login': (ctx) => const LoginScreen(),
      },
    );
  }
}
