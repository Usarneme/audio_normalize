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

// var formats = [...]string{"3g2", "3gp", "4xm", "a64", "aa", "aac", "ac3", "acm", "act", "adf", "adp", "ads", "adts", "adx", "aea", "afc", "aiff", "aix", "alaw", "alias_pix", "alp", "amr", "amrnb", "amrwb", "anm", "apc", "ape", "apm", "apng", "aptx", "aptx_hd", "aqtitle", "argo_asf", "asf", "asf_o", "asf_stream", "ass", "ast", "au", "av1", "avfoundation", "avi", "avm2", "avr", "avs", "avs2", "bethsoftvid", "bfi", "bfstm", "bin", "bink", "bit", "bmp_pipe", "bmv", "boa", "brender_pix", "brstm", "c93", "caf", "cavsvideo", "cdg", "cdxl", "cine", "codec2", "codec2raw", "concat", "crc", "dash", "data", "daud", "dcstr", "dds_pipe", "derf", "dfa", "dhav", "dirac", "dnxhd", "dpx_pipe", "dsf", "dsicin", "dss", "dts", "dtshd", "dv", "dvbsub", "dvbtxt", "dvd", "dxa", "ea", "ea_cdata", "eac3", "epaf", "exr_pipe", "f32be", "f32le", "f4v", "f64be", "f64le", "ffmetadata", "fifo", "fifo_test", "film_cpk", "filmstrip", "fits", "flac", "flic", "flv", "framecrc", "framehash", "framemd5", "frm", "fsb", "fwse", "g722", "g723_1", "g726", "g726le", "g729", "gdv", "genh", "gif", "gif_pipe", "gsm", "gxf", "h261", "h263", "h264", "hash", "hca", "hcom", "hds", "hevc", "hls", "hnm", "ico", "idcin", "idf", "iff", "ifv", "ilbc", "image2", "image2pipe", "ingenient", "ipmovie", "ipod", "ircam", "ismv", "iss", "iv8", "ivf", "ivr", "j2k_pipe", "jacosub", "jpeg_pipe", "jpegls_pipe", "jv", "kux", "kvag", "latm", "lavfi", "live_flv", "lmlm4", "loas", "lrc", "lvf", "lxf", "m4v", "matroska", "md5", "mgsts", "microdvd", "mjpeg", "mjpeg_2000", "mkv", "mkvtimestamp_v2", "mlp", "mlv", "mm", "mmf", "mov", "mp4", "m4a", "3gp", "3g2", "mj2", "mp2", "mp4", "mpc", "mpc8", "mpeg", "mpeg1video", "mpeg2video", "mpegts", "mpegtsraw", "mpegvideo", "mpjpeg", "mpl2", "mpsub", "msf", "msnwctcp", "mtaf", "mtv", "mulaw", "musx", "mv", "mvi", "mxf", "mxf_d10", "mxf_opatom", "mxg", "nc", "nistsphere", "nsp", "nsv", "null", "nut", "nuv", "oga", "ogg", "ogv", "oma", "opus", "paf", "pam_pipe", "pbm_pipe", "pcx_pipe", "pgm_pipe", "pgmyuv_pipe", "pictor_pipe", "pjs", "pmp", "png_pipe", "pp_bnk", "ppm_pipe", "psd_pipe", "psp", "psxstr", "pva", "pvf", "qcp", "qdraw_pipe", "r3d", "rawvideo", "realtext", "redspark", "rl2", "rm", "roq", "rpl", "rsd", "rso", "rtp", "rtp_mpegts", "rtsp", "s16be", "s16le", "s24be", "s24le", "s32be", "s32le", "s337m", "s8", "sami", "sap", "sbc", "sbg", "scc", "sdl", "sdp", "sdr2", "sds", "sdx", "segment", "ser", "sgi_pipe", "shn", "siff", "singlejpeg", "sln", "smjpeg", "smk", "smoothstreaming", "smush", "sol", "sox", "spdif", "spx", "srt", "stl", "stream_segment", "streamhash", "subviewer", "subviewer1", "sunrast_pipe", "sup", "svag", "svcd", "svg_pipe", "swf", "tak", "tedcaptions", "tee", "thp", "tiertexseq", "tiff_pipe", "tmv", "truehd", "tta", "tty", "txd", "ty", "u16be", "u16le", "u24be", "u24le", "u32be", "u32le", "u8", "uncodedframecrc", "v210", "v210x", "vag", "vc1", "vc1test", "vcd", "vidc", "vividas", "vivo", "vmd", "vob", "vobsub", "voc", "vpk", "vplayer", "vqf", "w64", "wc3movie", "webm", "webm_chunk", "webm_dash_manifest", "webp", "webp_pipe", "webvtt", "wsaud", "wsd", "wsvqa", "wtv", "wv", "wve", "xa", "xbin", "xmv", "xpm_pipe", "xvag", "xwd_pipe", "xwma", "yop", "yuv4mpegpipe"}
var formats = [...]string{"avi", "mkv", "mp4"}

func collectFilenames(dir string) ([]string, error) {
	color.Cyan("[func] Collecting all filenames within '%s%s'", dir, string(os.PathSeparator))

	var files []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			length := len(info.Name())
			var shortName string
			if length > 20 {
				shortName = "..." + info.Name()[length%20:]
			} else {
				shortName = info.Name()
			}

			if info.IsDir() == true {
				color.Blue("\tSKIPPED \tDirectory\t'%s'", shortName)
			} else {
				lastIndex := strings.LastIndex(info.Name(), ".")
				ext := info.Name()[lastIndex+1:]
				if contains(formats, ext) {
					color.Magenta("\tAPPENDED\tAud/Vis File\t'%s'", shortName)
					files = append(files, path)
				} else {
					color.Blue("\tSKIPPED \tUnsupported\t'%s'", shortName)
				}
			}
			return nil
		})

	if err != nil {
		return nil, err
	}
	return files, nil
}

func doNormalization(wg *sync.WaitGroup, file string) {
	defer wg.Done()
	color.Cyan("[func] Preparing to normalize file at '%s", file)

	lastIndex := strings.LastIndex(file, string(os.PathSeparator))
	name := file[lastIndex+1:]
	prepend := "NA" // placeholder for cli flag
	separator := "_"
	outName := fmt.Sprintf("output%c%s%s%s", os.PathSeparator, prepend, separator, name)
	cmd := exec.Command("ffmpeg", "-loglevel", "error", "-i", file, "-filter:a", "loudnorm", "-c:v", "copy", outName)

	// stderr output is piped so we can use it for logging results
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
	path := fmt.Sprintf("output%slog.txt", string(os.PathSeparator))
	if fileExists(path) == true {
		// already exists, append to file
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
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
		f, err := os.Create(path)
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
		fmt.Println("Try again, such as: \nLinux/Mac: \t~ ./audio_normalize movies/ \nWindows\t audio_normalize.exe movies\\")
		os.Exit(1)
	}

	dir := os.Args[1]
	files, readErr := collectFilenames(dir)
	if readErr != nil {
		color.Red("[ERROR]: %s", readErr)
		os.Exit(1)
	}

	// worker waitgroup to ensure all goroutines are completed before exiting program
	var wg sync.WaitGroup

	color.Red("[main] Assigning worker tasks")
	for _, file := range files {
		wg.Add(1)
		go doNormalization(&wg, file)
	}

	wg.Wait()
	color.Red("[main] All workers finished and audio normalization complete!")
	color.Yellow("[main] Please check the file log.txt in the output/ directory for any error details.")
}
