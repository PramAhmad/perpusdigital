package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Kategori struct {
	Id           int    `json:"id"`
	NamaKategori string `json:"nama_kategori"`
}

func GetAllKategori(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM kategori_buku")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var kategoris []Kategori
	for rows.Next() {
		var kategori Kategori
		err := rows.Scan(&kategori.Id, &kategori.NamaKategori)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		kategoris = append(kategoris, kategori)
	}
	c.JSON(200, kategoris)

}
func GetKategoriById(c *gin.Context) {
	rows := DB.QueryRow("SELECT * FROM kategori_buku WHERE id = ?", c.Param("id"))
	var kategori Kategori
	err := rows.Scan(&kategori.Id, &kategori.NamaKategori)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, kategori)

}
func CreateKategori(c *gin.Context) {
	var kategori Kategori
	err := c.BindJSON(&kategori)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := DB.Exec("INSERT INTO kategori_buku (nama_kategori) VALUES (?)", kategori.NamaKategori)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}

func UpdateKategori(c *gin.Context) {
	var kategori Kategori
	err := c.BindJSON(&kategori)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := DB.Exec("UPDATE kategori_buku SET nama_kategori = ? WHERE id = ?", kategori.NamaKategori, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func DeleteKategori(c *gin.Context) {
	result, err := DB.Exec("DELETE FROM kategori_buku WHERE id = ?", c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
