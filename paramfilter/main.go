package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// paramRange include a range for a parameter of a
// certain name.
type paramRange struct {
	Name      string  `json:"name"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Increment float64 `json:"increment"`
}

// paramRange is a slice of paramRange values.
type paramRanges []paramRange

func main() {

	// Read in the parameter file.
	//file, e := ioutil.ReadFile("/pfs/parameters/parameters.json")
	file, err := ioutil.ReadFile("example_parameters.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the ranges.
	var pRanges paramRanges
	if err := json.Unmarshal(file, &pRanges); err != nil {
		log.Fatal(err)
	}

	// Prepare combinations.
	var sets [][]float64
	for _, rng := range pRanges {
		set := []float64{}
		currentVal := rng.Min
		for currentVal <= rng.Max {
			set = append(set, currentVal)
			currentVal += rng.Increment
		}
		sets = append(sets, set)
	}

	// Expand the combinations.
	combos := make(map[string]string)
	current := make([]float64, len(pRanges))
	expand(sets, current, 0, combos)

	// Gather and output the results.
	for k, v := range combos {
		fmt.Println(k, v)
		//data := []byte(v)
		//err = ioutil.WriteFile("/pfs/out/"+k, data, 0644)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
}

// expand computes the cartesian product of N float slices.
func expand(input [][]float64, current []float64, k int, combos map[string]string) {
	if k == len(input) {
		var combo string
		for i := 0; i < k; i++ {
			combo = combo + fmt.Sprintf("%.2f ", current[i])
		}
		fileName := strings.Replace(combo, " ", "", -1)
		fileName = strings.Replace(fileName, ".", "", -1)
		combos[fileName] = combo
	} else {
		for i := 0; i < len(input[k]); i++ {
			current[k] = input[k][i]
			expand(input, current, k+1, combos)
		}
	}
}
