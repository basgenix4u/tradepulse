package websocket
import ("encoding/json"; "net/http"; "sync"; "github.com/gorilla/websocket")
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
type Hub struct { clients map[*Client]bool; broadcast chan []byte; register, unregister chan *Client; mu sync.RWMutex }
type Client struct { hub *Hub; conn *websocket.Conn; send chan []byte }
func NewHub() *Hub { return &Hub{clients: make(map[*Client]bool), broadcast: make(chan []byte, 256), register: make(chan *Client), unregister: make(chan *Client)} }
func (h *Hub) Run() { for { select { case c := <-h.register: h.mu.Lock(); h.clients[c] = true; h.mu.Unlock(); case c := <-h.unregister: h.mu.Lock(); delete(h.clients, c); close(c.send); h.mu.Unlock(); case m := <-h.broadcast: h.mu.RLock(); for c := range h.clients { select { case c.send <- m: default: close(c.send); delete(h.clients, c) } }; h.mu.RUnlock() } } }
func (h *Hub) BroadcastPriceUpdate(d interface{}) { b, _ := json.Marshal(map[string]interface{}{"type": "price_update", "data": d}); h.broadcast <- b }
func ServeWS(h *Hub, w http.ResponseWriter, r *http.Request) { c, _ := upgrader.Upgrade(w, r, nil); client := &Client{hub: h, conn: c, send: make(chan []byte, 256)}; h.register <- client; go client.writePump(); go client.readPump() }
func (c *Client) readPump() { defer func() { c.hub.unregister <- c; c.conn.Close() }(); for { _, _, err := c.conn.ReadMessage(); if err != nil { break } } }
func (c *Client) writePump() { for m := range c.send { c.conn.WriteMessage(websocket.TextMessage, m) } }
