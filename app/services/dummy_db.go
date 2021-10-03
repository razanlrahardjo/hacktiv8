package services

// import (
// 	"github.com/razanlrahardjo/hacktiv8/app/migrations"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // DBConnectTest test
// func DBConnectTest() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	if nil != err {
// 		panic(err)
// 	}

// 	err = db.AutoMigrate(migrations.ModelMigrations...)
// 	if nil != err {
// 		panic(err)
// 	}
// 	DB = db

// 	return db
// }
