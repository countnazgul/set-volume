package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/itchyny/volume-go"
)

func main() {
	volumeToSet := 0

	args := os.Args

	if len(args) == 1 {
		fmt.Printf("Please provide new volume value")
		os.Exit(1)
	}

	// when setting specific volume value
	if !strings.Contains(args[1], ":") {
		argVolumeValue, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf(
				"Unable to convert the provided value to number. Hint: Use only integer values",
			)
			os.Exit(1)
		}

		volumeToSet = argVolumeValue
	}

	// when want to "cycle" between two possible values
	if strings.Contains(args[1], ":") {
		cycleValues := strings.Split(args[1], ":")

		v := make([]int, len(cycleValues))

		// convert the provided values to integer
		for i := 0; i < len(cycleValues); i++ {
			a, err := strconv.Atoi(cycleValues[i])
			if err != nil {
				fmt.Printf(
					"Unable to convert the MIN value to number. Hint: Use only integer values",
				)
				os.Exit(1)
			}

			if a < 0 || a > 100 {
				fmt.Printf("Volume values should be between 0 and 100")
				os.Exit(1)
			}

			v[i] = a
		}

		// sort the values in asc order ... just in case
		sort.Ints(v[:])

		// get the new volume value
		volumeToSet = findNextVolume(&v)
	}

	// set the volume to the new value
	err := volume.SetVolume(volumeToSet)
	if err != nil {
		fmt.Printf("set volume failed: %+v", err)
		os.Exit(1)
	}

	// just for the record get the actual volume and print it
	vol, err := volume.GetVolume()
	if err != nil {
		fmt.Printf("get volume failed: %+v", err)
		os.Exit(1)
	}

	fmt.Printf("New volume set: %v", vol)
}

// Find the next volume value in the provided list and return it
// if no value is found then return the first
func findNextVolume(values *[]int) int {
	// get the current actual volume
	vol, err := volume.GetVolume()
	if err != nil {
		fmt.Printf("get volume failed: %+v", err)
		os.Exit(1)
	}

	newVolume := 0

	// loop through the values and if the current volume value
	// is lower than the current iteration value then stop the loop
	// and return the current iteration value
	// on the next run of the program the next iteration will be returned
	for _, v := range *values {
		if vol < v {
			newVolume = v
			break
		}
	}

	if newVolume == 0 {
		newVolume = (*values)[0]
	}

	return newVolume
}
