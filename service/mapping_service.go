package service

import (
	"encoding/json"
	"fmt"
	"github.com/cobre420/mapper-go/domain"
	"github.com/savaki/jq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

func ProcessMapping(mappingFile, configFile string, state *domain.StateVector) {
	c := &domain.Config{}
	loadJsonFile(c, configFile)
	log.Infof("Loaded %v", *c)

	m := &domain.Mapping{}
	loadYamlFile(m, mappingFile)
	log.Infof("Loaded %v", *m)

	if m.Kind != "mapping" {
		panic("Invalid manifest type: " + m.Kind)
	}
	if m.Version != "v1" {
		panic("Invalid manifest version: " + m.Version)
	}
	for _, dsSpec := range *m.Spec.DataSources {
		service := _findDs(c, dsSpec.Name)
		if service == nil {
			panic(fmt.Sprintf("Cannot process mapping: datasource=%s not found", dsSpec.Name))
		}
		_evaluateDatasource(service, &dsSpec, state)
	}
}

func _evaluateDatasource(service *domain.Service, ds *domain.DataSource, state *domain.StateVector) {
	if !strings.EqualFold(ds.Kind, "rest") {
		panic(fmt.Sprintf("Unsupported service kind %s", ds.Kind))
	}

	request := _createRequest(service, ds, state)

	log.Infof("Created request %v", request)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	log.Infof("Response %v", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Infof("Response body %s", body)
}

func _createRequest(service *domain.Service, ds *domain.DataSource, state *domain.StateVector) (request *http.Request) {
	elements := strings.Split(ds.Rest.Request.Resource, "/")

	parameters := make(map[string]string, 0)
	for _, pe := range elements {
		if strings.HasPrefix(pe, "<") && strings.HasSuffix(pe, ">") {
			//log.Infof("Element %s hsa to be evaluated", pe)

			p := strings.TrimPrefix(pe, "<")
			p = strings.TrimSuffix(p, ">")
			parameters[p] = "TBD"

			if param := findParam(p, ds.Rest.Params); param != nil {
				if param.Value != "" {
					parameters[p] = param.Value
				} else if param.Expression != "" {
					parameters[p] = resolveParam(param.Expression, state)
				} else {
					panic(fmt.Errorf(fmt.Sprintf("No value nor expression defined for param [%s]", param)))
				}
			}
		}
	}

	log.Infof("Parameters resolved %v", parameters)

	urlResource := resolveUrlResource(ds.Rest.Request.Resource, parameters)

	log.Infof("URL resolved %v", urlResource)

	var err error
	switch strings.ToUpper(ds.Rest.Request.Method) {
	case "GET":
		request, err = http.NewRequest(
			http.MethodGet,
			service.Host+service.Prefix+urlResource,
			nil,
		)
	case "POST":
		request, err = http.NewRequest(
			http.MethodPost,
			service.Host+service.Prefix,
			strings.NewReader((*ds.Rest.Params)[0].Value),
		)
	case "PUT":
		request, err = http.NewRequest(
			http.MethodPut,
			service.Host+service.Prefix,
			nil,
		)
	case "DELETE":
		request, err = http.NewRequest(
			http.MethodDelete,
			service.Host+service.Prefix,
			nil,
		)
	default:
		panic(fmt.Sprintf("Unsupported request method %s", ds.Rest.Request.Method))
	}

	if err != nil {
		panic(err)
	}

	return request
}

func resolveUrlResource(resource string, parameters map[string]string) string {
	result := resource
	for k, v := range parameters {
		result = strings.Replace(result, "<"+k+">", v, 1)
	}
	return result
}

func resolveParam(expression string, sv *domain.StateVector) string {
	ex := strings.Replace(expression, "vector", "", 1)

	op, err := jq.Parse(ex)
	if err != nil {
		panic(err)
	}
	r, err := op.Apply([]byte((*sv).Vector("main").Text))

	if err != nil {
		panic(err)
	}

	s := string(r)
	if strings.HasPrefix(s, "\"") {
		r = r[1:]
	}
	if strings.HasSuffix(s, "\"") {
		r = r[:len(r)-1]
	}

	return string(r)
}

func findParam(paramName string, params *[]domain.Param) *domain.Param {
	for i, p := range *params {
		if strings.EqualFold(p.Name, paramName) {
			return &(*params)[i]
		}
	}
	return nil
}

func _findDs(c *domain.Config, dsName string) *domain.Service {
	for _, ds := range *c.Services {
		if ds.Name == dsName {
			return &ds
		}
	}
	return nil
}

func loadYamlFile(out interface{}, resourceFile string) {
	data, err := ioutil.ReadFile(resourceFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		panic(err)
	}
}

func loadJsonFile(out interface{}, resourceFile string) {
	data, err := ioutil.ReadFile(resourceFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		panic(err)
	}
}
