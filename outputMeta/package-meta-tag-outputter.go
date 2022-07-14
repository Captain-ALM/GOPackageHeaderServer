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
	PathLengthLimit    uint //The number of path entries in the go import paths
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaTags(pathIn string) string {
	return "<meta name=\"go-import\" content=\"" + pkgMTO.GetMetaContentForGoImport(pathIn) + "\">\r\n" +
		"<meta name=\"go-source\" content=\"" + pkgMTO.GetMetaContentForGoSource(pathIn) + "\">"
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaContentForGoImport(pathIn string) string {
	pathLoc := pkgMTO.GetPath(pathIn)
	return pkgMTO.getPrefix(pathLoc) + " git " + pkgMTO.getHomeURL(pathLoc)
}

func (pkgMTO *PackageMetaTagOutputter) GetMetaContentForGoSource(pathIn string) string {
	pathLoc := pkgMTO.GetPath(pathIn)
	return pkgMTO.getPrefix(pathLoc) + " " + pkgMTO.getHomeURL(pathLoc) + " " +
		pkgMTO.getDirectoryURL(pathLoc) + " " + pkgMTO.getFileURL(pathLoc)
}

func (pkgMTO *PackageMetaTagOutputter) GetPath(pathIn string) string {
	cleaned := path.Clean(pathIn)
	if cleaned == "/" || cleaned == "." {
		return cleaned
	}
	split := strings.Split(cleaned, "/")
	toReturn := ""
	for i := 1; i < len(split) && i < int(pkgMTO.PathLengthLimit)+1; i++ {
		toReturn += split[i] + "/"
	}
	return toReturn[:len(toReturn)-1]
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
