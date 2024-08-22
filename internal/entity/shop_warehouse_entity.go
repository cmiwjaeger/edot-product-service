package entity

import "github.com/google/uuid"

type ShopWarehouse struct {
	ID          uuid.UUID `gorm:"column:id"`
	WarehouseID uuid.UUID `gorm:"column:warehouse_id"`
	Warehouse   Warehouse `gorm:"references:WarehouseID"`
	ShopID      uuid.UUID `gorm:"column:shop_id"`
	Shop        Shop      `gorm:"references:ShopID"`
}

func (u *ShopWarehouse) TableName() string {
	return "shop_warehouses"
}
