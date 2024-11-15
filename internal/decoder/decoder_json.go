package decoder

import (
	"encoding/json"
	"io"
	"log"
)

// Decodes json parameters from request's body 
func Decode(r io.Reader, v interface{}){
	decoder := json.NewDecoder(r)
	err := decoder.Decode(v)
	if err != nil{
		log.Printf("error decoding request's body: %v", err)
		return
	}	
}
