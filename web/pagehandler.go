package web

import (
	"golang.captainalm.com/GOPackageHeaderServer/outputMeta"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type PageHandler struct {
	Name       string
	OutputPage bool
	MetaOutput *outputMeta.PackageMetaTagOutputter
}

func (pgh *PageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet || request.Method == http.MethodHead {
		thePage := "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n"
		if pgh.OutputPage && pgh.Name != "" {
			thePage += "<title>Go Package: " + pgh.Name + "</title>\r\n"
		}
		thePage += pgh.MetaOutput.GetMetaTags(request.URL.Path) + "\r\n</head>\r\n<body>\r\n"
		if pgh.OutputPage {
			if pgh.Name != "" {
				thePage += "<h1>Go Package: " + pgh.Name + "</h1>\r\n"
			}
			var theLink string
			if pgh.MetaOutput.Username == "" {
				theLink = pgh.MetaOutput.BasePrefixURL + "/" + strings.TrimLeft(path.Clean(request.URL.Path), "/")
			} else {
				theLink = pgh.MetaOutput.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pgh.MetaOutput.Username, request.URL.Path), "/")
			}
			thePage += "<a href=\"" + theLink + "\">" + theLink + "</a>\r\n"
		}
		thePage += "</body>\r\n</html>\r\n"
		writer.Header().Set("Content-Length", strconv.Itoa(len([]byte(thePage))))
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		if writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusOK, "") {
			_, _ = writer.Write([]byte(thePage))
		}
	} else {
		writer.Header().Set("Allow", http.MethodOptions+", "+http.MethodGet+", "+http.MethodHead)
		if request.Method == http.MethodOptions {
			writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusOK, "")
		} else {
			writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusMethodNotAllowed, "")
		}
	}
}
