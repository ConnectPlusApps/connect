package repositories

import (
    "testing"
    
    "github.com/connectplus/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
    suite.Suite
    db *gorm.DB
    repo UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
    var err error
    suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(suite.T(), err)
    
    // Migrate the schema
    err = suite.db.AutoMigrate(&models.User{})
    assert.NoError(suite.T(), err)
    
    suite.repo = NewUserRepository(suite.db)
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
    db, _ := suite.db.DB()
    db.Close()
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
    user := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    
    err := suite.repo.Create(user)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), user.ID)
}

func (suite *UserRepositoryTestSuite) TestFindByID() {
    // Create test user
    user := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    suite.db.Create(user)
    
    // Test finding the user
    foundUser, err := suite.repo.FindByID(user.ID)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), user.ID, foundUser.ID)
    assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepositoryTestSuite) TestFindByEmail() {
    // Create test user
    user := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    suite.db.Create(user)
    
    // Test finding the user
    foundUser, err := suite.repo.FindByEmail(user.Email)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), user.ID, foundUser.ID)
    assert.Equal(suite.T(), user.Email, foundUser.Email)

    // Test non-existent email
    _, err = suite.repo.FindByEmail("nonexistent@example.com")
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *UserRepositoryTestSuite) TestCreateUserDuplicateEmail() {
    // Create initial user
    user1 := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    err := suite.repo.Create(user1)
    assert.NoError(suite.T(), err)

    // Try to create user with same email
    user2 := &models.User{
        Email:        "test@example.com",
        PasswordHash: "different",
    }
    err = suite.repo.Create(user2)
    assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestUpdateNonExistentUser() {
    user := &models.User{
        ID:           999, // Non-existent ID
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    err := suite.repo.Update(user)
    assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestUpdateUser() {
    // Create test user
    user := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    suite.db.Create(user)
    
    // Update user
    user.Email = "updated@example.com"
    err := suite.repo.Update(user)
    assert.NoError(suite.T(), err)
    
    // Verify update
    updatedUser, err := suite.repo.FindByID(user.ID)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "updated@example.com", updatedUser.Email)
}

func (suite *UserRepositoryTestSuite) TestDeleteNonExistentUser() {
    err := suite.repo.Delete(999) // Non-existent ID
    assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestDeleteUser() {
    // Create test user
    user := &models.User{
        Email:        "test@example.com",
        PasswordHash: "testpass",
    }
    suite.db.Create(user)
    
    // Delete user
    err := suite.repo.Delete(user.ID)
    assert.NoError(suite.T(), err)
    
    // Verify deletion
    _, err = suite.repo.FindByID(user.ID)
    assert.Error(suite.T(), err)
}

func TestUserRepositorySuite(t *testing.T) {
    suite.Run(t, new(UserRepositoryTestSuite))
}
