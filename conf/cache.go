package conf

type CacheSettingsYaml struct {
	MaxAge                               uint `yaml:"maxAge"`
	NotModifiedResponseUsingLastModified bool `yaml:"notModifiedUsingLastModified"`
	NotModifiedResponseUsingETags        bool `yaml:"notModifiedUsingETags"`
}
