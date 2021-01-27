# audio_normalize
Normalize audio (remove silent troughs and explosively loud peaks) in videos; fast concurrent golang.


### Known Issues / Troubleshooting Guide

Duplicates - when ffmpeg prompts me if I want to overwrite an existing filename in the output/ folder, instead of accepting y or n after I hit enter, the terminal shows ^M

This is most likely a terminal problem. Quit the program with Ctrl+C and then type `stty sane` into your terminal. Re-run /audio-normalize and it should now accept Enter correctly. 
