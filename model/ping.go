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

}

func (this *Ping) GetResults() ([]byte,error) {
	
}