# GORM 相关操作


## 关系映射


### 多对多

~~~go
type AccountEntity struct {
	Id    int           `json:"id"`
	Name  string        `json:"name"`
	Roles []*RoleEntity `json:"roles" gorm:"many2many:account_role;foreignKey:Id;joinForeignKey:accountId;joinReferences:roleId;"`
}

type RoleEntity struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Accounts []*AccountEntity `json:"accounts" gorm:"many2many:account_role;foreignKey:Id;joinForeignKey:roleId;joinReferences:accountId"`
}

func (AccountEntity) TableName() string {
	return "account"
}

func (RoleEntity) TableName() string {
	return "role"
}

func init() {
	utils.GormDb.AutoMigrate(&AccountEntity{}, &RoleEntity{})
}

~~~

1.上面几个用到的字段解析说明,[官网地址](https://gorm.io/zh_CN/docs/associations.html#%E5%85%B3%E8%81%94%E6%A0%87%E7%AD%BE)

- many2many:user_profiles 定义中间表名为:user_profiles
- foreignKey:Id 使用当前表的id作为外键
- joinForeignKey:accountId 当前数据模型外键关联到中间件表的字段名叫accountId
- joinReferences:roleId 反向引用字段，如果是账号表就要写中间表的roleId
  
1. 查找账号列表顺便查询出角色信息
   
   ~~~go
   var accountList []AccountEntity
   utils.GormDb.Preload("Roles").Find(&accountList)
   fmt.Println(tools.MapToJson(accountList))
   ~~~

2. 根据角色查询到账号信息
   ~~~go
   var roleList []RoleEntity
   utils.GormDb.Preload("Accounts").Find(&roleList)
   fmt.Println(tools.MapToJson(roleList))
   ~~~


 这三个表如何通过article 查到所有信息