package mas_provider

import (
	"fmt"
	"github.com/MaisrForAdvancedSystems/biller-mas-provider/providers"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)



type MasProvider struct {
	providers.InfoProvider
	providers.DataProvider
	providers.TariffProvider
	providers.DataConsumer
}

func validateTypes() {
	var d billing.BillingDataProviderServer=&providers.DataProvider{}
	var t billing.BillingTariffProviderServer=&providers.TariffProvider{}
	var c billing.BillingDataCousumerServer=&providers.DataConsumer{}
	fmt.Println(d,t,c)
}
