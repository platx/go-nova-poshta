Nova Poshta Golang API client
=============================

[![Version](https://img.shields.io/github/go-mod/go-version/platx/go-nova-poshta)](https://go.dev/doc/go1.20)
[![Release](https://img.shields.io/github/v/release/platx/go-nova-poshta.svg)](https://github.com/platx/go-nova-poshta/releases)
[![codecov](https://codecov.io/gh/platx/go-nova-poshta/branch/main/graph/badge.svg?token=BIJNOMOK8V)](https://codecov.io/gh/platx/go-nova-poshta)
[![Go Report](https://goreportcard.com/badge/github.com/platx/go-nova-poshta)](https://goreportcard.com/badge/github.com/platx/go-nova-poshta)
[![GoDoc](https://godoc.org/github.com/platx/go-nova-poshta?status.svg)](https://pkg.go.dev/github.com/platx/go-nova-poshta)
[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://developers.novaposhta.ua/documentation)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/platx/go-nova-poshta/blob/master/LICENSE)

A simple API client for [Nova Poshta](https://developers.novaposhta.ua) written in Go.

# API documentation

https://developers.novaposhta.ua/documentation

# Go Versions Supported

We currently support Go versions 1.20 and newer.

# Installation

```
go get github.com/platx/go-nova-poshta
```

# Usage

## Initialize client

You should get `your_api_token` in the Nova Poshta [security settings page](https://new.novaposhta.ua/dashboard/settings/developers)

### Basic config
```Go
package main

import (
    "net/http"

    "github.com/platx/go-nova-poshta"
    "github.com/platx/go-nova-poshta/adapter"
)

func main() {
    client := go_nova_poshta.NewClient(adapter.NewConfig("your_api_token"))
}
```

### Custom config
You are able to customize optional settings, like this:

```Go
package main

import (
    "log"
    "net/http"

    "github.com/platx/go-nova-poshta"
    "github.com/platx/go-nova-poshta/adapter"
)

func main() {
    client := go_nova_poshta.NewClient(adapter.NewConfig(
        "your_api_token",                           // required
        adapter.WithHTTPClient(http.DefaultClient), // custom http client to send requests
        adapter.WithFormat("json"),                 // or "xml"
        adapter.WithDebug(true),                    // debug mode
        adapter.WithLogger(...),          // custom logger
    ))
}
```

## Call API methods

Client uses models (method groups) to call methods. Each model has a set of methods to work with it.

For example, you want to get list of cities from Nova Poshta API. You should use `Address` model and `GetCities` method, like this:
```Go
package main

import (
    "fmt"
	
    "github.com/platx/go-nova-poshta"
    "github.com/platx/go-nova-poshta/adapter"
    "github.com/platx/go-nova-poshta/api/address"
    "github.com/platx/go-nova-poshta/utils"
)

func main() {
    client := go_nova_poshta.NewClient(adapter.NewConfig("your_api_token"))

    cities, err := client.Address().GetCities(address.GetCitiesRequest{
        FindByString: utils.PTR("Київ"),
        Page: utils.PTR(1),
        Limit: utils.PTR(10),
    })
	
    if err != nil {
        panic(err)
    }
	
    fmt.Println(cities)
}
```

### All available models and methods

| Model             | Method                         | Request params                                      | Response data                                         | Usage                                                                                                            | Documentation                                                                                                                        |
|-------------------|--------------------------------|-----------------------------------------------------|-------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| Address           | SearchSettlements              | `address.SearchSettlementsReq`                      | `address.SearchSettlementRes`                         | `client.Address().SearchSettlements(address.SearchSettlementsReq{})`                                             | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a0eb83ab-8512-11ec-8ced-005056b2dbe1) |
| Address           | SearchSettlementStreets        | `address.SearchSettlementStreetsReq`                | `address.SearchSettlementStreetsRes`                  | `client.Address().SearchSettlementStreets(address.SearchSettlementStreetsReq{})`                                 | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1329635-8512-11ec-8ced-005056b2dbe1) |
| Address           | Save                           | `address.CreateReq`                                 | `address.SaveRes`                                     | `client.Address().Save(address.CreateReq{})`                                                                     | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a155d0d9-8512-11ec-8ced-005056b2dbe1) |
| Address           | Update                         | `address.UpdateReq`                                 | `address.SaveRes`                                     | `client.Address().Update(address.UpdateReq{})`                                                                   | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a19ba934-8512-11ec-8ced-005056b2dbe1) |
| Address           | Delete                         | `address.DeleteReq`                                 | `address.DeleteRes`                                   | `client.Address().Delete(address.DeleteReq{})`                                                                   | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a177069a-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetSettlements                 | `address.GetSettlementsReq`                         | `address.GetSettlementsRes`                           | `client.Address().GetSettlements(address.GetSettlementsReq{})`                                                   | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1c42723-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetCities                      | `address.GetCitiesReq`                              | `address.GetCitiesRes`                                | `client.Address().GetCities(address.GetCitiesReq{})`                                                             | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1e6f0a7-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetAreas                       | -                                                   | `address.GetAreasRes`                                 | `client.Address().GetAreas()`                                                                                    | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a20ee6e4-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetWarehouses                  | `address.GetWarehousesReq`                          | `address.GetWarehousesRes`                            | `client.Address().GetWarehouses(address.GetWarehousesReq{})`                                                     | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2322f38-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetWarehouseTypes              | -                                                   | `address.GetWarehouseTypesRes`                        | `client.Address().GetWarehouseTypes()`                                                                           | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2587b53-8512-11ec-8ced-005056b2dbe1) |
| Address           | GetStreet                      | `address.GetStreetReq`                              | `address.GetStreetRes`                                | `client.Address().GetStreet(address.GetStreetReq{})`                                                             | [Link](https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a27c20d7-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| Counterparty      | SavePrivatePerson              | `counterparty.SavePrivatePersonReq`                 | `counterparty.SaveRes`                                | `client.Counterparty().SavePrivatePerson(counterparty.SavePrivatePersonReq{})`                                   | [Link](https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/0ae5dd75-8a5f-11ec-8ced-005056b2dbe1) |
| Counterparty      | SaveThirdPerson                | `counterparty.SaveThirdPerson`                      | `counterparty.SaveRes`                                | `client.Counterparty().SaveThirdPerson(counterparty.SaveThirdPerson{})`                                          | [Link](https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/b0fdf818-8a8e-11ec-8ced-005056b2dbe1) |
| Counterparty      | SaveOrganization               | `counterparty.SaveOrganizationReq`                  | `counterparty.SaveRes`                                | `client.Counterparty().SaveOrganization(counterparty.SaveOrganizationReq{})`                                     | [Link](https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/bc3c44c7-8a8a-11ec-8ced-005056b2dbe1) |
| Counterparty      | Update                         | `counterparty.UpdateReq`                            | `counterparty.SaveRes`                                | `client.Counterparty().Update(counterparty.UpdateReq{})`                                                         | [Link](https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2c3c708-8512-11ec-8ced-005056b2dbe1) |
| Counterparty      | Delete                         | `counterparty.DeleteReq`                            | `counterparty.DeleteRes`                              | `client.Counterparty().Delete(counterparty.DeleteReq{})`                                                         | [Link](https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2eb27e8-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| ContactPerson     | Save                           | `contactperson.CreateReq`                           | `contactperson.SaveRes`                               | `client.ContactPerson().Save(contactperson.CreateReq{})`                                                         | [Link](https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3a25bda-8512-11ec-8ced-005056b2dbe1) |
| ContactPerson     | Update                         | `contactperson.UpdateReq`                           | `contactperson.SaveRes`                               | `client.ContactPerson().Update(contactperson.UpdateReq{})`                                                       | [Link](https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3c5a577-8512-11ec-8ced-005056b2dbe1) |
| ContactPerson     | Delete                         | `contactperson.DeleteReq`                           | `contactperson.DeleteRes`                             | `client.ContactPerson().Delete(contactperson.DeleteReq{})`                                                       | [Link](https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3ea91c8-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| ScanSheet         | GetScanSheetList               | -                                                   | `scansheet.GetScanSheetListRes`                       | `client.ScanSheet().GetScanSheetList()`                                                                          | [Link](https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a4d93a89-8512-11ec-8ced-005056b2dbe1) |
| ScanSheet         | GetScanSheet                   | `scansheet.GetScanSheetReq`                         | `scansheet.GetScanSheetRes`                           | `client.ScanSheet().GetScanSheet(scansheet.GetScanSheetReq{})`                                                   | [Link](https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a4abdd36-8512-11ec-8ced-005056b2dbe1) |
| ScanSheet         | InsertDocuments                | `scansheet.InsertDocumentsReq `                     | `scansheet.InsertDocumentsRes `                       | `client.ScanSheet().InsertDocuments(scansheet.InsertDocumentsReq{})`                                             | [Link](https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a482293c-8512-11ec-8ced-005056b2dbe1) |
| ScanSheet         | RemoveDocuments                | `scansheet.RemoveDocumentsReq `                     | `scansheet.RemoveDocumentsRes `                       | `client.ScanSheet().RemoveDocuments(scansheet.RemoveDocumentsReq{})`                                             | [Link](https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a53dea8a-8512-11ec-8ced-005056b2dbe1) |
| ScanSheet         | DeleteScanSheet                | `scansheet.DeleteScanSheetReq `                     | `scansheet.DeleteScanSheetRes `                       | `client.ScanSheet().DeleteScanSheet(scansheet.DeleteScanSheetReq{})`                                             | [Link](https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a50e049b-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| Common            | GetTimeIntervals               | `common.GetTimeIntervalsReq`                        | `common.GetTimeIntervalsRes`                          | `client.Common().GetTimeIntervals(contactperson.GetTimeIntervalsReq{})`                                          | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a56d5c1c-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetCargoTypes                  | -                                                   | `common.ListItemsRes[enum.CargoType]`                 | `client.Common().GetCargoTypes()`                                                                                | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5912a1e-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetBackwardDeliveryCargoTypes  | -                                                   | `common.ListItemsRes[enum.CargoType]`                 | `client.Common().GetBackwardDeliveryCargoTypes()`                                                                | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5b46873-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetPalletsList                 | -                                                   | `common.GetPalletsListRes`                            | `client.Common().GetPalletsList()`                                                                               | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5dd575e-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetTypesOfPayersForRedelivery  | -                                                   | `common.ListItemsRes[string]`                         | `client.Common().GetTypesOfPayersForRedelivery()`                                                                | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6247f2f-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetPackList                    | -                                                   | `common.GetPackListRes`                               | `client.Common().GetPackList()`                                                                                  | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6492db4-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetTiresWheelsList             | -                                                   | `common.GetTiresWheelsListRes`                        | `client.Common().GetTiresWheelsList()`                                                                           | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a66fada0-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetCargoDescriptionList        | -                                                   | `common.GetCargoDescriptionListRes`                   | `client.Common().GetCargoDescriptionList()`                                                                      | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a697db47-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetMessageCodeText             | -                                                   | `common.GetMessageCodeTextRes`                        | `client.Common().GetMessageCodeText()`                                                                           | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6bce5a1-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetServiceTypes                | -                                                   | `common.ListItemsRes[enum.ServiceType]`               | `client.Common().GetServiceTypes()`                                                                              | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6e189f7-8512-11ec-8ced-005056b2dbe1) |
| Common            | GetOwnershipFormsList          | -                                                   | `common.GetOwnershipFormsListRes`                     | `client.Common().GetOwnershipFormsList()`                                                                        | [Link](https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a754ff0d-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| AdditionalService | CheckPossibilityCreateReturn   | `additionalservice.CheckPossibilityCreateReturnReq` | `additionalservice.CheckPossibilityCreateReturnRes`   | `client.AdditionalService().CheckPossibilityCreateReturn(additionalservice.CheckPossibilityCreateReturnReq{})`   | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a778f519-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | GetReturnReasons               | -                                                   | `additionalservice.GetReturnReasonsRes`               | `client.AdditionalService().GetReturnReasons()`                                                                  | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7a6bacb-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | GetReturnReasonsSubtypes       | `additionalservice.GetReturnReasonsSubtypesReq`     | `additionalservice.GetReturnReasonsSubtypesRes`       | `client.AdditionalService().GetReturnReasonsSubtypes(additionalservice.GetReturnReasonsSubtypesReq{})`           | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7cb69ee-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | CheckPossibilityForRedirecting | `additionalservice.CheckPossibilityCreateReturnReq` | `additionalservice.CheckPossibilityForRedirectingRes` | `client.AdditionalService().CheckPossibilityForRedirecting(additionalservice.CheckPossibilityCreateReturnReq{})` | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8d29fc2-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | CheckPossibilityChangeEW       | `additionalservice.CheckPossibilityChangeEWReq`     | `additionalservice.CheckPossibilityChangeEWRes`       | `client.AdditionalService().CheckPossibilityChangeEW(additionalservice.CheckPossibilityChangeEWReq{})`           | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a886b776-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | SaveReturn                     | `additionalservice.SaveReturnReq`                   | `additionalservice.SaveRes`                           | `client.AdditionalService().SaveReturn(additionalservice.SaveReturnReq{})`                                       | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7fb4a3a-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | SaveReturnNewAddress           | `additionalservice.SaveReturnNewAddressReq`         | `additionalservice.SaveRes`                           | `client.AdditionalService().SaveReturnNewAddress(additionalservice.SaveReturnNewAddressReq{})`                   | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/175baec3-8f0d-11ec-8ced-005056b2dbe1) |
| AdditionalService | SaveReturnNewWarehouse         | `additionalservice.SaveReturnNewWarehouseReq`       | `additionalservice.SaveRes`                           | `client.AdditionalService().SaveReturnNewWarehouse(additionalservice.SaveReturnNewWarehouseReq{})`               | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/5a64f960-e7fa-11ec-a60f-48df37b921db) |
| AdditionalService | SaveRedirecting                | `additionalservice.SaveRedirectingReq`              | `additionalservice.SaveRes`                           | `client.AdditionalService().SaveRedirecting(additionalservice.SaveRedirectingReq{})`                             | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/98acb0f6-8f0b-11ec-8ced-005056b2dbe1) |
| AdditionalService | SaveChangeEW                   | `additionalservice.SaveChangeEWReq`                 | `additionalservice.SaveRes`                           | `client.AdditionalService().SaveChangeEW(additionalservice.SaveChangeEWReq{})`                                   | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/c09f1b02-8a66-11ec-8ced-005056b2dbe1) |
| AdditionalService | Delete                         | `additionalservice.DeleteReq`                       | `additionalservice.DeleteRes`                         | `client.AdditionalService().Delete(additionalservice.DeleteReq{})`                                               | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a85bb34b-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | GetReturnOrdersList            | `additionalservice.GetOrdersListReq`                | `additionalservice.GetOrdersListRes`                  | `client.AdditionalService().GetReturnOrdersList(additionalservice.GetOrdersListReq{})`                           | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a82d087c-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | GetChangeEWOrdersList          | `additionalservice.GetOrdersListReq`                | `additionalservice.GetOrdersListRes`                  | `client.AdditionalService().GetChangeEWOrdersList(additionalservice.GetOrdersListReq{})`                         | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8ad1dc7-8512-11ec-8ced-005056b2dbe1) |
| AdditionalService | GetRedirectionOrdersList       | `additionalservice.GetOrdersListReq`                | `additionalservice.GetOrdersListRes`                  | `client.AdditionalService().GetRedirectionOrdersList(additionalservice.GetOrdersListReq{})`                      | [Link](https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8faa2e6-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| InternetDocument  | GetDocumentPrice               | `internetdocument.GetDocumentPriceReq`              | `internetdocument.GetDocumentPriceRes`                | `client.InternetDocument().GetDocumentPrice(internetdocument.GetDocumentPriceReq{})`                             | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a91f115b-8512-11ec-8ced-005056b2dbe1) |
| InternetDocument  | GetDocumentDeliveryDate        | `internetdocument.GetDocumentDeliveryDateReq`       | `internetdocument.GetDocumentDeliveryDateRes`         | `client.InternetDocument().GetDocumentDeliveryDate(internetdocument.GetDocumentDeliveryDateReq{})`               | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a941c714-8512-11ec-8ced-005056b2dbe1) |
| InternetDocument  | Save                           | `internetdocument.SaveReq`                          | `internetdocument.SaveRes`                            | `client.InternetDocument().Save(internetdocument.SaveReq{})`                                                     | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a965630e-8512-11ec-8ced-005056b2dbe1) |
| InternetDocument  | SavePostomat                   | `internetdocument.PostomatSaveReq`                  | `internetdocument.SaveRes`                            | `client.InternetDocument().SavePostomat(internetdocument.PostomatSaveReq{})`                                     | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/0227072e-8f38-11ec-8ced-005056b2dbe1) |
| InternetDocument  | SaveWarehouse                  | `internetdocument.WarehouseSaveReq`                 | `internetdocument.SaveRes`                            | `client.InternetDocument().SaveWarehouse(internetdocument.WarehouseSaveReq{})`                                   | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a19ced79-8f32-11ec-8ced-005056b2dbe1) |
| InternetDocument  | SaveAddress                    | `internetdocument.AddressSaveReq`                   | `internetdocument.SaveRes`                            | `client.InternetDocument().SaveAddress(internetdocument.AddressSaveReq{})`                                       | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/f74a0918-8f18-11ec-8ced-005056b2dbe1) |
| InternetDocument  | Update                         | `internetdocument.UpdateReq`                        | `internetdocument.SaveRes`                            | `client.InternetDocument().Update(internetdocument.UpdateReq{})`                                                 | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a98a4354-8512-11ec-8ced-005056b2dbe1) |
| InternetDocument  | Delete                         | `internetdocument.DeleteReq`                        | `internetdocument.DeleteRes`                          | `client.InternetDocument().Delete(internetdocument.DeleteReq{})`                                                 | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a9f43ff1-8512-11ec-8ced-005056b2dbe1) |
| InternetDocument  | GetDocumentList                | `internetdocument.GetDocumentListReq`               | `internetdocument.GetDocumentListRes`                 | `client.InternetDocument().GetDocumentList(internetdocument.GetDocumentListReq{})`                               | [Link](https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a9d22b34-8512-11ec-8ced-005056b2dbe1) |
|                   |                                |                                                     |                                                       |                                                                                                                  |                                                                                                                                      |
| TrackingDocument  | GetStatusDocuments             | `trackingdocument.GetStatusDocumentsReq`            | `trackingdocument.GetStatusDocumentsRes`              | `client.TrackingDocument().GetStatusDocuments(trackingdocument.GetStatusDocumentsReq{})`                         | [Link](https://developers.novaposhta.ua/view/model/a99d2f28-8512-11ec-8ced-005056b2dbe1/method/a9ae7bc9-8512-11ec-8ced-005056b2dbe1) |

### Custom API calls
If you are not ok with the provided API models/methods, you can always make your own custom calls. For example, if you want to make a call to the `Address` model with `getCities` method, you can do it like this:
```go
package main

import (
    "fmt"
	
    "github.com/platx/go-nova-poshta"
    "github.com/platx/go-nova-poshta/adapter"
)

func main() {
    client := go_nova_poshta.NewClient(adapter.NewConfig("your_api_token"))

    req := map[string]any{
        "FindByString": "Київ",
        "Limit":        10,
        "Page":         1,
    }

    var res []map[string]any

    err := client.Custom().Call("Address", "getCities", req, &res)

    if err != nil {
        panic(err)
    }

    fmt.Println(res)
}
```

_P.S. In this case you have take care for the request/response structs by yourself._