package main

import (
	"fetch-go/fetch"
	"fetch-go/utils"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

var printMetadata bool
var validURLs []string

func init() {
	flag.BoolVar(&printMetadata, "metadata", false, "Display Metadata or not")
}

func main() {
	flag.Parse()
	args := flag.Args()
	validURLs = utils.ParseURI(args)

	//	fmt.Println("metadata?", printMetadata)
	//	fmt.Println(validURLs)

	// extract links from content
	// fetch and save date from above links
	// rewrite links to suitable for local file system
	// save the rewited content to index.html file

	var wg sync.WaitGroup

	for _, url := range validURLs {

		url := url

		instance := fetch.Fetch{
			WG: &sync.WaitGroup{},
			Input: &fetch.FetchInput{
				BaseURL: url,
				Time:    time.Now(),
			},
		}

		if printMetadata {
			instance.MPrint()
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()

				instance.FetchALL()
				instance.SavePage()
				instance.Wait()
			}()

		}
	}

	if !printMetadata {
		wg.Wait()
	}
}
