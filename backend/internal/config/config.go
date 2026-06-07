package config
import "os"
type Config struct {
	DatabaseURL string
	JWTSecret string
	FinnhubAPIKey string
	Port string
}
func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "dev-secret-change-in-prod"),
		FinnhubAPIKey: getEnv("FINNHUB_API_KEY", "demo"),
		Port: getEnv("PORT", "8080"),
	}
}
func getEnv(k, d string) string { if v := os.Getenv(k); v != "" { return v }; return d }
