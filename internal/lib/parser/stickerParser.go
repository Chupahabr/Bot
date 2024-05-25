package parser

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"skin-monkey/internal/entity"
	repository "skin-monkey/internal/repository/postgres"
	"strconv"
	"time"
)

const (
	steamSearchUri = "https://steamcommunity.com/market/search/render"
)

type StickerParser struct {
	log  *slog.Logger
	repo *repository.Repository
}

func NewStickerParser(log *slog.Logger, repo *repository.Repository) *StickerParser {
	return &StickerParser{
		log:  log,
		repo: repo,
	}
}

func (s *StickerParser) Run() {
	start := 0

	for {
		responseBody := s.Request(start)

		for _, val := range responseBody.Results {
			sticker := entity.Sticker{
				InstanceId:    val.AssetDescription.Classid,
				Name:          val.Name,
				HashName:      val.HashName,
				SellPrice:     val.SellPrice,
				SellPriceText: val.SellPriceText,
			}

			err := s.repo.Sticker.CreateSticker(&sticker)
			if err != nil {
				fmt.Println(err)
			}
		}

		if start >= responseBody.TotalCount {
			break
		}

		start = start + responseBody.PageSize
	}
}

type AssetDescription struct {
	Appid           int    `json:"appid"`
	Classid         string `json:"classid"`
	Instanceid      string `json:"instanceid"`
	BackgroundColor string `json:"background_color"`
	IconURL         string `json:"icon_url"`
	Tradable        int    `json:"tradable"`
	Name            string `json:"name"`
	NameColor       string `json:"name_color"`
	Type            string `json:"type"`
	MarketName      string `json:"market_name"`
	MarketHashName  string `json:"market_hash_name"`
	Commodity       int    `json:"commodity"`
}

type Result struct {
	Name             string           `json:"name"`
	HashName         string           `json:"hash_name"`
	SellListings     int              `json:"sell_listings"`
	SellPrice        int              `json:"sell_price"`
	SellPriceText    string           `json:"sell_price_text"`
	AppIcon          string           `json:"app_icon"`
	AppName          string           `json:"app_name"`
	AssetDescription AssetDescription `json:"asset_description"`
	SalePriceText    string           `json:"sale_price_text"`
}

type ResponseData struct {
	Success    bool `json:"success"`
	Start      int  `json:"start"`
	PageSize   int  `json:"pagesize"`
	TotalCount int  `json:"total_count"`
	SearchData struct {
		Query              string `json:"query"`
		SearchDescriptions bool   `json:"search_descriptions"`
		TotalCount         int    `json:"total_count"`
		PageSize           int    `json:"pagesize"`
		Prefix             string `json:"prefix"`
		ClassPrefix        string `json:"class_prefix"`
	} `json:"searchdata"`
	Results []Result `json:"results"`
}

func (s *StickerParser) Request(page int) ResponseData {
	queryParams := map[string]string{
		"appid":               "730",
		"start":               strconv.Itoa(page),
		"count":               "100",
		"category_730_Type[]": "tag_CSGO_Tool_Sticker",
		"norender":            "1",
		"sort_column":         "price",
		"sort_dir":            "asc",
	}

	req, err := http.NewRequest("GET", steamSearchUri, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ResponseData{}
	}

	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ResponseData{}
	}

	defer resp.Body.Close()

	var responseData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return ResponseData{}
	}

	if len(responseData.Results) == 0 {
		time.Sleep(1 * time.Second)
		return s.Request(page)
	}

	return responseData
}
