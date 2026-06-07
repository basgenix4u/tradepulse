package handlers
import ("database/sql"; "net/http"; "strings"; "github.com/gin-gonic/gin"; "github.com/tradepulse/backend/internal/auth"; "github.com/tradepulse/backend/internal/config"; "github.com/tradepulse/backend/internal/models")
type AuthHandler struct { db *sql.DB; cfg *config.Config }
func NewAuthHandler(db *sql.DB, cfg *config.Config) *AuthHandler { return &AuthHandler{db, cfg} }
func (h *AuthHandler) Register(c *gin.Context) {
	var r struct{ Email, Password, Name string }
	if err := c.ShouldBindJSON(&r); err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
	var exists bool
	h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", r.Email).Scan(&exists)
	if exists { c.JSON(409, gin.H{"error": "Email exists"}); return }
	hash, _ := auth.HashPassword(r.Password)
	var id string
	err := h.db.QueryRow("INSERT INTO users (email, name, password_hash) VALUES ($1,$2,$3) RETURNING id", strings.ToLower(r.Email), r.Name, hash).Scan(&id)
	if err != nil { c.JSON(500, gin.H{"error": "Failed to create user"}); return }
	token, _ := auth.GenerateToken(id, r.Email, h.cfg.JWTSecret)
	c.JSON(201, gin.H{"user": gin.H{"id": id, "email": r.Email, "name": r.Name}, "token": token})
}
func (h *AuthHandler) Login(c *gin.Context) {
	var r struct{ Email, Password string }
	c.ShouldBindJSON(&r)
	var u models.User; var hash string
	err := h.db.QueryRow("SELECT id,email,name,password_hash FROM users WHERE email=$1", strings.ToLower(r.Email)).Scan(&u.ID, &u.Email, &u.Name, &hash)
	if err != nil || !auth.CheckPasswordHash(r.Password, hash) { c.JSON(401, gin.H{"error": "Invalid"}); return }
	token, _ := auth.GenerateToken(u.ID, u.Email, h.cfg.JWTSecret)
	c.JSON(200, gin.H{"user": u, "token": token})
}
func (h *AuthHandler) Me(c *gin.Context) {
	uid := c.GetString("user_id")
	var u models.User
	h.db.QueryRow("SELECT id,email,name FROM users WHERE id=$1", uid).Scan(&u.ID, &u.Email, &u.Name)
	c.JSON(200, gin.H{"user": u})
}
func HealthCheck(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) }
