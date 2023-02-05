package auth

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type HcxUser struct {
	Username string
	Password string
	HcxUrl   string
}

func HcxLogin(s HcxUser) string {
	uri := "/ hybridity/api/sessions"
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(s.HcxUrl + uri)
	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}
	return response.Header().Get("X-VMWARE-VCLOUD-ACCESS-TOKEN")

}

//def authenticate(url, username, password):
//headers = {"Accept": "application/json", "Content-Type": "application/json"}
//payload = {"username": username, "password": password}
//endpoint = url + "/hybridity/api/sessions"
//try:
//response = requests.post(endpoint, headers=headers, data=json.dumps(payload), verify=False)
//token = response.headers["x-hm-authorization"]
//return token
//except HTTPError as http_err:
//logging.info("error", http_err)