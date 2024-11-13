package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// advertiser account
type AdvertiserAccount struct {
	AdvertiserId uint64
	Name         string
	CompanyId    uint64
	Status       uint
	CreatedId    uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var AdvertiserAccountDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	AdvertiserAccountDB = db
	//defer db.Close()
}
