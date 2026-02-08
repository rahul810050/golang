package main

import (
	"fmt"
	"net/url"
)

func main(){
	fmt.Println("handling url in golang")

	myurl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("type of the url: %T\n", myurl) // string

	parsedUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("cant parse url", err)
		return
	}
	fmt.Printf("type of the url: %T\n", parsedUrl) // *url.URL

	fmt.Println("scheme", parsedUrl.Scheme)
	fmt.Println("Host", parsedUrl.Host)
	fmt.Println("path", parsedUrl.Path)
	fmt.Println("RawQuery", parsedUrl.RawQuery)


	// Modifying the url components
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "me=cocane"

	// constructing a url string from a url object
	newUrl := parsedUrl.String()
	fmt.Println("new url", newUrl) // new url from the previous url
	fmt.Printf("type of url: %T\n", newUrl) // string
}