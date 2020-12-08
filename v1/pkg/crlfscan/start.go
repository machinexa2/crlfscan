package crlfscan
import (
	"fmt"
	"sync"
)
import (
	"github.com/machinexa2/gobasic"
	"github.com/machinexa2/crlfscan/v1/runner"
)

func StartRunner(){
        var wg sync.WaitGroup
	var globals = &runner.Globals
	var options = &runner.Settings
	var color = &runner.Color
	queue := make(chan string)
	targets := make(chan string)
	if options.Url != "" {
		go func(url string, c chan string){
			c <- url
			close(c)
		}(options.Url, targets)
	} else if options.Wordlist != "" {
		go gobasic.ReadFile(options.Wordlist, targets)
	}

	go Generator(targets, globals.Payloads, queue)
	wg.Add(options.Threads)
        for i := 0; i < options.Threads; i++ {
                go func() {
                        for target := range queue {
				vuln := runner.Scan(target)
				if vuln != "" {
					fmt.Printf("%s Found vulnerable url: %s\n", color.Vuln, color.Blue(vuln))
				}
				if options.Output != "" {
					options.OutputFile.WriteString(vuln + "\n")
				}
                        }
                        wg.Done()
                }()
        }
	wg.Wait()

}

