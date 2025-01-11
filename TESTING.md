# Connect+ Testing Documentation

## Running Tests

### Prerequisites
- Go 1.20+ installed
- SQLite3 (for in-memory testing)
- Testify package installed (`go get github.com/stretchr/testify`)

### Running All Tests
```bash
go test ./... -v
```

### Running Specific Test Suites
```bash
# Example: Run user repository tests
go test ./repositories -run TestUserRepositorySuite -v
```

### Test Coverage
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Organization

### Test Suites
Each repository has its own test suite with comprehensive test coverage:

#### UserRepositoryTestSuite
- Basic CRUD operations
- Error handling for non-existent users
- Duplicate email prevention
- Foreign key constraints with User model
- Input validation testing

#### ProfileRepositoryTestSuite
- Basic CRUD operations
- One-to-one relationship with users
- Duplicate profile prevention
- Error cases for non-existent profiles/users
- Foreign key constraints with User model

#### MatchRepositoryTestSuite
- Basic CRUD operations
- Bidirectional relationship testing (user1_id/user2_id)
- Match status updates and validation
- Error cases for non-existent matches
- Foreign key constraints with User model
- Duplicate match prevention

#### MessageRepositoryTestSuite
- Basic CRUD operations
- Conversation retrieval in both directions
- Message ordering by timestamp
- Read status updates
- Error cases for non-existent messages
- Foreign key constraints with User model

#### PreferenceRepositoryTestSuite
- Basic CRUD operations
- One-to-one relationship with users
- Validation of preference fields (age ranges, distances)
- Error cases for non-existent preferences
- Duplicate preference prevention
- Foreign key constraints with User model

### Test Structure
Each test suite follows this pattern:
1. `SetupTest()` - Initializes in-memory SQLite database with proper migrations
2. Test methods for success cases (e.g., `TestCreateUser`)
3. Test methods for error cases (e.g., `TestCreateDuplicateUser`)
4. Test methods for edge cases (e.g., `TestUpdateNonExistentUser`)
5. `TearDownTest()` - Cleans up test database

## Common Test Commands

### Run Tests with Race Detector
```bash
go test ./... -race
```

### Generate Test Coverage HTML
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Test Environment

### In-Memory SQLite Database
- Tests use an in-memory SQLite database
- Database is automatically migrated with all required models
- Foreign key constraints are enforced
- Database is cleaned up after each test

### Test Data
- Test data is created within each test method
- Data includes both valid and invalid test cases
- Edge cases are explicitly tested
- Data is automatically cleaned up after each test

## Troubleshooting

### Common Issues
1. **Database connection errors**: Ensure SQLite3 is installed
2. **Test failures**: Check for concurrent test runs
3. **Missing dependencies**: Run `go mod tidy`
4. **Foreign key violations**: Ensure proper model migrations
5. **Validation errors**: Check model constraints

### Debugging Tests
To debug a failing test:
1. Add `t.Log()` statements
2. Run with `-v` flag for verbose output
3. Use `go test -run` to isolate specific tests
4. Check error messages for constraint violations
5. Verify test data setup and cleanup
