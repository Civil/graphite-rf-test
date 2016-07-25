package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Civil/graphite-rf-test/chooser"
	"github.com/Civil/graphite-rf-test/rf1r"
	"github.com/Civil/graphite-rf-test/rf2"
	"github.com/Civil/graphite-rf-test/simulator"

	"github.com/dgryski/go-onlinestats"
)

func chooseReplicationMethod(clusterType string, servers int, gnuplot bool) chooser.Chooser {
	var method chooser.Chooser
	switch clusterType {
	case "rf2":
		if !gnuplot {
			fmt.Printf("Running simulation for Replication Factor 2\n")
		}
		method = simulator.New(rf2.New(), servers)
	case "rf1-2c":
		if !gnuplot {
			fmt.Printf("Running simulation for Replication Factor 1 with 2 clusters\n")
		}
		method = simulator.New(rf1r.New(false), servers)
	case "rf1-2c-rnd":
		if !gnuplot {
			fmt.Printf("Running simulation for Replication Factor 1 with 2 clusters and randomization\n")
		}
		method = simulator.New(rf1r.New(true), servers)
	default:
		log.Fatalf("unknown chooser: %s; known rf2, rf1-2c, rf1-2c-rnd", clusterType)
	}
	return method
}

func loadMetrics(filePath string, method chooser.Chooser) {
	var buckets []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buckets = append(buckets, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = method.LoadMetrics(buckets)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	workers := flag.Int("w", 4, "workers")
	servers := flag.Int("s", 4, "servers per cluster")
	experiments := flag.Int("e", 1000, "experiments per dead server count")
	clusterType := flag.String("type", "rf2", "\"rf2\", \"rf1-2c\", \"rf1-2c-rnd\"")
	filePath := flag.String("f", "metrics", "path to file with metric names")
	printRaw := flag.Bool("r", false, "Print raw stats in the end")
	gnuplot := flag.Bool("g", false, "Gnuplot friendly output")

	flag.Parse()

	method := chooseReplicationMethod(*clusterType, *servers, *gnuplot)

	loadMetrics(*filePath, method)

	resCh := make(chan chooser.Result, *workers)
	if *gnuplot {
		fmt.Printf("TotalSrv	DownSrv	MeanLoss	StdDev	CV	MaxLoss	MinLoss	ChanceToLooseData\n")
	}
	for srv := 2; srv < *servers; srv++ {
		if !*gnuplot {
			fmt.Printf("Simulating servers down: %v of %v, running %v experiments\n", srv, *servers, *experiments)
		}
		expsToRun := *experiments
		var res []chooser.Result
		for expsToRun > 0 {
			runs := *workers
			if expsToRun < runs {
				runs = expsToRun
			}

			for i := 0; i < runs; i++ {
				go func(ch chooser.Chooser, servers int) {
					resCh <- ch.Simulate(servers)
				}(method, srv)
			}

			for i := 0; i < runs; i++ {
				res = append(res, <-resCh)
			}

			expsToRun -= runs
		}

		stats := onlinestats.NewRunning()
		max := float64(0)
		min := float64(1)
		cnt := 0
		for _, r := range res {
			if r.MetricsLost > 0 {
				cnt++
			}
			prct := float64(r.MetricsLost) / float64(r.MetricsTotal)
			if prct > max {
				max = prct
			}
			if prct < min {
				min = prct
			}
			stats.Push(prct)
		}
		chanceToLooseData := float64(cnt*100) / float64(len(res))
		if !*gnuplot {
			fmt.Printf("mean: %f stddev: %f, cv: %f, max: %f, min: %f, chance: %f\n", stats.Mean()*100, stats.Stddev()*100, stats.Stddev()/stats.Mean()*100, max*100, min*100, chanceToLooseData)
			if *printRaw {
				fmt.Printf("Raw Stats: %+v\n", res)
			}
			fmt.Printf("\n")
		} else {
			//fmt.Printf("TotalSrv	DownSrv	MeanLoss	StdDev	CV	MaxLoss	MinLoss\n")
			fmt.Printf("%v	%v	%f	%f	%f	%f	%f	%f\n", *servers, srv, stats.Mean()*100, stats.Stddev()*100, stats.Stddev()/stats.Mean()*100, max*100, min*100, chanceToLooseData)
		}
	}
}
