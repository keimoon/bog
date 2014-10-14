# Bog

[![GoDoc](https://godoc.org/github.com/keimoon/bog?status.png)](http://godoc.org/github.com/keimoon/bog)

_Bog_ is a tool to archive directories or files into Go source code.
The directories and files data then will be built into a single Go executable
file, therefore simplify deploy process.

This tool comes with 2 parts:

* An executable command _bog_ for archiving and extracting your files.
* A library to access those files in your code. The API is nearly
  identical to _os_ package, therefore compatible with almost all io, ioutil
  functions.

## Installing

```
go get github.com/keimoon/bog
go install github.com/keimoon/bog/tool/bog
```

## Archiving

To create an archive of a directory or a file:

```
bog archive /path/to/directory
bog archive /path/to/file
```

Or shorter:

```
bog a /path/to/directory
bog a /path/to/file
```

## Extracting

To extract the data of archived source file to current directory:

```
bog extract directory-archive.go
```

Or:

```
bog e directory-archive.go
```

## Setting package name

By default, bog will use current directory for package name in generated files. To change that, use _-p switch_:

```
bog -p mypackage a /path/to/directory
```

If you set the package name, the generated file will be put in _mypackage_ folder. The only exception is if you use _main_ as package name. In this case, the generated file will be put in current folder.

## How to use generated file

In the generated file, a single public variable will be exported. This variable is normally named _MyFolderArchive_ if your directory's name is _my-folder_. You can find this variable at the very end of the generated file.

This variable is of type _bog.Archive_, with methods nearly identical with _os_ package.

## Open a file

To open a file, use _Open_ method:

```
f, err := MyFolderArchive.Open("path/to/my-file")
```

## Read data from an openned file

The return value from _Open_ method is of interface _bog.File_, which is identical with _os.File_ except you cannot write to that.

To read data, you can simply use _ioutil.ReadAll_, _bufio_ or whatever methods:

```
b, err := ioutil.ReadAll(f)
```

## List files in a folder

To list all files in a folder, use _ReadDir_ method, which is identical with _ioutil.ReadDir_:

```
fi, err := MyFolderArchive.ReadDir("path/to/subfolder")
fi, err := MyFolderArchive.ReadDir("") // will list files from root folder.
```

## Development mode

_Bog_ supports development mode, in which _bog_ will not archive files, and read data from real files directly. To enable development mode, put _-d switch_ when run _bog_ command:

```
bog -d a /path/to/directory
```

Beside from reading directly from real folder, there is no difference between development mode and normal mode.

## Setting root folder

When using development mode, bog relies on _root folder_, which is set to the path to the directory you want to archive. If for some reason the folder is moved, or you run on another machine where the path is different, the root folder can be set by using _SetRoot_ method:

```
MyFolderArchive.SetRoot("/path/to/new-directory")
```

Please notice that _SetRoot_ only works in development mode.

## Contributing

Contributions and pull requests are always welcome.

## License

Bog is available under [BSD License](http://opensource.org/licenses/BSD-3-Clause)
