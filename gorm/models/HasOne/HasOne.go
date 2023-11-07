package hasone

import "gormdemo/models"

type Animal struct {
	AnimalID int64 `gorm:"primary_key;column:animal_id"`
	Name     string
	OwnerID  int64
}

type Owner struct {
	OwnerID int64 `gorm:"primary_key;column:owner_id"`
	Name    string
	Animal  Animal
}

func HasOne() {

	//外键约束,Owner表的Animal是外键

	//这里需要把两个表都迁移
	models.Db.AutoMigrate(&Owner{}, &Animal{})
	//两个动物,依附于主键Owner表，没有数据则动物创建失败
	a1 := Animal{Name: "小黑"}
	a2 := Animal{Name: "小白"}

	//张三不需要动物也能创建成功
	o1 := Owner{Name: "张三"}

	o2 := Owner{Name: "李四", Animal: a1}
	o3 := Owner{Name: "王五", Animal: a2}

	models.Db.Create(&o1)
	models.Db.Create(&o2)
	models.Db.Create(&o3)
}
