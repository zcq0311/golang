package data

import (
	"student/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewStudentRepo, NewGormDB)

// 1.引入 *gorm.DB
type Data struct {
	// TODO wrapped database client
	gormDB *gorm.DB
}

// 2.初始化gorm
func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	dsn := "root:kiyosumishirakawa151?!@tcp(localhost:3306)/kratosUser?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db, err
}

// 3.初始化Data
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{gormDB: db}, cleanup, nil
}
