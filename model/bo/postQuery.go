package bo

type PostQuery struct {
	//参数都通过json形式传递
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int
	//
	Page int `json:"page"`
	Size int `json:"size"`
}
