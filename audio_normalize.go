package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

func collectFilenames(dir string) ([]os.FileInfo, error) {
	color.Cyan("Reading all files in the directory '%s/'", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if len(files) > 1 {
		fmt.Printf("Found %d files in dir '%s/'\n", len(files), dir)
	} else {
		fmt.Printf("Found %d file in dir '%s/'\n", len(files), dir)
	}

	if false {
		errStr := fmt.Sprintf("Problem reading directory %s and collection filenames.", dir)
		return nil, errors.New(errStr)
	}

	return files, nil
}

func main() {
	fmt.Print("\n\n")
	color.Red(`/\ |_| |) | ()   |\| () /? |\/| /\ |_ | ~/_ [-`)
	fmt.Print("\n")

	if len(os.Args) < 2 {
		fmt.Println("[usage error] - please provide a directory in which to look for video files.")
		fmt.Println("Try again, such as: ~ ./audio_normalize movies/")
		os.Exit(-1)
	}

	// create a slice to hold all of the subdirectory/filenames of files to be audio normalized
	filepaths := []string{}

	dir := os.Args[1]
	files, readErr := collectFilenames(dir)
	if readErr != nil {
		color.Red("[ERROR]: %s", readErr)
	}
	// color.Green("After collecting file names:")
	// for _, v := range files {
	// 	fmt.Printf("\n\t%+v\n", v)
	// }

	// example result of printf for a file:
	// &{name:big_buck_bunny_480p_stereo.avi
	// size:156506028
	// mode:420
	// modTime:{wall:26551519 ext:63747019898 loc:0x117d380}
	// sys:{Dev:16777220 Mode:33188 Nlink:1 Ino:3075165 Uid:501 Gid:20 Rdev:0 Pad_cgo_0:[0 0 0 0] Atimespec:{Sec:1611537984 Nsec:934362000} Mtimespec:{Sec:1611423098 Nsec:26551519} Ctimespec:{Sec:1611538228 Nsec:609836550} Birthtimespec:{Sec:1611423097 Nsec:926494001}
	// Size:156506028 Blocks:305680 Blksize:4096 Flags:0 Gen:0 Lspare:0 Qspare:[0 0]}}

	// example printf of a directory:
	// &{name:tmp size:64
	// mode:2147484141
	// modTime:{wall:362958889 ext:63747139758 loc:0x117d380}
	// sys:{Dev:16777220 Mode:16877 Nlink:2 Ino:3125540 Uid:501 Gid:20 Rdev:0 Pad_cgo_0:[0 0 0 0] Atimespec:{Sec:1611542958 Nsec:362958889} Mtimespec:{Sec:1611542958 Nsec:362958889} Ctimespec:{Sec:1611542958 Nsec:362958889} Birthtimespec:{Sec:1611542958 Nsec:362958889}
	// Size:64 Blocks:0 Blksize:4096 Flags:0 Gen:0 Lspare:0 Qspare:[0 0]}}

	fmt.Println("\nLooping through returned os.FileInfo slice for filenames...")
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			// recurse and run again in that directory
			color.Magenta("Folder found! %s", files[i].Name())
			// TODO => files, readErr := collectFilenames(files[i].Name())
		} else {
			// TODO verify filetype is avi, mov, etc...
			filepaths = append(filepaths, files[i].Name())
		}
	}
	// interface for os.FileInfo as returned by ioutil.ReadDir(), files is a slice of these:
	// type FileInfo interface {
	//   Name() string       // base name of the file
	//   Size() int64        // length in bytes for regular files; system-dependent for others
	//   Mode() FileMode     // file mode bits
	//   ModTime() time.Time // modification time
	//   IsDir() bool        // abbreviation for Mode().IsDir()
	//   Sys() interface{}   // underlying data source (can return nil)
	// }

	color.Green("After looping, created filepaths slice of: %s", filepaths)

	// cmd := fmt.Sprintf("-i files/big_buck_bunny_480p_stereo.avi -filter:a loudnorm files/big_buck_bunny_480p_stereo_NA.avi", path, path)
	// fmt.Printf("\nPreparing to execute command: 'ffmpeg %s'", cmd)

	// working code for executing a single ffmpeg routine
	// cmd := exec.Command("ffmpeg", "-i", "files/big_buck_bunny_480p_stereo.avi", "-filter:a", "loudnorm", "files/big_buck_bunny_480p_stereo_NA.avi")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// e := cmd.Run()
	// if e != nil {
	// 	color.Red("[ERROR]")
	// 	fmt.Printf("%s", e)
	// }
	// fmt.Println("After ffmpeg run")

	// if err != nil {
	// 	fmt.Printf("\n\nERROR encountered %s\n", err)
	// 	log.Fatal(err)
	// }

	// fmt.Printf("OUTPUT: %s", out)

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
