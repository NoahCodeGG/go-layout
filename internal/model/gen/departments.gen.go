// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"time"
)

const TableNameDepartment = "departments"

// Department mapped from table <departments>
type Department struct {
	ID          uint      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:Primary key" json:"id"`                                                // Primary key
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now();comment:Creation timestamp" json:"created_at"`              // Creation timestamp
	Name        string    `gorm:"column:name;type:character varying(100);not null;uniqueIndex:idx_departments_name,priority:1;comment:Department name" json:"name"` // Department name
	Description string    `gorm:"column:description;type:text;not null;comment:Department description" json:"description"`                                          // Department description
	Enabled     bool      `gorm:"column:enabled;type:boolean;not null;default:true;comment:Department enabled status" json:"enabled"`                               // Department enabled status
}

// TableName Department's table name
func (*Department) TableName() string {
	return TableNameDepartment
}
