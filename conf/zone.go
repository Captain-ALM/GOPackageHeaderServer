package conf

import "golang.captainalm.com/GOPackageHeaderServer/outputMeta"

type ZoneYaml struct {
	Name               string            `yaml:"name"`
	Domains            []string          `yaml:"domains"`
	CssURL             string            `yaml:"cssURL"`
	HavePageContents   bool              `yaml:"havePageContents"`
	BasePath           string            `yaml:"basePath"`
	UsernameProvided   bool              `yaml:"usernameProvided"` //If set, the outputter will do /{user}/{repo}/ for repos rather than /{repo}/ ; Should really be named usernameProvidedByRequest
	Username           string            `yaml:"username"`
	BasePrefixURL      string            `yaml:"basePrefixURL"`
	SuffixDirectoryURL string            `yaml:"suffixDirectoryURL"`
	SuffixFileURL      string            `yaml:"suffixFileURL"`
	RangeSupported     bool              `yaml:"rangeSupported"`
	PathLengthLimit    uint              `yaml:"pathLengthLimit"` //The length of the path (Number of entries in the path) to return in the responses; (If 0: defaults to 1, if the username is not expected to be provided by the request, otherwise defaulting to 2)
	CacheSettings      CacheSettingsYaml `yaml:"cacheSettings"`
}

func (zy ZoneYaml) GetPackageMetaTagOutputter() *outputMeta.PackageMetaTagOutputter {
	var theUsername string
	if !zy.UsernameProvided {
		theUsername = zy.Username
	}
	pthLength := zy.PathLengthLimit
	if pthLength == 0 {
		if zy.UsernameProvided {
			pthLength = 2
		} else {
			pthLength = 1
		}
	}
	return &outputMeta.PackageMetaTagOutputter{
		BasePath:           zy.BasePath,
		Username:           theUsername,
		BasePrefixURL:      zy.BasePrefixURL,
		SuffixDirectoryURL: zy.SuffixDirectoryURL,
		SuffixFileURL:      zy.SuffixFileURL,
		PathLengthLimit:    pthLength,
	}
}
