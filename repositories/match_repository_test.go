package repositories

import (
    "testing"
    
    "github.com/connectplus/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type MatchRepositoryTestSuite struct {
    suite.Suite
    db *gorm.DB
    repo MatchRepository
}

func (suite *MatchRepositoryTestSuite) SetupTest() {
    var err error
    suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(suite.T(), err)
    
    // Migrate the schema for both User and Match
    err = suite.db.AutoMigrate(&models.User{}, &models.Match{})
    assert.NoError(suite.T(), err)
    
    suite.repo = NewMatchRepository(suite.db)
}

func (suite *MatchRepositoryTestSuite) TearDownTest() {
    db, _ := suite.db.DB()
    db.Close()
}

func (suite *MatchRepositoryTestSuite) TestCreateMatch() {
    match := &models.Match{
        User1ID: 1,
        User2ID: 2,
        Status:  models.MatchPending,
    }
    
    err := suite.repo.Create(match)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), match.ID)
}

func (suite *MatchRepositoryTestSuite) TestFindByUserID() {
    // Create test matches where user is User1ID
    match1 := &models.Match{User1ID: 1, User2ID: 2, Status: models.MatchPending}
    match2 := &models.Match{User1ID: 1, User2ID: 3, Status: models.MatchAccepted}
    suite.db.Create(match1)
    suite.db.Create(match2)
    
    // Create test match where user is User2ID
    match3 := &models.Match{User1ID: 4, User2ID: 1, Status: models.MatchPending}
    suite.db.Create(match3)
    
    // Test finding all matches for user
    matches, err := suite.repo.FindByUserID(1)
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), matches, 3)
    
    // Test finding matches for non-existent user
    matches, err = suite.repo.FindByUserID(999)
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), matches, 0)
}

func (suite *MatchRepositoryTestSuite) TestFindByUsers() {
    // Create test match
    match := &models.Match{User1ID: 1, User2ID: 2, Status: models.MatchPending}
    suite.db.Create(match)
    
    // Test finding match with original order
    foundMatch, err := suite.repo.FindByUsers(1, 2)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), match.ID, foundMatch.ID)
    
    // Test finding match with reverse order
    foundMatch, err = suite.repo.FindByUsers(2, 1)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), match.ID, foundMatch.ID)
    
    // Test finding non-existent match
    _, err = suite.repo.FindByUsers(1, 999)
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *MatchRepositoryTestSuite) TestCreateDuplicateMatch() {
    // Create initial match
    match1 := &models.Match{
        User1ID: 1,
        User2ID: 2,
        Status:  models.MatchPending,
    }
    err := suite.repo.Create(match1)
    assert.NoError(suite.T(), err)

    // Try to create duplicate match with same users
    match2 := &models.Match{
        User1ID: 1,
        User2ID: 2,
        Status:  models.MatchPending,
    }
    err = suite.repo.Create(match2)
    assert.Error(suite.T(), err)

    // Try to create duplicate match with reversed users
    match3 := &models.Match{
        User1ID: 2,
        User2ID: 1,
        Status:  models.MatchPending,
    }
    err = suite.repo.Create(match3)
    assert.Error(suite.T(), err)
}

func (suite *MatchRepositoryTestSuite) TestUpdateStatusNonExistentMatch() {
    err := suite.repo.UpdateStatus(999, models.MatchAccepted)
    assert.Error(suite.T(), err)
}

func (suite *MatchRepositoryTestSuite) TestUpdateStatus() {
    // Create test match
    match := &models.Match{User1ID: 1, User2ID: 2, Status: models.MatchPending}
    suite.db.Create(match)
    
    // Update status
    err := suite.repo.UpdateStatus(match.ID, models.MatchAccepted)
    assert.NoError(suite.T(), err)
    
    // Verify update
    updatedMatch, err := suite.repo.FindByUsers(1, 2)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), models.MatchAccepted, updatedMatch.Status)
}

func (suite *MatchRepositoryTestSuite) TestDeleteNonExistentMatch() {
    err := suite.repo.Delete(999)
    assert.Error(suite.T(), err)
}

func (suite *MatchRepositoryTestSuite) TestDeleteMatch() {
    // Create test match
    match := &models.Match{User1ID: 1, User2ID: 2, Status: models.MatchPending}
    suite.db.Create(match)
    
    // Delete match
    err := suite.repo.Delete(match.ID)
    assert.NoError(suite.T(), err)
    
    // Verify deletion
    _, err = suite.repo.FindByUsers(1, 2)
    assert.Error(suite.T(), err)
}

func TestMatchRepositorySuite(t *testing.T) {
    suite.Run(t, new(MatchRepositoryTestSuite))
}
