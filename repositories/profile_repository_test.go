package repositories

import (
    "testing"
    
    "github.com/connectplus/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type ProfileRepositoryTestSuite struct {
    suite.Suite
    db *gorm.DB
    repo ProfileRepository
}

func (suite *ProfileRepositoryTestSuite) SetupTest() {
    var err error
    suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(suite.T(), err)
    
    // Migrate the schema for both User and Profile
    err = suite.db.AutoMigrate(&models.User{}, &models.Profile{})
    assert.NoError(suite.T(), err)
    
    suite.repo = NewProfileRepository(suite.db)
}

func (suite *ProfileRepositoryTestSuite) TearDownTest() {
    db, _ := suite.db.DB()
    db.Close()
}

func (suite *ProfileRepositoryTestSuite) TestCreateProfile() {
    profile := &models.Profile{
        UserID:   1,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    
    err := suite.repo.Create(profile)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), profile.ID)
}

func (suite *ProfileRepositoryTestSuite) TestFindByUserID() {
    // Create test profile
    profile := &models.Profile{
        UserID:   1,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    suite.db.Create(profile)
    
    // Test finding the profile
    foundProfile, err := suite.repo.FindByUserID(profile.UserID)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), profile.ID, foundProfile.ID)
    assert.Equal(suite.T(), profile.Bio, foundProfile.Bio)
    
    // Test non-existent user ID
    _, err = suite.repo.FindByUserID(999)
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *ProfileRepositoryTestSuite) TestCreateDuplicateProfile() {
    // Create initial profile
    profile1 := &models.Profile{
        UserID:   1,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    err := suite.repo.Create(profile1)
    assert.NoError(suite.T(), err)

    // Try to create another profile for same user
    profile2 := &models.Profile{
        UserID:   1,
        Bio:      "Another bio",
        Location: "Another Location",
    }
    err = suite.repo.Create(profile2)
    assert.Error(suite.T(), err)
}

func (suite *ProfileRepositoryTestSuite) TestUpdateNonExistentProfile() {
    profile := &models.Profile{
        ID:       999,
        UserID:   999,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    err := suite.repo.Update(profile)
    assert.Error(suite.T(), err)
}

func (suite *ProfileRepositoryTestSuite) TestUpdateProfile() {
    // Create test profile
    profile := &models.Profile{
        UserID:   1,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    suite.db.Create(profile)
    
    // Update profile
    profile.Bio = "Updated bio"
    err := suite.repo.Update(profile)
    assert.NoError(suite.T(), err)
    
    // Verify update
    updatedProfile, err := suite.repo.FindByUserID(profile.UserID)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "Updated bio", updatedProfile.Bio)
}

func (suite *ProfileRepositoryTestSuite) TestDeleteNonExistentProfile() {
    err := suite.repo.Delete(999) // Non-existent user ID
    assert.Error(suite.T(), err)
}

func (suite *ProfileRepositoryTestSuite) TestDeleteProfile() {
    // Create test profile
    profile := &models.Profile{
        UserID:   1,
        Bio:      "Test bio",
        Location: "Test Location",
    }
    suite.db.Create(profile)
    
    // Delete profile
    err := suite.repo.Delete(profile.UserID)
    assert.NoError(suite.T(), err)
    
    // Verify deletion
    _, err = suite.repo.FindByUserID(profile.UserID)
    assert.Error(suite.T(), err)
}

func TestProfileRepositorySuite(t *testing.T) {
    suite.Run(t, new(ProfileRepositoryTestSuite))
}
