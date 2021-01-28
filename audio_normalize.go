package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var formats = [...]string{"3g2", "3gp", "4xm", "a64", "aa", "aac", "ac3", "acm", "act", "adf", "adp", "ads", "adts", "adx", "aea", "afc", "aiff", "aix", "alaw", "alias_pix", "alp", "amr", "amrnb", "amrwb", "anm", "apc", "ape", "apm", "apng", "aptx", "aptx_hd", "aqtitle", "argo_asf", "asf", "asf_o", "asf_stream", "ass", "ast", "au", "av1", "avfoundation", "avi", "avm2", "avr", "avs", "avs2", "bethsoftvid", "bfi", "bfstm", "bin", "bink", "bit", "bmp_pipe", "bmv", "boa", "brender_pix", "brstm", "c93", "caf", "cavsvideo", "cdg", "cdxl", "cine", "codec2", "codec2raw", "concat", "crc", "dash", "data", "daud", "dcstr", "dds_pipe", "derf", "dfa", "dhav", "dirac", "dnxhd", "dpx_pipe", "dsf", "dsicin", "dss", "dts", "dtshd", "dv", "dvbsub", "dvbtxt", "dvd", "dxa", "ea", "ea_cdata", "eac3", "epaf", "exr_pipe", "f32be", "f32le", "f4v", "f64be", "f64le", "ffmetadata", "fifo", "fifo_test", "film_cpk", "filmstrip", "fits", "flac", "flic", "flv", "framecrc", "framehash", "framemd5", "frm", "fsb", "fwse", "g722", "g723_1", "g726", "g726le", "g729", "gdv", "genh", "gif", "gif_pipe", "gsm", "gxf", "h261", "h263", "h264", "hash", "hca", "hcom", "hds", "hevc", "hls", "hnm", "ico", "idcin", "idf", "iff", "ifv", "ilbc", "image2", "image2pipe", "ingenient", "ipmovie", "ipod", "ircam", "ismv", "iss", "iv8", "ivf", "ivr", "j2k_pipe", "jacosub", "jpeg_pipe", "jpegls_pipe", "jv", "kux", "kvag", "latm", "lavfi", "live_flv", "lmlm4", "loas", "lrc", "lvf", "lxf", "m4v", "matroska", "md5", "mgsts", "microdvd", "mjpeg", "mjpeg_2000", "mkvtimestamp_v2", "mlp", "mlv", "mm", "mmf", "mov", "mp4", "m4a", "3gp", "3g2", "mj2", "mp2", "mp3", "mp4", "mpc", "mpc8", "mpeg", "mpeg1video", "mpeg2video", "mpegts", "mpegtsraw", "mpegvideo", "mpjpeg", "mpl2", "mpsub", "msf", "msnwctcp", "mtaf", "mtv", "mulaw", "musx", "mv", "mvi", "mxf", "mxf_d10", "mxf_opatom", "mxg", "nc", "nistsphere", "nsp", "nsv", "null", "nut", "nuv", "oga", "ogg", "ogv", "oma", "opus", "paf", "pam_pipe", "pbm_pipe", "pcx_pipe", "pgm_pipe", "pgmyuv_pipe", "pictor_pipe", "pjs", "pmp", "png_pipe", "pp_bnk", "ppm_pipe", "psd_pipe", "psp", "psxstr", "pva", "pvf", "qcp", "qdraw_pipe", "r3d", "rawvideo", "realtext", "redspark", "rl2", "rm", "roq", "rpl", "rsd", "rso", "rtp", "rtp_mpegts", "rtsp", "s16be", "s16le", "s24be", "s24le", "s32be", "s32le", "s337m", "s8", "sami", "sap", "sbc", "sbg", "scc", "sdl", "sdp", "sdr2", "sds", "sdx", "segment", "ser", "sgi_pipe", "shn", "siff", "singlejpeg", "sln", "smjpeg", "smk", "smoothstreaming", "smush", "sol", "sox", "spdif", "spx", "srt", "stl", "stream_segment", "streamhash", "subviewer", "subviewer1", "sunrast_pipe", "sup", "svag", "svcd", "svg_pipe", "swf", "tak", "tedcaptions", "tee", "thp", "tiertexseq", "tiff_pipe", "tmv", "truehd", "tta", "tty", "txd", "ty", "u16be", "u16le", "u24be", "u24le", "u32be", "u32le", "u8", "uncodedframecrc", "v210", "v210x", "vag", "vc1", "vc1test", "vcd", "vidc", "vividas", "vivo", "vmd", "vob", "vobsub", "voc", "vpk", "vplayer", "vqf", "w64", "wav", "wc3movie", "webm", "webm_chunk", "webm_dash_manifest", "webp", "webp_pipe", "webvtt", "wsaud", "wsd", "wsvqa", "wtv", "wv", "wve", "xa", "xbin", "xmv", "xpm_pipe", "xvag", "xwd_pipe", "xwma", "yop", "yuv4mpegpipe"}

func collectFilenames(dir string) ([]string, error) {
	color.Cyan("[func] Collecting all filenames within '%s/'", dir)

	var files []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			length := len(info.Name())
			idx := 0
			var shortName string
			if length > 20 {
				idx = length % 20
				shortName = "..." + info.Name()[idx:]
			} else {
				shortName = info.Name()[idx:]
			}

			if info.IsDir() == true {
				color.Blue("\tSKIPPED \t folder\t '%s'", shortName)
			} else {
				lastIndex := strings.LastIndex(info.Name(), ".")
				ext := info.Name()[lastIndex+1:]
				if contains(formats, ext) {
					color.Magenta("\tAPPENDED\t a/v file\t '%s'", shortName)
					files = append(files, path)
				} else {
					color.Blue("\tSKIPPED \t non-a/v file\t '%s'", shortName)
				}
			}
			return nil
		})
	// after walking all directories and subdirectories...
	if err != nil {
		log.Println(err)
		return nil, err
	}
	color.Green("[func] Returning from collectFilenames - %v", files)
	return files, nil
}

func doNormalization(wg *sync.WaitGroup, file string) {
	defer wg.Done()
	color.Cyan("[func] Preparing to normalize file at '%s", file)

	lastIndex := strings.LastIndex(file, "/")
	name := file[lastIndex+1:]
	outName := fmt.Sprintf("output/%s", name)

	// command with stderr passed to a temp file descriptor
	// cmd := exec.Command("exec", "3>&1;")
	//  "sout=$(ffmpeg -i files/tmp/some.avi -filter:a loudnorm -c:v copy NA_big_buck_bunny.avi -loglevel error 2>&1 1>&3); exec 3>&-; echo $sout;")

	cmd := exec.Command("ffmpeg", "-loglevel", "error", "-i", file, "-filter:a", "loudnorm", "-c:v", "copy", outName)

	// duplicate output files are ignored/not overwritten by default, so we don't need any interactivity via stdin
	// '-loglevel error' suppresses stdout logging of successful processing
	// stderr output is piped so we can use it for logging results/informing the user
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		color.Red("  [ERROR] %v", err)
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		errMsg := scanner.Text()
		color.Magenta("[func] Errors encoutered: Calling update log")
		updateLog(errMsg)
	}
	color.Green("[func] Returning from doNormalize - '%s'", outName)
}

func updateLog(msg string) {
	color.Cyan("[func] Outputting error msg to file: '%s'", msg)
	if fileExists("output/log.txt") == true {
		// already exists, append to file
		f, err := os.OpenFile("output/log.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		dt := time.Now()
		_, err = fmt.Fprintln(f, dt.String())
		if err != nil {
			fmt.Println(err)
		}

		_, err = fmt.Fprintln(f, msg)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// create the log file
		f, err := os.Create("output/log.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		dt := time.Now()
		f.WriteString(dt.String())
		fmtStr := fmt.Sprintf("\n%s\n", msg)
		f.WriteString(fmtStr)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func contains(arr [len(formats)]string, query string) bool {
	for _, val := range arr {
		if val == query {
			return true
		}
	}
	return false
}

func main() {
	color.Red(`/\ |_| |) | ()   |\| () /? |\/| /\ |_ | ~/_ [-`)

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

	color.Red("[main] Assigning worker tasks")
	for _, file := range files {
		wg.Add(1)
		go doNormalization(&wg, file)
	}

	wg.Wait()
	color.Red("[main] Workers completed. Exiting.")
	color.Yellow("[main] Please check output/log.txt for any error details.")
}
