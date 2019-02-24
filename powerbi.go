package main

import (
	con "PowerBiGo/contracts"
	"encoding/json"
	"fmt"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	commits := map[string]string{"AZURE_TENANT_ID": "470cec91-5a0e-47c7-87a9-2fcaf82d5d90", "AZURE_CLIENT_ID": "bcd57285-ddd1-4ae8-a8ad-cb72f7d24aaf", "AZURE_CLIENT_SECRET": "0$Hv3b/SB?hH1S^3?m9Qlne:_sYWirrD8KG%4U_H;yRGRd"}
	tokenx, err := NewServicePrincipalTokenFromCredentials(commits, "https://analysis.windows.net/powerbi/api")
	if err != nil {
		log.Fatal("Err")
	}

	err = tokenx.Refresh()
	if err != nil {
		log.Fatal("Err")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.powerbi.com/v1.0/myorg/groups", nil)

	toke := tokenx.Token().AccessToken

	req.Header.Add("authorization", "Bearer "+toke)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Err")
	}
	body, _ := ioutil.ReadAll(res.Body)
	dda := string(body)
	fmt.Println(dda)
	err = res.Body.Close()

	if err != nil {
		log.Fatalf("Error closing response body")
	}
	var results = &con.GroupData{}
	_ = json.Unmarshal(body, results)
	fmt.Println(results)
}

func NewServicePrincipalTokenFromCredentials(c map[string]string, scope string) (*adal.ServicePrincipalToken, error) {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, c["AZURE_TENANT_ID"])
	if err != nil {
		panic(err)
	}
	return adal.NewServicePrincipalToken(*oauthConfig, c["AZURE_CLIENT_ID"], c["AZURE_CLIENT_SECRET"], scope)
}
