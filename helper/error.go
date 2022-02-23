package helper

import "log"

func ErrorNotNil (err error){
	if err != nil{
		log.Printf("Error: %s", err)
	}
}
