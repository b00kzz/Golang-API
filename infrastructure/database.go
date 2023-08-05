package infrastructure

import (
	"context"
	"fmt"
	"ticket/goapi/internal/core/port"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		DryRun: false,//เปิดปิดเพื่อทดลองการทำงาน
		Logger: &SqlLogger{},//แสดงคำสั่งของ SQL
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	//การสร้างตารางในฐานข้อมูลด้วยโมเดล
	// db.AutoMigrate(&port.Bill{},)
	// db.AutoMigrate(&port.Customer{})
	// db.AutoMigrate(&port.Role{})
	db.AutoMigrate(&port.Payment{})
	db.AutoMigrate(&port.Review{})
	db.AutoMigrate(&port.User{})
	db.AutoMigrate(&port.Ticket{})
	db.AutoMigrate(&port.UserDetail{})
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n", sql)
}
