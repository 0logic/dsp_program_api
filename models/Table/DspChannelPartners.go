package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// DSP Channel Partner
type DspChannelPartner struct {
	ChannelPartnerId   uint64    `gorm:"channel_partner_id" json:"channel_partner_id"`
	ChannelPartnerName string    `gorm:"channel_partner_name" json:"channel_partner_name"`
	ChannelPartnerType string    `gorm:"channel_partner_type" json:"channel_partner_type"`
	ChannelPartnerOpen uint      `gorm:"channel_partner_open" json:"channel_partner_open"`
	MediaSpecialistId  uint64    `gorm:"media_specialist_id" json:"media_specialist_id"`
	PartnerPositions   int       `gorm:"partner_positions" json:"partner_positions"`
	Status             uint      `gorm:"status" json:"status"`
	CreatedId          uint64    `gorm:"created_id" json:"created_id"`
	CreatedAt          time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"updated_at" json:"updated_at"`
}

var DspChannelPartnerDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	DspChannelPartnerDB = db
	//defer db.Close()
}

func (DspChannelPartner) TableName() string {
	return "dsp_channel_partners"
}
