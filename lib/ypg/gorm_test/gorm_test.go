package gorm_test

import (
	"fmt"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
	"testing"
	"time"
)

type (
	PersonX struct {
		ID        uint32    `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
		Name      string    `gorm:"size:255;not null;unique" json:"Name"`
		Nickname  string    `gorm:"size:255" json:"nickname"`
		Email     string    `gorm:"size:255" json:"email"`
		Password  string    `gorm:"size:255" json:"password"`
		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	}
)

func (PersonX) TableName() string {
	return "person"
}

func NewDBCnt() *gorm.DB {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"127.0.0.1", //viper.GetString("DB_HOST"),
		"root",      //viper.GetString("DB_USER"),
		"test",      //viper.GetString("DB_NAME"),
		"password",  ///viper.GetString("DB_PASS"),
	)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxOpenConns(500)
	return db
}

func Test_cnt(t *testing.T) {
	bean := &PersonX{}

	db := NewDBCnt()

	log.Println("after cnt ")
	db.CreateTable(bean)
	db.AutoMigrate(bean)
	log.Println("create table okt ")
}

func Test_insert(t *testing.T) {
	db := NewDBCnt()

	h := 100000
	var wg sync.WaitGroup
	wg.Add(h)

	for i := 0; i < h; i++ {
		go func() {
			bean := &PersonX{
				Name: fmt.Sprint("aaaaa", i, "55ssssss5", i),
			}

			if err := db.Save(bean).Error; err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("append ok")
			}
		}()
	}

	wg.Wait()

}
