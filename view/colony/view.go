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

// 获取蜜罐集群列表
func GetColony(c *gin.Context) {
	result, err := dbUtil.DB().Table("hfish_colony").OrderBy("id desc").Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取蜜罐集群列表失败", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
}

// 删除集群
func PostColonyDel(c *gin.Context) {
	id := c.PostForm("id")

	_, err := dbUtil.DB().Table("hfish_colony").Where("id", "=", id).Delete()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "删除集群失败", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}
