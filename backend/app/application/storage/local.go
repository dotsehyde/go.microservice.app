package storage

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/noelyahan/mergi"
)

type local struct {
	ctx  context.Context
	task string
}

func newLocal() storageDriverService {
	return &local{
		ctx:  context.Background(),
		task: "delete_file",
	}
}

func (l *local) uploadFile(dir string, fh *multipart.FileHeader) (string, error) {
	filename := l.getFilename(fh)
	var fpath string
	if strings.HasPrefix(dir, "public/") {
		fpath = strings.Replace(dir, "public/", "", 1)
		go l.save(dir, filename, fh)
		return fmt.Sprintf("%s/%s/%s", config.App().AppURL, fpath, filename), nil
	}
	fpath = dir
	go l.save(dir, filename, fh)
	return fmt.Sprintf("%s/%s", fpath, filename), nil
}
func (l *local) uploadCSVFile(dir string, fh *multipart.FileHeader) (string, error) {
	filename := strings.Replace(l.getFilename(fh), ".txt", ".csv", 1)
	var fpath string
	if strings.HasPrefix(dir, "public/") {
		fpath = strings.Replace(dir, "public/", "", 1)
		l.save(dir, filename, fh)
		return fmt.Sprintf("%s/%s/%s", config.App().AppURL, fpath, filename), nil
	}
	fpath = dir
	l.save(dir, filename, fh)
	return fmt.Sprintf("%s/%s", fpath, filename), nil
}

func (l *local) uploadFiles(dir string, files []*multipart.FileHeader) ([]string, error) {
	urls := make([]string, 0, len(files))
	for _, file := range files {
		url, err := l.uploadFile(dir, file)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}
func (l *local) uploadResizeImage(dir string, fh *multipart.FileHeader, width, height int) (string, error) {
	filename := l.getFilename(fh)
	var fpath string
	if strings.HasPrefix(dir, "public/") {
		fpath = strings.Replace(dir, "public/", "", 1)
	}
	go l.resizeImage(dir, filename, fh, width, height)
	return fmt.Sprintf("%s/%s", fpath, filename), nil
}
func (l *local) deleteFile(path any) error {
	fPath := path.(string)
	if err := os.Remove(fmt.Sprintf("./storage/%s", fPath)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("./storage/%s.fiber.gz", fPath)); err != nil {
		return err
	}
	return nil
}

func (l *local) deleteFiles(paths any) error {
	if fPaths, ok := paths.([]string); ok {
		for _, fPath := range fPaths {
			if err := l.deleteFile(fPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *local) getFilename(fh *multipart.FileHeader) string {
	return fmt.Sprintf("%s%s", strings.ReplaceAll(uuid.New().String(), "-", ""), l.getFileExt(fh))
}
func (l *local) getFileExt(fh *multipart.FileHeader) string {
	file, err := fh.Open()
	if err != nil {
		log.Panicln(err)
	}
	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Panicln(err)
	}
	mimetype.SetLimit(0)
	return mimetype.Detect(buffer).Extension()
}

func (l *local) save(destination, filename string, fh *multipart.FileHeader) {

	file, err := fh.Open()
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	basePath := fmt.Sprintf("./storage/%s", destination)
	err = os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
	dst, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, filename), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Panicln(err)
	}
}

func (l *local) resizeImage(destination, filename string, fh *multipart.FileHeader, width, height int) {
	ext := l.getFileExt(fh)
	file, err := fh.Open()
	if err != nil {
		log.Println(err)
	}
	file.Close()
	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, _ = jpeg.Decode(file)
	} else if ext == ".png" {
		img, _ = png.Decode(file)
	} else {
		log.Println(err)
	}
	basePath := fmt.Sprintf("./mnt/storage/%s", destination)
	err = os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	dst, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, filename), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	rImage, _ := mergi.Resize(img, uint(width), uint(height))
	if ext == ".jpg" || ext == ".jpeg" {
		_ = jpeg.Encode(dst, rImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if ext == ".png" {
		_ = png.Encode(dst, rImage)
	}

}
