package conf

type ConfigYaml struct {
	Listen ListenYaml `yaml:"listen"`
	Zones  []ZoneYaml `yaml:"zones"`
}
