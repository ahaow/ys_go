package model

import "time"

type Helper struct {
	BaseModel
	Name      string     `gorm:"type:varchar(100);not null;comment:姓名"`     // 姓名
	StartTime *time.Time `gorm:"type:datetime;comment:开始时间"`                // 可空开始时间
	EndTime   *time.Time `gorm:"type:datetime;comment:结束时间"`                // 可空结束时间
	Intro     string     `gorm:"type:varchar(200);comment:简介"`              // 简介
	Role      int        `gorm:"type:int;default:1;comment:'1表示需要,2表示不需要'"` // 登录需求
	Status    int        `gorm:"type:int;default:1;comment:'1表示启用,2表示停用'"`  // 登录需求
	Address   string     `gorm:"type:varchar(100);comment:地址"`              // 地址
}
