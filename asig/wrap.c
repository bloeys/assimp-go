#include "wrap.h"
#include <stdio.h>

extern C_STRUCT aiFile *go_aiFileOpenProc(C_STRUCT aiFileIO *io, char *name, char *flags);
C_STRUCT aiFile *cgo_aiFileOpenProc(C_STRUCT aiFileIO *io, const char *name, const char *flags)
{
	return go_aiFileOpenProc(io, (char *)name, (char *)flags);
}

extern void go_aiFileCloseProc(C_STRUCT aiFileIO *io, C_STRUCT aiFile *file);
void cgo_aiFileCloseProc(C_STRUCT aiFileIO *io, C_STRUCT aiFile *file)
{
	go_aiFileCloseProc(io, file);
}

extern size_t go_aiFileWriteProc(C_STRUCT aiFile *file, const char *buffer, size_t size, size_t count);
size_t cgo_aiFileWriteProc(C_STRUCT aiFile *file, const char *buffer, size_t size, size_t count)
{
	return go_aiFileWriteProc(file, (char *)buffer, size, count);
}

extern size_t go_aiFileReadProc(C_STRUCT aiFile *, char *, size_t, size_t);
size_t cgo_aiFileReadProc(C_STRUCT aiFile *file, char *buffer, size_t size, size_t count)
{
	return go_aiFileReadProc(file, buffer, size, count);
}

extern size_t go_aiFileTellProc(C_STRUCT aiFile *file);
size_t cgo_aiFileTellProc(C_STRUCT aiFile *file)
{
	return go_aiFileTellProc(file);
}

extern size_t go_aiFileFileSizeProc(C_STRUCT aiFile *file);
size_t cgo_aiFileFileSizeProc(C_STRUCT aiFile *file)
{
	return go_aiFileFileSizeProc(file);
}

extern void go_aiFileFlushProc(C_STRUCT aiFile *file);
void cgo_aiFileFlushProc(C_STRUCT aiFile *file)
{
	go_aiFileFlushProc(file);
}

extern C_ENUM aiReturn go_aiFileSeekProc(C_STRUCT aiFile *file, size_t offset, C_ENUM aiOrigin origin);
C_ENUM aiReturn cgo_aiFileSeekProc(C_STRUCT aiFile *file, size_t offset, C_ENUM aiOrigin origin)
{
	return go_aiFileSeekProc(file, offset, origin);
}
