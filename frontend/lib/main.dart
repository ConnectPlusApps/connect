import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'services/api_service.dart';
import 'providers/auth_provider.dart';
import 'screens/login_screen.dart';

void main() async {
  await dotenv.load(fileName: ".env");
  final apiService = ApiService();

  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => AuthProvider(apiService)),
      ],
      child: MaterialApp(
        title: 'Connect+',
        theme: ThemeData(
          primarySwatch: Colors.pink,
          visualDensity: VisualDensity.adaptivePlatformDensity,
        ),
        home: const LoginScreen(),
        routes: {
          '/login': (ctx) => const LoginScreen(),
        },
      ),
    ),
  );
}
