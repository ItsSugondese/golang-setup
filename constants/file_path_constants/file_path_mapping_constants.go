package filepathconstants

import (
	"wabustock/enums/struct-enums/project_module"
)

// FilePathMapping represents the file path mapping struct
type FilePathMapping struct {
	Path     string
	Location string
}

// Define the file path mappings
var (
	TemporaryFile = FilePathMapping{
		Path:     "image" + FileSeparator + "file" + FileSeparator + "temporary" + FileSeparator,
		Location: "mahabira-tempdocument" + FileSeparator + "doc" + FileSeparator,
	}
	UserFile = FilePathMapping{
		Path:     "image" + FileSeparator + "file" + FileSeparator + "user" + FileSeparator,
		Location: "user-document" + FileSeparator + "doc" + FileSeparator,
	}
)

// FilePathMappings map for easy lookup
var FilePathMappings = map[string]FilePathMapping{
	project_module.ModuleNameEnums.TEMPORARY_ATTACHMENTS: TemporaryFile,
	project_module.ModuleNameEnums.BASE_USER:             UserFile,
}
