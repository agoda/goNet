
package model


type Command interface{
	Run() (bool,error)
	GetResults() ([]byte,error)
}