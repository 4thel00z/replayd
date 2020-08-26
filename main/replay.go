package main

import (
	"encoding/json"
	"fmt"
	"github.com/4thel00z/replayd"
	"io/ioutil"
	"os"
)

func main() {

	b, _ := ioutil.ReadAll(os.Stdin)
	deserialized, err := replayd.Deserialize(string(b))
	if err != nil {
		os.Exit(1)
	}
	payload, err := json.Marshal(deserialized)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print(string(payload))

}
