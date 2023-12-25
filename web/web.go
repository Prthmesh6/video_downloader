package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Downloader interface {
	Download(uri string) (*io.PipeReader, error)
}

type VideoDownloader struct {
	urls      *[]string
	readers   []*io.PipeReader
	fileNames *[]string
}

func NewDownloader(urls *[]string) *VideoDownloader {
	readers := make([]*io.PipeReader, len(*urls))
	fileNames := make([]string, len(*urls))

	return &VideoDownloader{
		urls:      urls,
		readers:   readers,
		fileNames: &fileNames,
	}
}

func (v *VideoDownloader) Download(uri string) (r *io.PipeReader, err error) {
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		result, err := http.Get(uri)
		if err != nil {
			fmt.Println("Error requesting URL:", err)
			return
		}
		defer result.Body.Close()
		byteArray, err := ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		_, err = writer.Write(byteArray)
		if err != nil {
			fmt.Println("Error writing to pipe writer:", err)
		}
	}()

	return reader, err
}

func (v *VideoDownloader) DownloadVideos() (*[]*io.PipeReader, *[]string) {
	for index := range *v.urls {
		fileName := "file" + strconv.Itoa(index) + ".mp4"
		r, err := v.Download((*v.urls)[index])
		if err != nil {
			errMsg := fmt.Errorf("there was an error with file %v, the error is %v", (*v.urls)[index], err)
			fmt.Println(errMsg)
		}
		v.readers[index] = r
		(*v.fileNames)[index] = fileName
	}

	return &v.readers, v.fileNames
}
