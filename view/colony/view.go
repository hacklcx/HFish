package colony

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/error"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "colony.html", gin.H{})
}

// 获取蜜罐分类信息
func GetColony(c *gin.Context) {
	sql := `
		SELECT
			id,
			agent_name,
			agent_ip,
			web_status,
			deep_status,
			ssh_status,
			redis_status,
			mysql_status,
			http_status,
			telnet_status,
            ftp_status,
            mem_cache_status,
            plug_status,
			last_update_time
		FROM
			hfish_colony
		ORDER BY
			id DESC;
	`

	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

// 删除集群
func PostColonyDel(c *gin.Context) {
	id := c.PostForm("id")

	sqlDel := `delete from hfish_colony where id=?;`
	dbUtil.Delete(sqlDel, id)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}
