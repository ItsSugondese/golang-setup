package temporary_attachments

import "mime/multipart"

type TemporaryAttachmentsDetailRequest struct {
	Attachments []*multipart.FileHeader
}
