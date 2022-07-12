package outputMeta

import (
	"path"
	"strings"
)

type PackageMetaTagOutputter struct {
	BasePath           string
	Username           string //If set, the outputter will do /{repo}/ for repos rather than /{user}/{repo}/
	BasePrefixURL      string
	SuffixDirectoryURL string
	SuffixFileURL      string
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaTags(pathIn string) string {
	return "<meta name=\"go-import\" content=\"" + pkgMTO.GetMetaContentForGoImport(pathIn) + "\">\r\n" +
		"<meta name=\"go-source\" content=\"" + pkgMTO.GetMetaContentForGoSource(pathIn) + "\">"
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaContentForGoImport(pathIn string) string {
	return pkgMTO.getPrefix(pathIn) + " git " + pkgMTO.getHomeURL(pathIn)
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaContentForGoSource(pathIn string) string {
	return pkgMTO.getPrefix(pathIn) + " " + pkgMTO.getHomeURL(pathIn) + " " +
		pkgMTO.getDirectoryURL(pathIn) + " " + pkgMTO.getFileURL(pathIn)
}

func (pkgMTO *PackageMetaTagOutputter) assureBasePrefixURL() (failed bool) {
	if pkgMTO.BasePrefixURL == "" {
		if pkgMTO.BasePath == "" {
			return true
		}
		pkgMTO.BasePrefixURL = "http://" + pkgMTO.BasePath
	}
	return false
}

func (pkgMTO *PackageMetaTagOutputter) getPrefix(pathIn string) string {
	if pkgMTO.BasePath == "" {
		return "_"
	}
	return path.Join(pkgMTO.BasePath, pathIn)
}

func (pkgMTO *PackageMetaTagOutputter) getHomeURL(pathIn string) string {
	if pkgMTO.assureBasePrefixURL() {
		return "_"
	}

	if pkgMTO.Username == "" {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Clean(pathIn), "/")
	} else {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pkgMTO.Username, pathIn), "/")
	}
}

func (pkgMTO *PackageMetaTagOutputter) getDirectoryURL(pathIn string) string {
	if pkgMTO.assureBasePrefixURL() || pkgMTO.SuffixDirectoryURL == "" {
		return "_"
	}

	if pkgMTO.Username == "" {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pathIn, pkgMTO.SuffixDirectoryURL), "/")
	} else {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pkgMTO.Username, pathIn, pkgMTO.SuffixDirectoryURL), "/")
	}
}

func (pkgMTO *PackageMetaTagOutputter) getFileURL(pathIn string) string {
	if pkgMTO.assureBasePrefixURL() || pkgMTO.SuffixFileURL == "" {
		return "_"
	}

	if pkgMTO.Username == "" {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pathIn, pkgMTO.SuffixFileURL), "/")
	} else {
		return pkgMTO.BasePrefixURL + "/" + strings.TrimLeft(path.Join(pkgMTO.Username, pathIn, pkgMTO.SuffixFileURL), "/")
	}
}
