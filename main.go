package main

import (
	"fmt"

	"github.com/bloeys/assimp-go/aig"
)

func main() {

	scene := aig.AiImportFile("obj.obj", uint(aig.AiProcess_OptimizeMeshes))
	meshes := scene.MMeshes()

	verts := meshes.Get(0).MVertices()
	for i := 0; i < int(verts.Size()); i++ {
		v := verts.Get(i)
		fmt.Printf("V%v: (%v, %v, %v)\n", i, v.GetX(), v.GetY(), v.GetZ())
	}

	scene = aig.AiImportFile("obj.fbx", uint(aig.AiProcess_OptimizeMeshes))
	meshes = scene.MMeshes()

	verts = meshes.Get(0).MVertices()
	for i := 0; i < int(verts.Size()); i++ {
		v := verts.Get(i)
		fmt.Printf("V%v: (%v, %v, %v)\n", i, v.GetX(), v.GetY(), v.GetZ())
	}
}
