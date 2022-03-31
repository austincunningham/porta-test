package main

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
	err = threescaleClient.DeleteProduct(4) // err should not be nil when the product exists
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
}