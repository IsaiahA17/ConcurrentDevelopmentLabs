// Copyright (C) 2025  Isaiah Andres

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"sync"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
