package dto

type ChapterBookInfo struct {
	BookId     int    `json:"book_id"`
	Name       string `json:"name"`
	IsFinished int    `json:"is_finished"`
}

type ChapterInfoDto struct {
	Book         ChapterBookInfo `json:"book"`
	ChapterTitle string          `json:"chapter_title"`
	Content      string          `json:"content"`
	IsPreview    int             `json:"is_preview"`
	LastChapter  int             `json:"last_chapter"`
	NextChapter  string          `json:"next_chapter"`
	ChapterId    int             `json:"chapter_id"`
	UpdateTime   string          `json:"update_time"`
}

type ChaptersFrontLists struct {
	BookId       int    `json:"book_id"`
	ChapterId    int    `json:"chapter_id"`
	ChapterTitle string `json:"chapter_title"`
	Words        int    `json:"words"`
	IsVip        int    `json:"is_vip"`
	CanRead      int    `json:"can_read"`
	IsPreview    int    `json:"is_preview"`
	Tag          struct {
		Tab   string `json:"tab"`
		Color string `json:"color"`
	} `json:"tag"`
	UpdateTime   string `json:"update_time"`
	DisplayOrder int `json:"display_order"`
	LastChapter  int `json:"last_chapter"`
	NextChapter  int `json:"next_chapter"`
}
type ChaptersFrontListDto struct {
	BookId       int                  `json:"book_id"`
	Name         string               `json:"name"`
	Cover        string               `json:"cover"`
	Description  string               `json:"description"`
	TotalChapter int                  `json:"total_chapter"`
	ChapterList  []ChaptersFrontLists `json:"chapter_list"`
}
