package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func newurl(url string) string {
	u := strings.Split(url, "?")
	return u[0]
}

func reqUrl(url string) [3]string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:96.0) Gecko/20100101 Firefox/96.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	sb := string(body)
	contentType := resp.Header.Get("Content-type")
	statusCode := strconv.Itoa(resp.StatusCode)

	var r [3]string

	r[0] = sb
	r[1] = contentType
	r[2] = statusCode

	return r

}

func checkJs(url string) bool {
	req := reqUrl(url)
	contentType := strings.Split(req[1], ";")
	statusCode, err := strconv.Atoi(req[2])

	if err != nil {
		log.Fatal(err)
	}

	if statusCode >= 200 && statusCode <= 299 && strings.Contains(contentType[0], "application/javascript") {
		return true
	} else {
		return false
	}
}

func checkMap(url string) bool {
	req := reqUrl(url)
	body := req[0]
	contentType := strings.Split(req[1], ";")
	statusCode, err := strconv.Atoi(req[2])

	if err != nil {
		log.Fatal(err)
	}

	if statusCode >= 200 && statusCode <= 299 && contentType[0] == "application/json" && strings.Contains(body, "version:") {
		return true
	} else {
		return false
	}

}

func readStdin() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		u := newurl(line)
		umap := u + ".map"
		if checkMap(umap) == true {
			fmt.Println(umap)
		} else {
			if checkJs(u) == true {
				fmt.Println(u)
			}
		}
	}

}

func main() {
	readStdin()
}
