package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	fmt.Print("\n\n")
	color.Red(`/\ |_| |) | ()   |\| () /? |\/| /\ |_ | ~/_ [-`)
	fmt.Print("\n\n")

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

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("DIR: %s", path)

	cmd := fmt.Sprintf("-i %s/files/big_buck_bunny_480p_stereo.avi -filter:a loudnorm %s/files/big_buck_bunny_480p_stereo_NA.avi", path, path)
	fmt.Printf("\nPreparing to execute command: 'ffmpeg %s'", cmd)

	out, err := exec.Command("ffmpeg", cmd).Output()

	if err != nil {
		fmt.Printf("\n\nERROR encountered %s\n", err)
		log.Fatal(err)
	}

	fmt.Printf("OUTPUT: %s", out)

	// osd := fmt.Sprintf(" * Reading %d files...", len(files))

	// bar := progressbar.NewOptions(len(files),
	// 	progressbar.OptionEnableColorCodes(true),
	// 	progressbar.OptionShowBytes(true),
	// 	// progressbar.OptionSetWidth(15),
	// 	progressbar.OptionSetDescription(osd),
	// 	progressbar.OptionSetTheme(progressbar.Theme{
	// 		Saucer:        "[green]=[reset]",
	// 		SaucerHead:    "[green]>[reset]",
	// 		SaucerPadding: " ",
	// 		BarStart:      "[",
	// 		BarEnd:        "]",
	// 	}))

	// names := make([]string, len(files))

	// for i := 0; i < len(files); i++ {
	// 	// fmt.Printf("\t\tFound video file '%s'", files[i].Name())
	// 	bar.Add(1)
	// 	names = append(names, files[i].Name())
	// 	name := strings.Split(files[i].Name(), ".")[0]
	// 	ext := strings.Split(files[i].Name(), ".")[1]
	// 	output := name + "_NA"
	// 	rawPath := fmt.Sprintf("%s/", dir)
	// 	osPath := filepath.FromSlash(rawPath)

	// 	// ffmpeg -i COPY_unchanged_HTTHD.mp4 -filter:a loudnorm COPY_loudnorm_1.mp4
	// 	fileInfo := fmt.Sprintf("-i %s%s.%s -filter:a loudnorm %s%s.%s", osPath, name, ext, osPath, output, ext)
	// 	fmt.Printf("\n\nPreparing to execute command: 'ffmpeg %s'", fileInfo)

	// 	cmd := exec.Command("ffmpeg", fileInfo)
	// 	err := cmd.Run() // spawns a goroutine
	// 	if err != nil {
	// 		fmt.Printf("\nERROR encountered %s\n", err)
	// 		log.Fatal(err)
	// 	}
	// 	time.Sleep(140 * time.Millisecond)
	// }

	// fmt.Printf("\n * All files updated: %v\n", names)

}
