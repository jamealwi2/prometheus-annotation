package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type CreateAnnotationRequest struct {
	DashboardID uint     `json:"dashboardId,omitempty"`
	PanelID     uint     `json:"panelId,omitempty"`
	Time        int64    `json:"time,omitempty"`
	TimeEnd     int64    `json:"timeEnd,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Text        string   `json:"text,omitempty"`
}

var promEndPoint string = "<PROM URL>/api/annotations" //TO BE REPLACED WITH VALID VALUE
var promToken string = "Bearer Token" // TO BE REPLACED WITH VALID VALUE

func main() {

	timeBegin := time.Now().Unix()
	timeEnd := timeBegin + 1

	dashID := flag.Int("dashID", -1, "Dashboard ID")
	panelID := flag.Int("panelID", -1, "Panel ID")
	message := flag.String("m", "", "Annotation Message")
	flag.Parse()

	if *dashID == -1 {
		fmt.Println("Please provide Dashboard ID")
		os.Exit(1)
	}
	if *panelID == -1 {
		fmt.Println("Please provide Panel ID")
		os.Exit(1)
	}

	if *message == "" {
		fmt.Println("Please provide some annotation message")
		os.Exit(1)
	}

	promAnnotation := &CreateAnnotationRequest{DashboardID: uint(*dashID), PanelID: uint(*panelID), Time: (timeBegin - 30) * 1000, TimeEnd: (timeEnd - 30) * 1000, Text: *message}
	promAnnotationJSON, err := json.Marshal(promAnnotation)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := bytes.NewReader(promAnnotationJSON)

	req, _ := http.NewRequest("POST", promEndPoint, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", promToken)

	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		log.Print(err2.Error())
	}
	defer resp.Body.Close()
}
