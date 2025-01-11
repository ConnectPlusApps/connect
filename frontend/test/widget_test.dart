import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/screens/login_screen.dart';

void main() {
  testWidgets('Login screen smoke test', (WidgetTester tester) async {
    await tester.pumpWidget(const MaterialApp(
      home: LoginScreen(),
    ));

    expect(find.text('Connect+'), findsOneWidget);
    expect(find.text('Login'), findsOneWidget);
    expect(find.text('Create an account'), findsOneWidget);
  });
}
