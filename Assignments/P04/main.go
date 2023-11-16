package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// adding counts to give unique file names
var seqCounter, conCounter int

// Sequential version of the image downloader.
// Takes a slice of strings called urls. runs through the slice
// increments the count and concatenates the incremented number at the end
// of said img.
// Calls the downloadImage function which takes a url and the filename.
// Does error checking if the image was not downloaded properly.
func downloadImagesSequential(urls []string) {
	for _, url := range urls {
		seqCounter++ // increment count
		filename := fmt.Sprintf("seqImg%d.jpg", seqCounter)
		err := downloadImage(url, filename)
		if err != nil {
			fmt.Printf("Error downloading %s: %s\n", url, err)
		}
	}
}

// Concurrent version of the image downloader.
// This function will create a dynamic array of strings pass them to ch.
// It will then run through each image sequentially(but not really). Then
// hopping into said go routine and allows for each image to be ran
// concurrently in the background. It then uses some concatentation like before
// to increment the count of the file name. Checks for errs in the download and if
// there is an err and nothing is found return an empty string.
//
// completed will keep track of downloads. When the filename is downloaded it is passes to
// `downloadedFilename` and prints the name of the file. The last if statement checks
// if the images were downloaded and closes the channel.
func downloadImagesConcurrent(urls []string) {
	ch := make(chan string) // making dynamic array of strings passing it to ch

	for _, url := range urls {
		go func(url string) { // run concurrently key word `go` allows the function to run in background
			conCounter++ // incrememnt count
			filename := fmt.Sprintf("conImg%d.jpg", conCounter)
			err := downloadImage(url, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %s\n", url, err)
				ch <- "" // Sending an empty string on error
				return
			}
			ch <- filename
		}(url)
	}

	// Counter to keep track of the completed downloads
	completed := 0
	for range urls {
		downloadedFilename := <-ch
		if downloadedFilename != "" {
			fmt.Printf("Downloaded: %s\n", downloadedFilename)
		}
		completed++
		if completed == len(urls) { // Check if all downloads are completed
			close(ch)
		}
	}
}

// Function to download the image. Helper Function!!!
// downloadImage downloads an image from a URL and saves it with a specified filename.
// Make an HTTP GET request to the provided URL
// If there's an error in the HTTP request, return the error immediately
// Ensure that the response body will be closed when this function returns
// Create a new file with the specified filename
// If there's an error creating the file, return the error
// Ensure that the file will be closed when this function returns
// Copy the contents of the HTTP response body (image data) into the created file
// If there's an error while copying data, return the error
// Return nil to indicate that no errors occurred during the process
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
	start := time.Now() // start timer
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now() // start time
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}
