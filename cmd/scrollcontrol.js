function scroll(delta) {
    app.stage.scale *= (100 - delta) / 100
}

document.addEventListener("wheel", (e) => scroll(e.deltaY))