package engine

var (
	scenes map[string]Renderer
	scene  Renderer
)

func init() {
	scenes = make(map[string]Renderer)
}

// AddScene registers a scene with the engine
func AddScene(name string, r Renderer) {
	scenes[name] = r
}

// SetScene sets the current scene the engine is rendering
func SetScene(name string) {
	scene = scenes[name]
}
