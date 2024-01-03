package domain

import (
	"fmt"
	"encoding/json"
	"golang.org/x/net/http"
)

func GetIncidents(token string) map[string]interface{} {
	url := fmt.Sprintf("https://%s/connect/api/v1/incidents?tenant_id=817418af76d44358922636e34be9627c&sort=incident_score&order=desc&limit=10", HOST)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}
