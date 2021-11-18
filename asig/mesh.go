package asig

import "github.com/bloeys/gglm/gglm"

const (
	MaxColorSets = 8
	MaxTexCoords = 8
)

type Mesh struct {

	//Bitwise combination of PrimitiveType enum
	PrimitiveTypes PrimitiveType
	Vertices       []gglm.Vec3
	Normals        []gglm.Vec3
	Tangents       []gglm.Vec3
	BitTangents    []gglm.Vec3

	//ColorSets vertex color sets where each set is either empty or has length=len(Vertices), with max number of sets=MaxColorSets
	ColorSets [MaxColorSets][]gglm.Vec4

	//TexCoords (aka UV channels) where each TexCoords[i] has NumUVComponents[i] channels, and is either empty or has length=len(Vertices), with max number of TexCoords per vertex = MaxTexCoords
	TexCoords            [MaxTexCoords][]gglm.Vec3
	TexCoordChannelCount [MaxTexCoords]uint

	Faces       []Face
	Bones       []*Bone
	AnimMeshes  []*AnimMesh
	AABB        AABB
	MorphMethod MorphMethod

	MaterialIndex uint
	Name          string
}

type Face struct {
	Indices []uint
}

type AnimMesh struct {
	Name string

	/** Replacement for Mes.Vertices. If this array is non-NULL,
	 *  it *must* contain mNumVertices entries. The corresponding
	 *  array in the host mesh must be non-NULL as well - animation
	 *  meshes may neither add or nor remove vertex components (if
	 *  a replacement array is NULL and the corresponding source
	 *  array is not, the source data is taken instead)*/
	Vertices    []gglm.Vec3
	Normals     []gglm.Vec3
	Tangents    []gglm.Vec3
	BitTangents []gglm.Vec3
	Colors      [MaxColorSets][]gglm.Vec4
	TexCoords   [MaxTexCoords][]gglm.Vec3

	Weight float32
}

type AABB struct {
	Min gglm.Vec3
	Max gglm.Vec3
}

type Bone struct {
	Name string
	//The influence weights of this bone
	Weights []VertexWeight

	/** Matrix that transforms from bone space to mesh space in bind pose.
	 *
	 * This matrix describes the position of the mesh
	 * in the local space of this bone when the skeleton was bound.
	 * Thus it can be used directly to determine a desired vertex position,
	 * given the world-space transform of the bone when animated,
	 * and the position of the vertex in mesh space.
	 *
	 * It is sometimes called an inverse-bind matrix,
	 * or inverse bind pose matrix.
	 */
	OffsetMatrix gglm.Mat4
}

type VertexWeight struct {
	VertIndex uint
	//The strength of the influence in the range (0...1). The total influence from all bones at one vertex is 1
	Weight float32
}
