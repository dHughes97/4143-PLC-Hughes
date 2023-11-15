package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for _, url := range urls {
		filename := filepath.Base(url)
		err := downloadImage(url, filename)
		if err != nil {
			fmt.Printf("Error downloading %s: %s\n", url, err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			filename := filepath.Base(url)
			err := downloadImage(url, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %s\n", url, err)
				ch <- "" // Sending an empty string on error
				return
			}
			ch <- filename
		}(url)
	}

	for range urls {
		downloadedFilename := <-ch
		if downloadedFilename != "" {
			fmt.Printf("Downloaded: %s\n", downloadedFilename)
		}
	}
	close(ch) // Close the channel after all downloads are handled
}

func downloadImage(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	urls := []string{
		"https://img.freepik.com/premium-photo/corgi-puppy-is-wooden-table_106368-1108.jpg",
		"https://images.freeimages.com/images/large-previews/77c/nemo-the-horse-1339807.jpg",
		"https://cdn.pixabay.com/photo/2017/07/27/13/07/electric-car-2545290_1280.png",
		"https://cdn.stocksnap.io/img-thumbs/960w/simple-pumpkin_WPXAMOFWKY.jpg",
		"https://images.pexels.com/photos/18111476/pexels-photo-18111476/free-photo-of-traditional-door-number.jpeg",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}
