package globaldto

import "wabustock/constants/file_type_constants"

type FileDetails struct {
	FilePath string
	Size     int64
	FileType file_type_constants.FileType
}
