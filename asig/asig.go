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
import (
	"unsafe"

	"github.com/bloeys/gglm/gglm"
)

type Node struct {
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

	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	cs := C.aiImportFile(cstr, C.uint(postProcessFlags))
	scene := parseScene(cs)

	return scene
}

func parseScene(cs *C.struct_aiScene) *Scene {

	s := &Scene{}
	s.Meshes = parseMeshes(cs.mMeshes, uint(cs.mNumMeshes))

	return s
}

func parseMeshes(cm **C.struct_aiMesh, count uint) []*Mesh {

	if cm == nil {
		return []*Mesh{}
	}

	meshes := make([]*Mesh, count)
	cmeshes := unsafe.Slice(cm, count)

	for i := 0; i < int(count); i++ {

		m := &Mesh{}

		cmesh := cmeshes[i]
		vertCount := uint(cmesh.mNumVertices)

		m.Vertices = parseVec3s(cmesh.mVertices, vertCount)
		m.Normals = parseVec3s(cmesh.mNormals, vertCount)
		m.Tangents = parseVec3s(cmesh.mTangents, vertCount)
		m.BitTangents = parseVec3s(cmesh.mBitangents, vertCount)

		//Color sets
		m.ColorSets = parseColorSet(cmesh.mColors, vertCount)

		//Tex coords
		m.TexCoords = parseTexCoords(cmesh.mTextureCoords, vertCount)
		m.TexCoordChannelCount = [8]uint{}
		for j := 0; j < len(cmesh.mTextureCoords); j++ {

			//If a color set isn't available then it is nil
			if cmesh.mTextureCoords[j] == nil {
				continue
			}

			m.TexCoordChannelCount[j] = uint(cmeshes[j].mNumUVComponents[j])
		}

		//Faces
		cFaces := unsafe.Slice(cmesh.mFaces, cmesh.mNumFaces)
		m.Faces = make([]Face, cmesh.mNumFaces)
		for j := 0; j < len(m.Faces); j++ {

			m.Faces[j] = Face{
				Indices: parseUInts(cFaces[j].mIndices, uint(cFaces[j].mNumIndices)),
			}
		}

		//Other
		m.Bones = parseBones(cmesh.mBones, uint(cmesh.mNumBones))
		m.AnimMeshes = parseAnimMeshes(cmesh.mAnimMeshes, uint(cmesh.mNumAnimMeshes))
		m.AABB = AABB{
			Min: parseVec3(&cmesh.mAABB.mMin),
			Max: parseVec3(&cmesh.mAABB.mMax),
		}

		m.MorphMethod = MorphMethod(cmesh.mMethod)
		m.MaterialIndex = uint(cmesh.mMaterialIndex)
		m.Name = parseAiString(cmesh.mName)

		meshes[i] = m
	}

	return meshes
}

func parseVec3(cv *C.struct_aiVector3D) gglm.Vec3 {

	if cv == nil {
		return gglm.Vec3{}
	}

	return gglm.Vec3{
		Data: [3]float32{
			float32(cv.x),
			float32(cv.y),
			float32(cv.z),
		},
	}
}

func parseAnimMeshes(cam **C.struct_aiAnimMesh, count uint) []*AnimMesh {

	if cam == nil {
		return []*AnimMesh{}
	}

	animMeshes := make([]*AnimMesh, count)
	cAnimMeshes := unsafe.Slice(cam, count)

	for i := 0; i < int(count); i++ {

		m := cAnimMeshes[i]
		animMeshes[i] = &AnimMesh{
			Name:        parseAiString(m.mName),
			Vertices:    parseVec3s(m.mVertices, uint(m.mNumVertices)),
			Normals:     parseVec3s(m.mNormals, uint(m.mNumVertices)),
			Tangents:    parseVec3s(m.mTangents, uint(m.mNumVertices)),
			BitTangents: parseVec3s(m.mBitangents, uint(m.mNumVertices)),
			Colors:      parseColorSet(m.mColors, uint(m.mNumVertices)),
			TexCoords:   parseTexCoords(m.mTextureCoords, uint(m.mNumVertices)),
			Weight:      float32(m.mWeight),
		}
	}

	return animMeshes
}

func parseTexCoords(ctc [MaxTexCoords]*C.struct_aiVector3D, vertCount uint) [MaxTexCoords][]gglm.Vec3 {

	texCoords := [MaxTexCoords][]gglm.Vec3{}

	for j := 0; j < len(ctc); j++ {

		//If a color set isn't available then it is nil
		if ctc[j] == nil {
			continue
		}

		texCoords[j] = parseVec3s(ctc[j], vertCount)
	}

	return texCoords
}

func parseColorSet(cc [MaxColorSets]*C.struct_aiColor4D, vertCount uint) [MaxColorSets][]gglm.Vec4 {

	colorSet := [MaxColorSets][]gglm.Vec4{}
	for j := 0; j < len(cc); j++ {

		//If a color set isn't available then it is nil
		if cc[j] == nil {
			continue
		}

		colorSet[j] = parseColors(cc[j], vertCount)
	}

	return colorSet
}

func parseBones(cbs **C.struct_aiBone, count uint) []*Bone {

	if cbs == nil {
		return []*Bone{}
	}

	bones := make([]*Bone, count)
	cbones := unsafe.Slice(cbs, count)

	for i := 0; i < int(count); i++ {

		cBone := cbones[i]
		bones[i] = &Bone{
			Name:         parseAiString(cBone.mName),
			Weights:      parseVertexWeights(cBone.mWeights, uint(cBone.mNumWeights)),
			OffsetMatrix: parseMat4(&cBone.mOffsetMatrix),
		}
	}

	return bones
}

func parseMat4(cm4 *C.struct_aiMatrix4x4) gglm.Mat4 {

	if cm4 == nil {
		return gglm.Mat4{}
	}

	return gglm.Mat4{
		Data: [4][4]float32{
			{float32(cm4.a1), float32(cm4.b1), float32(cm4.c1), float32(cm4.d1)},
			{float32(cm4.a2), float32(cm4.b2), float32(cm4.c2), float32(cm4.d2)},
			{float32(cm4.a3), float32(cm4.b3), float32(cm4.c3), float32(cm4.d3)},
			{float32(cm4.a4), float32(cm4.b4), float32(cm4.c4), float32(cm4.d4)},
		},
	}
}

func parseVertexWeights(cWeights *C.struct_aiVertexWeight, count uint) []VertexWeight {

	if cWeights == nil {
		return []VertexWeight{}
	}

	vw := make([]VertexWeight, count)
	cvw := unsafe.Slice(cWeights, count)

	for i := 0; i < int(count); i++ {

		vw[i] = VertexWeight{
			VertIndex: uint(cvw[i].mVertexId),
			Weight:    float32(cvw[i].mWeight),
		}
	}

	return vw
}

func parseAiString(aiString C.struct_aiString) string {
	return C.GoStringN(&aiString.data[0], C.int(aiString.length))
}

func parseUInts(cui *C.uint, count uint) []uint {

	if cui == nil {
		return []uint{}
	}

	uints := make([]uint, count)
	cUInts := unsafe.Slice(cui, count)
	for i := 0; i < len(cUInts); i++ {
		uints[i] = uint(cUInts[i])
	}

	return uints
}

func parseVec3s(cv *C.struct_aiVector3D, count uint) []gglm.Vec3 {

	if cv == nil {
		return []gglm.Vec3{}
	}

	carr := unsafe.Slice(cv, count)
	verts := make([]gglm.Vec3, count)

	for i := 0; i < int(count); i++ {
		verts[i] = gglm.Vec3{
			Data: [3]float32{
				float32(carr[i].x),
				float32(carr[i].y),
				float32(carr[i].z),
			},
		}
	}

	return verts
}

func parseColors(cv *C.struct_aiColor4D, count uint) []gglm.Vec4 {

	if cv == nil {
		return []gglm.Vec4{}
	}

	carr := unsafe.Slice(cv, count)
	verts := make([]gglm.Vec4, count)

	for i := 0; i < int(count); i++ {
		verts[i] = gglm.Vec4{
			Data: [4]float32{
				float32(carr[i].r),
				float32(carr[i].g),
				float32(carr[i].b),
				float32(carr[i].a),
			},
		}
	}

	return verts
}
