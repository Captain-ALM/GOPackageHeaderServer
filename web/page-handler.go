package web

import (
	_ "embed"
	"golang.captainalm.com/GOPackageHeaderServer/conf"
	"golang.captainalm.com/GOPackageHeaderServer/outputMeta"
	"golang.captainalm.com/GOPackageHeaderServer/web/utils"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"time"
)

type PageHandler struct {
	Name           string
	CSS            string
	OutputPage     bool
	RangeSupported bool
	CacheSettings  conf.CacheSettingsYaml
	MetaOutput     *outputMeta.PackageMetaTagOutputter
}

var startTime = time.Now()

//go:embed output-page.html
var outputPage string

var pageTemplateFuncMap = template.FuncMap{
	"isNotEmpty": func(stringIn string) bool {
		return stringIn != ""
	},
}

func (pgh *PageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet || request.Method == http.MethodHead {
		tmpl, err := template.New("page-handler").Funcs(pageTemplateFuncMap).Parse(outputPage)
		if err != nil {
			utils.WriteResponseHeaderCanWriteBody(request.Method, writer, http.StatusInternalServerError, "Page Template Parsing Failure")
			return
		}
		tm := handlerTemplateMarshal{
			PageHandler: *pgh,
			RequestPath: request.URL.Path,
		}
		theBuffer := &utils.BufferedWriter{}
		err = tmpl.Execute(theBuffer, tm)
		if err != nil {
			utils.WriteResponseHeaderCanWriteBody(request.Method, writer, http.StatusInternalServerError, "Page Template Execution Failure")
			return
		}
		writer.Header().Set("Content-Length", strconv.Itoa(len(theBuffer.Data)))
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		utils.SetLastModifiedHeader(writer.Header(), startTime)
		utils.SetCacheHeaderWithAge(writer.Header(), pgh.CacheSettings.MaxAge, startTime)
		theETag := utils.GetValueForETagUsingBufferedWriter(theBuffer)
		writer.Header().Set("ETag", theETag)
		if utils.ProcessSupportedPreconditionsForNext(writer, request, startTime, theETag, pgh.CacheSettings.NotModifiedResponseUsingLastModified, pgh.CacheSettings.NotModifiedResponseUsingETags) {
			httpRangeParts := utils.ProcessRangePreconditions(int64(len(theBuffer.Data)), writer, request, startTime, theETag, pgh.RangeSupported)
			if httpRangeParts != nil {
				if len(httpRangeParts) <= 1 {
					var theWriter io.Writer = writer
					if len(httpRangeParts) == 1 {
						theWriter = utils.NewPartialRangeWriter(theWriter, httpRangeParts[0])
					}
					_, _ = theWriter.Write(theBuffer.Data)
				} else {
					multWriter := multipart.NewWriter(writer)
					writer.Header().Set("Content-Type", "multipart/byteranges; boundary="+multWriter.Boundary())
					for _, currentPart := range httpRangeParts {
						mimePart, err := multWriter.CreatePart(textproto.MIMEHeader{
							"Content-Range": {currentPart.ToField(int64(len(theBuffer.Data)))},
							"Content-Type":  {"text/plain; charset=utf-8"},
						})
						if err != nil {
							break
						}
						_, err = mimePart.Write(theBuffer.Data[currentPart.Start : currentPart.Start+currentPart.Length])
						if err != nil {
							break
						}
					}
					_ = multWriter.Close()
				}
			}
		}
	} else {
		writer.Header().Set("Allow", http.MethodOptions+", "+http.MethodGet+", "+http.MethodHead)
		if request.Method == http.MethodOptions {
			utils.WriteResponseHeaderCanWriteBody(request.Method, writer, http.StatusOK, "")
		} else {
			utils.WriteResponseHeaderCanWriteBody(request.Method, writer, http.StatusMethodNotAllowed, "")
		}
	}
}
