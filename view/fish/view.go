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
	typex, _ := c.GetQuery("type")
	soText, _ := c.GetQuery("so_text")

	pInt, _ := strconv.ParseInt(p, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)

	pageStart := page.Start(pInt, pageSizeInt)

	sql := `select id,type,project_name,agent,ip,ip_info,create_time from hfish_info where 1=1`
	sqlx := `select count(1) as sum from hfish_info where 1=1`
	sqlStatus := 0

	if typex != "all" {
		sql = sql + ` and type=?`
		sqlx = sqlx + ` and type=?`
		sqlStatus = 1
	}

	if soText != "" {
		sql = sql + ` and (project_name like ? or ip like ?)`
		sqlx = sqlx + ` and type=?`
		if sqlStatus == 1 {
			sqlStatus = 3
		} else {
			sqlStatus = 2
		}
	}

	sql = sql + ` ORDER BY id desc LIMIT ?,?;`

	if sqlStatus == 0 {
		result := dbUtil.Query(sql, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx)
		pageCount := resultx[0]["sum"].(int64)
		pageCount = page.TotalPage(pageCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":      result,
			"pageCount": pageCount,
			"page":      p,
		})
	} else if sqlStatus == 1 {
		result := dbUtil.Query(sql, typex, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex)
		pageCount := resultx[0]["sum"].(int64)
		pageCount = page.TotalPage(pageCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":      result,
			"pageCount": pageCount,
			"page":      p,
		})
	} else if sqlStatus == 2 {
		result := dbUtil.Query(sql, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, "%"+soText+"%", "%"+soText+"%")
		pageCount := resultx[0]["sum"].(int64)
		pageCount = page.TotalPage(pageCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":      result,
			"pageCount": pageCount,
			"page":      p,
		})
	} else if sqlStatus == 3 {
		result := dbUtil.Query(sql, typex, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex, "%"+soText+"%", "%"+soText+"%")
		pageCount := resultx[0]["sum"].(int64)
		pageCount = page.TotalPage(pageCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":      result,
			"pageCount": pageCount,
			"page":      p,
		})
	}
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

// 获取钓鱼分类信息
func GetFishTypeInfo(c *gin.Context) {
	sql := `select type from hfish_info GROUP BY type;`
	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}
