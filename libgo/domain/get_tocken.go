package domain

import (
	"fmt"
	"encoding/base64"
	"encoding/json"
	"golang.org/x/net/http"
	ent"github.com/alphamystic/profiler/libgo/entities"
)

type GetToken interface {
	/* should implement a way to get a chronicle token:
				Chronicle
				Stellar
				Slack
	*/
	GetAccessToken() (string,error)
}

func GetAccessToken(userId ent.UserID,ref_tkn  ent.RefreshToken, host string) string {
	auth := base64.StdEncoding.EncodeToString([]byte(userId + ":" + ref_tkn))
	url := fmt.Sprintf("https://%s/connect/api/v1/access_token", host)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["access_token"].(string)
}
