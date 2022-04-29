package page

type Page struct {
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
	PageNo   int `json:"page_no"`
}

// PageTool 分页器
func PageTool(len, pageSize, pageNo int) (bool, int, int) {
	lastIndex := len
	if len == 0 {
		return true, 0, 0
	}
	if pageSize*pageNo > lastIndex {
		if (pageNo-1)*pageSize < lastIndex {
			return true, (pageNo - 1) * pageSize, lastIndex
		}
		return false, 0, 0
	}
	return true, (pageNo - 1) * pageSize, pageSize * pageNo
}
