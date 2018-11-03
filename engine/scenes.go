package engine

import "errors"

var (
	scenes      map[string]Renderer
	scene       Renderer
	changeScene chan string

	// ErrSceneUndefined is returned if tryign to use a scene
	// that does not exist
	ErrSceneUndefined = errors.New("scene not defined")
)

func init() {
	scenes = make(map[string]Renderer)
	changeScene = make(chan string, 1)
}

// AddScene registers a scene with the engine
func AddScene(name string, r Renderer) {
	scenes[name] = r
}

// SetScene sets the current scene the engine is rendering
func SetScene(name string) error {
	if _, ok := scenes[name]; ok {
		changeScene <- name
		return nil
	}

	return ErrSceneUndefined
}
