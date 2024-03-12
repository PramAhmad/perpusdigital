package main

import (
	"ditaperpus/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	repository.ConnectDb()

	r := gin.Default()
	// kategori
	r.GET("/kategori", repository.GetAllKategori)
	r.GET("/kategori/:id", repository.GetKategoriById)
	r.POST("/kategori", repository.CreateKategori)
	r.PUT("/kategori/:id", repository.UpdateKategori)
	r.DELETE("/kategori/:id", repository.DeleteKategori)
	// rak
	r.GET("/rak", repository.GetAllRak)
	r.GET("/rak/:id", repository.GetRakById)
	r.POST("/rak", repository.CreateRak)
	r.PUT("/rak/:id", repository.UpdateRak)
	r.DELETE("/rak/:id", repository.DeleteRak)
	// buku
	r.GET("/buku", repository.GetAllBuku)
	r.GET("/buku/:id", repository.GetBukuById)
	r.POST("/buku", repository.CreateBuku)
	r.PUT("/buku/:id", repository.UpdateBuku)
	r.DELETE("/buku/:id", repository.DeleteBuku)

	r.Run(":8080")
}
