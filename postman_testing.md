# Connect+ API Testing with Postman

## Setup Instructions

1. Install Postman from https://www.postman.com/downloads/
2. Import the Connect+ API collection (see below)
3. Set up environment variables in Postman:
   - `base_url`: http://localhost:8080
   - `jwt_token`: (leave blank, will be set after login)

## API Endpoints

### 1. Create User
- **Method**: POST
- **URL**: `{{base_url}}/user/create`
- **Headers**:
  - Content-Type: application/json
- **Body** (raw JSON):
```json
{
  "username": "testuser",
  "email": "test@example.com",
  "password": "testpass123"
}
Expected Response:
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
2. Get Server Status
Method: GET
URL: {{base_url}}/
Expected Response:
{
  "message": "Welcome to Connect+! The API is running!"
}
3. Test User Endpoint
Method: GET
URL: {{base_url}}/user
Expected Response:
{
  "message": "User endpoint hit!"
}
Testing Tips
Start the server: go run main.go
Use environment variables to manage different configurations
Save successful responses as examples
Test error cases with invalid inputs

Please manually create/update these files since the automated tools are failing. Let me know if you need any clarification or assistance with the manual implementation.