package web

import (
	_ "embed"
	"golang.captainalm.com/GOPackageHeaderServer/outputMeta"
	"html/template"
	"net/http"
	"strconv"
)

type PageHandler struct {
	Name       string
	CSS        string
	OutputPage bool
	MetaOutput *outputMeta.PackageMetaTagOutputter
}

//go:embed output-page.html
var outputPage string

var pageTemplateFuncMap template.FuncMap = template.FuncMap{
	"isNotEmpty": func(stringIn string) bool {
		return stringIn != ""
	},
}

func (pgh *PageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet || request.Method == http.MethodHead {
		tmpl, err := template.New("page-handler").Funcs(pageTemplateFuncMap).Parse(outputPage)
		if err != nil {
			writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusInternalServerError, "Page Template Parsing Failure")
			return
		}
		tm := handlerTemplateMarshal{
			PageHandler: *pgh,
			RequestPath: request.URL.Path,
		}
		theBuffer := &BufferedWriter{}
		err = tmpl.Execute(theBuffer, tm)
		if err != nil {
			writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusInternalServerError, "Page Template Execution Failure")
			return
		}
		writer.Header().Set("Content-Length", strconv.Itoa(len(theBuffer.Data)))
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		if writeResponseHeaderCanWriteBody(request.Method, writer, http.StatusOK, "") {
			_, _ = writer.Write(theBuffer.Data)
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
