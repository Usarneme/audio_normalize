package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var formats = [...]string{"3g2", "3gp", "4xm", "a64", "aa", "aac", "ac3", "acm", "act", "adf", "adp", "ads", "adts", "adx", "aea", "afc", "aiff", "aix", "alaw", "alias_pix", "alp", "amr", "amrnb", "amrwb", "anm", "apc", "ape", "apm", "apng", "aptx", "aptx_hd", "aqtitle", "argo_asf", "asf", "asf_o", "asf_stream", "ass", "ast", "au", "av1", "avfoundation", "avi", "avm2", "avr", "avs", "avs2", "bethsoftvid", "bfi", "bfstm", "bin", "bink", "bit", "bmp_pipe", "bmv", "boa", "brender_pix", "brstm", "c93", "caf", "cavsvideo", "cdg", "cdxl", "cine", "codec2", "codec2raw", "concat", "crc", "dash", "data", "daud", "dcstr", "dds_pipe", "derf", "dfa", "dhav", "dirac", "dnxhd", "dpx_pipe", "dsf", "dsicin", "dss", "dts", "dtshd", "dv", "dvbsub", "dvbtxt", "dvd", "dxa", "ea", "ea_cdata", "eac3", "epaf", "exr_pipe", "f32be", "f32le", "f4v", "f64be", "f64le", "ffmetadata", "fifo", "fifo_test", "film_cpk", "filmstrip", "fits", "flac", "flic", "flv", "framecrc", "framehash", "framemd5", "frm", "fsb", "fwse", "g722", "g723_1", "g726", "g726le", "g729", "gdv", "genh", "gif", "gif_pipe", "gsm", "gxf", "h261", "h263", "h264", "hash", "hca", "hcom", "hds", "hevc", "hls", "hnm", "ico", "idcin", "idf", "iff", "ifv", "ilbc", "image2", "image2pipe", "ingenient", "ipmovie", "ipod", "ircam", "ismv", "iss", "iv8", "ivf", "ivr", "j2k_pipe", "jacosub", "jpeg_pipe", "jpegls_pipe", "jv", "kux", "kvag", "latm", "lavfi", "live_flv", "lmlm4", "loas", "lrc", "lvf", "lxf", "m4v", "matroska", "md5", "mgsts", "microdvd", "mjpeg", "mjpeg_2000", "mkvtimestamp_v2", "mlp", "mlv", "mm", "mmf", "mov", "mp4", "m4a", "3gp", "3g2", "mj2", "mp2", "mp3", "mp4", "mpc", "mpc8", "mpeg", "mpeg1video", "mpeg2video", "mpegts", "mpegtsraw", "mpegvideo", "mpjpeg", "mpl2", "mpsub", "msf", "msnwctcp", "mtaf", "mtv", "mulaw", "musx", "mv", "mvi", "mxf", "mxf_d10", "mxf_opatom", "mxg", "nc", "nistsphere", "nsp", "nsv", "null", "nut", "nuv", "oga", "ogg", "ogv", "oma", "opus", "paf", "pam_pipe", "pbm_pipe", "pcx_pipe", "pgm_pipe", "pgmyuv_pipe", "pictor_pipe", "pjs", "pmp", "png_pipe", "pp_bnk", "ppm_pipe", "psd_pipe", "psp", "psxstr", "pva", "pvf", "qcp", "qdraw_pipe", "r3d", "rawvideo", "realtext", "redspark", "rl2", "rm", "roq", "rpl", "rsd", "rso", "rtp", "rtp_mpegts", "rtsp", "s16be", "s16le", "s24be", "s24le", "s32be", "s32le", "s337m", "s8", "sami", "sap", "sbc", "sbg", "scc", "sdl", "sdp", "sdr2", "sds", "sdx", "segment", "ser", "sgi_pipe", "shn", "siff", "singlejpeg", "sln", "smjpeg", "smk", "smoothstreaming", "smush", "sol", "sox", "spdif", "spx", "srt", "stl", "stream_segment", "streamhash", "subviewer", "subviewer1", "sunrast_pipe", "sup", "svag", "svcd", "svg_pipe", "swf", "tak", "tedcaptions", "tee", "thp", "tiertexseq", "tiff_pipe", "tmv", "truehd", "tta", "tty", "txd", "ty", "u16be", "u16le", "u24be", "u24le", "u32be", "u32le", "u8", "uncodedframecrc", "v210", "v210x", "vag", "vc1", "vc1test", "vcd", "vidc", "vividas", "vivo", "vmd", "vob", "vobsub", "voc", "vpk", "vplayer", "vqf", "w64", "wav", "wc3movie", "webm", "webm_chunk", "webm_dash_manifest", "webp", "webp_pipe", "webvtt", "wsaud", "wsd", "wsvqa", "wtv", "wv", "wve", "xa", "xbin", "xmv", "xpm_pipe", "xvag", "xwd_pipe", "xwma", "yop", "yuv4mpegpipe"}

func contains(arr [len(formats)]string, query string) bool {
	for _, val := range arr {
		if val == query {
			return true
		}
	}
	return false
}

func collectFilenames(dir string) ([]string, error) {
	color.Cyan("  [func] Reading all files in the directory '%s/'", dir)

	var files []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == true {
				color.Blue("\t'%s' skipped - folder.", info.Name())
			} else {
				lastIndex := strings.LastIndex(info.Name(), ".")
				ext := info.Name()[lastIndex+1:]
				if contains(formats, ext) {
					color.Magenta("\t'%s' appended to workgroup list - a/v file.", path)
					files = append(files, path)
				} else {
					color.Blue("\t'%s' skipped - non-a/v file.", info.Name())
				}
			}
			return nil
		})
	// after walking all directories and subdirectories...
	if err != nil {
		log.Println(err)
		return nil, err
	}
	color.Green("  [func] Returning: %v", files)
	return files, nil
}

func doNormalization(wg *sync.WaitGroup, file string) {
	defer wg.Done()
	color.Cyan("  [func] Preparing to normalize file at '%s", file)

	lastIndex := strings.LastIndex(file, "/")
	name := file[lastIndex+1:]
	outName := fmt.Sprintf("output/%s", name)

	cmd := exec.Command("ffmpeg", "-loglevel", "quiet", "-i", file, "-filter:a", "loudnorm", "-c:v", "copy", outName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	e := cmd.Run()
	if e != nil {
		color.Red("  [ERROR] - %v\n", e)
	}
	color.Green("  [func] Returning from doNormalize on '%s'", outName)
}

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
	files, readErr := collectFilenames(dir)
	if readErr != nil {
		color.Red("[ERROR]: %s", readErr)
	}

	// worker waitgroup to ensure all goroutines are completed before exiting program
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go doNormalization(&wg, file)
	}

	color.Red("  [main] Waiting for workers to finish")
	wg.Wait()
	color.Red("  [main] Completed. Exiting.")

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
	// 	// fmt.Printf("\t\tFound video file "%s'", files[i].Name())
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
