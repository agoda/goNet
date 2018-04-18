package model

import (
"encoding/json"
"os/exec"
)



type Ping struct {
	Output string `json:"output"`
	options []string
}

func NewPing(options []string) *Ping {
	return &Ping{
		options: options,
	}
}

func (this *Ping) Run() (bool,error){

	cmd := this.options[0]
	countFlag := this.options[1]
	count := this.options[2]
	target := this.options[3]

	os_cmd := exec.Command(cmd,countFlag,count,target)
	out,err := os_cmd.Output()

	if err != nil {
		return false, err
	} else {
		this.Output = string(out)
		return true, nil
	}
}

func (this *Ping) GetResults() ([]byte,error) {
	return json.Marshal(this)
}