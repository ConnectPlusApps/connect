import 'package:http/http.dart' as http;
import 'dart:convert';

class ApiService {
  final String baseUrl = 'https://8080-idx-connect-1736511427003.cluster-fu5knmr55rd44vy7k7pxk74ams.cloudworkstations.dev';

  Future<dynamic> post(String endpoint, Map<String, dynamic> body) async {
    final response = await http.post(
      Uri.parse('$baseUrl/$endpoint'),
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonEncode(body),
    );

    return _handleResponse(response);
  }

  Future<dynamic> get(String endpoint, {Map<String, String>? headers}) async {
    final response = await http.get(
      Uri.parse('$baseUrl/$endpoint'),
      headers: headers,
    );

    return _handleResponse(response);
  }

  dynamic _handleResponse(http.Response response) {
    if (response.statusCode >= 200 && response.statusCode < 300) {
      return jsonDecode(response.body);
    } else {
      final errorBody = jsonDecode(response.body);
      throw Exception(errorBody is String ? errorBody : errorBody['message'] ?? 'Request failed');
    }
  }
}
