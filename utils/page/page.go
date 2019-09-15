package page

// 分页从1开始
func Start(currentPage int, pageSize int) int {
	return (currentPage - 1) * pageSize
}

// 分页结束
func End(currentPage int64, pageSize int64) int64 {
	return currentPage * pageSize
}

// 分页总页数
func TotalPage(count int, pageSize int) int {
	result := count / pageSize
	yu := count % pageSize
	if yu > 0 {
		result = result + 1
	}
	return result
}
