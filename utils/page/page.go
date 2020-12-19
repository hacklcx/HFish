package page

// Paging starts from 1
func Start(currentPage int, pageSize int) int {
	return (currentPage - 1) * pageSize
}

// End of page
func End(currentPage int64, pageSize int64) int64 {
	return currentPage * pageSize
}

// Total number of pages
func TotalPage(count int, pageSize int) int {
	result := count / pageSize
	yu := count % pageSize
	if yu > 0 {
		result = result + 1
	}
	return result
}
