package web

import (
	"path"
	"strings"
)

type handlerTemplateMarshal struct {
	PageHandler PageHandler
	RequestPath string
}

func (htm handlerTemplateMarshal) GetGoImportMetaContent() string {
	return htm.PageHandler.MetaOutput.GetMetaContentForGoImport(htm.RequestPath)
}

func (htm handlerTemplateMarshal) GetGoSourceMetaContent() string {
	return htm.PageHandler.MetaOutput.GetMetaContentForGoSource(htm.RequestPath)
}

func (htm handlerTemplateMarshal) GetLink() string {
	if htm.PageHandler.MetaOutput.Username == "" {
		return htm.PageHandler.MetaOutput.BasePrefixURL + "/" + strings.TrimLeft(path.Clean(htm.PageHandler.MetaOutput.GetPath(htm.RequestPath)), "/")
	} else {
		return htm.PageHandler.MetaOutput.BasePrefixURL + "/" + strings.TrimLeft(path.Join(htm.PageHandler.MetaOutput.Username, htm.PageHandler.MetaOutput.GetPath(htm.RequestPath)), "/")
	}
}
