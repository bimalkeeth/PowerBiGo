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
	req, _ := http.NewRequest("GET", common.POWERBI_REST_GET_GROUPS, nil)

	q := req.URL.Query()
	q.Add("filter", "id eq '"+groupId+"'")
	req.URL.RawQuery = q.Encode()
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

//---------------------------------------------------
// Get group by name
//---------------------------------------------------
func (*ResCall) GetGroupByName(credential *con.Credential, name string) (*con.Groups, error) {

	var auth au.IAuth = &Auth{}
	authConfig, err := auth.Authorize(credential)

	ErrorHandle(err, "Error in authorising")

	err = authConfig.Refresh()
	client := &http.Client{}
	req, err := http.NewRequest("GET", common.POWERBI_REST_GET_GROUPS, nil)
	query := req.URL.Query()
	query.Add("filter", "name eq '"+name+"'")
	req.URL.RawQuery = query.Encode()

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

//-------------------------------------------------
//Create Group
//-------------------------------------------------
func (*ResCall) CreateGroup(credential *con.Credential, request con.GroupCreateRequest) (bool, error) {

	var auth au.IAuth = &Auth{}
	authConfig, err := auth.Authorize(credential)
	if err != nil {
		return false, err
	}
	err = authConfig.Refresh()
	//client := &http.Client{}
	//req, err := http.NewRequest("POST", common.POWERBI_REST_POST_GROUPS, nil)

	return false, nil
}
