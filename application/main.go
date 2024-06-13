package main

import {
	"github.com/sotaheavymetal21/go-party-box/models"
}
func main() {
	// itemsは、商品データのスライスです。
	items := []models.Item{
		{ID: 1, Name: "バッタ", Price: 1000, Description: "食用", SoldOut: false},
		{ID: 2, Name: "カナブン", Price: 5000, Description: "観賞用", SoldOut: true},
		{ID: 3, Name: "どんぐり", Price: 20000, Description: "投擲用", SoldOut: true},
	}
}