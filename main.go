package main

import (
	"github.com/bloeys/assimp-go/asig"
)

func main() {

	scene := asig.ImportFile("obj.obj", 0)
	println(scene)
	// meshes := scene.MMeshes()

	// verts := meshes.Get(0).MVertices()
	// for i := 0; i < int(verts.Size()); i++ {
	// 	v := verts.Get(i)
	// 	fmt.Printf("V%v: (%v, %v, %v)\n", i, v.GetX(), v.GetY(), v.GetZ())
	// }

	// scene = asig.AiImportFile("obj.fbx", uint(0))
	// meshes = scene.MMeshes()

	// verts = meshes.Get(0).MVertices()
	// for i := 0; i < int(verts.Size()); i++ {
	// 	v := verts.Get(i)
	// 	fmt.Printf("V%v: (%v, %v, %v)\n", i, v.GetX(), v.GetY(), v.GetZ())
	// }
}
