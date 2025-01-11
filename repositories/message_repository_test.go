package repositories

import (
    "testing"
    "time"
    
    "github.com/connectplus/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type MessageRepositoryTestSuite struct {
    suite.Suite
    db *gorm.DB
    repo MessageRepository
}

func (suite *MessageRepositoryTestSuite) SetupTest() {
    var err error
    suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(suite.T(), err)
    
    // Migrate the schema for both User and Message
    err = suite.db.AutoMigrate(&models.User{}, &models.Message{})
    assert.NoError(suite.T(), err)
    
    suite.repo = NewMessageRepository(suite.db)
}

func (suite *MessageRepositoryTestSuite) TearDownTest() {
    db, _ := suite.db.DB()
    db.Close()
}

func (suite *MessageRepositoryTestSuite) TestCreateMessage() {
    message := &models.Message{
        SenderID:   1,
        ReceiverID: 2,
        Content:    "Test message",
    }
    
    err := suite.repo.Create(message)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), message.ID)
}

func (suite *MessageRepositoryTestSuite) TestGetEmptyConversation() {
    messages, err := suite.repo.GetConversation(1, 2)
    assert.NoError(suite.T(), err)
    assert.Empty(suite.T(), messages)
}

func (suite *MessageRepositoryTestSuite) TestGetConversationOrdered() {
    // Create test messages with specific timestamps
    message1 := &models.Message{
        SenderID:   1,
        ReceiverID: 2,
        Content:    "First message",
        CreatedAt:  time.Now().Add(-2 * time.Hour),
    }
    message2 := &models.Message{
        SenderID:   2,
        ReceiverID: 1,
        Content:    "Second message",
        CreatedAt:  time.Now().Add(-1 * time.Hour),
    }
    message3 := &models.Message{
        SenderID:   1,
        ReceiverID: 2,
        Content:    "Third message",
        CreatedAt:  time.Now(),
    }
    
    // Create in random order
    suite.db.Create(message2)
    suite.db.Create(message3)
    suite.db.Create(message1)
    
    // Test getting conversation
    messages, err := suite.repo.GetConversation(1, 2)
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), messages, 3)
    
    // Verify order
    assert.Equal(suite.T(), "First message", messages[0].Content)
    assert.Equal(suite.T(), "Second message", messages[1].Content)
    assert.Equal(suite.T(), "Third message", messages[2].Content)
}

func (suite *MessageRepositoryTestSuite) TestGetConversation() {
    // Create test messages in both directions
    messages := []*models.Message{
        {SenderID: 1, ReceiverID: 2, Content: "Hello"},
        {SenderID: 2, ReceiverID: 1, Content: "Hi"},
        {SenderID: 1, ReceiverID: 2, Content: "How are you?"},
        {SenderID: 2, ReceiverID: 1, Content: "Good, thanks!"},
    }
    
    for _, msg := range messages {
        suite.db.Create(msg)
    }
    
    // Test getting conversation
    conversation, err := suite.repo.GetConversation(1, 2)
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), conversation, 4)
    
    // Test getting conversation with reversed user order
    conversationReversed, err := suite.repo.GetConversation(2, 1)
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), conversationReversed, 4)
}

func (suite *MessageRepositoryTestSuite) TestMarkNonExistentMessageAsRead() {
    err := suite.repo.MarkAsRead(999)
    assert.Error(suite.T(), err)
}

func (suite *MessageRepositoryTestSuite) TestMarkAsRead() {
    // Create test message
    message := &models.Message{SenderID: 1, ReceiverID: 2, Content: "Test message"}
    suite.db.Create(message)
    
    // Mark as read
    err := suite.repo.MarkAsRead(message.ID)
    assert.NoError(suite.T(), err)
    
    // Verify update
    updatedMessage, err := suite.repo.GetConversation(1, 2)
    assert.NoError(suite.T(), err)
    assert.True(suite.T(), updatedMessage[0].IsRead)
}

func (suite *MessageRepositoryTestSuite) TestDeleteNonExistentMessage() {
    err := suite.repo.Delete(999)
    assert.Error(suite.T(), err)
}

func (suite *MessageRepositoryTestSuite) TestDeleteMessage() {
    // Create test message
    message := &models.Message{SenderID: 1, ReceiverID: 2, Content: "Test message"}
    suite.db.Create(message)
    
    // Delete message
    err := suite.repo.Delete(message.ID)
    assert.NoError(suite.T(), err)
    
    // Verify deletion
    messages, err := suite.repo.GetConversation(1, 2)
    assert.NoError(suite.T(), err)
    assert.Empty(suite.T(), messages)
}

func TestMessageRepositorySuite(t *testing.T) {
    suite.Run(t, new(MessageRepositoryTestSuite))
}
