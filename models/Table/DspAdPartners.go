package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// DSP AD Partners
type DspAdPartner struct {
	AdPartnerId        uint64    `gorm:"ad_partner_id" json:"ad_partner_id"`
	AdPartnerName      string    `gorm:"ad_partner_name" json:"ad_partner_name"`
	AdPartnerShortName string    `gorm:"ad_partner_short_name" json:"ad_partner_short_name"`
	AdPartnerStyle     string    `gorm:"ad_partner_style" json:"ad_partner_style"`
	AdPartnerType      string    `gorm:"ad_partner_type" json:"ad_partner_type"`
	AdPartnerBalance   uint64    `gorm:"ad_partner_balance" json:"ad_partner_balance"`
	SellerId           uint64    `gorm:"seller_id" json:"seller_id"`
	Status             uint      `gorm:"status" json:"status"`
	CreatedId          uint64    `gorm:"created_id" json:"created_id"`
	CreatedAt          time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"updated_at" json:"updated_at"`
}

var DspAdPartnerDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	DspAdPartnerDB = db
	//defer db.Close()
}

func (DspAdPartner) TableName() string {
	return "dsp_ad_partners"
}
