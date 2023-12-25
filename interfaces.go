package main

import "io"

type Archiver interface {
	Archive() (err error)
}

type Downloader interface {
	Download() (*[]*io.PipeReader, *[]string)
}
