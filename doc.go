// Copyright 2014 keimoon. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package bog is a tool to archive directories or files into Go source code. 
The directories and files data then will be built into a single Go executable
file, therefore simplify deploy process.

This tool comes with 2 parts: 
  - An executable command "bog" for archiving and extracting your files.
  - A library to access those files in your code. The API is nearly 
    identical to "os" package, therefore compatible with almost all io, ioutil
    functions.

Installing

Install "bog" tool with:
  go get github.com/keimoon/bog
  go install github.com/keimoon/bog/tool/bog

Archiving

To create an archive of a directory or a file:
  bog archive /path/to/directory
  bog archive /path/to/file

Or shorter:
  bog a /path/to/directory
  bog a /path/to/file

Extracting

To extract the data of archived source file to current directory:
  bog extract directory-archive.go

Or:
  bog e directory-archive.go

Ignore files

Bog supports the use of special file called .bogignore to make the generator ignore certain files
or folders. It is nearly identical to .gitignore.

Setting package name

By default, bog will use current directory for package name in generated files. To change that, use -p switch:

  bog -p mypackage a /path/to/directory

If you set the package name, the generated file will be put in "mypackage" folder. The only exception is if 
you use "main" as package name. In this case, the generated file will be put in current folder.

How to use generated file

In the generated file, a single public variable will be exported. This variable is normally named "MyFolderArchive"
if your directory's name is "my-folder". You can find this variable at the very end of the generated file.

This variable is of type bog.Archive, with methods nearly identical with "os" package.

Open a file

To open a file, use Open method:
   f, err := MyFolderArchive.Open("path/to/my-file")

Read data from an openned file

The return value from Open method is of interface bog.File, which is identical with os.File except you cannot
write to that.

To read data, you can simply use ioutil.ReadAll, bufio or whatever methods:
   b, err := ioutil.ReadAll(f)

List files in a folder

To list all files in a folder, use ReadDir method, which is identical with ioutil.ReadDir:
   fi, err := MyFolderArchive.ReadDir("path/to/subfolder")
   fi, err := MyFolderArchive.ReadDir("") // will list files from root folder.

Development mode

Bog supports development mode, in which bog will not archive files, and read data from the real files directly.
To enable development mode, put -d switch when run bog command:

   bog -d a /path/to/directory

Beside from reading directly from real folder, there is no difference between development mode and normal mode.

Setting root folder

When using development mode, bog relies on "root folder", which is set to the path to the directory you want
to archive. If for some reason the folder is moved, or you run on another machine where the path is different, 
the root folder can be set by using SetRoot method:

   MyFolderArchive.SetRoot("/path/to/new-directory")

Please notice that SetRoot only works in development mode.

*/
package bog
