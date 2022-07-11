package conf

import "golang.captainalm.com/GOPackageHeaderServer/outputMeta"

type ZoneYaml struct {
	Name               string   `yaml:"name"`
	Domains            []string `yaml:"domains"`
	HavePageContents   bool     `yaml:"havePageContents"`
	BasePath           string   `yaml:"basePath"`
	UsernameProvided   bool     `yaml:"usernameProvided"`
	Username           string   `yaml:"username"`
	BasePrefixURL      string   `yaml:"basePrefixURL"`
	SuffixDirectoryURL string   `yaml:"suffixDirectoryURL"`
	SuffixFileURL      string   `yaml:"suffixFileURL"`
}

func (zy ZoneYaml) GetPackageMetaTagOutputter() *outputMeta.PackageMetaTagOutputter {
	var theUsername string
	if !zy.UsernameProvided {
		theUsername = zy.Username
	}
	return &outputMeta.PackageMetaTagOutputter{
		BasePath:           zy.BasePath,
		Username:           theUsername,
		BasePrefixURL:      zy.BasePrefixURL,
		SuffixDirectoryURL: zy.SuffixDirectoryURL,
		SuffixFileURL:      zy.SuffixFileURL,
	}
}