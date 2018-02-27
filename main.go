package main

import (
	"fmt"
	"io/ioutil"
	"log"
	 "net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/content/v2"
)

type merchantInfo struct {
	MerchantID             uint64 `json:"merchantId,omitempty"`
	WebsiteURL             string `json:"-"`
	AccountSampleUser      string `json:"accountSampleUser,omitempty"`
	AccountSampleAdwordsID uint64 `json:"accountSampleAdWordsCID,omitempty"`
	IsMCA                  bool   `json:"-"`
	Path                   string `json:"-"`
}


func authWithGoogle(ctx context.Context, samplesConfig merchantInfo) *http.Client {
	json, err := ioutil.ReadFile("sendo-setting-dev-5b23259509e0.json")

	if err != nil {
		log.Fatal(err)
	}
	config, err := google.JWTConfigFromJSON(json, content.ContentScope)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Service account credentials for user %s found.\n", config.Email)
	return config.Client(ctx)
}

const (
	merchantID=118052097
)

func main() {
	ctx := context.Background()
	samplesConfig := merchantInfo{}
	client := authWithGoogle(ctx, samplesConfig)

	contentService, err := content.New(client)

	if err != nil {
		log.Fatal(err)
	}

	//listCall := contentService.Accounts.List(merchantID)
	//listCall.MaxResults(100)
	//
	//err = listCall.Pages(ctx, printAccountsPage)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//shopName := "Shop Test"
	//shopUserName := "shop-tre"
	//
	//account := &content.Account{
	//	Name: shopName,
	//	WebsiteUrl: fmt.Sprintf("https://www.sendo.vn/shop/%s/", shopUserName),
	//}
	//
	//account, err = contentService.Accounts.Insert(merchantID, account).Do()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Subaccount added with ID %d.\n", account.Id)
	//account, err = contentService.Accounts.Get(merchantID, account.Id).Do()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//printAccount(account)

	listCall := contentService.Datafeeds.List(105359420)
	listCall.MaxResults(100)
	err = listCall.Pages(ctx, printFeedsPage)
}

func printAccountsPage(res *content.AccountsListResponse) error {
	for _, account := range res.Resources {
		printAccount(account)
	}
	return nil
}

func printAccount(account *content.Account) {
	fmt.Printf("Information for account %d:\n", account.Id)
	fmt.Printf("- Display name: %s\n", account.Name)
	if len(account.Users) == 0 {
		fmt.Println("- No registered users.")
	} else {
		fmt.Println("- Registered users:")
		for _, user := range account.Users {
			fmt.Print("  - ")
			if user.Admin != nil && *user.Admin {
				fmt.Print("(ADMIN) ")
			}
			fmt.Println(user.EmailAddress)
		}
	}
	if len(account.AdwordsLinks) == 0 {
		fmt.Println("- No linked Adwords accounts.")
	} else {
		fmt.Println("- Linked Adwords accounts:")
		for _, link := range account.AdwordsLinks {
			fmt.Printf("  - %d: %s\n", link.AdwordsId, link.Status)
		}
	}
}

func printFeedsPage(res *content.DatafeedsListResponse) error {
	for _, feed := range res.Resources {
		fmt.Printf(" - Id %d, Name %s\n", feed.Id, feed.Name)
	}
	return nil
}

func createSampleProduct(config *merchantInfo, offerID string) *content.Product {
	websiteURL := config.WebsiteURL
	if websiteURL == "" {
		websiteURL = "http://my-book-shop.com"
	}
	productPrice := content.Price{Currency: "USD", Value: "2.50"}
	shippingPrice := content.Price{Currency: "USD", Value: "0.99"}
	shippingWeight := content.ProductShippingWeight{
		Value: 200.0,
		Unit:  "grams",
	}
	shippingInfo := content.ProductShipping{
		Country: "US",
		Service: "Standard shipping",
		Price:   &shippingPrice,
	}
	product := content.Product{
		OfferId:               offerID,
		Title:                 "A Tale of Two Cities",
		Description:           "A classic novel about the French Revolution",
		Link:                  websiteURL + "/tale-of-two-cities.html",
		ImageLink:             websiteURL + "/tale-of-two-cities.jpg",
		ContentLanguage:       "en",
		TargetCountry:         "US",
		Channel:               "online",
		Availability:          "in stock",
		Condition:             "new",
		GoogleProductCategory: "123456",
		Gtin:           "9780007350896",
		Price:          &productPrice,
		Shipping:       [](*content.ProductShipping){&shippingInfo},
		ShippingWeight: &shippingWeight,
	}
	return &product
}