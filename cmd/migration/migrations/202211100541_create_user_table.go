package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202211100541_create_user_table",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID             uint64    `gorm:"primaryKey"`
				FirstName      string    `valid:"required" gorm:"varchar(100)"`
				LastName       string    `valid:"required" gorm:"varchar(100)"`
				Email          string    `valid:"required" gorm:"varchar(255)"`
				DDD            string    `valid:"-" gorm:"varchar(2)"`
				Phone          string    `valid:"-" gorm:"varchar(9)"`
				UserName       string    `valid:"required" gorm:"varchar(100)"`
				Password       string    `valid:"required" gorm:"-"`
				HashedPassword string    `valid:"required" gorm:"varchar(64)"`
				CreatedAt      time.Time `valid:"-" gorm:"autoCreateTime"`
				UpdatedAt      time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
				LastLogin      time.Time `valid:"-"`
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
