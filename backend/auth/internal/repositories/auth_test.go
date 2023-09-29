package auth_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
	auth "github.com/zhosyaaa/foodDeliverySystems-auth/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "user=yourusername password=yourpassword dbname=test_auth sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}

func clearTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}

func TestAuthService_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	authService := auth.NewAuthService(db)

	testUser := models.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}

	createdUser, err := authService.CreateUser(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdUser.ID)

	// Проверьте, что пользователь сохранен в базе данных
	var retrievedUser models.User
	db.First(&retrievedUser, createdUser.ID)
	assert.Equal(t, testUser.Username, retrievedUser.Username)
	assert.Equal(t, testUser.Password, retrievedUser.Password)
	assert.Equal(t, testUser.Email, retrievedUser.Email)
}

func TestAuthService_GetUserById(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	authService := auth.NewAuthService(db)

	testUser := models.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}
	db.Create(&testUser)

	retrievedUser, err := authService.GetUserById(int(testUser.ID))
	assert.NoError(t, err)
	assert.Equal(t, testUser.Username, retrievedUser.Username)
	assert.Equal(t, testUser.Password, retrievedUser.Password)
	assert.Equal(t, testUser.Email, retrievedUser.Email)
}

func TestAuthService_GetUserByEmail(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	authService := auth.NewAuthService(db)

	testUser := models.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}
	db.Create(&testUser)

	retrievedUser, err := authService.GetUserByEmail(testUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Username, retrievedUser.Username)
	assert.Equal(t, testUser.Password, retrievedUser.Password)
	assert.Equal(t, testUser.Email, retrievedUser.Email)
}
