// Trade-A-Plane Listing API – Go example
// Usage: PILOTERR_API_KEY=your-key go run listing.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	params := url.Values{}
	params.Set("query", "https://www.trade-a-plane.com/search?category_level1=Piston+Engine&make=ROLLS+ROYCE&model=PACKARD+V-1650&listing_id=2419556&s-type=engine")

	req, _ := http.NewRequest("GET", "https://api.piloterr.com/v2/tradeaplane/listing?"+params.Encode(), nil)
	req.Header.Set("x-api-key", os.Getenv("PILOTERR_API_KEY"))

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Piloterr error %d: %s\n", resp.StatusCode, body)
		os.Exit(1)
	}

	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println(pretty.String())
}
