package intermediate

import (
	"fmt"
	"net/url"
)

// A url (uniform resource locator) is structured as so:
// [scheme://] [userinfo@]host[:port][/path][?querry][#fragment]

func main() {
	exampleurl := "https://example.com:8000/path?querry=param#fragment"

	parsedUrl, err := url.Parse(exampleurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Querry:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	rawURL := "https://example.com/path?name=John&age=30"

	parseURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	query := parseURL.Query()

	fmt.Println("Name:", query.Get("name"))

	// Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query = baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "25")
	baseURL.RawQuery = query.Encode()

	fmt.Println(baseURL)
}
