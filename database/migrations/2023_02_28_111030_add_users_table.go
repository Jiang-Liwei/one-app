package migrations

import (
    "database/sql"
    "forum/app/models"
    "forum/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type User struct {
        models.BaseModel

        Name     string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.TimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        err := migrator.AutoMigrate(&User{})
        if err != nil {
            return
        }
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        err := migrator.DropTable(&User{})
        if err != nil {
            return
        }
    }

    migrate.Add("2023_02_28_111030_add_users_table", up, down)
}
