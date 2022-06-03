package epub

import (
	"errors"
	"fmt"
	"github.com/Seann-Moser/go-img-to-pdf-epub"
	epub "github.com/bmaupin/go-epub"
	"go.uber.org/multierr"
	"io/ioutil"
	"os"
	"strings"
)

type epubConverter struct {
	epub           *epub.Epub
	chapterSection string
	coverImage     string
	overwrite      bool
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

func (e *epubConverter) SetBookCover(img string) error {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) SetPageSize(pageSize go_img_to_pdf_epub.PageSize) {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) SetChapterCover(coverName string) {
	//TODO implement me
	e.coverImage = coverName
}

func (e *epubConverter) ConvertImage(path string) error {
	//TODO implement me
	panic("implement me")
}

func (e *epubConverter) ConvertChapter(name, dir, output string) error {
	// Add a section
	chapterSection, err := e.createChapter(name, dir, true)
	if err != nil {
		return multierr.Combine(err, fmt.Errorf("failed creating chapter"))
	}
	_, err = e.epub.AddSection(chapterSection, name, "", "")
	if err != nil {
		return multierr.Combine(err, fmt.Errorf("failed getting pages from chapter to convert"))
	}
	if len(strings.TrimSpace(output)) == 0 {
		output = name
	}
	if strings.HasSuffix(output, "/") {
		output += name
	}
	if !strings.HasSuffix(output, ".epub") {
		output += ".epub"
	}
	if exists(output) && !e.overwrite {
		return fmt.Errorf("file already exists")
	}
	return e.epub.Write(output)
}

func (e *epubConverter) ConvertBook(name, dir, output string) error {
	//TODO implement me
	panic("implement me")
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

func (e *epubConverter) createChapter(chapterName, chapterDir string, setCover bool) (string, error) {
	files, err := ioutil.ReadDir(chapterDir)
	if err != nil {
		return "", err
	}
	chapterSectionBody := fmt.Sprintf(`<h3 style="margin:auto;text-align:center;">%s</h3>`, chapterName)
	for i, page := range files {
		if page.IsDir() {
			continue
		}
		pageLocation := fmt.Sprintf("%s%s", chapterDir, page.Name())
		fileLocation, err := fixPages(pageLocation)
		if err != nil {
			continue
		}
		ext := strings.Split(fileLocation, ".")
		internalPath := fmt.Sprintf("%d-%d.%s", i, i+1, ext[len(ext)-1])
		if !exists(fileLocation) {
			continue
		}
		imgPath, err := e.epub.AddImage(fileLocation, internalPath)
		if err != nil {
			return "", err
		}
		if (e.coverImage == "" && (i == 0 && setCover)) || (e.coverImage != "" && page.Name() == e.coverImage) {
			e.epub.SetCover(imgPath, "")
		} else {
			chapterSectionBody += strings.ReplaceAll(e.chapterSection, "{image}", imgPath)
		}
	}
	return chapterSectionBody, nil
}
