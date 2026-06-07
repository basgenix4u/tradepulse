package middleware
import ("net/http"; "strings"; "github.com/gin-gonic/gin"; "github.com/tradepulse/backend/internal/auth")
func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" { c.JSON(401, gin.H{"error": "No auth"}); c.Abort(); return }
		t := strings.TrimPrefix(h, "Bearer ")
		claims, err := auth.ValidateToken(t, secret)
		if err != nil { c.JSON(401, gin.H{"error": "Invalid"}); c.Abort(); return }
		c.Set("user_id", claims.UserID); c.Next()
	}
}
