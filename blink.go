// This program achieves LED blink on Raspberry Pi with Go lang.
// This is implemented with hard-coding and uses only main function.

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Initialize GPIO25
	fmt.Println("Initialize GPIO25")
	fd, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open /sys/class/gpio/export fails")
		fmt.Println(err)
		return
	}
	fmt.Fprint(fd, "25")
	fd.Close()

	// Check iinitialization result
	fmt.Println("Check initialization result")
	_, err = os.Stat("/sys/class/gpio/gpio25")
	if err != nil {
		fmt.Println("Export GPIO25 fails")
		fmt.Println(err)
	} else {
		fmt.Println("Export GPIO25 succeeds")
	}

	// Set direction to out
	fmt.Println("Set direction of GPIO25 to out")
	fd, err = os.OpenFile("/sys/class/gpio/gpio25/direction", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open /sys/class/gpio/gpio25/direction fails")
		fmt.Println(err)
	}
	fmt.Fprint(fd, "out")
	fd.Close()

	// Start blink
	fmt.Println("Start blink")
	for i := 0; i < 20; i++ {
		fd, err := os.OpenFile("/sys/class/gpio/gpio25/value", os.O_WRONLY|os.O_SYNC, 0666)
		if err != nil {
			fmt.Println("open /sys/class/gpio/gpio25/value fails")
			fmt.Println(err)
		}
		if i%2 == 1 {
			// Turn on LED
			fmt.Fprint(fd, 1)
		} else {
			// Turn off LED
			fmt.Fprint(fd, 0)
		}
		fd.Close()
		time.Sleep(100 * time.Millisecond)
	}

	// End
	fmt.Println("Start finalizing")
	fd, err = os.OpenFile("/sys/class/gpio/gpio25/value", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open /sys/class/gpio/gpio25/value fails")
		fmt.Println(err)
	}
	// Turn off LED
	fmt.Println("Turn off LED")
	fmt.Fprint(fd, 0)

	fd.Close()

	fd, err = os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("open /sys/class/gpio/unexport fails")
		fmt.Println(err)
		return
	}
	fmt.Fprint(fd, "25")
	fd.Close()
	fmt.Println("End finalizing")

}
