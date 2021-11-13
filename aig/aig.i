%module aig

//NOTE: Add this above the 'C' package in aig_wrap.go `#cgo LDFLAGS: -L ./staticLibs -l zlibstatic -l IrrXML -l assimp` after generating

// SWIG helpers for std::string and std::vector wrapping.
%include <std_string.i>
%include <std_vector.i>

//Needed defines
#define AI_NO_EXCEPT noexcept
#define C_STRUCT struct
#define AI_FORCE_INLINE inline
#define C_ENUM enum
#define ASSIMP_API
#define PACK_STRUCT

//Macros
%define ASSIMP_ARRAY(CLASS, TYPE, NAME, LENGTH)
%newobject CLASS::NAME;
%extend CLASS {
  std::vector<TYPE > *NAME() const {
    std::vector<TYPE > *result = new std::vector<TYPE >;
    result->reserve(LENGTH);

    for (unsigned int i = 0; i < LENGTH; ++i) {
      result->push_back($self->NAME[i]);
    }

    return result;
  }
}
%ignore CLASS::NAME;
%enddef

%define ASSIMP_POINTER_ARRAY(CLASS, TYPE, NAME, LENGTH)
%newobject CLASS::NAME;
%extend CLASS {
  std::vector<TYPE *> *NAME() const {
    std::vector<TYPE *> *result = new std::vector<TYPE *>;
    result->reserve(LENGTH);

    TYPE *currentValue = $self->NAME;
    TYPE *valueLimit = $self->NAME + LENGTH;
    while (currentValue < valueLimit) {
      result->push_back(currentValue);
      ++currentValue;
    }

    return result;
  }
}
%ignore CLASS::NAME;
%enddef

%define ASSIMP_POINTER_ARRAY_ARRAY(CLASS, TYPE, NAME, OUTER_LENGTH, INNER_LENGTH)
%newobject CLASS::NAME;
%extend CLASS {
  std::vector<std::vector<TYPE *> > *NAME() const {
    std::vector<std::vector<TYPE *> > *result = new std::vector<std::vector<TYPE *> >;
    result->reserve(OUTER_LENGTH);

    for (unsigned int i = 0; i < OUTER_LENGTH; ++i) {
      std::vector<TYPE *> currentElements;

      if ($self->NAME[i] != 0) {
        currentElements.reserve(INNER_LENGTH);

        TYPE *currentValue = $self->NAME[i];
        TYPE *valueLimit = $self->NAME[i] + INNER_LENGTH;
        while (currentValue < valueLimit) {
          currentElements.push_back(currentValue);
          ++currentValue;
        }
      }

      result->push_back(currentElements);
    }

    return result;
  }
}
%ignore CLASS::NAME;
%enddef

//We need these otherwise swig won't generate interfaces for these types correctly
//because swig gets confused when there is a typedef and a templated class, so we put the typedefs here
//and a template directive at the end
typedef float ai_real;

typedef aiVector3t<ai_real> aiVector3D;
typedef aiVector2t<ai_real> aiVector2D;
typedef aiMatrix3x3t<ai_real> aiMatrix3x3;
typedef aiMatrix4x4t<ai_real> aiMatrix4x4;

%{
  #include "assimp/cimport.h"
  #include "assimp/scene.h"
  #include "assimp/mesh.h"
  #include "assimp/vector2.h"
  #include "assimp/vector3.h"
  #include "assimp/matrix3x3.h"
  #include "assimp/matrix4x4.h"
  #include "assimp/Defines.h"
  #include "assimp/color4.h"
  #include "assimp/postprocess.h"
  #include "assimp/types.h"
  #include "assimp/texture.h"
  #include "assimp/light.h"
  #include "assimp/camera.h"
  #include "assimp/material.h"
  #include "assimp/anim.h"
  #include "assimp/metadata.h"
  
  #include "zlib/zconf.h"
  #include "zlib/zlib.h"

  #include "irrxml/irrXML.h"

%}

//Features
%feature("d:stripprefix", "aiProcess_") aiPostProcessSteps;

//Ignores
%ignore aiString::Set(const std::string& pString);

//aiScene macros
ASSIMP_ARRAY(aiScene, aiAnimation*, mAnimations, $self->mNumAnimations);
ASSIMP_ARRAY(aiScene, aiCamera*, mCameras, $self->mNumCameras);
ASSIMP_ARRAY(aiScene, aiLight*, mLights, $self->mNumLights);
ASSIMP_ARRAY(aiScene, aiMaterial*, mMaterials, $self->mNumMaterials);
ASSIMP_ARRAY(aiScene, aiMesh*, mMeshes, $self->mNumMeshes);
ASSIMP_ARRAY(aiScene, aiTexture*, mTextures, $self->mNumTextures);

ASSIMP_ARRAY(aiNode, aiNode*, mChildren, $self->mNumChildren);
ASSIMP_ARRAY(aiNode, unsigned int, mMeshes, $self->mNumMeshes);

//aiMesh macros
ASSIMP_ARRAY(aiFace, unsigned int, mIndices, $self->mNumIndices);

ASSIMP_POINTER_ARRAY(aiBone, aiVertexWeight, mWeights, $self->mNumWeights);

ASSIMP_POINTER_ARRAY(aiAnimMesh, aiVector3D, mVertices, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiAnimMesh, aiVector3D, mNormals, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiAnimMesh, aiVector3D, mTangents, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiAnimMesh, aiVector3D, mBitangents, $self->mNumVertices);
ASSIMP_POINTER_ARRAY_ARRAY(aiAnimMesh, aiVector3D, mTextureCoords, AI_MAX_NUMBER_OF_TEXTURECOORDS, $self->mNumVertices);
ASSIMP_POINTER_ARRAY_ARRAY(aiAnimMesh, aiColor4D, mColors, AI_MAX_NUMBER_OF_COLOR_SETS, $self->mNumVertices);

ASSIMP_ARRAY(aiMesh, aiAnimMesh*, mAnimMeshes, $self->mNumAnimMeshes);
ASSIMP_ARRAY(aiMesh, aiBone*, mBones, $self->mNumBones);
ASSIMP_ARRAY(aiMesh, unsigned int, mNumUVComponents, AI_MAX_NUMBER_OF_TEXTURECOORDS);
ASSIMP_POINTER_ARRAY(aiMesh, aiVector3D, mVertices, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiMesh, aiVector3D, mNormals, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiMesh, aiVector3D, mTangents, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiMesh, aiVector3D, mBitangents, $self->mNumVertices);
ASSIMP_POINTER_ARRAY(aiMesh, aiFace, mFaces, $self->mNumFaces);
ASSIMP_POINTER_ARRAY_ARRAY(aiMesh, aiVector3D, mTextureCoords, AI_MAX_NUMBER_OF_TEXTURECOORDS, $self->mNumVertices);
ASSIMP_POINTER_ARRAY_ARRAY(aiMesh, aiColor4D, mColors, AI_MAX_NUMBER_OF_COLOR_SETS, $self->mNumVertices);

//Camera macros
ASSIMP_ARRAY(aiMaterial, aiMaterialProperty*, mProperties, $self->mNumProperties)

//Material settings
%include <typemaps.i>
%apply enum SWIGTYPE *OUTPUT { aiTextureMapping* mapping };
%apply unsigned int *OUTPUT { unsigned int* uvindex };
%apply float *OUTPUT { float* blend };
%apply enum SWIGTYPE *OUTPUT { aiTextureOp* op };
%apply unsigned int *OUTPUT { unsigned int* flags };

//Final includes
%include "assimp/cimport.h" // Plain-C interface
%include "assimp/scene.h"   // Output data structure
%include "assimp/mesh.h"
%include "assimp/vector2.h"
%include "assimp/vector3.h"
%include "assimp/matrix3x3.h"
%include "assimp/matrix4x4.h"
%include "assimp/Defines.h"
%include "assimp/color4.h"
%include "assimp/types.h"
%include "assimp/texture.h"
%include "assimp/light.h"
%include "assimp/camera.h"
%include "assimp/material.h"
%include "assimp/anim.h"
%include "assimp/metadata.h"
%include "assimp/postprocess.h"

%include "zlib/zconf.h"
%include "zlib/zlib.h"

%include "irrxml/irrXML.h"

// We have to "instantiate" the templates used by the ASSSIMP_*_ARRAY macros
// here at the end to avoid running into forward reference issues (SWIG would
// spit out the helper functions before the header includes for the element
// types otherwise).

%template(UintVector) std::vector<unsigned int>;
%template(aiAnimationVector) std::vector<aiAnimation *>;
%template(aiAnimMeshVector) std::vector<aiAnimMesh *>;
%template(aiBonesVector) std::vector<aiBone *>;
%template(aiCameraVector) std::vector<aiCamera *>;
%template(aiColor4DVector) std::vector<aiColor4D *>;
%template(aiColor4DVectorVector) std::vector<std::vector<aiColor4D *> >;
%template(aiFaceVector) std::vector<aiFace *>;
%template(aiLightVector) std::vector<aiLight *>;
%template(aiMaterialVector) std::vector<aiMaterial *>;
%template(aiMaterialPropertyVector) std::vector<aiMaterialProperty *>;
%template(aiMeshAnimVector) std::vector<aiMeshAnim *>;
%template(aiMeshVector) std::vector<aiMesh *>;
%template(aiNodeVector) std::vector<aiNode *>;
%template(aiNodeAnimVector) std::vector<aiNodeAnim *>;
%template(aiTextureVector) std::vector<aiTexture *>;
%template(aiVector3DVector) std::vector<aiVector3D *>;
%template(aiVector3DVectorVector) std::vector<std::vector<aiVector3D *> >;
%template(aiVertexWeightVector) std::vector<aiVertexWeight *>;
%template(GetInteger) aiMaterial::Get<int>;
%template(GetFloat) aiMaterial::Get<float>;
%template(GetColor4D) aiMaterial::Get<aiColor4D>;
%template(GetColor3D) aiMaterial::Get<aiColor3D>;
%template(GetString) aiMaterial::Get<aiString>;
%template(aiVector2D) aiVector2t<ai_real>;
%template(aiVector3D) aiVector3t<ai_real>;
%template(aiMatrix3x3) aiMatrix3x3t<ai_real>;
%template(aiMatrix4x4) aiMatrix4x4t<ai_real>;

//Material settings
%clear unsigned int* flags;
%clear aiTextureOp* op;
%clear float *blend;
%clear unsigned int* uvindex;
%clear aiTextureMapping* mapping;

%apply int &OUTPUT { int &pOut };
%apply float &OUTPUT { float &pOut };

%clear int &pOut;
%clear float &pOut;