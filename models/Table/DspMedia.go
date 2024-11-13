package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// DSP Media
type DspMedia struct {
	MediaId       uint64    `gorm:"media_id" json:"media_id"`
	MediaName     string    `gorm:"media_name" json:"media_name"`
	MediaIndustry string    `gorm:"media_industry" json:"media_industry"`
	MediaType     string    `gorm:"media_type" json:"media_type"`
	MediaCampaign string    `gorm:"media_campaign" json:"media_campaign"`
	MediaAdCount  int       `gorm:"media_ad_count" json:"media_ad_count"`
	IsSupport     uint      `gorm:"is_support" json:"is_support"`
	Status        uint      `gorm:"status" json:"status"`
	CreatedId     uint64    `gorm:"created_id" json:"created_id"`
	CreatedAt     time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
}

var DspMediaDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	DspMediaDB = db
	//defer db.Close()
}

func (DspMedia) TableName() string {
	return "dsp_medias"
}
