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
	ConvertChapter(chapter *Chapter, output string) error
	ConvertBook(book *Book, output string) error
	GetChapters(book *Book) ([]*Chapter, error)
	SetTitle(title string)
}
type Chapter struct {
	ChapterName      string `json:"chapter_name"`
	ChapterDirectory string `json:"chapter_directory"`
	ChapterNumber    int    `json:"chapter_number"`
}
type Book struct {
	BookName string     `json:"book_name"`
	Chapters []*Chapter `json:"chapters"`
	Dir      string     `json:"dir"`
}
