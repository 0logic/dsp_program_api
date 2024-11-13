package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// advertiser campaign
type AdvertiserCampaign struct {
	AdvertiserCampaignId uint64
	AdvertiserId         uint64
	CampaignName         string
	MarketingGoal        uint64
	BudgetMode           uint
	Budget               uint64
	CompanyId            uint64
	Status               uint
	CreatedId            uint64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

var AdvertiserCampaignDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	AdvertiserCampaignDB = db
	defer db.Close()
}
