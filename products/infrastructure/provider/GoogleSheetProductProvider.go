package provider

import (
	"context"
	"fmt"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
	"github.com/prodriguezval/delicaria_products/products/domain/provider/model"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"strconv"
	"strings"
)

type GoogleSheetProductProvider struct{}

func NewSheetProductProvider() provider.ProductProvider {
	return GoogleSheetProductProvider{}
}

func (p GoogleSheetProductProvider) GetByProviderName(providerName string) []model.ProductProviderResponse {
	allProducts := p.GetAll()
	var response []model.ProductProviderResponse

	for _, product := range allProducts {
		if !strings.Contains(strings.ToLower(product.Provider), strings.ToLower(providerName)) {
			continue
		}
		response = append(response, product)
	}

	return response
}

func (p GoogleSheetProductProvider) GetAll() []model.ProductProviderResponse {
	var products []model.ProductProviderResponse
	_, err := google.FindDefaultCredentials(context.Background(), sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error authenticating with credentials: %v", err)
	}

	client, err := google.DefaultClient(context.Background(), sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error creating GSheet client: %v", err)
	}

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating service: %v", err)
	}

	spreadsheetId := "14vCuyJo-oPDoWwj-1NCOV_pNIdxYO0OjW-WtWmIKQeA"
	sheetName := "Productos"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, sheetName).Do()
	if err != nil {
		log.Fatalf("Can't read data: %v", err)
	}

	if len(resp.Values) == 0 {
		log.Printf("Spreadsheet %s No data found on sheet %s", spreadsheetId, sheetName)
		return products
	}

	for i, row := range resp.Values {
		if i == 0 {
			continue
		}

		productName := row[1].(string)
		if productName == "" {
			log.Printf("Invalid row in line %d, please check %v", i+1, row)
			continue
		}

		id, _ := strconv.Atoi(fmt.Sprintf("%v", row[0]))
		var buyPrice, _ = strconv.ParseFloat(fmt.Sprintf("%v", row[5]), 64)
		var tax, _ = strconv.ParseFloat(fmt.Sprintf("%v", row[7]), 64)
		var salePrice, _ = strconv.ParseFloat(fmt.Sprintf("%v", row[9]), 64)
		var earningAmount, _ = strconv.ParseFloat(fmt.Sprintf("%v", row[6]), 64)

		products = append(
			products,
			model.ProductProviderResponse{
				Id:            id,
				Name:          productName,
				Provider:      fmt.Sprintf("%v", row[2]),
				BuyPrice:      buyPrice,
				Tax:           tax,
				SalePrice:     salePrice,
				EarningAmount: earningAmount,
			},
		)
	}

	return products
}
