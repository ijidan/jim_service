// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import "time"

const TableNameGoadminPermissions = "goadmin_permissions"

// GoadminPermissions mapped from table <goadmin_permissions>
type GoadminPermissions struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Slug       string    `gorm:"column:slug;not null" json:"slug"`
	HTTPMethod string    `gorm:"column:http_method" json:"http_method"`
	HTTPPath   string    `gorm:"column:http_path;not null" json:"http_path"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminPermissions's table name
func (*GoadminPermissions) TableName() string {
	return TableNameGoadminPermissions
}
