package main

import (
	"github.com/darkliquid/mindthegap/engine"
	"github.com/darkliquid/mindthegap/scenes"
)

func main() {
	err := engine.Init()
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	engine.AddScene("title", &scenes.Title{})
	engine.AddScene("intro", &scenes.Intro{})
	engine.SetScene("title")
	engine.Loop()
}
