package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// DSP Media Position
type DspMediaPosition struct {
	MediaPositionId      uint64    `gorm:"media_position_id" json:"media_position_id"`
	MediaId              uint64    `gorm:"media_id" json:"media_id"`
	PositionName         string    `gorm:"position_name" json:"position_name"`
	PositionType         string    `gorm:"position_type" json:"position_type"`
	ConnectType          string    `gorm:"connect_type" json:"connect_type"`
	PositionSize         string    `gorm:"position_size" json:"position_size"`
	IsAccurate           uint      `gorm:"is_accurate" json:"is_accurate"`
	PositionMaterialType uint      `gorm:"position_material_type" json:"position_material_type"`
	IsRing               uint      `gorm:"is_ring" json:"is_ring"`
	IsAuto               uint      `gorm:"is_auto" json:"is_auto"`
	AutoTime             uint      `gorm:"auto_time" json:"auto_time"`
	Status               uint      `gorm:"status" json:"status"`
	CreatedId            uint64    `gorm:"created_id" json:"created_id"`
	CreatedAt            time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"updated_at" json:"updated_at"`
}

var DspMediaPositionDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	DspMediaPositionDB = db
	//defer db.Close()
}

func (DspMediaPosition) TableName() string {
	return "dsp_media_positions"
}
