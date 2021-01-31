## audio_normalize
Adjust too-quiet troughs and too-loud audio peaks in videos. Written in fast, concurrent Go.

---

#### Install/Use
* Git clone and cd into the directory
* `go build audio_normalize.go`
Mac/Linux
* `./audio_normalize directory/` where directory/ contains some video files you want normalized
Windows
* `audio_normalize.exe directory/`

#### Requirements
* Go programming language (https://golang.org/dl/)
* ffmpeg (https://ffmpeg.org/download.html)

---

For a detailed writeup of my process in creating this please see here. 
* https://github.com/Usarneme/audio_normalize/blob/main/article.md

---

#### Quick Profiling
```
# Files   SIZE Gb   TIME Sec   Gb/Sec       Mb/Sec
9         1.3       61.1       = 0.0212765  ~21.2765
4         7.8       454.19     = 0.0171734  ~17.1734
12        9.5       699.17     = 0.0135875  ~13.5875

```


### No Warranty - As Is
This program is intended to do no harm but does a lot of reading and writing so things can happen. I have it setup to create copies of the files with normalized audio so you have complete control over the originals and what you want to do with them. This program should not delete anything from your computer but it will make copies of anything you point it at which may be a lot of large files.

This program is designed to maximize CPU usage for processing these audio changes as quickly as possible. Trying to use your computer when all cores are pinned at 100% may not be a good idea. I don't have any way of pausing the process so be ready/know how many files you are targeting. 

