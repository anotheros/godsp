package godsp

import (
	"encoding/json"
	"fmt"

	"github.com/bsm/openrtb"
)

var BidHandler = func(bidrequest openrtb.BidRequest) openrtb.BidResponse {

	b, _ := json.Marshal(bidrequest.Device)
	fmt.Println("Device", string(b))

	b, _ = json.Marshal(bidrequest.App)
	fmt.Println("App", string(b))
	b, _ = json.Marshal(bidrequest.User)
	fmt.Println("user", string(b))

	//req.BAdv
	b, _ = json.Marshal(bidrequest.BAdv)
	fmt.Println("badv", string(b))
	//req.Bcat
	b, _ = json.Marshal(bidrequest.Bcat)
	fmt.Println("bcat", string(b))
	//req.Cur
	b, _ = json.Marshal(bidrequest.Cur)
	fmt.Println("Cur", string(b))
	//req.Site
	b, _ = json.Marshal(bidrequest.Site)
	fmt.Println("Site", string(b))
	//req.WSeat
	b, _ = json.Marshal(bidrequest.WSeat)
	fmt.Println("WSeat", string(b))
	//req.AllImps
	b, _ = json.Marshal(bidrequest.AllImps)
	fmt.Println("AllImps", string(b))
	//req.Pmp
	b, _ = json.Marshal(bidrequest.Pmp)
	fmt.Println("Pmp", string(b))
	//req.Imp
	b, _ = json.Marshal(bidrequest.Imp)
	fmt.Println("Imp", string(b))
	//b, _ = json.Marshal(req)

	//fmt.Println(string(b))
	return openrtb.BidResponse{}
}
