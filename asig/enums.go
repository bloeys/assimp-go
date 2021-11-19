package asig

type PrimitiveType int32

//Specifies types of primitives that can be present in a mesh
const (
	PrimitiveTypePoint    = 1 << 0
	PrimitiveTypeLine     = 1 << 1
	PrimitiveTypeTriangle = 1 << 2
	PrimitiveTypePolygon  = 1 << 3
)

type MorphMethod int32

//Supported methods of mesh morphing
const (
	//Interpolation between morph targets
	MorphMethodVertexBlend = 0x1
	//Normalized morphing between morph targets
	MorphMethodMorphNormalized = 0x2
	//Relative morphing between morph targets
	MorphMethodMorphRelative = 0x3
)

//PostProcess defines the flags for all possible post processing steps.
type PostProcess int64

const (
	PostProcessCalcTangentSpace         PostProcess = 0x1
	PostProcessJoinIdenticalVertices    PostProcess = 0x2
	PostProcessMakeLeftHanded           PostProcess = 0x4
	PostProcessTriangulate              PostProcess = 0x8
	PostProcessRemoveComponent          PostProcess = 0x10
	PostProcessGenNormals               PostProcess = 0x20
	PostProcessGenSmoothNormals         PostProcess = 0x40
	PostProcessSplitLargeMeshes         PostProcess = 0x80
	PostProcessPreTransformVertices     PostProcess = 0x100
	PostProcessLimitBoneWeights         PostProcess = 0x200
	PostProcessValidateDataStructure    PostProcess = 0x400
	PostProcessImproveCacheLocality     PostProcess = 0x800
	PostProcessRemoveRedundantMaterials PostProcess = 0x1000
	PostProcessFixInfacingNormals       PostProcess = 0x2000
	PostProcessSortByPType              PostProcess = 0x8000
	PostProcessFindDegenerates          PostProcess = 0x10000
	PostProcessFindInvalidData          PostProcess = 0x20000
	PostProcessGenUVCoords              PostProcess = 0x40000
	PostProcessTransformUVCoords        PostProcess = 0x80000
	PostProcessFindInstances            PostProcess = 0x100000
	PostProcessOptimizeMeshes           PostProcess = 0x200000
	PostProcessOptimizeGraph            PostProcess = 0x400000
	PostProcessFlipUVs                  PostProcess = 0x800000
	PostProcessFlipWindingOrder         PostProcess = 0x1000000
	PostProcessSplitByBoneCount         PostProcess = 0x2000000
	PostProcessDebone                   PostProcess = 0x4000000
	PostProcessGlobalScale              PostProcess = 0x8000000
	PostProcessEmbedTextures            PostProcess = 0x10000000
	PostProcessForceGenNormals          PostProcess = 0x20000000
	PostProcessDropNormals              PostProcess = 0x40000000
	PostProcessGenBoundingBoxes         PostProcess = 0x80000000
)
