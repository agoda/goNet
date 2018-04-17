
package model


type Command interface{
	func Run() (bool,error)
	func GetResults() ([]byte,error)
}