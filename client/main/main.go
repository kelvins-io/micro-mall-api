package main

import "log"

func main() {
	token, err := NewJWT().CreateToken(1000009, "15532", 30*24*3600)
	if err != nil {
		log.Printf("createtoken err: %v\n", err)
		return
	}
	log.Printf("token ==%v", token)

	//var now = time.Now()
	//log.Println("0", now.Unix())
	//log.Println("0", now.AddDate(0,0,-1).Unix())
	//log.Println("0", now.AddDate(0,0,-2).Unix())
	//log.Println("0", now.AddDate(0,0,-3).Unix())
	//log.Println("0", now.AddDate(0,0,-4).Unix())
}
