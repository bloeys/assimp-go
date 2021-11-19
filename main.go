package main

import (
	"fmt"

	"github.com/bloeys/assimp-go/asig"
)

func main() {

	scene, err := asig.ImportFile("obj.obj", asig.PostProcessTriangulate)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(scene.Meshes); i++ {

		println("Mesh:", i, "; Verts:", len(scene.Meshes[i].Vertices), "; Normals:", len(scene.Meshes[i].Normals))
		for j := 0; j < len(scene.Meshes[i].Vertices); j++ {
			fmt.Printf("V(%v): (%v, %v, %v)\n", j, scene.Meshes[i].Vertices[j].X(), scene.Meshes[i].Vertices[j].Y(), scene.Meshes[i].Vertices[j].Z())
		}
	}

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
