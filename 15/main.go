package main

import (
	"fmt"
)

// AudioPlayer - интерфейс для воспроизведения аудио
type AudioPlayer interface {
	PlayAudio(filename string)
}

// MP3Player - конкретная реализация интерфейса AudioPlayer
type MP3Player struct{}

func (p *MP3Player) PlayAudio(filename string) {
	fmt.Printf("Playing MP3 file: %s\n", filename)
}

// VLCPlayer - плеер с другим интерфейсом
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(filename string) {
	fmt.Printf("Playing VLC file: %s\n", filename)
}

// VLCPlayerAdapter - адаптер для VLCPlayer, реализующий интерфейс AudioPlayer
type VLCPlayerAdapter struct {
	vlcPlayer *VLCPlayer
}

func (v *VLCPlayerAdapter) PlayAudio(filename string) {
	v.vlcPlayer.PlayVLC(filename)
}

func main() {
	// Используем MP3Player, который соответствует интерфейсу AudioPlayer
	mp3Player := &MP3Player{}
	mp3Player.PlayAudio("song.mp3")

	// Используем VLCPlayer через адаптер, чтобы он соответствовал интерфейсу AudioPlayer
	vlcPlayer := &VLCPlayer{}
	vlcAdapter := &VLCPlayerAdapter{vlcPlayer: vlcPlayer}
	vlcAdapter.PlayAudio("video.vlc")
}
