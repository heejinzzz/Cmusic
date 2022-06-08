package main

import "Cmusic"

func main() {
	Cmusic.GetSongList()
	Cmusic.DownloadSong("ViViD-heejin.flac", "./")
}
