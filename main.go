package main

import (
	"ditaperpus/controller"
	_ "ditaperpus/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title PerpusDigital
// @version		0.6
// @description	Backend Perpustakaan Digital
// @schemes		http
// @host			localhost:8080
// @basePath		/
// @produces		json
// @consumes		json
func main() {

	controller.ConnectDb()

	r := gin.Default()
	// kategori
	// route group
	kategori := r.Group("/kategori")
	{
		kategori.GET("", controller.GetAllKategori)
		kategori.GET("/:id", controller.GetKategoriById)
		kategori.POST("", controller.CreateKategori)
		kategori.PUT("/:id", controller.UpdateKategori)
		kategori.DELETE("/:id", controller.DeleteKategori)

	}
	// rak
	rak := r.Group("/rak")
	{
		rak.GET("", controller.GetAllRak)
		rak.GET("/:id", controller.GetRakById)
		rak.POST("", controller.CreateRak)
		rak.PUT("/:id", controller.UpdateRak)
		rak.DELETE("/:id", controller.DeleteRak)
	}

	// buku
	buku := r.Group("/buku")
	{
		buku.GET("", controller.GetAllBuku)
		buku.GET("/:id", controller.GetBukuById)
		buku.POST("", controller.CreateBuku)
		buku.PUT("/:id", controller.UpdateBuku)
		buku.DELETE("/:id", controller.DeleteBuku)
	}

	// swagger yaml handle
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
