package fish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/page"
	"strconv"
)

// 钓鱼 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "fish.html", gin.H{})
}

// 获取钓鱼列表
func GetFishList(c *gin.Context) {
	p, _ := c.GetQuery("page")
	pageSize, _ := c.GetQuery("pageSize")

	pInt, _ := strconv.ParseInt(p, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)

	pageStart := page.Start(pInt, pageSizeInt)

	sql := `select id,type,project_name,ip,create_time from hfish_info ORDER BY id desc LIMIT ?,?;`
	result := dbUtil.Query(sql, pageStart, pageSizeInt)

	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

// 删除钓鱼
func PostFishDel(c *gin.Context) {
	id := c.PostForm("id")
	sqlDel := `delete from hfish_info where id=?;`
	dbUtil.Delete(sqlDel, id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

// 获取钓鱼信息
func GetFishInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	sql := `select info from hfish_info where id=?;`
	result := dbUtil.Query(sql, id)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}
