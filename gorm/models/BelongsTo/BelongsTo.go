package belongsto

import "gormdemo/models"

type Animal struct {
	AnimalID int64 `gorm:"primary_key;column:animal_id"`
	Name     string
	OwnerID  int64
	//不能加  `gorm:"foreignkey:OwnerID"` 会报错,gorm会自动创建
	Owner Owner
}

type Owner struct {
	OwnerID int64 `gorm:"primary_key;column:owner_id"`
	Name    string
}

func BelongsTo() {
	//这里迁移会自动创建两个表
	models.Db.AutoMigrate(&Animal{})

	o1 := Owner{Name: "张三"}
	o2 := Owner{Name: "李四"}
	a1 := Animal{Name: "小黑", Owner: o1}
	a2 := Animal{Name: "小白", Owner: o2}
	models.Db.Create(&a1)
	models.Db.Create(&a2)
}
