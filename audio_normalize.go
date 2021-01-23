package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Print("\n")
	color.Red(`/\ |_| |) | ()   |\| () /? |\/| /\ |_ | ~/_ [-`)
	fmt.Print("\n")

	if len(os.Args) < 2 {
		fmt.Println("[usage error] - please provide a directory in which to look for video files.")
		fmt.Println("Try again, such as: ~ ./audio_normalize movies/")
		os.Exit(-1)
	}

	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 1 {
		fmt.Printf("Found %d files in dir '%s/'\n", len(files), dir)
	} else {
		fmt.Printf("Found %d file in dir '%s/'\n", len(files), dir)
	}

	osd := fmt.Sprintf(" * Reading %d files...", len(files))

	bar := progressbar.NewOptions(len(files),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		// progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(osd),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	names := make([]string, len(files))

	for i := 0; i < len(files); i++ {
		// fmt.Printf("\t\tFound video file '%s'", files[i].Name())
		bar.Add(1)
		names = append(names, files[i].Name())
		time.Sleep(140 * time.Millisecond)
	}

	fmt.Printf("\n * All files: %v\n", names)

}
