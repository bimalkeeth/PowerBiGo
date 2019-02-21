package interfaces

import (
	con "PowerBi/contracts"
	"github.com/Azure/go-autorest/autorest/adal"
)
type IAuth interface {
	Authorize(credential *con.Credential)(*adal.ServicePrincipalToken, error)
}