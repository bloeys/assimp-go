package main

import (
	"fmt"

	"github.com/bloeys/assimp-go/asig"
)

func main() {

	scene, release, err := asig.ImportFile("tex-cube.fbx", asig.PostProcessTriangulate)
	defer release()

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(scene.Meshes); i++ {

		println("Mesh:", i, "; Verts:", len(scene.Meshes[i].Vertices), "; Normals:", len(scene.Meshes[i].Normals))
		for j := 0; j < len(scene.Meshes[i].Vertices); j++ {
			fmt.Printf("V(%v): (%v, %v, %v)\n", j, scene.Meshes[i].Vertices[j].X(), scene.Meshes[i].Vertices[j].Y(), scene.Meshes[i].Vertices[j].Z())
		}
	}

	for i := 0; i < len(scene.Materials); i++ {

		println("Mesh:", i, "; Props:", len(scene.Materials[i].Properties))
		for j := 0; j < len(scene.Materials[i].Properties); j++ {

			p := scene.Materials[i].Properties[j]
			fmt.Printf("Data Type: %v; Len Bytes: %v; Texture Type: %v\n", p.TypeInfo.String(), len(p.Data), p.Semantic.String())

			fmt.Println("Texture count:", asig.GetMaterialTextureCount(scene.Materials[i], asig.TextureTypeDiffuse))
		}
	}
}
