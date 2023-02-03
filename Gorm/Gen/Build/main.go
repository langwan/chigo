package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(localhost:3306)/chigo?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../app/orm/dal",
		ModelPkgPath: "../app/orm/model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	// 设置目标 db
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateModel("accounts", gen.FieldType("level", "Level")))
	g.Execute()
}
