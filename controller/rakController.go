package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Rak struct {
	Id      int    `json:"id"`
	NamaRak string `json:"nama_rak"`
}

// Get All Rak
// @Summary Get all rak
// @Description Get all rak
// @Produce json
// @Success 200 {array} Rak
// @Router /rak [get]
// @Tags rak

func GetAllRak(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM rak")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var raks []Rak
	for rows.Next() {
		var rak Rak
		err := rows.Scan(&rak.Id, &rak.NamaRak)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		raks = append(raks, rak)
	}
	c.JSON(200, raks)
}

// Get Rak by id
// @Summary Get rak by id
// @Description Get rak by id
// @Produce json
// @Param id path int true "Rak ID"
// @Success 200 {object} Rak
// @Router /rak/{id} [get]
// @Tags rak
func GetRakById(c *gin.Context) {
	rows := DB.QueryRow("SELECT * FROM rak WHERE id = ?", c.Param("id"))
	var rak Rak
	err := rows.Scan(&rak.Id, &rak.NamaRak)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Rak tidak ditemukan"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rak)
}

// Create Rak
// @Summary Create rak
// @Description Create rak
// @Accept json
// @Produce json
// @Param rak body Rak true "Rak"
// @Success 201 {object} Rak
// @Router /rak [post]
// @Tags rak
func CreateRak(c *gin.Context) {
	var rak Rak
	err := c.BindJSON(&rak)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := DB.Exec("INSERT INTO rak (nama_rak) VALUES (?)", rak.NamaRak)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	rak.Id = int(id)
	c.JSON(201, rak)
}

// Update Rak
// @Summary Update rak
// @Description Update rak
// @Accept json
// @Produce json
// @Param id path int true "Rak ID"
// @Param rak body Rak true "Rak"
// @Success 200 {object} Rak
// @Router /rak/{id} [put]
// @Tags rak
func UpdateRak(c *gin.Context) {
	var rak Rak
	err := c.BindJSON(&rak)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := DB.Exec("UPDATE rak SET nama_rak = ? WHERE id = ?", rak.NamaRak, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if affected == 0 {
		c.JSON(404, gin.H{"error": "Rak tidak ditemukan"})
		return
	}
	c.JSON(200, rak)
}

// Delete Rak
// @Summary Delete rak
// @Description Delete rak
// @Produce json
// @Param id path int true "Rak ID"
// @Success 200 {object} Rak
// @Router /rak/{id} [delete]
// @Tags rak
func DeleteRak(c *gin.Context) {
	result, err := DB.Exec("DELETE FROM rak WHERE id = ?", c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if affected == 0 {
		c.JSON(404, gin.H{"error": "Rak tidak ditemukan"})
		return
	}
	c.JSON(200, gin.H{"message": "Rak berhasil dihapus"})
}
