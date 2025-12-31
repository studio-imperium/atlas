package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Imperium-Studios-LLC/atlas"
	"github.com/gorilla/websocket"
)

var world *atlas.World
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	visited     map[atlas.Point]bool
	seen        map[atlas.Point]bool
	lastRequest time.Time
}

type Data struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func sendChunks(conn *websocket.Conn) {
	client := Client{
		visited:     make(map[atlas.Point]bool),
		seen:        make(map[atlas.Point]bool),
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

		if time.Since(client.lastRequest) > 100*time.Millisecond {
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

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go sendChunks(conn)
}

// Tiles
var DEEPWATER int8 = 0
var WATER int8 = 1
var GRASS int8 = 2
var STONE int8 = 3
var SAND int8 = 4
var SANDSTONE int8 = 5
var DRYGRASS int8 = 6
var RUBBLE int8 = 7
var DARKSTONE int8 = 8
var SNOW int8 = 9
var ICE int8 = 10

var OceanMap []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
}

var IslandMap []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
	atlas.NewBiome(
		atlas.NewFill(SAND),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
}

var DesertMountainsMap []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewPattern(27, RUBBLE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, RUBBLE),
		atlas.NewBorder(ICE),
		atlas.NewSelectiveExternalBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, DARKSTONE, DARKSTONE, SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(DARKSTONE),
	),
	atlas.NewBiome(
		atlas.NewCropCircle(40, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, SAND, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
	),
	atlas.NewBiome(
		atlas.NewPattern(5, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewPattern(3.2, GRASS, STONE),
		atlas.NewBorder(STONE),
	),
	atlas.NewBiome(
		atlas.NewCropCircle(20, WATER, GRASS),
		atlas.NewSelectiveBorder(STONE, WATER),
	),
}

func main() {
	world = atlas.NewTemplateWorld(500)
	world.Infect(DesertMountainsMap, 0.7)

	fmt.Println("Listening on 8082")
	http.HandleFunc("/atlas", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
