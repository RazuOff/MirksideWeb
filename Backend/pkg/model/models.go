package model

type ShortNews struct {
	ID                 int
	ContentHeader      string
	ContentDescription string
	ImageUrl           string
}

type ExactNewsBlock struct {
	Text   string
	ImgUrl string
}

type ExactNews struct {
	ID             int
	ExactNewsBlock []ExactNewsBlock
}
