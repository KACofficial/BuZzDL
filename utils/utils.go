package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

var ascii_art = ` ____        _____    ____  _     
| __ ) _   _|__  /___|  _ \| |    
|  _ \| | | | / /|_  / | | | |    
| |_) | |_| |/ /_ / /| |_| | |___ 
|____/ \__,_/____/___|____/|_____|`

func printASCII() {
	art := strings.TrimSuffix(string(ascii_art), "\n")

	// Print the ASCII art in yellow color
	yellow := color.New(color.FgYellow, color.Bold)
	yellow.Println(art)
}

func StartProgram(url string, format string, output string) {
	printASCII()
	fmt.Print("\n")
	fmt.Printf("Starting BuZzDL on '%s' with the file format '%s'\n", url, format)
	video_title := getVideoTitle(url)
	video_src, err := getVideoSrc(url, video_title)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	video_mp4_filename := convertUrlToFilename(url)

	fmt.Printf("Found Video '%s', Starting Download.\n", video_title)

	fmt.Print("Downloading video...")

	mp4_location, err := downloadFile(video_mp4_filename, video_src, output)


	if err != nil {
		fmt.Println("Faild: ", err)
		return
	}
	fmt.Println("Done!")
	formatted_location := strings.ReplaceAll(mp4_location, ".mp4", "."+format)
	if format != "mp4" {
		fmt.Printf("Converting to %s(Slow)...", format)
		err := convertAndDelete(mp4_location, format)
		if err != nil {
			fmt.Println("Failed: ", err)
			return
		}
		fmt.Println("Done!")
	}

	fmt.Printf("Video '%s' has been downloaded to '%s'\n", video_title, formatted_location)
}



func convertUrlToFilename(url string) string {
	// Find the last slash and extract the substring from there
	splitUrl := strings.Split(url, "/")
	filename := splitUrl[len(splitUrl)-1]

	// remove any get parameters
	filename = strings.Split(filename, "?")[0]

	// Replace .html with .mp4
	mp4Filename := strings.Replace(filename, ".html", ".mp4", 1)

	return mp4Filename
}

func getVideoTitle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	title := doc.Find("title").Text()
	return title
}


func getVideoSrc(url string, title string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	var videoSrc string
	if strings.Contains(url, "shorts") {
		found := false
		doc.Find("div.swiper-slide.overflow-hidden").Each(func(i int, s *goquery.Selection) {
			divTitle, exists := s.Attr("data-title")
			if exists && strings.TrimSpace(divTitle) == title {
				video := s.Find("video").First()
				source := video.Find("source").First()
				src, exists := source.Attr("src")
				if exists {
					videoSrc = src
					found = true
					// fmt.Printf("Found video source: '%s'\n", videoSrc)
					return
				}
			}
		})
		if !found {
			return "", fmt.Errorf("no video found with matching title '%s'", title)
		}
	} else {
		video := doc.Find("video").First()
		source := video.Find("source").First()
		src, exists := source.Attr("src")
		if exists {
			videoSrc = src
		}
	}
	return videoSrc, nil
}

func downloadFile(fileName, url, location string) (string, error) {
	// Determine if the location is a directory or a file
	info, err := os.Stat(location)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	// Check if location is a directory
	isDir := false
	if err == nil && info.IsDir() {
		isDir = true
	} else if os.IsNotExist(err) {
		// If location does not exist, check if it's a directory based on trailing slash
		if strings.HasSuffix(location, "/") || !strings.Contains(filepath.Base(location), ".") {
			isDir = true
			if err := os.MkdirAll(location, os.ModePerm); err != nil {
				return "", err
			}
		} else {
			// It's a file, create necessary directories
			if err := os.MkdirAll(filepath.Dir(location), os.ModePerm); err != nil {
				return "", err
			}
		}
	}

	// If location is a directory, construct the full file path
	if isDir {
		location = filepath.Join(location, fileName)
	}

	// Create the file
	out, err := os.Create(location)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return location, nil
}

func convertAndDelete(fileName string, format string) error {
	// Convert from MP4
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("input file %s does not exist", fileName)
	}

	// Prepare the new file name with the desired format
	newFileName := strings.TrimSuffix(fileName, ".mp4") + "." + format

	cmd := exec.Command("ffmpeg", "-i", fileName, newFileName)
	err := cmd.Run()

	if err != nil {
		os.Remove(newFileName)
		return fmt.Errorf("ffmpeg error: %v", err)
	}

	err = os.Remove(fileName)

	if err != nil {
		return fmt.Errorf("error removing original file: %v", err)
	}

	return nil
}
