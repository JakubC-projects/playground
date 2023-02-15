package main

import "gorm.io/gen"

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/
	})

	g.UseDB(DB)

	g.ApplyBasic(g.GenerateModel("people"), g.GenerateModelAs("temp_user", "User"))

	g.Execute()
}
