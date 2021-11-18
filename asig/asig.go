package asig

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ./libs -l assimp_windows_amd64 -l IrrXML_windows_amd64 -l zlib_windows_amd64

#include <stdlib.h> //Needed for C.free

#include <assimp/scene.h>

//Functions
struct aiScene* aiImportFile(const char* pFile, unsigned int pFlags);
*/
import "C"
import "unsafe"

type Node struct {
}

type Mesh struct {
	Vertices uint
}

type Material struct {
}

type Animation struct {
}

type Texture struct {
}

type Light struct {
}

type Camera struct {
}

type Metadata struct {
}

type Scene struct {
	RootNode   *Node
	Meshes     []*Mesh
	Materials  []*Material
	Animations []*Animation
	Textures   []*Texture
	Lights     []*Light
	Cameras    []*Camera
}

func ImportFile(file string, postProcessFlags uint) *Scene {

	// C_STRUCT aiScene* aiImportFile(const char* pFile, unsigned int pFlags)

	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	cs := C.aiImportFile(cstr, C.uint(postProcessFlags))
	scene := parseScene(cs)
	println("Num verts:", scene.Meshes[0].Vertices, "; Meshes:", len(scene.Meshes))

	return &Scene{}
}

func parseScene(cs *C.struct_aiScene) *Scene {

	s := &Scene{}
	s.Meshes = parseMeshes(*cs.mMeshes, uint(cs.mNumMeshes))

	return s
}

func parseMeshes(cm *C.struct_aiMesh, count uint) []*Mesh {

	meshes := make([]*Mesh, count)
	cmeshes := unsafe.Slice(&cm, count)

	for i := 0; i < int(count); i++ {

		m := &Mesh{}
		m.Vertices = uint(cmeshes[i].mNumVertices)

		meshes[i] = m
	}

	return meshes
}
