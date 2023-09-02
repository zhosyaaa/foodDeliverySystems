package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "postgresql://postgres:1079localhost:8080/auth"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}

func clearTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}

func TestAuthService_RegisterUser(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	authService := NewAuthService(db)

	username := "testuser"
	password := "testpassword"

	err := authService.RegisterUser(username, password)
	assert.NoError(t, err)

	var user models.User
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "username = ?", username)
	assert.NoError(t, result.Error)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, password, user.Password)
}

func TestAuthService_AuthenticateUser(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	authService := NewAuthService(db)

	username := "testuser"
	password := "testpassword"

	err := authService.RegisterUser(username, password)
	assert.NoError(t, err)

	authenticated, err := authService.AuthenticateUser(username, password)
	assert.NoError(t, err)
	assert.True(t, authenticated)

	authenticated, err = authService.AuthenticateUser(username, "wrongpassword")
	assert.NoError(t, err)
	assert.False(t, authenticated)

	authenticated, err = authService.AuthenticateUser("nonexistentuser", "password")
	assert.NoError(t, err)
	assert.False(t, authenticated)
}
