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
	SetChapterCover(coverName string)
	Overwrite(overwrite bool)
	ConvertImage(path string) error
	ConvertChapter(name, dir, output string) error
	ConvertBook(name, dir, output string) error
	SetTitle(title string)
}
