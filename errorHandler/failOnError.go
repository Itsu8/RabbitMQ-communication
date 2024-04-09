package errorHandler

import "log"

func PanicOnError(msg string, err error){
	if err != nil{
		log.Panicf("%s: %s", msg, err)
	}
}
func LogOnError(msg string, err error){
	if err != nil{
		log.Printf("%s: %s", msg, err)
	}
}