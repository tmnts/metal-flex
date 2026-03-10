package handler // Vercel любит этот пакет для функций

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// ВАЖНО: Это точка входа для Vercel.
// Она заменяет func main() и r.Run()
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Инициализируем Gin в режиме релиза (чтобы не спамить в логи)
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(cors.Default())

	// 2. Твои роуты (пути должны совпадать с тем, что шлет фронтенд)

	// Главная (опционально, для теста /api/calculate/)
	app.GET("/api/calculate", func(c *gin.Context) {
		apiKey := os.Getenv("METAL_API_KEY")

		// Проверка токена
		if apiKey == "" {
			log.Println("ERROR: METAL_API_KEY is not set")
			c.JSON(500, gin.H{
				"error":  "Missing API key",
				"status": "Бамбаклаат! Забыли токен на Vercel",
			})
			return
		}

		// --- ТВОЯ ЛОГИКА РАСЧЕТА ---

		// 1. Сбор данных из URL
		goldStr := c.DefaultQuery("gold", "0")
		silverStr := c.DefaultQuery("silver", "1")
		platStr := c.DefaultQuery("plat", "0")
		pallStr := c.DefaultQuery("pall", "0")
		cuStr := c.DefaultQuery("cu", "0")
		feStr := c.DefaultQuery("fe", "0")
		tiStr := c.DefaultQuery("ti", "0")

		// 2. Превращаем в High-Precision Decimals
		goldPrice, _ := decimal.NewFromString(goldStr)
		silverPrice, _ := decimal.NewFromString(silverStr)
		platPrice, _ := decimal.NewFromString(platStr)
		pallPrice, _ := decimal.NewFromString(pallStr)
		cuPrice, _ := decimal.NewFromString(cuStr)
		fePrice, _ := decimal.NewFromString(feStr)
		tiPrice, _ := decimal.NewFromString(tiStr)

		// 3. Расчеты
		totalPrecious := goldPrice.Add(silverPrice).Add(platPrice).Add(pallPrice)
		gsRatio := goldPrice.Div(silverPrice)
		cuPerOunce := cuPrice.Div(decimal.NewFromInt(32150))
		gcRatio := goldPrice.Div(cuPerOunce)

		// 4. ФИНАЛЬНЫЙ JSON
		c.JSON(http.StatusOK, gin.H{
			"total_precious": totalPrecious.StringFixed(2),
			"gs_ratio":       gsRatio.StringFixed(2),
			"gold_copper":    gcRatio.StringFixed(0),
			"industrial": gin.H{
				"iron":     fePrice,
				"titanium": tiPrice,
			},
			"api_status": "Syncing with " + apiKey[0:4] + "****",
			"message":    "U a go be Industriaman! 🚀 Сила в точности.",
		})
	})

	// 3. МАГИЯ: Передаем запрос от Vercel в Gin
	app.ServeHTTP(w, r)
}
