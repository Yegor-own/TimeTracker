package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"time-tracker/model"

	"github.com/go-faker/faker/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (db *gorm.DB) {
	// Connecting to database
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

func migrateModels(db *gorm.DB) {
	// Migrating User model
	db.AutoMigrate(&model.User{})
	// Generating 6 Users with fake data
	var users = []model.User{
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
		{
			Name:     faker.FirstName(),
			Surname:  faker.FirstName(),
			Address:  fmt.Sprint(faker.GetRealAddress()),
			Passport: int(rand.Uint32()),
		},
	}
	db.Create(&users)

	// Migrating Task model
	db.AutoMigrate(&model.Task{})
	// Generating 6 Tasks with fake data
	var tasks = []model.Task{
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
		{
			Name:        faker.Sentence(),
			Description: faker.Paragraph(),
		},
	}
	db.Create(&tasks)

	// Migrating Track model
	db.AutoMigrate(&model.Track{})
	// Generating 6 Tracks with fake data
	var ts [12]time.Time
	for i := 0; i < 12; i++ {
		t, _ := time.Parse("2006-01-02 15:04:05", faker.Timestamp())
		ts[i] = t
	}
	var tracks = []model.Track{
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[0],
			Stop:   ts[1],
		},
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[2],
			Stop:   ts[3],
		},
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[4],
			Stop:   ts[5],
		},
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[6],
			Stop:   ts[7],
		},
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[8],
			Stop:   ts[9],
		},
		{
			UserID: uint(rand.Int31n(7)),
			TaskID: uint(rand.Int31n(7)),
			Start:  ts[10],
			Stop:   ts[11],
		},
	}
	db.Create(&tracks)
}
