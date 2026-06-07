package auth
import ("time"; "github.com/golang-jwt/jwt/v5"; "golang.org/x/crypto/bcrypt")
type Claims struct { UserID, Email string; jwt.RegisteredClaims }
func HashPassword(p string) (string, error) { b, e := bcrypt.GenerateFromPassword([]byte(p), 12); return string(b), e }
func CheckPasswordHash(p, h string) bool { return bcrypt.CompareHashAndPassword([]byte(h), []byte(p)) == nil }
func GenerateToken(uid, email, secret string) (string, error) {
	c := &Claims{UserID: uid, Email: email, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour))}}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c); return t.SignedString([]byte(secret))
}
func ValidateToken(ts, secret string) (*Claims, error) {
	t, e := jwt.ParseWithClaims(ts, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if e != nil { return nil, e }
	if c, ok := t.Claims.(*Claims); ok && t.Valid { return c, nil }
	return nil, e
}
