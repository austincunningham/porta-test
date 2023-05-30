package main
// update porta client to use a branch
// go get <path-to-repo>@<branch>
// go get github.com/3scale/3scale-porta-go-client/client@master
import (
	"encoding/json"
	"fmt"

	"github.com/3scale/3scale-porta-go-client/client"
)	

func jsonPrint(obj interface{}) {
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
// delete a product of id 3
func main() {
	adminPortalURL := "https://3scale-admin.apps.aucunnin.***/"
	threescaleAccessToken := "****"

	adminPortal, err := client.NewAdminPortalFromStr(adminPortalURL)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	threescaleClient := client.NewThreeScale(adminPortal, threescaleAccessToken, nil)
	// err = threescaleClient.DeleteProduct(4) // err should not be nil when the product exists
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	//staging promotion
	//_, err = threescaleClient.DeployProductProxy(5)
	//production promotion
	//_, err = threescaleClient.PromoteProxyConfig("5","sandbox","3","production")
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	//println(thejson)


	// proxyConfigList
	// host:="catfact-3scale-apicast-staging.apps.aucunnin.w03b.s1.devshift.org"
	// version:= "1"
    // proxyConfigList, err := threescaleClient.ListAccountProxyConfigs("sandbox",&version,&host)
    // if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	// jsonPrint(proxyConfigList)

	// listproducts
 
	// productlist, err := threescaleClient.ListProducts()
	// if err !=nil {
	// 	panic(err)
	// }
	// jsonPrint(productlist)

	// listbackends

	// backendList, err := threescaleClient.ListBackendApis()
	// if err !=nil {
	// 		panic(err)
	// 	}
	// jsonPrint(backendList)

	// listDeveloperAccounts

    // listDeveloperAccounts, err := threescaleClient.ListDeveloperAccounts()
	// if err !=nil {
	// 		panic(err)
	// 	}
	// jsonPrint(listDeveloperAccounts)
	var(
		backendID int64 = 613
		hitsID 	int64 = 618
	)

	listbackendsMethods, err := threescaleClient.ListBackendapiMethods(backendID,hitsID)
	if err !=nil {
			panic(err)
		}
	jsonPrint(listbackendsMethods)

}