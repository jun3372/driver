package tests

import (
	"testing"
	"time"
)

func TestMakeUser(t *testing.T) {
	_ = t.Run("make db", TestMakeDB)
	t.Run("make user table", func(t *testing.T) {
		if err := db.AutoMigrate(&User{}); err != nil {
			t.Fatal(err)
		}
		t.Parallel()
	})
	t.Run("make user save or update", func(t *testing.T) {
		var u = User{Username: "渣渣辉"}
		if err := db.Create(&u).Error; err != nil {
			t.Fatal(err)
		}

		u.Password = "123456-" + time.Now().String()
		if err := db.Omit("id").Where("id", u.ID).Updates(u).Error; err != nil {
			t.Fatal(err)
		}

		t.Parallel()
	})

	t.Run("select", func(t *testing.T) {
		var (
			u     []User
			end   = time.Now()
			start = end.Add(-time.Hour)
			where = map[string]interface{}{
				"created_at": []string{">", start.String()},
			}
		)

		if err := db.Where(where).Find(&u).Error; err != nil {
			t.Fatal(err)
		}

		t.Log("users", u)
		t.Parallel()
	})
}
