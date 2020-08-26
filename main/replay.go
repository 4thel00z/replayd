package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/4thel00z/replayd"
	"io/ioutil"
	"net/http/httputil"
	"os"
)

var (
	dry = flag.Bool("dry", false, "instead of replaying the request it is just being printed")
)

func main() {
	replayd.Init()
	flag.Parse()

	b, _ := ioutil.ReadAll(os.Stdin)
	request, err := replayd.Deserialize(string(b))
	if err != nil {
		os.Exit(1)
	}
	if *dry {
		payload, err := json.Marshal(request)
		if err != nil {
			os.Exit(1)
		}
		fmt.Print(string(payload))
		os.Exit(0)
	}

	response, err := replayd.Invoke(request)
	if err != nil {
		fmt.Println(request)
		os.Exit(1)
	}

	dump, err := httputil.DumpResponse(response, true)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%+v", dump)

}
