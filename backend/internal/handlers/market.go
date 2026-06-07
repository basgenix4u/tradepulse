package handlers
import ("math/rand"; "time"; "github.com/gin-gonic/gin"; "github.com/tradepulse/backend/internal/config"; "github.com/tradepulse/backend/internal/models"; "github.com/tradepulse/backend/internal/websocket")
type MarketService struct { hub *websocket.Hub; cfg *config.Config; prices map[string]*models.MarketData }
func NewMarketService(h *websocket.Hub, c *config.Config) *MarketService { return &MarketService{hub: h, cfg: c, prices: make(map[string]*models.MarketData)} }
func (m *MarketService) Start() { m.init(); go m.update() }
func (m *MarketService) init() {
	for _, a := range []struct{s,n string;p float64}{{"BTC","Bitcoin",67500},{"ETH","Ethereum",3850},{"SOL","Solana",165},{"AAPL","Apple",228},{"MSFT","Microsoft",415},{"NVDA","NVIDIA",875}} {
		m.prices[a.s] = &models.MarketData{Symbol: a.s, Name: a.n, Price: a.p, ChangePercent: (rand.Float64()-0.5)*4, Volume: rand.Float64()*1e9, UpdatedAt: time.Now().Unix()}
	}
}
func (m *MarketService) update() { t := time.NewTicker(2*time.Second); for range t.C { u := []*models.MarketData{}; for _, d := range m.prices { d.Price *= 1 + (rand.Float64()-0.5)*0.002; d.ChangePercent = (rand.Float64()-0.5)*4; u = append(u, d) }; m.hub.BroadcastPriceUpdate(u) } }
func (m *MarketService) GetMarketOverview(c *gin.Context) { c.JSON(200, gin.H{"data": m.prices}) }
func (m *MarketService) GetAssetDetails(c *gin.Context) { s := c.Param("symbol"); if d, ok := m.prices[s]; ok { c.JSON(200, d) } else { c.JSON(404, gin.H{"error": "not found"}) } }
