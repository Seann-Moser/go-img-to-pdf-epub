package go_img_to_pdf_epub

type ReadingDirection int

const (
	RightToLeft ReadingDirection = iota
	LeftToRight
)

type PageSize int

const (
	ImageSize PageSize = iota
)

type Convert interface {
	SetReadDirection(readDirection ReadingDirection)
	SetAuthor(author string)
	SetBookDescription(desc string)
	SetSetChapterSectionBody(chapterSectionBody string)
	SetBookCover(img string) error
	SetPageSize(pageSize PageSize)
	SetChapterCover(coverName string) error
	AddImage(path string) error
	AddChapter(dir string) error
	AddBook(dir string) error
}
