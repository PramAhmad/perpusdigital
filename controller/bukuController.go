package controller

import (
	"github.com/gin-gonic/gin"
)

type Buku struct {
	Id           int    `json:"id"`
	Cover        string `json:"cover"`
	CreatedAt    string `json:"created_at"`
	IdKategori   int    `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
	IdRak        int    `json:"id_rak"`
	NamaRak      string `json:"nama_rak"`
	Judul        string `json:"judul"`

	Penulis string `json:"penulis"`

	Penerbit string `json:"penerbit"`

	TahunTerbit string `json:"tahun_terbit"`

	JumlahHalaman int `json:"jml_hal"`

	Deskripsi string `json:"deskripsi"`
}

// Get All Buku
// @Summary Get all buku
// @Description Get all buku
// @Produce json
// @Success 200 {array} Buku
// @Router /buku [get]
// @Tags buku
// @Security ApiKeyAuth
// @Accept json
func GetAllBuku(c *gin.Context) {
	rows, err := DB.Query("SELECT buku.id, buku.cover, buku.created_at, buku.id_kategori, kategori_buku.nama, buku.id_rak, rak.nama_rak, buku.judul, buku.penulis, buku.penerbit, buku.tahun_terbit, buku.jml_hal, buku.deskripsi FROM buku JOIN kategori_buku ON buku.id_kategori = kategori_buku.id JOIN rak ON buku.id_rak = rak.id")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	if rows == nil {
		c.JSON(404, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	var bukus []Buku
	for rows.Next() {
		var buku Buku
		err := rows.Scan(&buku.Id, &buku.Cover, &buku.CreatedAt, &buku.IdKategori, &buku.NamaKategori, &buku.IdRak, &buku.NamaRak, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.TahunTerbit, &buku.JumlahHalaman, &buku.Deskripsi)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		bukus = append(bukus, buku)
	}
	c.JSON(200, bukus)
}

// Get Buku by id
// @Summary Get buku by id
// @Description Get buku by id
// @Produce json
// @Param id path int true "Buku ID"
// @Success 200 {object} Buku
// @Router /buku/{id} [get]
// @Tags buku
func GetBukuById(c *gin.Context) {
	rows := DB.QueryRow("SELECT buku.id, buku.cover, buku.created_at, buku.id_kategori, kategori_buku.nama, buku.id_rak, rak.nama_rak, buku.judul, buku.penulis, buku.penerbit, buku.tahun_terbit, buku.jml_hal, buku.deskripsi FROM buku JOIN kategori_buku ON buku.id_kategori = kategori_buku.id JOIN rak ON buku.id_rak = rak.id WHERE buku.id = ?", c.Param("id"))
	var buku Buku
	err := rows.Scan(&buku.Id, &buku.Cover, &buku.CreatedAt, &buku.IdKategori, &buku.NamaKategori, &buku.IdRak, &buku.NamaRak, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.TahunTerbit, &buku.JumlahHalaman, &buku.Deskripsi)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, buku)
}

// Create Buku
// @Summary Create buku
// @Description Create buku
// @Accept json
// @Produce json
// @Param cover formData file true "Cover buku"
// @Param id_kategori formData int true "ID Kategori"
// @Param id_rak formData int true "ID Rak"
// @Param judul formData string true "Judul buku"
// @Param penulis formData string true "Penulis buku"
// @Param penerbit formData string true "Penerbit buku"
// @Param tahun_terbit formData string true "Tahun terbit buku"
// @Param jml_hal formData int true "Jumlah halaman buku"
// @Param deskripsi formData string true "Deskripsi buku"
// @Success 201 {object} Buku
// @Router /buku [post]
// @Tags buku
func CreateBuku(c *gin.Context) {
	var buku Buku
	err := c.BindJSON(&buku)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// hanlde image for cover
	file, err := c.FormFile("cover")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, "images/"+file.Filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	bukuUrl := "http://localhost:8080/public/" + file.Filename
	buku.Cover = bukuUrl

	result, err := DB.Exec("INSERT INTO buku (cover, id_kategori, id_rak, judul, penulis, penerbit, tahun_terbit, jml_hal, deskripsi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", buku.Cover, buku.IdKategori, buku.IdRak, buku.Judul, buku.Penulis, buku.Penerbit, buku.TahunTerbit, buku.JumlahHalaman, buku.Deskripsi)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	buku.Id = int(id)
	c.JSON(201, buku)
}

// Update Buku
// @Summary Update buku
// @Description Update buku
// @Accept json
// @Produce json
// @Param id path int true "Buku ID"
// @Param buku body Buku true "Buku"
// @Success 200 {object} Buku
// @Router /buku/{id} [put]
// @Tags buku
func UpdateBuku(c *gin.Context) {
	var buku Buku
	err := c.BindJSON(&buku)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := DB.Exec("UPDATE buku SET cover = ?, id_kategori = ?, id_rak = ?, judul = ?, penulis = ?, penerbit = ?, tahun_terbit = ?, jml_hal = ?, deskripsi = ? WHERE id = ?", buku.Cover, buku.IdKategori, buku.IdRak, buku.Judul, buku.Penulis, buku.Penerbit, buku.TahunTerbit, buku.JumlahHalaman, buku.Deskripsi, buku.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// Delete Buku
// @Summary Delete buku
// @Description Delete buku
// @Produce json
// @Param id path int true "Buku ID"
// @Success 200 {object} Buku
// @Router /buku/{id} [delete]
// @Tags buku
func DeleteBuku(c *gin.Context) {
	result, err := DB.Exec("DELETE FROM buku WHERE id = ?", c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}
