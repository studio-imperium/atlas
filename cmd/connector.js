const Application = PIXI.Application
const Texture = PIXI.Texture
const Assets = PIXI.Assets

const SCALE = 6
const PIXEL = 8

let app = new Application
app.stage.scale = SCALE

let tileVariations = [
    10,
    6,
    3,
    13,
    8,
    6, 
    6,
    7,
    7,
]

function randi() {
	return Math.floor(1000 * Math.random())
}

function force2Digits(n) {
	let str = n.toString()
	return str.padStart(2, "0")
}

function codify(tileInt, variation = randi() % tileVariations[tileInt]) {
	let code = force2Digits(tileInt)
	return code + "_" + force2Digits(variation)
}

async function render(cell) {
    const tileSheet = await Assets.load("tiles")
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
            X,
            Y,
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
    
    app.stage.addChild(tilesContainer)
}

async function init(scale) {
    await app.init(
    {
        background : "#1f1f1f",
        width:  window.innerWidth,
        height: window.innerHeight,
    })
    document.body.appendChild(app.canvas)

    let tileSheet = await Assets.load("assets/tiles.png")
    tileSheet.source.scaleMode = "nearest"
    Assets.add({
        alias: 'tiles',
        src: 'assets/tiles.json',
        data: { texture: tileSheet }
    })
    
    app.stage.scale = scale
}

init(scale)
