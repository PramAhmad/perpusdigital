package repository

import "github.com/gin-gonic/gin"

type Buku struct {
	Id            int    `json:"id"`
	Cover         string `json:"cover"`
	CreatedAt     string `json:"created_at"`
	IdKategori    int    `json:"id_kategori"`
	NamaKategori  string `json:"nama_kategori"`
	IdRak         int    `json:"id_rak"`
	NamaRak       string `json:"nama_rak"`
	Judul         string `json:"judul"`
	Penulis       string `json:"penulis"`
	Penerbit      string `json:"penerbit"`
	TahunTerbit   string `json:"tahun_terbit"`
	JumlahHalaman int    `json:"jml_hal"`
	Deskripsi     string `json:"deskripsi"`
}

func GetAllBuku(c *gin.Context) {
	rows, err := DB.Query("SELECT buku.id, buku.cover, buku.created_at, buku.id_kategori, kategori_buku.nama, buku.id_rak, rak.nama_rak, buku.judul, buku.penulis, buku.penerbit, buku.tahun_terbit, buku.jml_hal, buku.deskripsi FROM buku JOIN kategori_buku ON buku.id_kategori = kategori_buku.id JOIN rak ON buku.id_rak = rak.id")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
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

func DeleteBuku(c *gin.Context) {
	result, err := DB.Exec("DELETE FROM buku WHERE id = ?", c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}
