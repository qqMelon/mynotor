package call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	a []interface{}
	m map[string]interface{}
)

type Client struct {
	ClientId      string `json:"client_id"`
	ClientVersion string `json:"client_version"`
}

type ThreatInfo struct {
	ThreatTypes      []string `json:"threat_types"`
	PlatformTypes    []string `json:"platform_types"`
	ThreatEntryTypes []string `json:"threat_entry_types"`
	ThreatEntries    []string `json:"threat_entries"`
}

type RequestBody struct {
	client     Client
	threatInfo ThreatInfo
}

func CheckUrl(url string) {
	apiUrl := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=AIzaSyB5kRkYVX-S-_m5Q7mdIiHlxuRV8Ii9xlM"

	newReq := RequestBody{
		client: Client(struct {
			ClientId      string
			ClientVersion string
		}{ClientId: "mynotor", ClientVersion: "1.5.2"}),
		threatInfo: ThreatInfo(struct {
			ThreatTypes      []string
			PlatformTypes    []string
			ThreatEntryTypes []string
			ThreatEntries    []string
		}{
			ThreatTypes:      []string{"MALWARE", "SOCIAL_ENGINEERING"},
			PlatformTypes:    []string{"WINDOWS"},
			ThreatEntryTypes: []string{"URL"},
			ThreatEntries:    []string{url},
		})}

	jsonData, err := json.Marshal(newReq)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code status: ", response.Status)
	fmt.Println("Response: ", body)
}
