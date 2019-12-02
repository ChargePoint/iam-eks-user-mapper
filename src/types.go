package main

type MapUserConfig struct {
	UserArn  string   `yaml:"userarn"`
	Username string   `yaml:"username"`
	Groups   []string `yaml:"groups"`
}

type groupMapping struct {
	IAMGroup string   `yaml:"iamGroup"`
	K8sCaps  []string `yaml:"k8sCaps"`
}

type groupMappings struct {
	Mappings []groupMapping `yaml:"mappings"`
}
