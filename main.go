package main

import (
	routes "example/scrapper/router"

	"github.com/gin-gonic/gin"
)





func main() {
	router := gin.Default()
	routes.RouteIndex(router)

	router.Run("localhost:3000")
}

func coinmarketcap(){
	// fName := "cryptocoinmarketcap-new.csv"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s\n", fName, err)
	// 	return
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// // Write CSV header
	// writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	
}



