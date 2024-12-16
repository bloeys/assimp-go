# assimp-go

[![Build](https://github.com/bloeys/assimp-go/actions/workflows/run-assimp-go.yml/badge.svg)](https://github.com/bloeys/assimp-go/actions/workflows/run-assimp-go.yml)

A Handcrafted Open Asset Import Library (AssImp) wrapper for Go.

## Features

The following features are already implemented:

* Loading all supported model formats into a Scene object
* Mesh data
* Materials
* Textures and embedded textures
* Error reporting
* Enums relevant to the above operations

Unimplemented (yet) AssImp Scene objects:

* Animation
* Lights
* Camera

## Using assimp-go

### Requirements

To run the project you need:

* A 64-bit machine (32-bit machines should be supportable if we get help adding the needed libs)
* A recent version of [Go](https://golang.org/) installed (1.17+)
* A C/C++ compiler installed and in your path
  * **Windows**: [MingW](https://www.mingw-w64.org/downloads/#mingw-builds) or similar
  * **Mac/Linux**: Should be installed by default, but if not try [GCC](https://gcc.gnu.org/) or [Clang](https://releases.llvm.org/download.html)

Now to be able to run assimp-go you will need the AssImp shared libraries (DLLs/DyLibs), which you can
download from the GitHub releases page.

### Installing on Windows

Download the **.dll** of the [release you want](https://github.com/bloeys/assimp-go/releases), and place it in the **root** of your Go project.

### Installing on MacOS

First, download the appropriate **.dylib** for your device (`_amd64` for Intel CPUs and `_arm64` for Apple CPUs).
Next you will need to rename the lib to `libassimp.5.dylib` and move it to `/usr/local/lib` or `/usr/lib`.

You can use this command to do it: `sudo mkdir -p /usr/local/lib && sudo cp libassimp_darwin*.dylib /usr/local/lib/libassimp.5.dylib`

### Installing on Linux

Download the the AssImp package for our distro or build from source and add it to your path.
> NOTE: Insatall assimp >= 3.1 for bindings to work as expected. 
>       Though getting the latest version is always recommended. 

#### Installing on Ubuntu
You can install the Asset-Importer-Lib via apt:
```
sudo apt-get update
sudo apt-get install libassimp-dev
```

#### Installing on Arch 
You can install the Asset-Importer-Lib via pacman
```
sudo pacman -S assimp
```
#### Building From Source 
To build the Asset Importer Package from sorce read the [sorce build guide](https://github.com/assimp/assimp/blob/master/Build.md)

### Running assimp-go

Use `go run .` to run the simple example in `main.go` ;)

> Note: that it might take a while to run the first time because of downloading/compiling dependencies.

### Getting Started

```Go
func main() {

    //Load this .fbx model with the following post processing flags
    scene, release, err := asig.ImportFile("my-cube.fbx", asig.PostProcessTriangulate | asig.PostProcessJoinIdenticalVertices)
    if err != nil {
            panic(err)
    }

    for i := 0; i < len(scene.Materials); i++ {

        m := scene.Materials[i]

        //Check how many diffuse textures are attached to this material
        texCount := asig.GetMaterialTextureCount(m, asig.TextureTypeDiffuse)
        fmt.Println("Texture count:", texCount)

        //If we have at least 1 diffuse texture attached to this material, load the first diffuse texture (index 0)
        if texCount > 0 {

            texInfo, err := asig.GetMaterialTexture(m, asig.TextureTypeDiffuse, 0)
            if err != nil {
                panic(err)
            }

            fmt.Printf("%v\n", texInfo)
        }
    }

    //Now that we are done with all our `asig.XYZ` calls we can release underlying C resources. 
    //
    //NOTE: Our Go objects (like scene, scene.Materials etc) will remain intact ;), but we must NOT use asig.XYZ calls on this scene and its children anymore
    release()
}
```

The `release()` function is used to free underlying C resources and should be called after all processing that requires C code is done.
`release()` Will not affect the returned Go structs like `Scene` or `Mesh`. Returned Go data will remain valid.

`asig.X` functions call into C and therefore should not be used on released objects. Calling any `asig.X` function after `release()` is **undefined**.

While `asig` functions should NOT be called on a Scene (or its objects) after they have been released, methods on structs (e.g. `myScene.XYZ`, `myMesh.ABCD()`) are **safe** even after release.

## Developing assimp-go

We link against assimp libraries that are built for each platform and the `.a`/`.dylib` files are added to the `asig/libs` package.
At build time, the `#cgo` directive choose the appropriate libs and links against them.

The general steps are:

* Copy assimp includes into `asig/assimp`
* Copy libraries and DLL import libraries into `asig/libs`

> Note: When dealing with libraries the compiler will probably (e.g. MinGW does this) ignore `lib` prefixes and `.a`/`.lib` suffixes.
So if your lib name is `libassimp.a` you need to pass it to CGO as `-l assimp`, otherwise you will get an error about library not found.

For platform specific steps:

**Windows**:

> Note: You must compile with the same C/C++ compiler you use with Go (e.g. if you use MinGW with Go, then compile assimp with MinGW by specifying the correct `-G` option to cMake)
---
> Note: If you get compilation errors with things like `File too big` or `can't write 166 bytes to section` then cmake isn't detecting you are using MinGW, so add this flag `-D CMAKE_COMPILER_IS_MINGW=TRUE`

Now assuming you are using MinGW on windows:

* Clone wanted release of assimp and run `cmake CMakeLists.txt -D ASSIMP_BUILD_ZLIB=ON -D ASSIMP_BUILD_ASSIMP_TOOLS=OFF -G "MinGW Makefiles"` in the root folder
* Run `cmake --build . --parallel 6`
* Copy the generated `*.lib` (or `*.a`) files from the `lib` folder and into `asig/libs`, and copy the generated dll from AssImp `bin` folder into the root of `assimp-go`.
* Copy the generated `libzlibstatic.a` file from `contrib/zlib` and into the `asig/libs` folder.

**MacOS**:

* Clone wanted release of assimp and run `cmake CMakeLists.txt -D ASSIMP_BUILD_ZLIB=ON -D ASSIMP_BUILD_ASSIMP_TOOLS=OFF` in the root folder
* Run `cmake --build . --parallel 6`
* Copy the generated `*.dylib` files from the `bin` folder and into both `asig/libs` and `/usr/local/lib`
