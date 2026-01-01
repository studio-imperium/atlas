package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/studio-imperium/atlas"
)

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

func sendChunks(conn *websocket.Conn, world *atlas.World) {
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

func newHandler(world *atlas.World) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		go sendChunks(conn, world)
	}
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
		atlas.NewFill(WATER),
	),
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
}

var IslandMap []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(SAND),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
}

var DesertMountainsMap []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(20, WATER, GRASS),
		atlas.NewSelectiveBorder(STONE, WATER),
	),
	atlas.NewBiome(
		atlas.NewPattern(3.2, GRASS, STONE),
		atlas.NewBorder(STONE),
	),
	atlas.NewBiome(
		atlas.NewPattern(5, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, SAND, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
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
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(DARKSTONE),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, DARKSTONE, DARKSTONE, SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, RUBBLE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, RUBBLE),
		atlas.NewBorder(ICE),
		atlas.NewSelectiveExternalBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, ICE, WATER),
		atlas.NewSelectiveBorder(WATER, ICE),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
}

var Sandy []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(3.2, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, SAND, SANDSTONE, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, DRYGRASS),
	),
	atlas.NewBiome(
		atlas.NewCropCircle(40, DRYGRASS, SAND),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SAND, DRYGRASS),
		atlas.NewSelectiveBorder(SANDSTONE, SAND),
	),
}

var Snowy []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(DARKSTONE),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, RUBBLE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(40, RUBBLE, DARKSTONE, DARKSTONE, SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewBorder(SNOW),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, RUBBLE, SNOW),
		atlas.NewSelectiveBorder(SNOW, ICE),
		atlas.NewSelectiveBorder(DARKSTONE, RUBBLE),
		atlas.NewBorder(ICE),
		atlas.NewSelectiveExternalBorder(SNOW, ICE),
	),
	atlas.NewBiome(
		atlas.NewPattern(27, ICE, WATER),
		atlas.NewSelectiveBorder(WATER, ICE),
	),
	atlas.NewBiome(
		atlas.NewFill(WATER),
	),
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
}

var PatternDemo []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewPattern(5, WATER,DEEPWATER),
	),
}

var CropCircleDemo []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewCropCircle(5, WATER,DEEPWATER),
	),
}

var VoronoiDemo []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewVoronoi(100, WATER,DEEPWATER),
	),
}

// Islands
var Islands1 []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewVoronoi(100, GRASS,DEEPWATER),
		atlas.NewSelectiveExternalBorder(SAND,GRASS),
		atlas.NewSelectiveBorder(WATER,SAND),
		atlas.NewSelectiveBorder(SAND,GRASS),
	),
}


// Big islands
var BorderDemo2 []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
		atlas.NewBorder(WATER),
	),
}


// Big islands
var BorderDemo []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
		atlas.NewBorder(DEEPWATER),
	),
}

// Funky islands
var FunkyDemo []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewPattern(7, DEEPWATER,GRASS),
		atlas.NewSelectiveExternalBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
	),
}

// Final demo
var Final []atlas.Biome = []atlas.Biome{
	atlas.NewBiome(
		atlas.NewFill(GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
		atlas.NewBorder(DEEPWATER),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
		atlas.NewBorder(DEEPWATER),
	),
	atlas.NewBiome(
		atlas.NewFill(GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
		atlas.NewBorder(DEEPWATER),
	),
	atlas.NewBiome(
		atlas.NewVoronoi(10, GRASS,DEEPWATER),
		atlas.NewSelectiveExternalBorder(SAND,GRASS),
		atlas.NewSelectiveBorder(WATER,SAND),
		atlas.NewSelectiveBorder(SAND,GRASS),
	),
	atlas.NewBiome(
		atlas.NewPattern(7, DEEPWATER,GRASS),
		atlas.NewSelectiveExternalBorder(WATER,GRASS),
		atlas.NewSelectiveBorder(SAND,GRASS),
	),
	atlas.NewBiome(
		atlas.NewFill(DEEPWATER),
	),
}


func main() {
	world1 := atlas.NewWorld(200,200,21)
	world1.Infect(Sandy, 0.2)
	world2 := atlas.NewWorld(300,100,21)
	world2.Infect(Snowy, 0.5)
	world3 := atlas.NewWorld(50,1,21)
	world3.Infect(Islands1, 0.5)
	world4 := atlas.NewWorld(50,50,21)
	world4.Infect(BorderDemo, 0.5)
	world5 := atlas.NewWorld(50,50,21)
	world5.Infect(FunkyDemo, 0.5)
	world6 := atlas.NewWorld(100,100,3)
	world6.Infect(Final, 1)

	fmt.Println("Listening on 8082")
	http.HandleFunc("/world1", newHandler(world1))
	http.HandleFunc("/world2", newHandler(world2))
	http.HandleFunc("/world3", newHandler(world3))
	http.HandleFunc("/world4", newHandler(world4))
	http.HandleFunc("/world5", newHandler(world5))
	http.HandleFunc("/world6", newHandler(world6))
	log.Fatal(http.ListenAndServe(":8082", nil))
}
