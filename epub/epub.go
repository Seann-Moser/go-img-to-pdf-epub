package epub

import (
	"errors"
	"fmt"
	"github.com/Seann-Moser/go-img-to-pdf-epub"
	epub "github.com/bmaupin/go-epub"
	"github.com/google/uuid"
	"go.uber.org/multierr"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type epubConverter struct {
	epub              *epub.Epub
	chapterSection    string
	chapterCoverImage string
	bookCoverImage    string
	overwrite         bool
}

func (e *epubConverter) AddPage(file string) error {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) Convert(outputFile string) error {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) Overwrite(overwrite bool) {
	//TODO implement me
	e.overwrite = overwrite
}

func NewEpubConverter() go_img_to_pdf_epub.Convert {
	e := epub.NewEpub("")
	return &epubConverter{
		epub:           e,
		chapterSection: `<img style='margin-left:auto;margin-right:auto;text-align:center;height:"95%";padding:0px;margin-top:10px' src='{image}'/>`,
		overwrite:      false,
	}
}

func (e *epubConverter) SetTitle(title string) {
	e.epub.SetTitle(title)
}

func (e *epubConverter) SetReadDirection(readDirection go_img_to_pdf_epub.ReadingDirection) {
	switch readDirection {
	case go_img_to_pdf_epub.RightToLeft:
		e.epub.SetPpd("rtl")
	case go_img_to_pdf_epub.LeftToRight:
		e.epub.SetPpd("ltr")
	default:
		e.epub.SetPpd("ltr")
	}
}

func (e *epubConverter) SetAuthor(author string) {
	//TODO implement me
	e.epub.SetAuthor(author)
}

func (e *epubConverter) SetBookDescription(desc string) {
	e.epub.SetDescription(desc)
}

func (e *epubConverter) SetSetChapterSectionBody(chapterSectionBody string) {
	e.chapterSection = chapterSectionBody
}

func (e *epubConverter) SetBookCover(img string) {
	e.bookCoverImage = img
}

func (e *epubConverter) SetChapterCover(coverName string) {
	e.chapterCoverImage = coverName
}

func (e *epubConverter) SetPageSize(pageSize go_img_to_pdf_epub.PageSize) {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) ConvertImage(path string) error {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) ConvertChapter(chapter *go_img_to_pdf_epub.Chapter, output string) (string, error) {
	// Add a section
	chapterSection, err := e.createChapter(chapter.ChapterName, chapter.ChapterDirectory, true)
	if err != nil {
		return "", multierr.Combine(err, fmt.Errorf("failed creating chapter"))
	}
	_, err = e.epub.AddSection(chapterSection, chapter.ChapterName, "", "")
	if err != nil {
		return "", multierr.Combine(err, fmt.Errorf("failed getting pages from chapter to convert"))
	}
	if len(strings.TrimSpace(output)) == 0 {
		output = chapter.ChapterName
	}
	if strings.HasSuffix(output, "/") {
		output += chapter.ChapterName
	}
	if !strings.HasSuffix(output, ".epub") {
		output += ".epub"
	}
	if exists(output) && !e.overwrite {
		return output, fmt.Errorf("file already exists")
	}
	return output, e.epub.Write(output)
}

func (e *epubConverter) ConvertBook(book *go_img_to_pdf_epub.Book, output string) (string, error) {
	e.SetTitle(book.BookName)
	for i, chapter := range book.Chapters {
		setCover := false
		if i == 0 || chapter.ChapterNumber == 1 {
			setCover = true
		}
		chapterSection, err := e.createChapter(chapter.ChapterName, chapter.ChapterDirectory, setCover)
		if err != nil {
			return "", multierr.Combine(err, fmt.Errorf("failed creating chapter"))
		}
		_, err = e.epub.AddSection(chapterSection, chapter.ChapterName, "", "")
		if err != nil {
			return "", multierr.Combine(err, fmt.Errorf("failed creating chapter %s -> %s", chapter.ChapterName, chapter.ChapterDirectory))
		}
	}
	if len(strings.TrimSpace(output)) == 0 {
		output = book.BookName
	}
	if strings.HasSuffix(output, "/") {
		output += book.BookName
	}
	if !strings.HasSuffix(output, ".epub") {
		output += ".epub"
	}
	if exists(output) && !e.overwrite {
		return output, fmt.Errorf("file already exists")
	}
	return output, e.epub.Write(output)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
func fixPages(fileLocation string) (string, error) {
	ext := fileLocation[strings.LastIndex(fileLocation, ".")+1:]
	switch ext {
	case "jpg":
		fallthrough
	case "jpeg":
		fallthrough
	case "png":
		return fileLocation, nil
	}
	return "", fmt.Errorf("invalid file format %s", ext)
}
func orderPages(files []fs.FileInfo) []fs.FileInfo {
	re := regexp.MustCompile("[0-9]+")
	sort.Slice(files, func(i, j int) bool {
		left := re.FindAllString(files[i].Name(), -1)
		right := re.FindAllString(files[j].Name(), -1)
		if len(left) == 0 {
			return true
		}
		if len(right) == 0 {
			return false
		}
		leftNum, err := strconv.Atoi(left[0])
		if err != nil {
			return true
		}
		rightNum, err := strconv.Atoi(right[0])
		if err != nil {
			return false
		}
		return leftNum < rightNum
	})
	return files
}
func (e *epubConverter) createChapter(chapterName, chapterDir string, setCover bool) (string, error) {
	if !strings.HasSuffix(chapterDir, "/") {
		chapterDir += "/"
	}
	files, err := ioutil.ReadDir(chapterDir)
	if err != nil {
		return "", err
	}
	chapterSectionBody := fmt.Sprintf(`<h3 style="margin:auto;text-align:center;">%s</h3>`, chapterName)
	chapterID := uuid.New().String()
	for i, page := range orderPages(files) {

		if page.IsDir() {
			continue
		}
		pageLocation := fmt.Sprintf("%s%s", chapterDir, page.Name())
		fileLocation, err := fixPages(pageLocation)
		if err != nil {
			continue
		}
		ext := strings.Split(fileLocation, ".")
		internalPath := fmt.Sprintf("%s-%d.%s", chapterID, i+1, ext[len(ext)-1])
		if !exists(fileLocation) {
			continue
		}
		imgPath, err := e.epub.AddImage(fileLocation, internalPath)
		if err != nil {
			return "", err
		}
		if (e.chapterCoverImage == "" && (i == 0 && setCover)) || (e.chapterCoverImage != "" && page.Name() == e.chapterCoverImage) {
			e.epub.SetCover(imgPath, "")
		} else {
			chapterSectionBody += strings.ReplaceAll(e.chapterSection, "{image}", imgPath)
		}
	}
	return chapterSectionBody, nil
}

func (e *epubConverter) GetChapters(book *go_img_to_pdf_epub.Book) ([]*go_img_to_pdf_epub.Chapter, error) {
	if book.Dir == "" {
		return nil, nil
	}
	files, err := ioutil.ReadDir(book.Dir)
	if err != nil {
		return nil, err
	}
	var chapterList []*go_img_to_pdf_epub.Chapter
	for i, file := range orderPages(files) {
		if !file.IsDir() {
			continue
		}
		chapter := &go_img_to_pdf_epub.Chapter{
			ChapterName:      "",
			ChapterDirectory: fmt.Sprintf("%s%s", book.Dir, file.Name()),
			ChapterNumber:    i + 1,
		}
		if number, err := strconv.Atoi(file.Name()); err == nil {
			chapter.ChapterNumber = number
		}
		chapterList = append(chapterList, chapter)
	}
	if book.Chapters == nil {
		book.Chapters = chapterList
	} else {
		book.Chapters = append(book.Chapters, chapterList...)
	}
	return chapterList, nil
}
