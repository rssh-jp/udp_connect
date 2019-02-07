package crypto

import (
	"fmt"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	src := []byte("testtest")
	key := []byte("01")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			dest := Cryption(src, key)

			dec := Decryption(dest, key)
			fmt.Println(src, dest, dec)

			for index, val := range dec {
				if val != src[index] {
					t.Fatalf("Could match data. want : %v, got : %v", src, dec)
				}
			}
		}(i)
	}
	wg.Wait()
}
