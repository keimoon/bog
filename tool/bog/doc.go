// Copyright 2014 keimoon. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Bog command is used for archiving and extracting your files.

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

Development mode

Bog supports development mode, in which bog will not archive files, and read data from the real files directly.
To enable development mode, put -d switch when run bog command:

   bog -d a /path/to/directory

Beside from reading directly from real folder, there is no difference between development mode and normal mode.
*/
package main
