package godownloader

import (
	"flag"
	"fmt"
	"log"
)

var MAXROUTINES uint64 = 10
var SUCCESS []string
var FAILURE []string

func main() {

	file_name := flag.String("file", "", "file path of the file which contains url and destination file path")
	concurrent := flag.Uint64("concurrent", MAXROUTINES, "number of concurrent processess")
	print_log := flag.Bool("log", false, "print log")
	flag.Parse()

	if *concurrent > MAXROUTINES {
		log.Fatal("Number of concurrent routines is more than the maximum limit")
	}

	urls := ReadUrlFile(*file_name)
	data, no_exist := ExtractHttpResponse(urls)

	total_files := len(urls)

	routines := uint64(len(data))
	if *concurrent < MAXROUTINES {
		routines = *concurrent
	}

	for r := uint64(0); r < uint64(len(data)); r += routines {

		end := r + routines
		if end >= uint64(len(data)) {
			end = uint64(len(data))
		}
		OneSet(data[r:end])

		fmt.Printf("(%d/%d) completed...\n\n", end, total_files)
	}

	fmt.Println("Download Finished")

	if *print_log {
		fmt.Print("\nLOG\n*********\n\n")
		fmt.Printf("%d files not exist\n----------\n", len(no_exist))
		for _, file := range no_exist {
			fmt.Println(file)
		}
		fmt.Println()

		fmt.Printf("Sucessfully downloaded %d files\n----------\n", len(SUCCESS))
		for _, file := range SUCCESS {
			fmt.Println(file)
		}
		fmt.Println()

		fmt.Printf("Failed to download %d files\n----------\n", len(FAILURE))
		for _, file := range FAILURE {
			fmt.Println(file)
		}
	}

}
