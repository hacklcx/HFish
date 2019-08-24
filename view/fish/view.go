package fish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/page"
	"strconv"
	"strings"
)

// 蜜罐 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "fish.html", gin.H{})
}

// 获取蜜罐列表
func GetFishList(c *gin.Context) {
	p, _ := c.GetQuery("page")
	pageSize, _ := c.GetQuery("pageSize")
	typex, _ := c.GetQuery("type")
	colony, _ := c.GetQuery("colony")
	soText, _ := c.GetQuery("so_text")

	pInt, _ := strconv.ParseInt(p, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)

	pageStart := page.Start(pInt, pageSizeInt)

	sql := `select id,type,project_name,agent,ip,country,region,city,create_time,info from hfish_info where 1=1`
	sqlx := `select count(1) as sum from hfish_info where 1=1`
	sqlStatus := 0

	if typex != "all" {
		sql = sql + ` and type=?`
		sqlx = sqlx + ` and type=?`
		sqlStatus = 1
	}

	if colony != "all" {
		sql = sql + ` and agent=?`
		sqlx = sqlx + ` and agent=?`
		if sqlStatus == 1 {
			sqlStatus = 3
		} else {
			sqlStatus = 2
		}
	}

	if soText != "" {
		sql = sql + ` and (project_name like ? or ip like ?)`
		sqlx = sqlx + ` and type=?`

		if sqlStatus == 1 {
			sqlStatus = 4
		} else if sqlStatus == 2 {
			sqlStatus = 5
		} else if sqlStatus == 3 {
			sqlStatus = 6
		} else {
			sqlStatus = 7
		}
	}

	sql = sql + ` ORDER BY id desc LIMIT ?,?;`

	if sqlStatus == 0 {
		result := dbUtil.Query(sql, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx)
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 1 {
		result := dbUtil.Query(sql, typex, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex)
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 2 {
		result := dbUtil.Query(sql, colony, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, colony)
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 3 {
		result := dbUtil.Query(sql, typex, colony, pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex, colony)
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 4 {
		result := dbUtil.Query(sql, typex, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex, "%"+soText+"%", "%"+soText+"%")
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 5 {
		result := dbUtil.Query(sql, colony, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, colony, "%"+soText+"%", "%"+soText+"%")
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 6 {
		result := dbUtil.Query(sql, typex, colony, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, typex, colony, "%"+soText+"%", "%"+soText+"%")
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	} else if sqlStatus == 7 {
		result := dbUtil.Query(sql, "%"+soText+"%", "%"+soText+"%", pageStart, pageSizeInt)
		resultx := dbUtil.Query(sqlx, "%"+soText+"%", "%"+soText+"%")
		totalCount := resultx[0]["sum"].(int64)
		pageCount := page.TotalPage(totalCount, pageSizeInt)

		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pageCount":  pageCount,
			"totalCount": totalCount,
			"page":       p,
		})
	}
}

// 删除蜜罐
func PostFishDel(c *gin.Context) {
	id := c.PostForm("id")

	idx := strings.Split(id, ",")

	// 暂时此种方式，待优化
	for _, x := range idx {
		sqlDel := `delete from hfish_info where id in (?);`
		dbUtil.Delete(sqlDel, x)
	}

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

// 获取蜜罐信息
func GetFishInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	sql := `select info from hfish_info where id=?;`
	result := dbUtil.Query(sql, id)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

// 获取蜜罐分类信息,集群信息
func GetFishTypeInfo(c *gin.Context) {
	sqlInfoType := `select type from hfish_info GROUP BY type;`
	resultInfoType := dbUtil.Query(sqlInfoType)

	sqlColonyName := `select agent_name from hfish_colony GROUP BY agent_name;`
	resultColonyName := dbUtil.Query(sqlColonyName)

	c.JSON(http.StatusOK, gin.H{
		"resultInfoType":   resultInfoType,
		"resultColonyName": resultColonyName,
	})
}
