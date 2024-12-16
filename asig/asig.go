package asig

/*
#cgo linux CFLAGS:
#cgo windows,amd64 CFLAGS: -I .
#cgo darwin,amd64 CFLAGS: -I .
#cgo darwin,arm64 CFLAGS: -I .

#cgo linux LDFLAGS: -lassimp
#cgo windows,amd64 LDFLAGS: -L libs -l assimp_windows_amd64
#cgo darwin,amd64 LDFLAGS: -L libs -l assimp_darwin_amd64
#cgo darwin,arm64 LDFLAGS: -L libs -l assimp_darwin_arm64

#include "wrap.c"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/bloeys/gglm/gglm"
)

type Node struct {
	Name string

	//The transformation relative to the node's parent
	Transformation *gglm.Mat4

	//Parent node. NULL if this node is the root node
	Parent *Node

	//The child nodes of this node. NULL if mNumChildren is 0
	Children []*Node

	//Each entry is an index into the mesh list of the scene
	MeshIndicies []uint

	/** Metadata associated with this node or NULL if there is no metadata.
	 *  Whether any metadata is generated depends on the source file format. See the
	 * @link importer_notes @endlink page for more information on every source file
	 * format. Importers that don't document any metadata don't write any.
	 */
	Metadata map[string]Metadata
}

type Animation struct {
}

type EmbeddedTexture struct {
	cTex *C.struct_aiTexture

	/** Width of the texture, in pixels
	 *
	 * If mHeight is zero the texture is compressed in a format
	 * like JPEG. In this case mWidth specifies the size of the
	 * memory area pcData is pointing to, in bytes.
	 */
	Width uint

	/** Height of the texture, in pixels
	 *
	 * If this value is zero, pcData points to an compressed texture
	 * in any format (e.g. JPEG).
	 */
	Height uint

	/** A hint from the loader to make it easier for applications
	 *  to determine the type of embedded textures.
	 *
	 * If Height != 0 this member is show how data is packed. Hint will consist of
	 * two parts: channel order and channel bitness (count of the bits for every
	 * color channel). For simple parsing by the viewer it's better to not omit
	 * absent color channel and just use 0 for bitness. For example:
	 * 1. Image contain RGBA and 8 bit per channel, achFormatHint == "rgba8888";
	 * 2. Image contain ARGB and 8 bit per channel, achFormatHint == "argb8888";
	 * 3. Image contain RGB and 5 bit for R and B channels and 6 bit for G channel, achFormatHint == "rgba5650";
	 * 4. One color image with B channel and 1 bit for it, achFormatHint == "rgba0010";
	 * If mHeight == 0 then achFormatHint is set set to '\\0\\0\\0\\0' if the loader has no additional
	 * information about the texture file format used OR the
	 * file extension of the format without a trailing dot. If there
	 * are multiple file extensions for a format, the shortest
	 * extension is chosen (JPEG maps to 'jpg', not to 'jpeg').
	 * E.g. 'dds\\0', 'pcx\\0', 'jpg\\0'.  All characters are lower-case.
	 * The fourth character will always be '\\0'.
	 */
	FormatHint string

	/** Data of the texture.
	 * Points to an array of Width * Height (or just len=Width if Height=0, which happens when data is compressed, like if the data is a PNG).
	 * The format of the texture data is always ARGB8888.
	 */
	Data []byte

	IsCompressed bool
	Filename     string
}

type Light struct {
}

type Camera struct {
}

type Metadata struct {
	Type  MetadataType
	Value interface{}
}

type MetadataEntry struct {
	Data []byte
}

type Scene struct {
	cScene *C.struct_aiScene
	Flags  SceneFlag

	RootNode  *Node
	Meshes    []*Mesh
	Materials []*Material

	/** Helper structure to describe an embedded texture
	 *
	 * Normally textures are contained in external files but some file formats embed
	 * them directly in the model file. There are two types of embedded textures:
	 * 1. Uncompressed textures. The color data is given in an uncompressed format.
	 * 2. Compressed textures stored in a file format like png or jpg. The raw file
	 * bytes are given so the application must utilize an image decoder (e.g. DevIL) to
	 * get access to the actual color data.
	 *
	 * Embedded textures are referenced from materials using strings like "*0", "*1", etc.
	 * as the texture paths (a single asterisk character followed by the
	 * zero-based index of the texture in the aiScene::mTextures array).
	 */
	Textures []*EmbeddedTexture

	// Animations []*Animation
	// Lights     []*Light
	// Cameras    []*Camera
}

func (s *Scene) releaseCResources() {
	C.aiReleaseImport(s.cScene)
}

//
// Assimp API
//

func ImportFile(file string, postProcessFlags PostProcess) (s *Scene, release func(), err error) {

	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	cs := C.aiImportFile(cstr, C.uint(postProcessFlags))
	if cs == nil {
		return nil, func() {}, getAiErr()
	}

	s = parseScene(cs)
	return s, func() { s.releaseCResources() }, nil
}

func getAiErr() error {
	return errors.New("asig error: " + C.GoString(C.aiGetErrorString()))
}

//
// Parsers
//

func parseScene(cs *C.struct_aiScene) *Scene {

	s := &Scene{cScene: cs}
	s.Flags = SceneFlag(cs.mFlags)
	s.RootNode = parseRootNode(cs.mRootNode)
	s.Meshes = parseMeshes(cs.mMeshes, uint(cs.mNumMeshes))
	s.Materials = parseMaterials(cs.mMaterials, uint(cs.mNumMaterials))
	s.Textures = parseTextures(cs.mTextures, uint(s.cScene.mNumTextures))

	return s
}

func parseRootNode(cNodesIn *C.struct_aiNode) *Node {

	rn := &Node{
		Name:           parseAiString(cNodesIn.mName),
		Transformation: parseMat4(&cNodesIn.mTransformation),
		Parent:         nil,
		MeshIndicies:   parseUInts(cNodesIn.mMeshes, uint(cNodesIn.mNumMeshes)),
		Metadata:       parseMetadata(cNodesIn.mMetaData),
	}

	rn.Children = parseNodes(cNodesIn.mChildren, rn, uint(cNodesIn.mNumChildren))
	return rn
}

func parseNodes(cNodesIn **C.struct_aiNode, parent *Node, parentChildrenCount uint) []*Node {

	if cNodesIn == nil {
		return []*Node{}
	}

	nodes := make([]*Node, parentChildrenCount)
	cNodes := unsafe.Slice(cNodesIn, parentChildrenCount)

	for i := 0; i < len(nodes); i++ {

		n := cNodes[i]

		//Fill basic node info
		nodes[i] = &Node{
			Name:           parseAiString(n.mName),
			Transformation: parseMat4(&n.mTransformation),
			Parent:         parent,
			MeshIndicies:   parseUInts(n.mMeshes, uint(n.mNumMeshes)),
			Metadata:       parseMetadata(n.mMetaData),
		}

		//Parse node's children
		nodes[i].Children = parseNodes(n.mChildren, nodes[i], uint(n.mNumChildren))
	}

	return nodes
}

func parseMetadata(cMetaIn *C.struct_aiMetadata) map[string]Metadata {

	if cMetaIn == nil {
		return map[string]Metadata{}
	}

	meta := make(map[string]Metadata, cMetaIn.mNumProperties)
	cKeys := unsafe.Slice(cMetaIn.mKeys, cMetaIn.mNumProperties)
	cVals := unsafe.Slice(cMetaIn.mValues, cMetaIn.mNumProperties)

	for i := 0; i < int(cMetaIn.mNumProperties); i++ {

		meta[parseAiString(cKeys[i])] = parseMetadataEntry(cVals[i])
	}

	return meta
}

func parseMetadataEntry(cv C.struct_aiMetadataEntry) Metadata {

	m := Metadata{Type: MetadataType(cv.mType)}

	if cv.mData == nil {
		return m
	}

	switch m.Type {
	case MetadataTypeBool:
		m.Value = *(*bool)(cv.mData)
	case MetadataTypeFloat32:
		m.Value = *(*float32)(cv.mData)
	case MetadataTypeFloat64:
		m.Value = *(*float64)(cv.mData)
	case MetadataTypeInt32:
		m.Value = *(*int32)(cv.mData)
	case MetadataTypeUint64:
		m.Value = *(*uint64)(cv.mData)
	case MetadataTypeString:
		m.Value = parseAiString(*(*C.struct_aiString)(cv.mData))
	case MetadataTypeVec3:
		m.Value = parseVec3((*C.struct_aiVector3D)(cv.mData))
	}

	return m
}

func parseTextures(cTexIn **C.struct_aiTexture, count uint) []*EmbeddedTexture {

	if cTexIn == nil {
		return []*EmbeddedTexture{}
	}

	textures := make([]*EmbeddedTexture, count)
	cTex := unsafe.Slice(cTexIn, count)

	for i := 0; i < int(count); i++ {

		textures[i] = &EmbeddedTexture{
			cTex:         cTex[i],
			Width:        uint(cTex[i].mWidth),
			Height:       uint(cTex[i].mHeight),
			FormatHint:   C.GoString(&cTex[i].achFormatHint[0]),
			Filename:     parseAiString(cTex[i].mFilename),
			Data:         parseTexels(cTex[i].pcData, uint(cTex[i].mWidth), uint(cTex[i].mHeight)),
			IsCompressed: cTex[i].mHeight == 0,
		}
	}

	return textures
}

func parseTexels(cTexelsIn *C.struct_aiTexel, width, height uint) []byte {

	//e.g. like a png. Otherwise we have pure color data
	isCompressed := height == 0

	texelCount := width
	if !isCompressed {
		texelCount *= height
	}
	texelCount /= 4

	data := make([]byte, texelCount*4)
	cTexels := unsafe.Slice(cTexelsIn, texelCount)

	for i := 0; i < int(texelCount); i++ {

		//Order here is important as in a compressed format the order will represent arbitrary bytes, not colors.
		//In aiTexel the struct field order is {b,g,r,a}, which puts A in the high bits and leads to a format of ARGB8888, therefore it must be maintained here
		index := i * 4
		data[index] = byte(cTexels[i].b)
		data[index+1] = byte(cTexels[i].g)
		data[index+2] = byte(cTexels[i].r)
		data[index+3] = byte(cTexels[i].a)
	}

	return data
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
			OffsetMatrix: *parseMat4(&cBone.mOffsetMatrix),
		}
	}

	return bones
}

func parseMat4(cm4 *C.struct_aiMatrix4x4) *gglm.Mat4 {

	if cm4 == nil {
		return &gglm.Mat4{}
	}

	return &gglm.Mat4{
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
	if aiString.length == 0 {
		return ""
	}

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

func parseMaterials(cMatsIn **C.struct_aiMaterial, count uint) []*Material {

	mats := make([]*Material, count)
	cMats := unsafe.Slice(cMatsIn, count)

	for i := 0; i < int(count); i++ {

		mats[i] = &Material{
			cMat:             cMats[i],
			Properties:       parseMatProperties(cMats[i].mProperties, uint(cMats[i].mNumProperties)),
			AllocatedStorage: uint(cMats[i].mNumAllocated),
		}
	}

	return mats
}

func parseMatProperties(cMatPropsIn **C.struct_aiMaterialProperty, count uint) []*MaterialProperty {

	matProps := make([]*MaterialProperty, count)
	cMatProps := unsafe.Slice(cMatPropsIn, count)

	for i := 0; i < int(count); i++ {

		cmp := cMatProps[i]

		matProps[i] = &MaterialProperty{
			name:     parseAiString(cmp.mKey),
			Semantic: TextureType(cmp.mSemantic),
			Index:    uint(cmp.mIndex),
			TypeInfo: MatPropertyTypeInfo(cmp.mType),
			Data:     C.GoBytes(unsafe.Pointer(cmp.mData), C.int(cmp.mDataLength)),
		}
	}

	return matProps
}
