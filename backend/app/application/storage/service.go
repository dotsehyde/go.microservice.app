package storage

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/config"
)

var drivers = map[string]string{
	"local":      "local",
	"uploadcare": "uploadcare",
}

type storageDriverService interface {
	uploadFile(dir string, f *multipart.FileHeader) (string, error)
	uploadCSVFile(dir string, f *multipart.FileHeader) (string, error)
	uploadFiles(dir string, files []*multipart.FileHeader) ([]string, error)
	uploadResizeImage(dir string, f *multipart.FileHeader, width, height int) (string, error)
	deleteFile(path any) error
	deleteFiles(paths any) error
}

type storage struct {
	WG        *sync.WaitGroup
	DataChan  chan any
	DoneChan  chan bool
	ErrorChan chan error
	Task      string
	diskType  storageDriverService
}

func NewService(wg *sync.WaitGroup) gateways.StorageService {
	dataChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)
	var diskType storageDriverService

	switch config.App().FilesystemDriver {
	default:
		diskType = newLocal()
	}
	return &storage{
		WG:        wg,
		DataChan:  dataChan,
		DoneChan:  doneChan,
		ErrorChan: errorChan,
		Task:      "delete_file",
		diskType:  diskType,
	}

}

func (s *storage) Disk(disk string) gateways.StorageService {
	switch disk {
	default:
		s.diskType = newLocal()
	}
	return s
}

func (s *storage) Listen() {
	for {
		select {
		case data := <-s.DataChan:
			go s.runTask(data)
		case err := <-s.ErrorChan:
			fmt.Println("error:: ", err)
		case <-s.DoneChan:
			return
		}
	}
}

func (s *storage) ExecuteTask(data any, task string) {
	s.WG.Add(1)
	s.DataChan <- data
	if task != "" {
		s.Task = task
	}
}

func (s *storage) Done() {
	s.DoneChan <- true
}

func (s *storage) Close() {
	close(s.DataChan)
	close(s.ErrorChan)
	close(s.DoneChan)
}

func (s *storage) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	return s.diskType.uploadFile(dir, f)
}
func (s *storage) UploadCSVFile(dir string, f *multipart.FileHeader) (string, error) {
	return s.diskType.uploadCSVFile(dir, f)
}
func (s *storage) UploadFiles(dir string, files []*multipart.FileHeader) ([]string, error) {
	return s.diskType.uploadFiles(dir, files)
}
func (s *storage) UploadResizeImage(dir string, f *multipart.FileHeader, width, height int) (string, error) {
	return s.diskType.uploadResizeImage(dir, f, width, height)
}

func (s *storage) runTask(data any) {
	defer s.WG.Done()
	switch s.Task {
	case "delete_file":
		if err := s.deleteFile(data); err != nil {
			s.ErrorChan <- err
		}
	case "delete_files":
		if err := s.deleteFiles(data); err != nil {
			s.ErrorChan <- err
		}
	}
}

func (s *storage) deleteFile(path any) error {
	return s.diskType.deleteFile(path)
}

func (s *storage) deleteFiles(paths any) error {
	return s.diskType.deleteFiles(paths)
}
