package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Package-level struct definitions for use with methods
type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Rating struct {
	Rating      int
	ProductName string
	Comments    []string
	WhoRated    string
}

func main() {

	color.Green("Hello, World!")
	str := "Line1 \nLine2 \nLine3"
	fmt.Println(str)

	fmt.Println("--------------------------------")

	str1 := `
Line1
Line2
Line3
`
	fmt.Println("Multiline string:", str1)

	fmt.Println("--------------------------------")

	str2 := "Hello, World! 👋 🌍"
	fmt.Println(str2)

	fmt.Println("--------------------------------")

	var productID = "item-123"
	var productRating = 4
	var ratingComment = "Great product!"
	var whoRated = "user-456"

	fmt.Println("Product:", productID, "Rating:", productRating, "Rated by:", whoRated, "Comment:", ratingComment)

	fmt.Println("--------------------------------")

	// Map with string keys and values that can be string, int, or float
	fruits := make(map[string]int)

	fruits["apple"] = 1
	fruits["banana"] = 2
	fruits["cherry"] = 3

	fmt.Println(fruits)

	_, exists := fruits["mango"]
	if exists {
		fmt.Println("Mango", "found")
	} else {
		fmt.Println("Mango not found")
	}

	v, exists := fruits["apple"]
	if exists {
		fmt.Println(v, "price of apple")
	} else {
		fmt.Println("price of apple not found")
	}

	product1 := Product{
		ID:       "123",
		Name:     "Apple",
		Price:    1.99,
		Quantity: 10,
	}

	fmt.Println("details of product1", product1)

	var product2 Product
	product2.Name = "Banana"

	fmt.Println("name of product2", product2.Name)

	fmt.Println("--------------------------------")

	// Example ratings
	rating1 := Rating{
		Rating:      5,
		ProductName: "Laptop",
		Comments:    []string{"Excellent product!", "Fast delivery", "Great quality"},
		WhoRated:    "user-123",
	}

	rating2 := Rating{
		Rating:      4,
		ProductName: "Phone",
		Comments:    []string{"Good value for money", "Nice features"},
		WhoRated:    "user-456",
	}

	rating3 := Rating{
		Rating:      3,
		ProductName: "Tablet",
		Comments:    []string{"It's okay", "Could be better"},
		WhoRated:    "user-789",
	}

	rating4 := Rating{
		Rating:      2,
		ProductName: "Headphones",
		Comments:    []string{"Not satisfied", "Poor quality"},
		WhoRated:    "user-321",
	}

	fmt.Println("=== Rating System ===")
	ratings := []Rating{rating1, rating2, rating3, rating4}
	for _, r := range ratings {
		var category string
		switch r.Rating {
		case 5:
			category = "best"
		case 3, 4:
			category = "good"
		default:
			category = "bad"
		}
		fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating, "("+category+")", "| Rated by:", r.WhoRated)
		fmt.Println("Comments:", r.Comments)
		fmt.Println()
	}

	fmt.Println("--------------------------------")

	for _, r := range ratings {
		if r.Rating < 3 {
			fmt.Println("Product rating category: bad")
			fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating)
			fmt.Println()
		}
		if r.Rating >= 3 && r.Rating < 5 {
			fmt.Println("Product rating category: good")
			fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating)
			fmt.Println()
		}
		if r.Rating == 5 {
			fmt.Println("Product rating category: best")
			fmt.Println("Product:", r.ProductName, "| Rating:", r.Rating)
			fmt.Println()
		}
	}

	// ========================================
	// NEW CODE: Same functionality using structs, functions, and methods
	// ========================================

	fmt.Println("--------------------------------")
	fmt.Println("\n========== REFACTORED APPROACH USING FUNCTIONS ==========")
	fmt.Println("--------------------------------")

	// Rating system using structs and methods
	newRating1 := Rating{
		Rating:      5,
		ProductName: "Laptop",
		Comments:    []string{"Excellent product!", "Fast delivery", "Great quality"},
		WhoRated:    "user-123",
	}

	newRating2 := Rating{
		Rating:      4,
		ProductName: "Phone",
		Comments:    []string{"Good value for money", "Nice features"},
		WhoRated:    "user-456",
	}

	newRating3 := Rating{
		Rating:      3,
		ProductName: "Tablet",
		Comments:    []string{"It's okay", "Could be better"},
		WhoRated:    "user-789",
	}

	newRating4 := Rating{
		Rating:      2,
		ProductName: "Headphones",
		Comments:    []string{"Not satisfied", "Poor quality"},
		WhoRated:    "user-321",
	}

	newRatings := []Rating{newRating1, newRating2, newRating3, newRating4}

	fmt.Println("=== Rating System ===")
	for _, r := range newRatings {
		DisplayRatingInfo(r)
	}

	fmt.Println("--------------------------------")

	fmt.Println("Products with 'best' category:")
	DisplayRatingByCategory(newRatings, "best")

	fmt.Println("Products with 'good' category:")
	DisplayRatingByCategory(newRatings, "good")

	fmt.Println("Products with 'bad' category:")
	DisplayRatingByCategory(newRatings, "bad")
}
