package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gordonklaus/portaudio"
)

func main() {
	bufferSize := 32
	buf := &audio.FloatBuffer{
		Data: make([]float64, bufferSize),
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  11250,
		},
	}
	currentNote := 220.0
	osc := generator.NewOsc(generator.WaveSine, currentNote, buf.Format.SampleRate)
	osc.Amplitude = 1

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]float32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(0, 1, float64(buf.Format.SampleRate), len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	done := make(chan struct{})

	go func() {
		fmt.Println("started")
		time.Sleep(time.Second)

		message := "hello"
		if len(os.Args) != 1 {
			message = os.Args[1]
		}
		fmt.Println("sending:", message)
		pack := PackFromBytes([]byte(message))

		for _, mote := range pack {
			osc.SetFreq(mote.Freq())
			fmt.Print(mote.String())
			time.Sleep(time.Second / 60)
		}

		fmt.Println("")
		fmt.Println("sent")

		done <- struct{}{}
	}()

	for {
		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		transforms.Gain(buf, 1)

		f64ToF32Copy(out, buf.Data)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

		select {
		case <-done:
			return
		default:
		}
	}
}

// portaudio doesn't support float64 so we need to copy our data over to the
// destination buffer.
func f64ToF32Copy(dst []float32, src []float64) {
	for i := range src {
		dst[i] = float32(src[i])
	}
}
