package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// advertiser guest
type AdvertiserGuest struct {
	AdvertiserGuestId uint64
	GuestName         string
	GuestEmail        string
	GuestPhone        string
	GuestCompany      string
	Country           string
	Reason            string
	GuestIp           string
	RelateId          uint64
	Status            uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

var AdvertiserGuestDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	AdvertiserGuestDB = db
	//defer db.Close()
}

// 通过TableName方法指定表名
func (AdvertiserGuest) TableName() string {
	return "advertiser_guests"
}
