package models
import "time"
type User struct { ID, Email, Name, AvatarURL string; EmailVerified bool; CreatedAt time.Time }
type MarketData struct { Symbol, Name string; Price, Change, ChangePercent, High24h, Low24h, Volume, MarketCap float64; UpdatedAt int64 }
