package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func getStats(url string) {
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

	resultInMH := metrics.Data.Currenthashrate / 1000000

	fmt.Printf("%+v\n", &metrics)

	fmt.Println("Current Hashrate in MH/s: ", resultInMH)
}

func main() {
	const ethAddress = ""
	url := fmt.Sprintf("https://api.ethermine.org/miner/%s/currentStats", ethAddress)

	getStats(url)
}