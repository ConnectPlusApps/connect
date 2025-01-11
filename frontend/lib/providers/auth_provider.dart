import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../services/api_service.dart';

class AuthProvider with ChangeNotifier {
  final ApiService _apiService;
  String? _token;
  bool _isAuthenticated = false;

  AuthProvider(this._apiService) {
    _initializeAuth();
  }

  bool get isAuthenticated => _isAuthenticated;

  Future<void> _initializeAuth() async {
    final prefs = await SharedPreferences.getInstance();
    _token = prefs.getString('auth_token');
    _isAuthenticated = _token != null;
    notifyListeners();
  }

  Future<void> login(String email, String password) async {
    try {
      final response = await _apiService.post('user/login', {
        'email': email,
        'password': password,
      });

      _token = response['token'];
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('auth_token', _token!);
      _isAuthenticated = true;
      notifyListeners();
    } catch (error) {
      throw Exception('Failed to login: ${error.toString()}');
    }
  }

  Future<void> logout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('auth_token');
    _token = null;
    _isAuthenticated = false;
    notifyListeners();
  }

  String? get token => _token;
}
