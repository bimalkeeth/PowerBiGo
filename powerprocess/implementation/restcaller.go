package implementation

import (
	"PowerBi/common"
	con "PowerBi/contracts"
	au "PowerBi/powerprocess/interfaces"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
type ResCall struct{}

func(*ResCall)GetAllGroups(credential *con.Credential)(*[]con.Groups,error){
    var auth au.IAuth=&Auth{}
    authConfig,err:= auth.Authorize(credential)
    if err!=nil{log.Fatal("Error in authorising")}

    err=authConfig.Refresh()
    client:=&http.Client{}
	req, _ := http.NewRequest("GET", common.POWERBI_REST_GET_GROUPS, nil)
	req.Header.Add("authorization","Bearer "+authConfig.Token().AccessToken)
	response,err:=client.Do(req)
	if err!=nil{
		log.Fatal("request is not processed successfully")
	}
	body, _ := ioutil.ReadAll(response.Body)
	err=response.Body.Close()
	if err!=nil{
		log.Fatal("Error in reading response")
	}
	var results map[string]interface{}
	_=json.Unmarshal(body,&results)

	resultgroups:=make([]con.Groups,0)


	return &resultgroups, nil
}