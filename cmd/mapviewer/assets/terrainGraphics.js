// Init
const PIXEL = 8
const SCALE = 6
const assets = {}
const tileLayer = new PIXI.RenderLayer()
const obstacleLayer = new PIXI.RenderLayer()
const app = new PIXI.Application
app.stage.scale = SCALE
app.stage.addChild(tileLayer)
app.stage.addChild(obstacleLayer)

async function init() {
    await app.init({
        background : "#1f1f1f",
        width:  window.innerWidth,
        height: window.innerHeight,
    })
    await initTiles()
    document.body.appendChild(app.canvas)
}

// Create the default terrain spritesheet
const tileRows = [
    [
        {
            id: "DEEPWATER [0]",
            variations: 4,
            weights: [7,1,1,1],
        },
        {
            id: "WATER [1]",
            variations: 4,
            weights: [7,1,1,1],
        },
    ],
    [
        {
            id: "GRASS [2]",
            variations: 7,
            weights: [5,5,5,1,1,1,1],
        },
        {
            id: "STONE [3]",
            variations: 7,
            weights: [1,1,1,1,1,1,1],
        },
    ],
    [
        {
            id: "SAND [4]",
            variations: 3,
            weights: [1,1,1],
        },
        {
            id: "SANDSTONE [5]",
            variations: 8,
            weights: [1,1,1,1,1,1,1,1],
        },
        {
            id: "DRYGRASS [6]",
            variations: 6,
            weights: [10,10,1,1,1,1],
        }
    ],
    [
        {
            id: "DARKSTONE [7]",
            variations: 6,
            weights: [1,1,1,1,1,1],
        },
        {
            id: "MEDSTONE [8]",
            variations: 7,
            weights: [1,1,1,1,1,1,1],
        },
        {
            id: "SNOW [9]",
            variations: 5,
            weights: [5,5,5,1,1],
        },
        {
            id: "ICE [10]",
            variations: 4,
            weights: [4,1,1,1],
        }
    ]
]
let tileVariations = {}

function randi() {
    return Math.floor(1000 * Math.random())
}

function force2Digits(n) {
    let str = n.toString()
    return str.padStart(2, "0")
}

function codify(value, variation = randi() % tileVariations[value]) {
    let code = force2Digits(value)
    return code + "_" + force2Digits(variation)
}

async function initTiles(){
    let tileData = {
        frames : {},
        meta : {
            image: "tiles.png",
            format: "RGBA8888",
            size: { w: 256, h: 256},
            scale: 1
        }
    }
    let id = 0
    let x = 0
    let y = 0
    
    for (let row of tileRows) { 
        for (let tile of row) {
            let variationCount = 0
            for (let i = 0; i < tile.variations; i++) {
                for (let j = 0; j < tile.weights[i]; j++) {
                    const code = force2Digits(id) + "_" + force2Digits(variationCount)

                    tileData.frames[code] = {
                        frame: {x: x * PIXEL, y: y * PIXEL, w: PIXEL, h: PIXEL},
                        sourceSize: {w: PIXEL, h: PIXEL},
                        spriteSourceSize: {x: 0, y: 0, w: PIXEL, h: PIXEL},
                    }
                    
                    variationCount++
                }
                x++
            }
            tileVariations[id] = variationCount
            id++
        }
        x = 0
        y++
    }

    const tilesTexture = await PIXI.Assets.load("assets/tiles.png")
    const tilesSheet = new PIXI.Spritesheet(tilesTexture, tileData)
    tilesSheet.parse()
    tilesSheet.textureSource.source.scaleMode = "nearest"
    assets.tiles = tilesSheet
}


// Render a whole cell (tiles + environment)
async function render(cell) {
    const tileSheet = assets.tiles
    const tileContainer = new PIXI.ParticleContainer({
        dynamicProperties: {
            position: false,
            vertex: false,
            rotation: false,
            color: false,
        }
    })

    function addTile(x, y, value) {
        const code = codify(value)
        const base = codify(value, 0)
        const X = x * PIXEL
        const Y = y * PIXEL
        
        const blend1 = new PIXI.Particle({
            texture: tileSheet.textures[base],
            x : X - 2,
            y : Y - 2,
            scaleX: 1.5,
            scaleY: 1.5,
            alpha: 0.2,
        })
        const blend2 = new PIXI.Particle({
            texture: tileSheet.textures[base],
            x : X - 1,
            y : Y - 1,
            scaleX: 1.25,
            scaleY: 1.25,
            alpha: 0.3,
        })
        const tileParticle = new PIXI.Particle({
            texture: tileSheet.textures[code],
            x : X,
            y : Y,
        })

        tileContainer.addParticle(blend1)
        tileContainer.addParticle(blend2)
        tileContainer.addParticle(tileParticle)
    }

    for (let tile of cell.tiles) {
        let x = tile.x
        let y = tile.y
        let value = tile.value

        addTile(x, y, value)
    }
    
    app.stage.addChild(tileContainer)
    tileLayer.attach(tileContainer)
}

init()