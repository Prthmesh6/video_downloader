package main

import (
	web "GO_EXERCISE/web"
	zip "GO_EXERCISE/zip"
	"fmt"
	"time"
)

func main() {

	// zipper := zip.New()
	time_start := time.Now()
	uri1 := "https://filesamples.com/samples/video/mp4/sample_1280x720_surfing_with_audio.mp4"
	uri2 := "https://filesamples.com/samples/video/mp4/sample_960x400_ocean_with_audio.mp4"

	urls := []string{uri1, uri2}

	downloader := web.NewDownloader(&urls)

	pipeReaders, filenames := downloader.DownloadVideos()

	zipper := zip.New(*pipeReaders, filenames)

	zipR, err := zipper.Archive()
	if err != nil {
		msg := fmt.Errorf("error while archiving :- %v", err)
		fmt.Println(msg)
	}

	err = zipper.CreateZip(zipR)
	if err != nil {
		msg := fmt.Errorf("error while creating zip file :- %v", err)
		fmt.Println(msg)
	}

	time_end := time.Now()
	fmt.Println("Time taken to complete the program ", time_end.Sub(time_start))

}
