package web

import (
	"github.com/gorilla/mux"
	"golang.captainalm.com/GOPackageHeaderServer/conf"
	"log"
	"net/http"
	"strings"
)

func New(yaml conf.ConfigYaml) (*http.Server, map[string]*PageHandler) {
	router := mux.NewRouter()
	var pages = make(map[string]*PageHandler)
	for _, zc := range yaml.Zones {
		currentPage := &PageHandler{
			Name:       zc.Name,
			OutputPage: zc.HavePageContents,
			MetaOutput: zc.GetPackageMetaTagOutputter(),
		}
		for _, d := range zc.Domains {
			ld := strings.ToLower(d)
			if _, exists := pages[ld]; !exists {
				pages[ld] = currentPage
				router.Host(ld).HandlerFunc(currentPage.ServeHTTP)
			}
		}
	}
	if yaml.Listen.Identify {
		router.Use(headerMiddleware)
	}
	if yaml.Listen.Web == "" {
		log.Fatalf("[Http] Invalid Listening Address")
	}
	s := &http.Server{
		Addr:         yaml.Listen.Web,
		Handler:      router,
		ReadTimeout:  yaml.Listen.GetReadTimeout(),
		WriteTimeout: yaml.Listen.GetWriteTimeout(),
	}
	go runBackgroundHttp(s)
	return s, pages
}

func runBackgroundHttp(s *http.Server) {
	err := s.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Println("The http server shutdown successfully")
		} else {
			log.Fatalf("[Http] Error trying to host the http server: %s\n", err.Error())
		}
	}
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Clerie Gilbert")
		w.Header().Set("X-Powered-By", "Love")
		w.Header().Set("X-Friendly", "True")
		next.ServeHTTP(w, r)
	})
}
