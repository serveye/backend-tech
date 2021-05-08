package config

import (
	"fmt"
	"github.com/serveye/backend-tech/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)
var DB *gorm.DB
func ConnectDataBase() {

	dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgres","docker", "dev_db", 5432)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	log.Println("database connected", DB)
	err = db.AutoMigrate(&models.Category{})
	if err != nil {
		log.Println(err)
	}
	DB = db

}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		//limit := 2
		page := 1
		//sort := "created_at asc"
		query := r.URL.Query()
		pageSize := 10
		for key, value := range query {
			queryValue := value[len(value)-1]
			switch key {
			//case "limit":
			//	limit, _ = strconv.Atoi(queryValue)
			//	break
			case "page":
				page, _ = strconv.Atoi(queryValue)
				break
			//case "sort":
			//	sort = queryValue
			//	break
			case "pageSize":
				pageSize, _ = strconv.Atoi(queryValue)
				break

			}
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

