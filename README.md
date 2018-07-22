# Paths
You're reading `<proj-root>/README.md`.

# Key points
- This repo is built and run by `go v1.10`.
- We are explicitly NOT using [G111Module](https://github.com/golang/go/wiki/Modules) at the moment. 
- The following hyperlinks assume that you have installed binaries [godoc](https://github.com/golang/tools/tree/master/godoc), and are running 
  - godoc -http=:6060
  - You're assumed to have already read [How to write Go codes](http://localhost:6060/doc/code.html), thus learned the followings. 
    - What the env variable `$GOPATH` is, how to get and set it. 
    - How to run a "Hello World" project under `$GOPATH/src` WITHOUT package management of custom packages. 
- We are explicitly NOT using [GO15VENDOREXPERIMENT](http://localhost:6060/cmd/go/#hdr-Vendor_Directories) either at the moment. 

# Relative import paths
We'll NOT be putting this `<proj-root>` under `$GOPATH/src`, therefore will be using the convenient [Relative Import Paths](http://localhost:6060/cmd/go/#hdr-Relative_import_paths) -- mind the last sentence in the linked godoc section.

```
To avoid ambiguity, Go programs cannot use relative import paths within a work space.
```

This is slightly different from the dir structure introduced in [How to write Go codes](http://localhost:6060/doc/code.html), and more similar to [the convention of NodeJs/npm](https://docs.npmjs.com/getting-started/installing-npm-packages-locally). 

## Build and run
```
user@proj-root> go build main.go
user@proj-root> ./main
```

# Package management
We use [Gopm Registry](https://gopm.io/) to download packages into `$GOPATH` but NOT its binary as one way of package management in China. 

[More options can be found here](https://github.com/golang/go/wiki/PackageManagementTools).


# Similarities to C/C++
Some of you come from a NodeJs or Python background where compiling is not emphasized. You are RECOMMENDED to read [this GCC tutorial](http://www3.ntu.edu.sg/home/ehchua/programming/cpp/gcc_make.html), or [a slightly modified one with some explicit introduction to the binary "ld" & env variable "LD_LIBRARY_PATH"](https://app.yinxiang.com/shard/s61/nl/13267014/430d27d8-4ce5-43a7-ae48-5c64d59282eb), to get a better understanding by comparison.
