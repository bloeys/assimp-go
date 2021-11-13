package main

import (
	"github.com/bloeys/assimp-go/aig"
)

func main() {

	scene := aig.AiImportFile("obj.obj", 0)

	meshNum := scene.GetMNumMeshes()
	println("Count:", meshNum)

	meshP := scene.MMeshes()
	println(meshP.Size())
	println(meshP.Get(0).GetMNumVertices())
	println(meshP.Get(0).MVertices().Get(0).GetX())
	// meshes := meshP

	// println(meshes)
	// println("F:", (*meshes[0]).GetMNumFaces())
	// println("UV:", (*meshes[0]).GetMNumUVComponents())
	// println("Color:", (*meshes[0]).GetNumColorChannels())
	// println("Vert:", (*meshes[0]).GetMNumVertices())

	// println("===========")
	// println("Vert:", meshes[0].Swig)

}
