package colony

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/log"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "colony.html", gin.H{})
}

// Get a list of honeypot clusters
func GetColony(c *gin.Context) {
	result, err := dbUtil.DB().Table("hfish_colony").OrderBy("id desc").Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to get the honeypot cluster list", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
}

// Delete cluster
func PostColonyDel(c *gin.Context) {
	id := c.PostForm("id")

	_, err := dbUtil.DB().Table("hfish_colony").Where("id", "=", id).Delete()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to delete cluster", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}
