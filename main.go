package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Structs
type logData struct {
	serviceName string
	date        int // we are assuming date will be in the format 20210916
	hour        int // we are assuming hour will be in the format 0 - 23
	content     string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func match(s string) bool {
	// format may need to be adjusted to allow for symbols in date (3rd group)
	format := "log-([a-zA-Z-9]+)-([0-9]+)-([0-9]+)"
	match, matchErr := regexp.MatchString(format, s)
	check(matchErr)
	return match
}

func aggregate(logFiles map[string]*logData) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		check(err)
		if match(info.Name()) {
			serviceName := strings.Split(info.Name(), "-")[1]
			content, readErr := os.ReadFile(path)
			check(readErr)

			// if logFiles contains the service already
			// then append the content of the log file to serviceName.content
			// else create a new map key for the service
			if s, ok := logFiles[serviceName]; ok {
				s.content += string(content)
			} else {
				date, dateErr := strconv.Atoi(strings.Split(info.Name(), "-")[2])
				check(dateErr)
				hour, hourErr := strconv.Atoi(strings.Split(info.Name(), "-")[3])
				check(hourErr)

				dat := logData{
					serviceName: serviceName,
					date:        date,
					hour:        hour,
					content:     string(content),
				}
				logFiles[serviceName] = &dat
			}

		}
		return nil
	}
}

func main() {
	var dir string
	var domain string
	logFiles := make(map[string]*logData)

	// Get directory as flag
	flag.StringVar(&dir, "dir", "", "root directory to aggregate logs")
	flag.StringVar(&dir, "d", "", "root directory to aggregate logs")
	// Get OpenSearch domain/url as flag
	flag.StringVar(&domain, "url", "", "OpenSearch Service domain (url)")
	flag.StringVar(&domain, "u", "", "OpenSearch Service domain (url)")

	flag.Parse()

	// Aggregate files
	aggErr := filepath.Walk(dir, aggregate(logFiles))
	check(aggErr)

	// Basic information for the Amazon OpenSearch Service domain
	index := "logs"
	// id := "1" // id is not needed for post requests
	endpoint := domain + "/" + index + "/" + "_doc"
	region := "" // e.g. us-east-1
	service := "opensearchservice"

	for _, v := range logFiles {
		// Log data to send
		json := fmt.Sprintf(`{ "serviceName": %s, "date": %d, "hour": %d, logs: %s }`, v.serviceName, v.date, v.hour, v.content)
		body := strings.NewReader(json)

		// Get credentials from environment variables and create the Signature Version 4 signer
		credentials := credentials.NewEnvCredentials()
		signer := v4.NewSigner(credentials)

		// An HTTP client for sending the request
		client := &http.Client{}

		// Form the HTTP request
		req, err := http.NewRequest(http.MethodPost, endpoint, body)
		check(err)

		// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
		req.Header.Add("Content-Type", "application/json")

		// Sign the request, send it, and print the response
		signer.Sign(req, body, service, region, time.Now())
		resp, err := client.Do(req)
		check(err)
		fmt.Print(resp.Status + "\n")
	}

}
