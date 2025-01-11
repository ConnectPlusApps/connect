package repositories

import (
    "testing"
    
    "github.com/connectplus/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type PreferenceRepositoryTestSuite struct {
    suite.Suite
    db *gorm.DB
    repo PreferenceRepository
}

func (suite *PreferenceRepositoryTestSuite) SetupTest() {
    var err error
    suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(suite.T(), err)
    
    // Migrate the schema for both User and Preference
    err = suite.db.AutoMigrate(&models.User{}, &models.Preference{})
    assert.NoError(suite.T(), err)
    
    suite.repo = NewPreferenceRepository(suite.db)
}

func (suite *PreferenceRepositoryTestSuite) TearDownTest() {
    db, _ := suite.db.DB()
    db.Close()
}

func (suite *PreferenceRepositoryTestSuite) TestCreatePreference() {
    preference := &models.Preference{
        UserID:           1,
        MatchDistance:    50,
        MinAge:          25,
        MaxAge:          35,
        NotifyNewMatches: true,
        NotifyMessages:   true,
        ShowOnlineStatus: true,
        ShowLastActive:   true,
        ShowDistance:     true,
    }
    
    err := suite.repo.Create(preference)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), preference.ID)
}

func (suite *PreferenceRepositoryTestSuite) TestFindByUserID() {
    // Create test preference
    preference := &models.Preference{
        UserID:           1,
        MatchDistance:    50,
        MinAge:          25,
        MaxAge:          35,
        NotifyNewMatches: true,
        NotifyMessages:   true,
        ShowOnlineStatus: true,
        ShowLastActive:   true,
        ShowDistance:     true,
    }
    suite.db.Create(preference)
    
    // Test finding preference
    foundPreference, err := suite.repo.FindByUserID(1)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), preference.ID, foundPreference.ID)
    assert.Equal(suite.T(), preference.MatchDistance, foundPreference.MatchDistance)
    
    // Test finding non-existent preference
    _, err = suite.repo.FindByUserID(999)
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *PreferenceRepositoryTestSuite) TestCreateDuplicatePreference() {
    // Create initial preference
    pref1 := &models.Preference{
        UserID:           1,
        MatchDistance:    50,
        MinAge:          25,
        MaxAge:          35,
        NotifyNewMatches: true,
    }
    err := suite.repo.Create(pref1)
    assert.NoError(suite.T(), err)

    // Try to create another preference for same user
    pref2 := &models.Preference{
        UserID:           1,
        MatchDistance:    100,
        MinAge:          20,
        MaxAge:          40,
        NotifyNewMatches: false,
    }
    err = suite.repo.Create(pref2)
    assert.Error(suite.T(), err)
}

func (suite *PreferenceRepositoryTestSuite) TestCreateInvalidPreference() {
    // Test invalid age range
    pref := &models.Preference{
        UserID:        1,
        MatchDistance: 50,
        MinAge:        35,
        MaxAge:        25, // Max age less than min age
    }
    err := suite.repo.Create(pref)
    assert.Error(suite.T(), err)

    // Test invalid match distance
    pref = &models.Preference{
        UserID:        1,
        MatchDistance: -1, // Negative distance
        MinAge:        25,
        MaxAge:        35,
    }
    err = suite.repo.Create(pref)
    assert.Error(suite.T(), err)
}

func (suite *PreferenceRepositoryTestSuite) TestUpdateNonExistentPreference() {
    preference := &models.Preference{
        ID:              999,
        UserID:          999,
        MatchDistance:   50,
        MinAge:         25,
        MaxAge:         35,
        NotifyNewMatches: true,
    }
    err := suite.repo.Update(preference)
    assert.Error(suite.T(), err)
}

func (suite *PreferenceRepositoryTestSuite) TestUpdatePreference() {
    // Create test preference
    preference := &models.Preference{
        UserID:           1,
        MatchDistance:    50,
        MinAge:          25,
        MaxAge:          35,
        NotifyNewMatches: true,
        NotifyMessages:   true,
        ShowOnlineStatus: true,
        ShowLastActive:   true,
        ShowDistance:     true,
    }
    suite.db.Create(preference)
    
    // Update preference
    preference.MatchDistance = 100
    err := suite.repo.Update(preference)
    assert.NoError(suite.T(), err)
    
    // Verify update
    updatedPreference, err := suite.repo.FindByUserID(1)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), 100, updatedPreference.MatchDistance)
}

func (suite *PreferenceRepositoryTestSuite) TestDeleteNonExistentPreference() {
    err := suite.repo.Delete(999)
    assert.Error(suite.T(), err)
}

func (suite *PreferenceRepositoryTestSuite) TestDeletePreference() {
    // Create test preference
    preference := &models.Preference{
        UserID:           1,
        MatchDistance:    50,
        MinAge:          25,
        MaxAge:          35,
        NotifyNewMatches: true,
        NotifyMessages:   true,
        ShowOnlineStatus: true,
        ShowLastActive:   true,
        ShowDistance:     true,
    }
    suite.db.Create(preference)
    
    // Delete preference
    err := suite.repo.Delete(preference.UserID)
    assert.NoError(suite.T(), err)
    
    // Verify deletion
    _, err = suite.repo.FindByUserID(1)
    assert.Error(suite.T(), err)
}

func TestPreferenceRepositorySuite(t *testing.T) {
    suite.Run(t, new(PreferenceRepositoryTestSuite))
}
