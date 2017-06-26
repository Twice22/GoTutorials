package main 

import (
	"fmt"
	"log"
	"os"
	"time"

	"Chapter4/github"
)

// use: go build Chapter4/exercices/ex4.10
// ./ex4.10.exe repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	
	now := time.Now()
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month old:")
	for _, item := range result.Items {
		// if diff of Hours/24 between now and created date is less than 30
		if now.Sub(item.CreatedAt).Hours()/24 <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}


	fmt.Println("Less than a year old:")
	for _, item := range result.Items {
		// if diff of Hours/24 between now and created date is less than 30
		if now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}


	fmt.Println("More than a year old:")
	for _, item := range result.Items {
		// if diff of Hours/24 between now and created date is less than 30
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

}