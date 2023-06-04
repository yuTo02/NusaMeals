package main

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"reglog/internal/common/config"
	"reglog/seeder/data"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: data.UserFaker(db)},
	}
}

func main() {
	config.InitMySQLDev()
	db := config.DB
	for _, seeder := range RegisterSeeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			logrus.Fatal("Cannot run seeder")
		} else {
			logrus.Info("Success run seeder")
		}
	}
}
