package runner

import "flag"
import "os"
import "fmt"

type Options struct {
	Url			string
	Wordlist		string
	Output			string
	Banner			bool
	Version			bool
	Verbose			bool
	Stdin			bool
	Threads			int
	OutputFile		*os.File
}

var Settings Options = Options{}

func init(){
	options := &Settings
	flag.StringVar(&options.Url, "url", "", "Single url to scan")
	flag.StringVar(&options.Wordlist, "wordlist", "", "List of urls to scan")
	flag.StringVar(&options.Output, "o", "", "File to write output to (optional)")
	flag.BoolVar(&options.Version, "version", false, "Show version")
	flag.BoolVar(&options.Verbose, "v", false, "Increase verbosity")
	flag.BoolVar(&options.Banner, "banner", false, "Show banner")
	flag.IntVar(&options.Threads, "c", 20, "Concurrent requests to make")
	flag.Parse()
	options.validateOptions()
}

func (options *Options) validateOptions(){
	color := &Color
	if options.Version {
		fmt.Printf("%s Current Version: %s\n", color.Good, color.Blue(Globals.Version))
		os.Exit(0)
	}
	if options.Banner {
		fmt.Println(Globals.Banner)
		fmt.Printf("%s Current Version: %s\n", color.Good, color.Blue(Globals.Version))
		os.Exit(0)
	}
	if options.Url == "" && options.Wordlist == "" {
		fmt.Printf("%s Use --help\n", color.Bad)
		os.Exit(0)
	}
	if options.Output != "" {
		options.OutputFile = returnFile(options.Output)
	}
}
