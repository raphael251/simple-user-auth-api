package main

import "github.com/raphael251/simple-user-auth-api/configs"

func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	println(configs.DBDriver)
}
