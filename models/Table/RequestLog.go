package Table

import (
	"dsp_program_api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// request log
type RequestLog struct {
	ID          uint64
	Name        string
	Desc        string
	RequestBody string
	Ip          string
	Method      string
	CreatedAt   time.Time
}

var RequestLogDB *gorm.DB

func init() {
	// connect mysql
	db, err := gorm.Open("mysql", config.MYSQL_DEFAULT_DB)
	if err != nil {
		panic(err)
	}

	RequestLogDB = db
}
