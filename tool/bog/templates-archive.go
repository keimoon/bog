package main

import (
	"github.com/keimoon/bog"
	"time"
)



var v_1413190984_templates_main_go_tmpl = bog.NewBogFile([]byte{0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x7b, 0x7b, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0xa, 0xa, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0xa, 0x9, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x69, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x62, 0x6f, 0x67, 0x22, 0xa, 0x9, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa, 0x29, 0xa, 0xa, 0x7b, 0x7b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x69, 0x66, 0x20, 0x2e, 0x49, 0x73, 0x44, 0x69, 0x72, 0x7d, 0x7d, 0xa, 0x76, 0x61, 0x72, 0x20, 0x7b, 0x7b, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x20, 0x3d, 0x20, 0x62, 0x6f, 0x67, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x67, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x28, 0x5b, 0x5d, 0x62, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x7b, 0x7b, 0x22, 0x7b, 0x22, 0x7d, 0x7d, 0x7b, 0x7b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x7d, 0x7d, 0x7b, 0x7b, 0x2e, 0x7d, 0x7d, 0x2c, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0x7b, 0x7b, 0x22, 0x7d, 0x22, 0x7d, 0x7d, 0x2c, 0x20, 0x26, 0x62, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x2c, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x7d, 0x7d, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x7d, 0x7d, 0x2c, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x28, 0x7b, 0x7b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x7d, 0x7d, 0x2c, 0x20, 0x30, 0x29, 0x2c, 0xa, 0x7d, 0x29, 0xa, 0x7b, 0x7b, 0x65, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0xa, 0x76, 0x61, 0x72, 0x20, 0x7b, 0x7b, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x20, 0x3d, 0x20, 0x62, 0x6f, 0x67, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x28, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x7d, 0x7d, 0x2c, 0x20, 0x26, 0x62, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x7b, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x2c, 0x20, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x7d, 0x7d, 0x2c, 0x20, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x7d, 0x7d, 0x2c, 0x20, 0xa, 0x9, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x28, 0x7b, 0x7b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x7d, 0x7d, 0x2c, 0x20, 0x30, 0x29, 0x2c, 0xa, 0x7d, 0x29, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0xa, 0x76, 0x61, 0x72, 0x20, 0x7b, 0x7b, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x20, 0x3d, 0x20, 0x62, 0x6f, 0x67, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x28, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x7b, 0xa, 0x9, 0x7b, 0x7b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x7d, 0x7d, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x7d, 0x7d, 0x3a, 0x20, 0x7b, 0x7b, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x2c, 0xa, 0x9, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x7d, 0x2c, 0x20, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x44, 0x65, 0x76, 0x7d, 0x7d, 0x2c, 0x20, 0x7b, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x66, 0x20, 0x22, 0x25, 0x23, 0x76, 0x22, 0x20, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x7d, 0x7d, 0x29, 0xa}, &bog.FileInfo{
	FileName:"main.go.tmpl", 
	FileSize:871, 
	FileMode:0x1a4, 
	FileModTime:time.Unix(1413188989, 0),
})



var v_1413190984_templates = bog.NewBogFolder([]bog.File{v_1413190984_templates_main_go_tmpl,}, &bog.FileInfo{
        FileName:"templates",
	FileSize:102,
        FileMode:0x800001ed,
	FileModTime:time.Unix(1413188989, 0),
})



var TemplatesArchive = bog.NewArchive(map[string]bog.File{
	"/main.go.tmpl": v_1413190984_templates_main_go_tmpl,
	"/": v_1413190984_templates,
	
}, false, "templates")
