package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// ListBucketResult was generated 2025-01-05 13:01:30 by https://xml-to-go.github.io/ in Ukraine.
type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult" json:"listbucketresult,omitempty"`
	Text        string   `xml:",chardata" json:"text,omitempty"`
	Xmlns       string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     string   `xml:"MaxKeys"`
	IsTruncated string   `xml:"IsTruncated"`
	Contents    []struct {
		Text         string `xml:",chardata" json:"text,omitempty"`
		Key          string `xml:"Key"`
		LastModified string `xml:"LastModified"`
		ETag         string `xml:"ETag"`
		Size         string `xml:"Size"`
		StorageClass string `xml:"StorageClass"`
		Owner        struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			ID          string `xml:"ID"`
			DisplayName string `xml:"DisplayName"`
		} `xml:"Owner" json:"owner,omitempty"`
	} `xml:"Contents" json:"contents,omitempty"`
}

// Function of object Person
func (bucket ListBucketResult) String() ([]byte, error) {
	return json.Marshal(bucket)
}

func main() {
	var bucket ListBucketResult

	// Load
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	// Parse
	raw, err := base64.StdEncoding.DecodeString(
		os.Getenv("SEFTSFVOSUZPUk1SRVNPVVJDRUxPQ0FUT1IK"),
	)
	if err != nil {
		panic(err)
	}

	// Response
	response, err := http.Get(string(raw))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// Read body
	xmlData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	xml.Unmarshal(xmlData, &bucket)

	// Content
	for _, content := range bucket.Contents {
		param := url.PathEscape(content.Key)
		url.PathEscape(content.Key)
		fmt.Printf("Last Modified:\t%s\n", content.LastModified)
		downloadFile("contents/"+param, string(raw)+param)
		time.Sleep(time.Second * 5)
		break
	}

}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	o, e := os.Create(filepath)
	if e != nil {
		return e
	}
	defer o.Close()

	// Get the data
	r, e := http.Get(url)
	if e != nil {
		return e
	}
	defer r.Body.Close()

	// Writer the body to file
	_, e = io.Copy(o, r.Body)
	if e != nil {
		return e
	}

	return nil
}
