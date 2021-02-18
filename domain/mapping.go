package domain

type Mapping struct {
	Kind    string `yaml:"kind"`
	Version string `yaml:"version"`
	Spec    spec   `yaml:"spec"`
}

type spec struct {
	DataSources *[]DataSource `yaml:"dataSources"`
	DataSinks   *[]dataSink   `yaml:"dataSinks"`
	Mapping     *mapping      `yaml:"mapping"`
}

type DataSource struct {
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
	Service string `yaml:"service"`
	Rest    rest   `yaml:"rest"`
}
type rest struct {
	Request request  `yaml:"request"`
	Body    string   `yaml:"body"`
	Params  *[]Param `yaml:"params"`
}

type request struct {
	Method   string `yaml:"method"`
	Resource string `yaml:"resource"`
}

type Param struct {
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
