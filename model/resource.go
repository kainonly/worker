package model

type Resource struct {
	ID     uint64 `json:"id"`
	Status *bool  `gorm:"default:true" json:"status"`
	Parent string `gorm:"type:varchar(50);default:'root';comment:父节点" json:"parent"`
	Path   string `gorm:"type:varchar(50);unique;not null;comment:路径" json:"path"`
	Name   string `gorm:"type:varchar(20);not null;comment:资源名称" json:"name"`
	Nav    *bool  `gorm:"default:false;comment:是否为导航" json:"nav"`
	Router *bool  `gorm:"default:false;comment:是否为路由页面" json:"router"`
	Icon   string `gorm:"type:varchar(50);comment:导航节点的字体图标" json:"icon"`
	Sort   uint8  `gorm:"default:0;comment:导航节点排序" json:"sort"`
}
