package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// advertiser material
type AdvertiserMaterial struct {
	AdvertiserMaterialId uint64    `gorm:"advertiser_material_id" json:"advertiser_material_id"`
	AdvertiserId         uint64    `gorm:"advertiser_id" json:"advertiser_id"`
	MaterialName         string    `gorm:"material_name" json:"material_name"`
	MaterialAliasName    string    `gorm:"material_alias_name" json:"material_alias_name"`
	MaterialExtName      string    `gorm:"material_ext_name" json:"material_ext_name"`
	MaterialUrl          string    `gorm:"material_url" json:"material_url"`
	MaterialWidth        uint      `gorm:"material_width" json:"material_width"`
	MaterialHeight       uint      `gorm:"material_height" json:"material_height"`
	MaterialSize         uint64    `gorm:"material_size" json:"material_size"`
	MaterialSignature    string    `gorm:"material_signature" json:"material_signature"`
	CompanyId            uint64    `gorm:"company_id" json:"company_id"`
	Status               uint      `gorm:"status" json:"status"`
	IsDel                uint      `gorm:"is_del" json:"is_del"`
	CreatedId            uint64    `gorm:"created_id" json:"created_id"`
	CreatedAt            time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"updated_at" json:"updated_at"`
}

var AdvertiserMaterialDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	AdvertiserMaterialDB = db
	defer db.Close()
}
