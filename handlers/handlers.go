package handlers

import (
	"net/http"

	"github.com/GenaroNajera/benkyou/database"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Study(c *gin.Context) {
	c.HTML(http.StatusOK, "study.html", gin.H{})
}

func Api(c *gin.Context, d []database.Data) {
	c.JSON(http.StatusOK, d)
}
