#ifndef WRAP_H
#define WRAP_H

#include <assimp/cimport.h> // Plain-C interface
#include <assimp/cfileio.h> // Plain-C interface
#include <assimp/scene.h>	// Output data structure
#include <assimp/postprocess.h>
#include <stdlib.h>

C_STRUCT aiFile *cgo_aiFileOpenProc(C_STRUCT aiFileIO *io, const char *name, const char *flags);
void cgo_aiFileCloseProc(C_STRUCT aiFileIO *io, C_STRUCT aiFile *file);

size_t cgo_aiFileWriteProc(C_STRUCT aiFile *, const char *, size_t, size_t);
size_t cgo_aiFileReadProc(C_STRUCT aiFile *, char *, size_t, size_t);
size_t cgo_aiFileTellProc(C_STRUCT aiFile *);
size_t cgo_aiFileFileSizeProc(C_STRUCT aiFile *);
void cgo_aiFileFlushProc(C_STRUCT aiFile *);
C_ENUM aiReturn cgo_aiFileSeekProc(C_STRUCT aiFile *, size_t, C_ENUM aiOrigin);

#endif