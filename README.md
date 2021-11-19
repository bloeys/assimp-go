# assimp-go

## Using assimp-go

The `release()` function is used to free underlying C resources and should be called after all processing that requires C code is done.
`release()` Will not affect the returned Go structs like `Scene` or `Mesh`. Returned Go data will remain valid.

`asig.X` functions call into C and therefore should not be used on released objects. Calling any `asig.X` function after `release()` is **undefined**.

While `asig` functions should NOT be called on a Scene (or its objects) after they have been released, methods on structs (e.g. `myScene.XYZ`, `myMesh.ABCD()`) are **safe** even after release.

## Developing assimp-go

We link against static assimp libraries that are built for each platform and added to the `asig/libs` package.
Depending on the platform we select one of them and link against it when doing `go build`.

The general steps are:

- Copy assimp includes into `asig/assimp`
- Copy `zlib.h`, `zconf.h` and `irrXML.h` into `asig/zlib` and `asig/irrxml` respectively.
- Copy static libraries into `asig/libs`
- Generate the wrappers using `swig -go -c++ -intgosize 64 asig/asig.i`
- Add `#cgo LDFLAGS: -L ./staticLibs -l zlibstatic -l IrrXML -l assimp` at the top of the 'C' import in `asig.go`

> Note: When dealing with static libraries the compiler will probably (e.g. MinGW does this) ignore `lib` suffixes and `.a`/`.lib` suffixes.
So if your lib name is `libassimp.a` you need to pass it to CGO as `-l assimp`, otherwise you will get an error about library not found.

For platform specific steps:

**Windows**:

> Note: You must compile with the same C/C++ compiler you use with Go (e.g. if you use MinGW with Go, then compile assimp with MinGW by sepcifying the correct `-G` option)
---
> Note: If you get compilation errors with things like `File too big` or `can't write 166 bytes to section` then cmake isn't detecting you are using MinGW, so add this flag `-D CMAKE_COMPILER_IS_MINGW=TRUE`

Now assuming you are using MinGW on windows:

- Clone wanted release of assimp and run `cmake CMakeLists.txt -D BUILD_SHARED_LIBS=OFF -D ASSIMP_BUILD_ZLIB=ON -D ASSIMP_BUILD_ASSIMP_TOOLS=OFF -D ASSIMP_BUILD_TESTS=OFF -G "MinGW Makefiles"` in the root folder
- Run `cmake --build . --parallel 6`
- Copy the generated `*.lib` (or `*.a`) files into `asig/lib`
