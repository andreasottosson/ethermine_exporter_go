package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
)

func getStats(url string) []string {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Metrics struct {
		Status string `json:"status"`
		Data   struct {
			// Time             int         `json:"time"`
			// Lastseen         int         `json:"lastSeen"`
			Reportedhashrate int         `json:"reportedHashrate"`
			Currenthashrate  float64     `json:"currentHashrate"`
			Validshares      int         `json:"validShares"`
			Invalidshares    int         `json:"invalidShares"`
			Staleshares      int         `json:"staleShares"`
			Averagehashrate  float64     `json:"averageHashrate"`
			Activeworkers    int         `json:"activeWorkers"`
			// Unpaid           int64       `json:"unpaid"`
			// Unconfirmed      interface{} `json:"unconfirmed"`
			// Coinspermin      float64     `json:"coinsPerMin"`
			// Usdpermin        float64     `json:"usdPerMin"`
			// Btcpermin        float64     `json:"btcPerMin"`
		} `json:"data"`
	}
	
	var metrics Metrics

	err = json.Unmarshal(bytes, &metrics)
	if err != nil {
		fmt.Println("error:", err)
	}

	prefix := "ethermine"

	var metricsOut []string

	metricsOut = append(metricsOut, fmt.Sprintf("%v_%v %f", prefix, "current_hashrate", metrics.Data.Currenthashrate))

	return metricsOut
	
}

func MetricsHttp(w http.ResponseWriter, r *http.Request) {
	const ethAddress = ""
	url := fmt.Sprintf("https://api.ethermine.org/miner/%s/currentStats", ethAddress)

	allOut := getStats(url)
	fmt.Fprintln(w, strings.Join(allOut, "\n"))
}

func main() {
	port := "9118"
	http.HandleFunc("/metrics", MetricsHttp)
	panic(http.ListenAndServe("0.0.0.0:"+port, nil))
}