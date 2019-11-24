package persist

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()

	return out
}
