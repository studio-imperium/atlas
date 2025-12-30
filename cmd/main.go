package main

import (
	"log"
	"fmt"
	"net/http"
	"atlas/atlas"
	"github.com/gorilla/websocket"
	"time"
)

var world *atlas.World
var upgrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
}

type Client struct {
	visited map[atlas.Point]bool
	seen map[atlas.Point]bool
	lastRequest time.Time
}

type Data struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func sendChunks(conn *websocket.Conn) {
	client := Client{
		visited: make(map[atlas.Point]bool),
		seen: make(map[atlas.Point]bool),
		lastRequest: time.Now(),
	}
	
	for {
		json := Data{}
		err := conn.ReadJSON(&json)
		
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		
		if time.Since(client.lastRequest) > 100 * time.Millisecond {
			cell := world.GetNearestCell(atlas.Point{
				json.X,
				json.Y,
			})
			
			client.lastRequest = time.Now()
			if !client.visited[cell.Origin] {
				client.visited[cell.Origin] = true
				
				for _, adj := range append(cell.GetAdjacentCells(), cell) {
					if !client.seen[adj.Origin] {
						client.seen[adj.Origin] = true
						err = conn.WriteJSON(adj)
						
						if err != nil {
							fmt.Println(err)
							conn.Close()
							return
						}
					}
				}
			}
		}
	}
}

func handler (w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	go sendChunks(conn)
}

func main() {
	world = atlas.TemplateWorld(500)
	
	fmt.Println("Listening on 8082")
	http.HandleFunc("/atlas", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
