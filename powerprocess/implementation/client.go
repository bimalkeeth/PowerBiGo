package implementation

import (
	"PowerBi/common"
con "PowerBi/contracts"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"log"
)
type Auth struct {}
func (*Auth) Authorize(credential *con.Credential)(*adal.ServicePrincipalToken, error){
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, credential.TenantId)
	if err!=nil{log.Fatal("Error in retrieving configuration")}
	return adal.NewServicePrincipalToken(*oauthConfig,credential.ApplicationId,credential.SecretKey,common.POWERBI_RESORCE)
}