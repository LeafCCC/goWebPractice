package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var urls = []string{
	//"http://www.google.com/",
	"https://baidu.com/",
	"https://zhihu.com/",
}

func getStatusCode(urls []string) {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
		fmt.Println(url, "protocol:", resp.Proto)
	}
}

func httpFetch() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input the website address: ")
	url, err := inputReader.ReadString('\n')
	url = strings.TrimSpace(strings.TrimSuffix(url, "\r\n"))
	if err != nil {
		fmt.Println("Error namel reading,", err)
		return
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	fmt.Printf("Got: %q", string(data))
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	getStatusCode(urls)

	httpFetch()
}
