package main

import (
	"crypto/tls"
	"fmt"
	//"github.com/digitalocean/go-netbox/netbox/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"os"


	//"github.com/lotanaharoniP81/netbox_api/client"
	//"github.com/lotanaharoniP81/netbox_api"
	//"github.com/netbox-community/go-netbox/netbox/client/ipam"

	//"github.com/digitalocean/go-netbox/netbox/client"
	//"github.com/digitalocean/go-netbox/netbox/client/ipam"

	"log"
	"net/http"
)

func main() {

	fmt.Println("hello")

	// ignore expired / self signed SSL certificates
	httpClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}



	// The 'host', 'protocol' and the 'token' are taken from the 'model' file (locally)
	transport := httptransport.NewWithClient(host, client.DefaultBasePath, []string{protocol}, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", token)

	c := client.New(transport, nil)

	//req := dcim.NewDcimSitesListParams()
	//res, err := c.Dcim.DcimSitesList(req, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	////log.Infof("res: %v", res)
	//
	//fmt.Printf("%v\n", *(res.Payload.Count))
	//
	//
	// get all IPs
	//req2 := ipam.NewIpamIPAddressesListParams()
	//res2, err := c.Ipam.IpamIPAddressesList(req2, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", *(res2.Payload.Count))


	//// get all IPs
	//req2 := ipam.NewIpamIPAddressesListParams()
	//res2, err := c.Ipam.IpamIPAddressesList(req2, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", *(res2.Payload.Count))



	//// get all prefixes for a specific tenant
	//var prefixes []*models.Prefix
	//exampleTenant := "sx-sj1"
	//req5 := ipam.NewIpamPrefixesListParams()
	//res5, err := c.Ipam.IpamPrefixesList(req5, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	////fmt.Printf("%v\n", res5)
	//for _, p := range res5.Payload.Results {
	//	if p.Tenant != nil && p.Tenant.Display == exampleTenant {
	//		prefixes = append(prefixes, p)
	//	}
	//}
	//fmt.Printf("the prefixes: %v", prefixes)


	//// get the available ips per prefix
	//req3 := ipam.NewIpamPrefixesAvailableIpsListParams()
	//req3.SetID(2)
	//res3, err := c.Ipam.IpamPrefixesAvailableIpsList(req3, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res3.Payload[0].Address)

	// get the available ips per prefix
	req7 := ipam.NewIpamPrefixesAvailableIpsCreateParams()
	req7.SetID(2)
	for i := 0; i < 4; i++ {
		go func() {
			res10, err := c.Ipam.IpamPrefixesAvailableIpsCreate(req7, nil)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			fmt.Println(res10)
		}()
	}


	//res7, err := c.Ipam.IpamPrefixesAvailableIpsCreate(req7, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res7.Payload[0].Address)
	//fmt.Println()


	//for i := 0; i < 3; i ++ {
	//	go func() {
	//		req3.SetID(2)
	//		res3, err := c.Ipam.IpamPrefixesAvailableIpsList(req3, nil)
	//		if err != nil {
	//			fmt.Printf("%v\n", err)
	//			os.Exit(1)
	//		}
	//		fmt.Printf("%v\n", res3.Payload[0].Address)
	//	}()
	//}

	// todo:
	// Get tenant -> and returns Prefixes - V
	// How to avoid locks in case of multiple requests...
	// check work with Ip addresses instead of prefixes - V (we can go over all the IPs - but it seems not the correct way...)


	//req4 := tenancy.NewTenancyTenantsListParams()
	//res4, err := c.Tenancy.TenancyTenantsList(req4, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res4)







	//c.Ipam.IpamIPAddressesList()

	//req3 := ipam.NewIpamPrefixesAvailableIpsListParams()
	//req3.SetID(4)
	//res3, err := c.Ipam.IpamPrefixesAvailableIpsList(req3, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res3)






	//
	//params := &models.WritableIPAddress{}
	//
	//address := "10.0.0.32/24"
	//params.Address = &address
	//
	//description := "Description updated"
	//params.Description = description
	//
	//resource := ipam.NewIpamIPAddressesUpdateParams().WithData(params)
	//_, err := c.Ipam.IpamIPAddressesUpdate(resource, nil)
	//if err != nil {
	//	fmt.Println(err)	}


	ip, err := getNextAvailableIP(c, "MyIPAM Block")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
}

func getNextAvailableIP(nbClient *client.NetBoxAPI, blockName string)(string, error){
	params := ipam.NewIpamIPAddressesListParams()
	blocks, err := nbClient.Ipam.IpamIPAddressesList(params, nil)
	if err != nil {
		return "", err
	}
	fmt.Println(blocks)






	//var block *models.IPAddress
	//for _, b := range blocks.Payload.Results {
	//	if b.Prefix.Description == blockName {
	//		block = b
	//		break
	//	}
	//}

	//if block == nil {
	//	return "", fmt.Errorf("IPAM block with name '%s' not found", blockName)
	//}

	//if err != nil {
	//	return "", err
	//}

	//ipParams := ipam.NewIpamIPAddressesUpdateParams()
	//ipParams.SetID(ip.Payload.ID)
	//ipParams.Set
	//return ip.Payload.Address, nil

	return "", nil
}