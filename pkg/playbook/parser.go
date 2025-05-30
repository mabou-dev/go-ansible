package playbook

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func Load(path string) (*Playbook, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pb Playbook
	err = yaml.Unmarshal(data, &pb)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
