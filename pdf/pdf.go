package pdf

import "github.com/Seann-Moser/go-img-to-pdf-epub"

type pdf struct {
}

func (p pdf) GetChapters(book *go_img_to_pdf_epub.Book) ([]*go_img_to_pdf_epub.Chapter, error) {
	//TODO implement me
	panic("implement me")
}

func NewPDFConverter() go_img_to_pdf_epub.Convert {
	return &pdf{}
}
func (p pdf) Overwrite(overwrite bool) {
	//TODO implement me
	panic("implement me")
}
func (p pdf) SetTitle(title string) {
	//TODO implement me
	panic("implement me")
}
func (p pdf) Convert(path string) error {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetReadDirection(readDirection go_img_to_pdf_epub.ReadingDirection) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetAuthor(author string) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetBookDescription(desc string) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetSetChapterSectionBody(chapterSectionBody string) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetBookCover(img string) error {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetPageSize(pageSize go_img_to_pdf_epub.PageSize) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) SetChapterCover(coverName string) {
	//TODO implement me
	panic("implement me")
}

func (p pdf) ConvertImage(path string) error {
	//TODO implement me
	panic("implement me")
}
func (p pdf) ConvertChapter(chapter *go_img_to_pdf_epub.Chapter, output string) error {
	//TODO implement me
	panic("implement me")
}

func (p pdf) ConvertBook(book *go_img_to_pdf_epub.Book, output string) error {
	//TODO implement me
	panic("implement me")
}
