# assimp-go

## Developing assimp-go

`#cgo LDFLAGS: -L ./staticLibs -l zlibstatic -l IrrXML -l assimp`

We link against static assimp libraries that are built for each platform and added to the `aig/lib` package.
Depending on the platform we select one of them and link against it when doing `go build`.

**Windows**:

> Note: You must compile with the same C/C++ compiler you use with Go (e.g. if you use MinGW with Go, then compile assimp with MinGW by sepcifying the correct `-G` option)

> Note: If you get compilation errors with things like `File too big` or `can't write 166 bytes to section` then cmake isn't detecting you are using MinGW, so add this flag `-D CMAKE_COMPILER_IS_MINGW=TRUE`

Now assuming you are using MinGW on windows:

- Clone wanted release of assimp and run `cmake CMakeLists.txt -D BUILD_SHARED_LIBS=OFF -D ASSIMP_BUILD_ZLIB=ON -D ASSIMP_BUILD_ASSIMP_TOOLS=OFF -D ASSIMP_BUILD_TESTS=OFF -G "MinGW Makefiles"` in the root folder
- Run `cmake --build . --parallel 6`
- Copy the generated `*.lib` file into `aig/lib`
