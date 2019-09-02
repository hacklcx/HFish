package fish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/page"
	"strconv"
	"strings"
	"HFish/utils/log"
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

	// 统计攻击IP
	db := dbUtil.DB().Table("hfish_info").Fields("id", "type", "project_name", "agent", "ip", "country", "region", "city", "create_time", "info").Where("1", "=", "1")
	dbCount := dbUtil.DB().Table("hfish_info").Where("1", "=", "1")

	if typex != "all" {
		db.Where("type", "=", typex)
		dbCount.Where("type", "=", typex)
	}

	if colony != "all" {
		db.Where("agent", "=", colony)
		dbCount.Where("agent", "=", colony)
	}

	if soText != "" {
		db.Where("project_name", "like", "%"+soText+"%").OrWhere("ip", "like", "%"+soText+"%")
		dbCount.Where("project_name", "like", "%"+soText+"%").OrWhere("ip", "like", "%"+soText+"%")
	}

	// 统计查询数量
	totalCount, errCount := dbCount.Count()

	if errCount != nil {
		log.Pr("HFish", "127.0.0.1", "统计分页总数失败", errCount)
	}

	// 查询列表
	pInt, _ := strconv.Atoi(p)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageStart := page.Start(pInt, pageSizeInt)

	result, err := db.OrderBy("id desc").Limit(pageStart).Offset(pageSizeInt).Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询上钩信息列表失败", err)
	}

	totalCountString := strconv.FormatInt(totalCount, 10)
	totalCountInt, _ := strconv.Atoi(totalCountString)

	pageCount := page.TotalPage(totalCountInt, pageSizeInt)

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pageCount":  pageCount,
		"totalCount": totalCount,
		"page":       p,
	})
}

// 删除蜜罐
func PostFishDel(c *gin.Context) {
	id := c.PostForm("id")

	idx := strings.Split(id, ",")
	inId := make([]interface{}, 20)

	for _, x := range idx {
		inId = append(inId, x)
	}

	_, err := dbUtil.DB().Table("hfish_info").WhereIn("id", inId).Delete()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "删除蜜罐失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 获取蜜罐信息
func GetFishInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")

	result, err := dbUtil.DB().Table("hfish_info").Fields("info").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取蜜罐信息失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": result,
	})
}

// 获取蜜罐分类信息,集群信息
func GetFishTypeInfo(c *gin.Context) {
	resultType, errType := dbUtil.DB().Table("hfish_info").Fields("type").GroupBy("type").Get()

	if errType != nil {
		log.Pr("HFish", "127.0.0.1", "获取蜜罐分类失败", errType)
	}

	resultAgent, errAgent := dbUtil.DB().Table("hfish_info").Fields("agent").GroupBy("agent").Get()

	if errAgent != nil {
		log.Pr("HFish", "127.0.0.1", "获取集群分类失败", errAgent)
	}

	data := map[string]interface{}{
		"resultInfoType":   resultType,
		"resultColonyName": resultAgent,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}
