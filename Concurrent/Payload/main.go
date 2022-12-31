package main

import "github.com/langwan/chiab"

func main() {
	chiab.Run(func(id int64) bool {
		payload()
		return true
	}, 100, 100, "payload", false)
}
