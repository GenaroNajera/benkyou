package handlers

import (
	"net/http"

	"github.com/GenaroNajera/benkyou/database"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context, d []database.Data) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home", "class": "home", "text": d})
}
