package implementation

import (
	"PowerBiGo/common"
	con "PowerBiGo/contracts"
	au "PowerBiGo/powerprocess/interfaces"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//-------------------------------------------------
// Get PowerBI Groups
//-------------------------------------------------
func (*ResCall) GetAllGroups(credential *con.Credential) (*[]con.Groups, error) {
	var auth au.IAuth = &Auth{}
	authConfig, err := auth.Authorize(credential)
	if err != nil {
		log.Fatal("Error in authorising")
	}

	err = authConfig.Refresh()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", common.POWERBI_REST_GET_GROUPS, nil)
	req.Header.Add("authorization", "Bearer "+authConfig.Token().AccessToken)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("request is not processed successfully")
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		log.Fatal("Error in reading response")
	}
	resultGroups := make([]con.Groups, 0)
	_ = json.Unmarshal(body, &resultGroups)

	return &resultGroups, nil
}

//-------------------------------------------------
// Get PowerBI Group
//-------------------------------------------------
func (*ResCall) GetGroupById(credential *con.Credential, groupId string) (*con.Groups, error) {
	var auth au.IAuth = &Auth{}
	authConfig, err := auth.Authorize(credential)
	if err != nil {
		log.Fatal("Error in authorising")
	}
	err = authConfig.Refresh()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", common.POWERBI_REST_GET_GROUPS+"?id="+groupId, nil)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("request is not processed successfully")
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		log.Fatal("Error in reading response")
	}
	resultGroups := &con.Groups{}
	_ = json.Unmarshal(body, resultGroups)
	return resultGroups, nil
}

func (*ResCall) GetGroupByName(credential *con.Credential, groupId string) (*[]con.Groups, error) {

	var auth au.IAuth = &Auth{}
	authConfig, err := auth.Authorize(credential)
	if err != nil {
		log.Fatal("Error in authorising")
	}
	err = authConfig.Refresh()

	return nil, nil
}
