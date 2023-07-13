package models

import "gormdemo/models"

// HasMany 一对多关系: 一个动物有多个种类
//动物表
type Animal struct {
	AnimalID int64 `gorm:"primary_key;column:animal_id"`
	Name     string
	//foreignkey:AnimalID 指定外键,这里是AnimalSpecies表中的AnimalID字段
	AnimalSpecies []AnimalSpecies `gorm:"foreignkey:AnimalID"`
}

//动物种类表
type AnimalSpecies struct {
	AnimalSpeciesID int64 `gorm:"primary_key;column:animal_species_id"`
	AnimalID        int64
	Species         string
}

// type Animal struct {
// 	gorm.Model
// 	AnimalID      int64
// 	Name          string
// 	AnimalSpecies []AnimalSpecies `gorm:"foreignkey:AnimalID"`
// }

// type AnimalSpecies struct {
// 	AnimalSpeciesID int64
// 	AnimalID        int64
// 	Species         string
// 	gorm.Model
// }

// HasMany 一对多关系
func HasMany() {
	// 1. 创建表
	models.Db.AutoMigrate(&Animal{}, &AnimalSpecies{})
	// 2. 插入数据
	a1 := AnimalSpecies{AnimalSpeciesID: 1, Species: "猫"}
	a2 := AnimalSpecies{AnimalSpeciesID: 2, Species: "狗"}
	b1 := Animal{AnimalID: 1, Name: "小黑", AnimalSpecies: []AnimalSpecies{a1, a2}}
	models.Db.Create(&b1)
	// 3. 查询数据
	// 4. 更新数据
	// 5. 删除数据
}
