package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/dinesh-GDK/multibar"
)

func OneSet(data []http_file) {
	wg := new(sync.WaitGroup)
	wg.Add(len(data))

	Progress, _ := multibar.New()

	var ProgressBars []multibar.ProgressFunc
	for i := 0; i < len(data); i++ {
		ProgressBars = append(ProgressBars, Progress.MakeBar(100, data[i].file_name))
	}
	go Progress.Listen()

	for i := 0; i < len(data); i++ {
		go DownloadFile(data[i].http_response, data[i].file_name, wg, ProgressBars[i])
	}

	wg.Wait()
}

func DownloadFile(http *http.Response, filepath string, wg *sync.WaitGroup, progressBar multibar.ProgressFunc) {

	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		FAILURE = append(FAILURE, filepath)
		return
	}

	counter := &WriteCounter{Total: float64(http.ContentLength), ProgressBar: progressBar, UpdateFreq: 500}

	if _, err = io.Copy(out, io.TeeReader(http.Body, counter)); err != nil {
		out.Close()
		FAILURE = append(FAILURE, filepath)
		return
	}
	progressBar(100)
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		FAILURE = append(FAILURE, filepath)
		return
	}

	http.Body.Close()
	wg.Done()
	SUCCESS = append(SUCCESS, filepath)
}

func ExtractHttpResponse(urls [][]string) ([]http_file, []string) {

	var http_response []http_file
	var no_exist []string

	for _, url := range urls {
		resp, _ := http.Get(url[0])

		if resp.StatusCode != 200 {
			fmt.Printf("'%s' do not exist\n", url[0])
			no_exist = append(no_exist, url[0])
			continue
		}

		http_response = append(http_response, http_file{http_response: resp, file_name: url[1]})
	}

	fmt.Printf("\n")

	sort.Sort(http_file_list(http_response))

	return http_response, no_exist
}

func ReadUrlFile(url_file_path string) [][]string {

	fmt.Printf("Reading %s ...\n\n", url_file_path)

	file, err := os.Open(url_file_path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var urls [][]string

	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		urls = append(urls, split)
	}

	file.Close()

	return urls
}
