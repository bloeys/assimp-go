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
