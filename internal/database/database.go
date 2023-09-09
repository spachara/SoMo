package database

import (
	"fmt"
	"somo/service/internal/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

func Open(cfg Config) (*sqlx.DB, error) {
	sslmode := "require"
	if cfg.DisableTLS {
		sslmode = "disable"
	}
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)
	return sqlx.Connect("postgres", dataSoruce)
}

func OpenPostgreSQL() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=1234 dbname=first port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.BookChanel{})
	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Reservation{})
	db.AutoMigrate(&model.ReservationDetail{})

	return db, err
}

// INSERT INTO public.rooms(
// 	created_at, updated_at, name, price, hotel_id, room_status, no_guest, extra_price, no_extra_guest)
//    VALUES ( NOW(),  NOW(), 'Sleepy tree', 1800, 1, 0, 2, 700, 1);

// INSERT INTO public.hotels(
// 	created_at, updated_at, name, address, tel)
//    VALUES ( NOW(), NOW(), 'Tubtao Sleepy Hill', '', 000);
