// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"time"

	"github.com/lib/pq"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	ID          uint           `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:Primary key" json:"id"`                                    // Primary key
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now();comment:Creation timestamp" json:"created_at"`  // Creation timestamp
	Name        string         `gorm:"column:name;type:character varying(100);not null;uniqueIndex:idx_roles_name,priority:1;comment:Role name" json:"name"` // Role name
	Slug        string         `gorm:"column:slug;type:character varying(100);not null;comment:Role slug" json:"slug"`                                       // Role slug
	Description string         `gorm:"column:description;type:text;not null;comment:Role description" json:"description"`                                    // Role description
	Enabled     bool           `gorm:"column:enabled;type:boolean;not null;default:true;comment:Role enabled status" json:"enabled"`                         // Role enabled status
	Permissions pq.StringArray `gorm:"column:permissions;type:character varying(100)[];not null;comment:Role permissions" json:"permissions"`                // Role permissions
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
