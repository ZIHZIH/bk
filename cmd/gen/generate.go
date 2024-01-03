package main

func main() {
	//if err := config.InitConfig(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if err := infra.InitMysqlDB(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	//// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	//g := gen.NewGenerator(gen.Config{
	//	// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
	//	// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
	//	// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
	//	OutPath: "./dal/query",
	//	/* ModelPkgPath: "dal/model"*/
	//
	//	// gen.WithoutContext：禁用WithContext模式
	//	// gen.WithDefaultQuery：生成一个全局Query对象Q
	//	// gen.WithQueryInterface：生成Query接口
	//	Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	//})
	//g.UseDB(infra.DB)
	//
	//g.ApplyBasic(g.GenerateAllTable()...)
	//
	//// 执行并生成代码
	//g.Execute()
}
