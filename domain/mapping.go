package domain

type Mapping struct {
	Kind    string `yaml:"kind"`
	Version string `yaml:"version"`
	Spec    spec   `yaml:"spec"`
}

type spec struct {
	DataSources *[]dataSource `yaml:"dataSources"`
	DataSinks   *[]dataSink   `yaml:"dataSinks"`
	Mapping     *mapping    `yaml:"mapping"`
}

type dataSource struct {
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
	Service string `yaml:"service"`
	Rest    rest   `yaml:"rest"`
}
type rest struct {
	Request request  `yaml:"request"`
	Params  *[]param `yaml:"params"`
}

type request struct {
	Method   string `yaml:"method"`
	Resource string `yaml:"resource"`
}

type param struct {
	Name       string `yaml:"name"`
	Value      string `yaml:"value"`
	Expression string `yaml:"expression"`
}

type dataSink struct {
	Kind string `yaml:"kind"`
	Name string `yaml:"name"`
}

type mapping struct {
	Kind string `yaml:"kind"`
	Jq   string `yaml:"jq"`
}
