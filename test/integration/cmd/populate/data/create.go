package data

import (
	"fmt"
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/routedns"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rulesengine"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"os"
	"terraform-provider-edgecast/test/integration/cmd/populate/internal"
	"time"
)

func account() string {
	return os.Getenv("ACCOUNT_NUMBER")
}

func email() string {
	return fmt.Sprintf("devenablement+testing%d@edgecast.com", time.Now().Unix())
}

func unique(s string) string {
	return fmt.Sprintf("%s%d", s, time.Now().Unix())
}

func Fix(cfg edgecast.SDKConfig) {
	rulesEnginePolicyID := createPolicyV4(internal.Check(rulesengine.New(cfg)))
	fmt.Println("rules engine policy id:", rulesEnginePolicyID)
	/*
		AddPolicy: SubmitRequest: ecRequestBuilder.buildRequest: request.setAuthorization: failed to get authorization: EOF
	*/
	wafScopesID := createWAFScopes(internal.Check(waf.New(cfg)))
	fmt.Println("waf scopes id:", wafScopesID)
	/*
		ModifyAllScopes: SubmitRequest: sendRequest failed (HTTP StatusCode:400): {"errors":[{"code":400,"message":"Failed: wjc could not validate scope. Error validating configuration file with waf json compiler tool. Reason:"}],"success":false}
	*/
	svc := internal.Check(routedns.New(cfg))
	zoneID := createZone(svc)
	fmt.Println("zone id:", zoneID)
	/*
		AddZone: SubmitRequest: sendRequest failed (HTTP StatusCode:500): {"Message":"Operation Error. Contact Administrator"}
	*/
	groupID := createGroup(svc)
	fmt.Println("dns group id:", groupID)
	/*
		AddGroup: SubmitRequest: sendRequest failed (HTTP StatusCode:500): {"Message":"Operation Error. Contact Administrator"}
	*/
}

func Create(cfg edgecast.SDKConfig) {
	accountNumber, customerUser := createCustomerData(cfg)
	fmt.Println("account number:", accountNumber)
	fmt.Println("customer user:", customerUser)

	originID := createOriginData(cfg)
	fmt.Println("origin id:", originID)

	edgeCnameID := createEdgeCnameData(cfg)
	fmt.Println("edge cname id:", edgeCnameID)

	groupID, masterServerGroupID, masterServerA, masterServerB, secondaryServerGroupID, tsgID, zoneID := createDNSData(cfg)
	fmt.Println("dns group id:", groupID)
	fmt.Println("master server group id:", masterServerGroupID)
	fmt.Println("master server a id:", masterServerA)
	fmt.Println("master server b id:", masterServerB)
	fmt.Println("secondary server group id:", secondaryServerGroupID)
	fmt.Println("tsg id:", tsgID)
	fmt.Println("zone id:", zoneID)

	rulesEnginePolicyID := createRulesEnginePolicyData(cfg)
	fmt.Println("rules engine policy id:", rulesEnginePolicyID)

	wafRateRuleID, wafAccessRuleID, wafCustomRuleID, wafManagedRuleID, wafScopesID := createWAFData(cfg)
	fmt.Println("waf access rule id:", wafAccessRuleID)
	fmt.Println("waf custom rule id:", wafCustomRuleID)
	fmt.Println("waf managed rule b id:", wafManagedRuleID)
	fmt.Println("waf rate rule id:", wafRateRuleID)
	fmt.Println("waf scopes id:", wafScopesID)
}