# Go script for normalizing* audio in video files

### Learning Go while solving an audio volume problem.

##### *Technically it is called 'dynamic range compression' but for simplicity, normalization is: increasing the low and decreasing the high volume parts in a file (video or audio) to be closer to a median volume level.

##### [I don't need all this, I just want to run the script on my video library](i-just-want-to-run-this-on-my-video-library)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/fe/The_Thinker%2C_Mus%C3%A9e_Rodin%2C_Paris_September_2013_001.jpg/1600px-The_Thinker%2C_Mus%C3%A9e_Rodin%2C_Paris_September_2013_001.jpg" alt="The Thinker, Musée Rodin, Paris." license="https://creativecommons.org/licenses/by-sa/2.0" declaration="No changes were made to the license or image (other than resizing for viewing)." author="Tammy Lo from New York, NY" style="max-width: 722px;" >

----
##### Learning How To Learn

I love to learn. I get excited figuring out how things work and solving problems. I think a lot of us feel that way and it is what draws us to working in software. There's always a new technology (language, stack, framework, library, etc.) to learn about and try and use to implement something useful. Learning to use code to solve problems has been a fun adventure and I wanted to share some of my thought processes and actions when solving problems and picking up new skills. 

Quick motivations/tips for successful learning:
* Want to do it - do you have a growth mindset? Are you interested in things? Do you want to know more?
* Have a problem you would like to solve - necessity is the mother of invention. There's no impetus to scratch without an itch.
* Have some time and dedicate it to learning - use your free time to read, research, listen.
* Make things! Academic education is ephemeral; building something you use is a cement foundation (even better if other people use it, best if other people use it and tell you what they do and don't like about it)

We established I like tinkering and learning so on to bullet point two: what's the itch? 

---

##### The Annoying Thing Into a Problem Domain

For me I have been spending a lot of time at home lately (for obvious pandemic reasons) and so I've found myself consuming media at home a lot more. I've noticed when using my not-top-of-the-line audio equipment to watch movies and TV shows I have a volume issue. I'm sure many of you have experienced this: watching a movie, enjoying some popcorn, there's an intimate scene with people speaking in low tones so you reach to turn the volume up so you can hear what they are saying; and then the next scene starts with an explosion of sound that is suddenly air-raid-siren, fire-alarm, screaming-banshee-insanely loud! Volume goes down. In the midst of the action a couple of characters speak to each other and you can't hear it! This is insanely frustrating to me to have to hold the remote and volume up, volume down every two minutes. immersion? Gone. Paying attention to the acting is co-opted by the distraction of volume vigilance!

Ok there's a problem software created so it should be a problem software can solve! Now what? You guessed it, next bullet point. 

Spend some time learning about the problem domain. What kinds of things cause the problem? What steps could be taken to remediate them?

This is an audio encoding problem. Audio engineers are likely not unaware of this issue but like almost all modern problems the issues are complex and involve tradeoffs. 

---

Here's a mini-dump of concepts I researched during this phase (links in a footer below if you want to follow the bookmark thread):
* How is sound encoded in a vidual/audio file? 
* What audio formats are there? 
* What is volume? 
* What is a decibal? 
* Are decibals relative units or related to something in the real world? (turns out it is both)
* What were the loudness wars? 
* What is EBU R128? 
* What would I need to know to break up a movie or TV show into the audio and video components? 
* Can I create an audio stream from a video file? 
* Can I read the audio stream for decibal/volume information? 
* Can I modify the audio stream?

---

##### Getting Sidetracked

Many times while researching your problem you'll find that one thing you are researching leads to another (down the wikipedia rabbit hole!) as well as inflection points: areas where you may want to consider what is the problem you are solving and whether you are letting the scope of your original solution concept creep into other areas.

If you are learning something new it is very important to start simple, learn what you need to do to get a working product that is hopefully architected to be flexible and extensible, then extend it after it works to cover different use cases. For me at this point I was looking at signal processing and waveforms, doing sine math, and I realized I had strayed from what I originally wanted: fix it so I can watch a movie without having to adjust the volume. 

<img src="https://upload.wikimedia.org/wikipedia/commons/e/e7/Interpolation_beamforming.PNG" alt="diagram showing advanced calculus of signal processing and transforms" license="https://creativecommons.org/licenses/by-sa/4.0/deed.en" declaration="No changes were made to the license or image (other than resizing for viewing)." author="Btkramer9" caption="May be overthinking this..." style="max-width: 722px;">

`Is knowing this^^ helping me solve my problem?`

---
##### Refocus! Get Help & Make Decisions

We've learned a lot of things about our problem domain now. But we've also tightened up our constraints: I want to solve my problem but I don't need to know every single detail about the implementation of the processes that will help me solve my problem. What do we do with this realization? Find libraries! Look for open source work done by the community that can help you focus on your problem domain and let them work on the fine details. Getting lost in details and ending up with nothing is not what we want!

More research at this point was focused on finding an open source library for modifying audio streams. There's a few, many of them are not open source, or incomplete, or over-engineered for this problem domain, or built GUI-first when we want something CLI-oriented so we can call it programmatically from our own code. 

My top two results were avidemix and ffmpeg. FFmpeg ended up being the winner for having a more robust CLI interface and having it's own opinions about processing which further simplify what I need to understand to solve my problem and make my code work. 
##### RTFM: Read The ffmpeg Manual 

Look at the documentation to see what the library can accomplish. It turns out not only does ffmpeg break apart audio and video streams to read and output useful information, there is already a function called "loudnorm" for audio normalization built in! What a boon; this should take care of the bulk of the heavy lifting (if it works; we need to make sure it works before we rely on it in our project!).

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Go_gopher_bumper.png/1600px-Go_gopher_bumper.png" alt="Go Language Gopher and link to golang.org" license="https://creativecommons.org/licenses/by-sa/2.0" declaration="No changes were made to the license or image (other than resizing for viewing)." author="Renee French" style="max-width: 722px;">

---
##### Scripting A Solution

We haven't even written any code yet! Part of this project for me was to solve the problem and the other part was for me to learn about Go, Google's strongly-typed, compiled programming language. I chose GoLang because it is in a lot of ways the oppositve of my typical language choice (interpreted, dynamically-typed JS, for example) and because I wanted to take advantage of Go's lightweight concurrent goroutines as I expect audio processing to be computationally complex (decoding, modifying, and encoding 100s of thousands of frames of data seems like it would be CPU intensive so running on a single thread would take FOR. EV. ER.).

Here's a mini-dump of concepts I researched when learning about the Go Language (again, links will be at the bottom):
* slices
* structs
* types
* defer
* concurrent goroutines
* OS command execution (ie: calling a program from your program/spawn process/exec command)
* OS filepaths, separators, IO differences
* recursive folder and file reads
* string formatting (%s%w%d etc.)
* wait groups
* worker pools

This is not an exhaustive list of Go features; just some new concepts or different implementations than I'd seen working previously with JavaScript and other languages. 

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Craftmen_at_work%2C_bamboo_basket_weaving_and_textile_mobile_sculptures%2C_in_Heuan_Chan_heritage_house%2C_Luang_Prabang%2C_Laos.jpg/1600px-Craftmen_at_work%2C_bamboo_basket_weaving_and_textile_mobile_sculptures%2C_in_Heuan_Chan_heritage_house%2C_Luang_Prabang%2C_Laos.jpg" alt="Textile workers at work weaving" license="https://creativecommons.org/licenses/by-sa/4.0/deed.en" declaration="No changes were made to the license or image (other than resizing for viewing)." author="Basile Morin" style="max-width: 722px;">

---

##### Weaving Together The Pieces
Write some code! For now this is a brand new language for me so I am figuring out how to install the language on my system, how to initialize a project (Go uses modules which can encapsulate functionality and be shared and run elsewhere easily), how to write the boilerplate, how to compile the code into an executable, how to execute it, how to have the program fail and exit gracefully should it encounter an error, how to test for errors, how to do IO/file reads/file creation/directory reads/spawn processes/bind Stdin and Stdout to an interactive shell, how to log useful information to see how the program is progressing, testing, architecture, output, error states, user input, etc. etc. etc. Get your hands dirty and write a bunch of stuff. It probably won't work right away, or it won't compile, or it won't process the way you expect. Good news! You're learning!

---
##### Refocus! Back To The Problem

After spending time researching and implementing things in the new language, now is a good time to take a step back and remember we are writing code to solve a problem. What can we measure to ensure our problem is actually solved? We want to normalize the audio stream so the highs and lows are a bit closer together. Ffmpeg has a utility for detecting the volume at various points throughout a film to give you a range and idea of the peaks and troughs.

Running: 

`~
ffmpeg -i big_buck_bunny_480p_stereo.avi -filter:a volumedetect -f null /dev/null
`

Will output something like:

```sh
[Parsed_volumedetect_0 @ 0x7f97f6626000] n_samples: 57256704
[Parsed_volumedetect_0 @ 0x7f97f6626000] mean_volume: -22.5 dB
[Parsed_volumedetect_0 @ 0x7f97f6626000] max_volume: 0.0 dB
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_0db: 2221
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_1db: 3708
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_2db: 6927
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_3db: 12952
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_4db: 23453
[Parsed_volumedetect_0 @ 0x7f97f6626000] histogram_5db: 40929
```

There's quite a bit of other data I've left out since this is the salient part. We've got our numbers now: average, max, and some readings at different points throughout.

To test if this ffmpeg normalize function works let's run it on a file and see:

`~ 
ffmpeg -i big_buck_bunny_480p_stereo.avi -filter:a loudnorm -c:v copy NA_big_buck_bunny.avi
`

Flags: 
* `i` for input file
* `filter:a` select's the audio stream for 
* `loudnorm` which is the name of ffmpeg's function for normalizing audio (using ffmpeg defaults)
* `-c:v copy` has ffmpeg directly (c)opy the (v)ideo stream to the output file with no changes
* and finally create the output file with `NA_` (normalized audio) prepended to the filename

Give it a bit to run and create the new file. Then we can run the same audio volume detector on the new file via:

`~
ffmpeg -i NA_big_buck_bunny_480p_stereo.avi -filter:a volumedetect -f null /dev/null 
`

Which outputs: 

```sh
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] n_samples: 57259008
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] mean_volume: -27.5 dB
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] max_volume: -1.3 dB
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_1db: 1
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_2db: 296
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_3db: 438
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_4db: 784
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_5db: 1323
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_6db: 3174
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_7db: 4297
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_8db: 6225
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_9db: 9590
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_10db: 15289
[Parsed_volumedetect_0 @ 0x7fdf6ad0e8c0] histogram_11db: 26075
```

What does this mean? Compare the numbers from before and after doing audio normalization. What numbers changed? What was added or removed? The mean/average volume went from -22.5 dB to -27.5 dB. 22.5/27.5 = 0.82 or 82% so the change could be described as an increase in average volume of 18% across the clip. The max volume was also normalized down a bit from 0 dB to -1.3 dB so whatever point or points in the movie when the volume was very loud were brought down just a bit.

What do these numbers mean for a human experiencing the audio? The real problem isn't about changing numbers it's about enjoying the movie for a person. Take some time to watch a few clips especially looking for areas you may have noticed quiet speaking or loud action before. Open both movies and compare the listening experience between the two.

Big Buck Bunny doesn't get changed much but running this on some other video files you will see and experience a much greater difference in the listening experience. Ok ffmpeg's loudnorm func checks out.

---

##### Back To Coding! 
Architecting this I am going to be running it from the command line so there's no need for any graphical interfance (although that could be built to call on this Go module). However I do want this executable to be able to be run in both Linux/Mac and Windows environments so the code should be OS-agnostic. 

Breaking this down I want to take advantage of Go's worker pools using what they call WaitGroups. A WaitGroup allows you to add as many goroutines as you want each doing their own work and Go will do the coordination of splitting and giving work to your CPU cores. This means I can run a bunch of the ffmpeg library functions simultaneously on multiple files to ensure I am using all of the computer's processing power!

Go like many compiled languages uses a main() function that is called when it is run. My main function handles looking for video files in the directory provided by the user when they run the program, then using that list of files it creates goroutines to run off and process the audio normalization of each. That's it!

Boilerplate main func:

```go
  func main() {
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

```

The two functinos mentioned in main() are collectFilenames and doNormalization which hopefully are self-explanatory if you've read this far.

Main functionality:

```go
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
```

I also use a couple of helper functions just to keep the main functions focused and short/easier to reason about and test:
```go
func updateLog(msg string) // => creates or updates an existing log in output/log.txt w/any errors encountered during normalization

func fileExists(filename string) bool // => returns true or false if a fileExists of that name (used to check for output/log.txt file)

func contains(arr [len(formats)]string, query string) bool // => returns true or false if an array contains a specific string, used to filter accepted file formats
```

## That's it!

Go does all the heavy lifting of coordinating running tasks and taking advantage of all the power your computer can pump out! I did some back-of-an-envelope math after profiling a few test files and found this to be generally 6x (600%) faster than running a for loop bash script which calls ffmpeg. Way To Go!

I hope this very brief explanation helps you all in understanding my process when diving into solving a problem and learning some new programming paradigms. 

---

All the code is available on my GitHub repository [https://github.com/Usarneme/audio_normalize](https://github.com/Usarneme/audio_normalize). You should be able to compile and run this program from the command line, giving it a folder to look for folders and files containing video files. It will queue them up and normalize audio and save the updated copy as a new file. This way if something goes horribly wrong you still have your originals.

Let me know if you have any suggestions or corrections and feel free to open a pull request on GitHub!

---

As promised, the links:

Go
* https://tour.golang.org/welcome/1
* https://gobyexample.com/
* https://github.com/avelino/awesome-go
* https://golangbyexample.com/execute-os-system-command-golang/
* https://golangdocs.com/exec-in-golang
* https://golang.hotexamples.com/examples/os.exec/Cmd/SysProcAttr/golang-cmd-sysprocattr-method-examples.html
* https://yourbasic.org/golang/current-directory/

Understanding Audio/Video 
* https://en.wikipedia.org/wiki/Dynamic_range_compression
* https://en.wikipedia.org/wiki/Audio_normalization
* https://en.wikipedia.org/wiki/EBU_R_128
* https://en.wikipedia.org/wiki/LKFS
* https://community.plus.net/t5/Tech-Help-Software-Hardware-etc/Change-remove-default-audio-strean-in-an-avi/td-p/721935
* http://avidemux.sourceforge.net/doc/en/command.xml.html - runner up to ffmpeg for av processing library


Ffmpeg 
* https://trac.ffmpeg.org/wiki/AudioVolume 
* http://ffmpeg.org/ffmpeg-filters.html#loudnorm


Github Repo
* https://github.com/Usarneme/audio_normalize

#### I just want to run this on my video library
* https://github.com/Usarneme/audio_normalize/files/ez.md

---
##### Attribution & Credits

Please view image HTML tags for attribution; most are CC with attribution v2 or v4 and have been unchanged other than resizing to fit this page. 

Big Buck Bunny is also CC w/attribution v4: Source: Blender Foundation / www.bigbuckbunny.org. Only the file name is referenced in this article but I wanted to give credit to the source. 

---
##### ffmpeg supported video and audio file types: 3g2,3gp,4xm,a64,aa,aac,ac3,acm,act,adf,adp,ads,adts,adx,aea,afc,aiff,aix,alaw,alias_pix,alp,amr,amrnb,amrwb,anm,apc,ape,apm,apng,aptx,aptx_hd,aqtitle,argo_asf,asf,asf_o,asf_stream,ass,ast,au,av1,avfoundation,avi,avm2,avr,avs,avs2,bethsoftvid,bfi,bfstm,bin,bink,bit,bmp_pipe,bmv,boa,brender_pix,brstm,c93,caf,cavsvideo,cdg,cdxl,cine,codec2,codec2raw,concat,crc,dash,data,daud,dcstr,dds_pipe,derf,dfa,dhav,dirac,dnxhd,dpx_pipe,dsf,dsicin,dss,dts,dtshd,dv,dvbsub,dvbtxt,dvd,dxa,ea,ea_cdata,eac3,epaf,exr_pipe,f32be,f32le,f4v,f64be,f64le,ffmetadata,fifo,fifo_test,film_cpk,filmstrip,fits,flac,flic,flv,framecrc,framehash,framemd5,frm,fsb,fwse,g722,g723_1,g726,g726le,g729,gdv,genh,gif,gif_pipe,gsm,gxf,h261,h263,h264,hash,hca,hcom,hds,hevc,hls,hnm,ico,idcin,idf,iff,ifv,ilbc,image2,image2pipe,ingenient,ipmovie,ipod,ircam,ismv,iss,iv8,ivf,ivr,j2k_pipe,jacosub,jpeg_pipe,jpegls_pipe,jv,kux,kvag,latm,lavfi,live_flv,lmlm4,loas,lrc,lvf,lxf,m4v,matroska,md5,mgsts,microdvd,mjpeg,mjpeg_2000,mkvtimestamp_v2,mlp,mlv,mm,mmf,mov,mp4,m4a,3gp,3g2,mj2,mp2,mp3,mp4,mpc,mpc8,mpeg,mpeg1video,mpeg2video,mpegts,mpegtsraw,mpegvideo,mpjpeg,mpl2,mpsub,msf,msnwctcp,mtaf,mtv,mulaw,musx,mv,mvi,mxf,mxf_d10,mxf_opatom,mxg,nc,nistsphere,nsp,nsv,null,nut,nuv,oga,ogg,ogv,oma,opus,paf,pam_pipe,pbm_pipe,pcx_pipe,pgm_pipe,pgmyuv_pipe,pictor_pipe,pjs,pmp,png_pipe,pp_bnk,ppm_pipe,psd_pipe,psp,psxstr,pva,pvf,qcp,qdraw_pipe,r3d,rawvideo,realtext,redspark,rl2,rm,roq,rpl,rsd,rso,rtp,rtp_mpegts,rtsp,s16be,s16le,s24be,s24le,s32be,s32le,s337m,s8,sami,sap,sbc,sbg,scc,sdl,sdp,sdr2,sds,sdx,segment,ser,sgi_pipe,shn,siff,singlejpeg,sln,smjpeg,smk,smoothstreaming,smush,sol,sox,spdif,spx,srt,stl,stream_segment,streamhash,subviewer,subviewer1,sunrast_pipe,sup,svag,svcd,svg_pipe,swf,tak,tedcaptions,tee,thp,tiertexseq,tiff_pipe,tmv,truehd,tta,tty,txd,ty,u16be,u16le,u24be,u24le,u32be,u32le,u8,uncodedframecrc,v210,v210x,vag,vc1,vc1test,vcd,vidc,vividas,vivo,vmd,vob,vobsub,voc,vpk,vplayer,vqf,w64,wav,wc3movie,webm,webm_chunk,webm_dash_manifest,webp,webp_pipe,webvtt,wsaud,wsd,wsvqa,wtv,wv,wve,xa,xbin,xmv,xpm_pipe,xvag,xwd_pipe,xwma,yop,yuv4mpegpipe