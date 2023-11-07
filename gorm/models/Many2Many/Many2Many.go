package many2many

import "gormdemo/models"

// Many2Many 多对多关系

//如果使用gorm.Model, 择下面primary_key 都要删除，一张表只能一个主键
//中间表(AnimalAnimalSpecies)字段为 animal_id, animal_species_id
//如果不 使用gorm.Model, 记得设置一个主键
//中间表(AnimalAnimalSpecies)字段为 aimal_animal_id, aimal_animal_species_id

//动物种类表
type AnimalSpecies struct {
	AnimalSpeciesID int64 `gorm:"primary_key;column:animal_species_id;"`
	Species         string
	Animal          []Animal `gorm:"many2many:animal_animal_species;"` //多对多
}

//动物表
type Animal struct {
	AnimalID      int64 `gorm:"primary_key;column:animal_id;"`
	Name          string
	AnimalSpecies []AnimalSpecies `gorm:"many2many:animal_animal_species;"` //多对多
	AddressID     int64
}

type Address struct {
	AddressID int64 `gorm:"primary_key;column:address_id;"`
	Detail    string
	Animals   []Animal `gorm:"foreignkey:AddressID"` //(address)一对多(Animal)
}

type AnimalAnimalSpecies struct {
	AimalAnimalID        int64 `gorm:"primary_key;column:animal_id"`
	AimalAnimalSpeciesID int64 `gorm:"primary_key;column:animal_species_id"`
}

func Many2Many() {
	models.Db.AutoMigrate(&Animal{}, &AnimalSpecies{})

	As1 := AnimalSpecies{Species: "水里游得"}
	As2 := AnimalSpecies{Species: "地上跑的"}
	As3 := AnimalSpecies{Species: "天上飞的"}

	A1 := Animal{Name: "鸭子", AnimalSpecies: []AnimalSpecies{As1, As2}}
	A2 := Animal{Name: "海豚", AnimalSpecies: []AnimalSpecies{As1}}
	A3 := Animal{Name: "猫头鹰", AnimalSpecies: []AnimalSpecies{As3, As2}}
	A4 := Animal{Name: "鸽子", AnimalSpecies: []AnimalSpecies{As3, As2}}
	A5 := Animal{Name: "鳄鱼", AnimalSpecies: []AnimalSpecies{As1, As2}}

	models.Db.Create(&A1)
	models.Db.Create(&A2)
	models.Db.Create(&A3)
	models.Db.Create(&A4)
	models.Db.Create(&A5)

}

func Preload() {
	var Ad []Address
	// Address ->  Animal 一对多关系。 Animal->AnimalSpecies 多对多关系
	//s使用下面的语句直接查询三张表，从addressb包含nimal, Animal包含AnimalSpecies
	models.Db.Preload("Animal.AnimalSpecies").Find(&Ad)
}
