package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

type zipper struct {
	readers []*io.PipeReader
	names   *[]string
}

type Archiver interface {
	Archive() (outR io.Reader, err error)
}

func New(readers []*io.PipeReader, names *[]string) *zipper {
	return &zipper{
		readers: readers,
		names:   names,
	}
}

func (z zipper) Archive() (outR io.Reader, err error) {
	zipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipBuffer)

	for i, reader := range z.readers {
		fileName := (*z.names)[i]
		fileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer zipWriter.Close()

		_, err = io.Copy(fileWriter, reader)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return zipBuffer, nil
}

func (z zipper) CreateZip(zipR io.Reader) error {
	zipW, err := os.Create("result.zip")

	if err != nil {
		return err
	}

	_, err = io.Copy(zipW, zipR)
	defer zipW.Close()

	return err
}
