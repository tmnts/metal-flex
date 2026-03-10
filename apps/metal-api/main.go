package main

import (
	"log"
	"net/http" // Standard library for web stuff period ;)
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// This route now has the "Senior" status code AND the "Metal-Man" vibe
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     "Metal-Flex API",
			"status":  "Forging precision math...",
			"engine":  "Gin + Decimal",
			"message": "U a go be Metal-Man now!",
			"vibe":    "Jamaican-Brat-2-Hybrid",
		})
	})

	// 3. Precision Calculation Route
	// GET /calculate?gold=2500.50&silver=25.15
	r.GET("/calculate", func(c *gin.Context) {
		apiKey := os.Getenv("METAL_API_KEY")
		if apiKey == "" {
			// 1. Пишем в логи Vercel (ты увидишь это в Dashboard -> Logs)
			log.Println("ERROR: METAL_API_KEY is not set in environment variables")

			// 2. Отвечаем фронтенду понятной ошибкой
			c.JSON(500, gin.H{
				"error":  "Server configuration error: missing API key",
				"status": "Бамбаклаат! Забыли прокинуть токен в настройках Vercel",
			})
			return // Прерываем выполнение, чтобы не стучаться в Metalprice с пустым ключом
		}

		// 1. Collect strings from URL
		goldStr := c.DefaultQuery("gold", "0")
		silverStr := c.DefaultQuery("silver", "1")
		platStr := c.DefaultQuery("plat", "0")
		pallStr := c.DefaultQuery("pall", "0")
		cuStr := c.DefaultQuery("cu", "0")
		feStr := c.DefaultQuery("fe", "0")
		tiStr := c.DefaultQuery("ti", "0")

		// 2. Transform into High-Precision Decimals
		goldPrice, _ := decimal.NewFromString(goldStr)
		silverPrice, _ := decimal.NewFromString(silverStr)
		platPrice, _ := decimal.NewFromString(platStr)
		pallPrice, _ := decimal.NewFromString(pallStr)
		cuPrice, _ := decimal.NewFromString(cuStr)
		fePrice, _ := decimal.NewFromString(feStr)
		tiPrice, _ := decimal.NewFromString(tiStr)

		// 3. DO THE MATH (The "Portfolio" & "Ratios")
		// Total Precious Portfolio
		totalPrecious := goldPrice.Add(silverPrice).Add(platPrice).Add(pallPrice)

		// Gold-to-Silver Ratio
		gsRatio := goldPrice.Div(silverPrice)

		// Gold-to-Copper Ratio (The "Industrial Health" Index)
		// Note: Cu is per Tonne, so we convert it to Ounces to compare with Gold!
		// 1 Tonne = 32,150 Troy Ounces
		cuPerOunce := cuPrice.Div(decimal.NewFromInt(32150))
		gcRatio := goldPrice.Div(cuPerOunce)

		// 4. THE JSON RESPONSE
		c.JSON(http.StatusOK, gin.H{
			"total_precious": totalPrecious.StringFixed(2),
			"gs_ratio":       gsRatio.StringFixed(2),
			"gold_copper":    gcRatio.StringFixed(0), // Usually a big number!
			"industrial": gin.H{
				"iron":     fePrice,
				"titanium": tiPrice,
			},
			"api_status": "Syncing with " + apiKey[0:4] + "****",
			"message":    "U a go be Industriaman now! 🚀 Ну и сила в точности.",
		})
	})

	// 5. Vercel serverless run

}
