package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oopsguy/m3u8/dl"
)

var (
	url      string
	output   string
	encryptedVideoKey   string
	chanSize int
	enableMerge bool
)

func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required")
	flag.IntVar(&chanSize, "c", 25, "Maximum number of occurrences")
	flag.StringVar(&output, "o", "", "Output folder, required")
	flag.StringVar(&encryptedVideoKey, "k", "", "encryptedVideoKey")
	flag.BoolVar(&enableMerge, "m", true, "enableMerge")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if url == "" {
		panicParameter("u")
	}
	if output == "" {
		panicParameter("o")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	downloader, err := dl.NewTask(output, url, encryptedVideoKey)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize, enableMerge); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
