package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// advertiser creative
type AdvertiserCreative struct {
	AdvertiserCreativeId uint64
	AdvertiserCampaignId uint64
	AdvertiserId         uint64
	CreativeName         string
	Type                 uint
	BidType              uint
	BidPrice             uint64
	CreativeTitle        string
	VideoMaterialId      uint64
	ImageMaterialId      uint64
	CompanyId            uint64
	Status               uint
	CreatedId            uint64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

var AdvertiserCreativeDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	AdvertiserCreativeDB = db
	defer db.Close()
}
