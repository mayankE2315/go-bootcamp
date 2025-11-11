package main

import "fmt"

func GetCategory(r Rating) string {
	if r.Rating == 5 {
		return "best"
	}
	if r.Rating >= 3 && r.Rating < 5 {
		return "good"
	}
	return "bad"
}

func DisplayRatingInfo(r Rating) {
	category := GetCategory(r)
	fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating, "| Rated by:", r.WhoRated)
	fmt.Println("Comments:", r.Comments)
	fmt.Println("Rating category:", category)
	fmt.Println()
}

func DisplayRatingByCategory(ratings []Rating, category string) {
	for _, r := range ratings {
		if GetCategory(r) == category {
			fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating)
		}
	}
	fmt.Println()
}
