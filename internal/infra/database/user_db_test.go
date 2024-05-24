package database

import (
	"testing"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Tiago", "tiago@email.com", "123-password")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	require.Nil(t, err)

	var createdUser entity.User

	err = db.First(&createdUser, "id = ?", user.ID).Error
	require.Nil(t, err)
	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Name, user.Name)
	require.Equal(t, createdUser.Email, user.Email)
	require.NotEmpty(t, createdUser.Password)
}
