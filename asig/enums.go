package asig

type aiReturn int32

const (
	//Indicates that a function was successful
	aiReturnSuccess = 0x0

	//Indicates that a function failed
	aiReturnFailure = -0x1

	//Indicates that not enough memory was available to perform the requested operation
	aiReturnOutofMemory = -0x3
)

type SceneFlag int32

const (
	/**
	 * Specifies that the scene data structure that was imported is not complete.
	 * This flag bypasses some internal validations and allows the import
	 * of animation skeletons, material libraries or camera animation paths
	 * using Assimp. Most applications won't support such data.
	 */
	SceneFlagIncomplete SceneFlag = 1 << 0

	/**
	 * This flag is set by the validation postprocess-step (PostProcessValidateDataStructure)
	 * if the validation is successful. In a validated scene you can be sure that
	 * any cross references in the data structure (e.g. vertex indices) are valid.
	 */
	SceneFlagValidated SceneFlag = 1 << 1

	/**
	 * This flag is set by the validation postprocess-step (PostProcessValidateDataStructure)
	 * if the validation is successful but some issues have been found.
	 * This can for example mean that a texture that does not exist is referenced
	 * by a material or that the bone weights for a vertex don't sum to 1.0 ... .
	 * In most cases you should still be able to use the import. This flag could
	 * be useful for applications which don't capture Assimp's log output.
	 */
	SceneFlagValidationWarning SceneFlag = 1 << 2

	/**
	 * This flag is currently only set by the PostProcessJoinIdenticalVertices step.
	 * It indicates that the vertices of the output meshes aren't in the internal
	 * verbose format anymore. In the verbose format all vertices are unique,
	 * no vertex is ever referenced by more than one face.
	 */
	SceneFlagNonVerboseFormat SceneFlag = 1 << 3

	/**
	 * Denotes pure height-map terrain data. Pure terrains usually consist of quads,
	 * sometimes triangles, in a regular grid. The x,y coordinates of all vertex
	 * positions refer to the x,y coordinates on the terrain height map, the z-axis
	 * stores the elevation at a specific point.
	 *
	 * TER (Terragen) and HMP (3D Game Studio) are height map formats.
	 * @note Assimp is probably not the best choice for loading *huge* terrains -
	 * fully triangulated data takes extremely much free store and should be avoided
	 * as long as possible (typically you'll do the triangulation when you actually
	 * need to render it).
	 */
	SceneFlagTerrain SceneFlag = 1 << 4

	/**
	 * Specifies that the scene data can be shared between structures. For example:
	 * one vertex in few faces. SceneFlagNonVerboseFormat can not be
	 * used for this because SceneFlagNonVerboseFormat has internal
	 * meaning about postprocessing steps.
	 */
	SceneFlagAllowShared SceneFlag = 1 << 5
)

//aiGetErrorString specifies the types of primitives that can be present in a mesh
type PrimitiveType int32

const (
	PrimitiveTypePoint    = 1 << 0
	PrimitiveTypeLine     = 1 << 1
	PrimitiveTypeTriangle = 1 << 2
	PrimitiveTypePolygon  = 1 << 3
)

//MorphMethod specifies the Supported methods of mesh morphing
type MorphMethod int32

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

type TextureType int32

const (

	/** Dummy value.
	 *
	 *  No texture, but the value to be used as 'texture semantic'
	 *  (MaterialProperty.Semantic) for all material properties
	 *  *not* related to textures.
	 */
	TextureTypeNone TextureType = 0

	/** The texture is combined with the result of the diffuse
	 *  lighting equation.
	 */
	TextureTypeDiffuse TextureType = 1

	/** The texture is combined with the result of the specular
	 *  lighting equation.
	 */
	TextureTypeSpecular TextureType = 2

	/** The texture is combined with the result of the ambient
	 *  lighting equation.
	 */
	TextureTypeAmbient TextureType = 3

	/** The texture is added to the result of the lighting
	 *  calculation. It isn't influenced by incoming light.
	 */
	TextureTypeEmissive TextureType = 4

	/** The texture is a height map.
	 *
	 *  By convention, higher gray-scale values stand for
	 *  higher elevations from the base height.
	 */
	TextureTypeHeight TextureType = 5

	/** The texture is a (tangent space) normal-map.
	 *
	 *  Agn, there are several conventions for tangent-space
	 *  normal maps. Assimp does (intentionally) not
	 *  distinguish here.
	 */
	TextureTypeNormal TextureType = 6

	/** The texture defines the glossiness of the material.
	 *
	 *  The glossiness is in fact the exponent of the specular
	 *  (phong) lighting equation. Usually there is a conversion
	 *  function defined to map the linear color values in the
	 *  texture to a suitable exponent. Have fun.
	 */
	TextureTypeShininess TextureType = 7

	/** The texture defines per-pixel opacity.
	 *
	 *  Usually 'white' means opaque and 'black' means
	 *  'transparency'. Or quite the opposite. Have fun.
	 */
	TextureTypeOpacity TextureType = 8

	/** Displacement texture
	 *
	 *  The exact purpose and format is application-dependent.
	 *  Higher color values stand for higher vertex displacements.
	 */
	TextureTypeDisplacement TextureType = 9

	/** Lightmap texture (aka Ambient Occlusion)
	 *
	 *  Both 'Lightmaps' and dedicated 'ambient occlusion maps' are
	 *  covered by this material property. The texture contains a
	 *  scaling value for the final color value of a pixel. Its
	 *  intensity is not affected by incoming light.
	 */
	TextureTypeLightmap TextureType = 10

	/** Reflection texture
	 *
	 * Contains the color of a perfect mirror reflection.
	 * Rarely used, almost never for real-time applications.
	 */
	TextureTypeReflection TextureType = 11

	/** Unknown texture
	 *  A texture reference that does not match any of the definitions
	 *  above is considered to be 'unknown'. It is still imported,
	 *  but is excluded from any further post-processing.
	 */
	TextureTypeUnknown TextureType = 18
)

/** PBR Materials
 * PBR definitions from maya and other modelling packages now use this standard.
 * This was originally introduced around 2012.
 * Support for this is in game engines like Godot, Unreal or Unity3D.
 * Modelling packages which use this are very common now.
 */
const (
	TextureTypeBaseColor        TextureType = 12
	TextureTypeNormalCamera     TextureType = 13
	TextureTypeEmissionColor    TextureType = 14
	TextureTypeMetalness        TextureType = 15
	TextureTypeDiffuseRoughness TextureType = 16
	TextureTypeAmbientOcclusion TextureType = 17
)

func (tp TextureType) String() string {

	switch tp {
	case TextureTypeNone:
		return "None"
	case TextureTypeDiffuse:
		return "Diffuse"
	case TextureTypeSpecular:
		return "Specular"
	case TextureTypeAmbient:
		return "Ambient"
	case TextureTypeAmbientOcclusion:
		return "AmbientOcclusion"
	case TextureTypeBaseColor:
		return "BaseColor"
	case TextureTypeDiffuseRoughness:
		return "DiffuseRoughness"
	case TextureTypeDisplacement:
		return "Displacement"
	case TextureTypeEmissionColor:
		return "EmissionColor"
	case TextureTypeEmissive:
		return "Emissive"
	case TextureTypeHeight:
		return "Height"
	case TextureTypeLightmap:
		return "Lightmap"
	case TextureTypeMetalness:
		return "Metalness"
	case TextureTypeNormal:
		return "Normal"
	case TextureTypeNormalCamera:
		return "NormalCamera"
	case TextureTypeOpacity:
		return "Opacity"
	case TextureTypeReflection:
		return "Reflection"
	case TextureTypeShininess:
		return "Shininess"
	case TextureTypeUnknown:
		return "Unknown"
	default:
		return "Invalid"
	}
}

type MatPropertyTypeInfo int32

const (
	MatPropTypeInfoFloat32 MatPropertyTypeInfo = iota + 1
	MatPropTypeInfoFloat64
	MatPropTypeInfoString
	MatPropTypeInfoInt32

	//Simple binary buffer, content undefined. Not convertible to anything.
	MatPropTypeInfoBuffer
)

func (mpti MatPropertyTypeInfo) String() string {

	switch mpti {
	case MatPropTypeInfoFloat32:
		return "Float32"
	case MatPropTypeInfoFloat64:
		return "Float64"
	case MatPropTypeInfoString:
		return "String"
	case MatPropTypeInfoInt32:
		return "Int32"
	case MatPropTypeInfoBuffer:
		return "Buffer"
	default:
		return "Unknown"
	}
}

type MetadataType int32

const (
	MetadataTypeBool    MetadataType = 0
	MetadataTypeInt32   MetadataType = 1
	MetadataTypeUint64  MetadataType = 2
	MetadataTypeFloat32 MetadataType = 3
	MetadataTypeFloat64 MetadataType = 4
	MetadataTypeString  MetadataType = 5
	MetadataTypeVec3    MetadataType = 6
	MetadataTypeMAX     MetadataType = 7
)
