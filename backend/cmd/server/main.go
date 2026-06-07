package main
import (
	"log"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tradepulse/backend/internal/config"
	"github.com/tradepulse/backend/internal/db"
	"github.com/tradepulse/backend/internal/handlers"
	"github.com/tradepulse/backend/internal/middleware"
	"github.com/tradepulse/backend/internal/websocket"
)
func main() {
	godotenv.Load()
	cfg := config.Load()
	log.Printf("Starting TradePulse v2.0")
	
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer database.Close()
	
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	
	hub := websocket.NewHub()
	go hub.Run()
	
	marketService := handlers.NewMarketService(hub, cfg)
	go marketService.Start()
	
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	authHandler := handlers.NewAuthHandler(database, cfg)
	
	api := r.Group("/api/v1")
	api.GET("/health", handlers.HealthCheck)
	api.POST("/auth/register", authHandler.Register)
	api.POST("/auth/login", authHandler.Login)
	api.GET("/market/overview", marketService.GetMarketOverview)
	api.GET("/market/asset/:symbol", marketService.GetAssetDetails)
	
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	protected.GET("/auth/me", authHandler.Me)
	
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWS(hub, c.Writer, c.Request)
	})
	
	port := cfg.Port
	log.Printf("Server running on :%s", port)
	r.Run(":" + port)
}
