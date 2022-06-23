package sql

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

const ShortForm = "2006-01-02 15:04:05"

func DbWhereSearchDateRange(db *gorm.DB, info request.SearchDateRangeInfo, tableName string) *gorm.DB {
	if tableName != "" {
		tableName = tableName + "."
	}

	db = DbWhereDateRange(db, tableName+"created_at", info.CreatedAtStart, info.CreatedAtEnd)
	db = DbWhereDateRange(db, tableName+"updated_at", info.UpdatedAtStart, info.UpdatedAtEnd)
	return db
}

func DbWhereDateRange(db *gorm.DB, key string, start time.Time, end time.Time) *gorm.DB {

	if start.IsZero() {
		return db
	}

	if end.IsZero() {
		return db
	}

	return db.Where(
		fmt.Sprintf("%s >= ? and %s <= ?", key, key),
		start.Format(ShortForm),
		end.Format(ShortForm),
	)
}
