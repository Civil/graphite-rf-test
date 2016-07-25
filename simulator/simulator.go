package simulator

import (
	algorithm "github.com/Civil/graphite-rf-test/algorithm"
	"github.com/Civil/graphite-rf-test/chooser"

	"math/rand"
	"time"
)

type Simulator struct {
	algorithm        algorithm.Algorithm
	amountOfMetrics  int
	metricsToNodeMap [][]int
	nodeMax          int
}

func New(algorithm algorithm.Algorithm, servers int) *Simulator {
	res := Simulator{algorithm: algorithm, nodeMax: servers}
	return &res
}

// LoadMetrics load metrics
func (data *Simulator) LoadMetrics(buckets []string) error {
	data.amountOfMetrics = len(buckets)
	data.metricsToNodeMap = make([][]int, data.nodeMax)

	for _, metric := range buckets {
		metricsServers := data.algorithm.Choose(metric, data.nodeMax)
		data.metricsToNodeMap[metricsServers[0]] = append(data.metricsToNodeMap[metricsServers[0]], metricsServers[1])
		data.metricsToNodeMap[metricsServers[1]] = append(data.metricsToNodeMap[metricsServers[1]], metricsServers[0])
	}

	return nil
}

func (data *Simulator) Simulate(serversToKill int) chooser.Result {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	killedServers := make([]int, serversToKill)
	serversStatus := make([]bool, data.nodeMax)
	srvList := make([]int, data.nodeMax)
	for i := range srvList {
		serversStatus[i] = true
		srvList[i] = i
	}

	for i := 0; i < serversToKill; i++ {
		s := int(rnd.Int31n(int32(len(srvList))))
		killedServers[i] = srvList[s]
		serversStatus[srvList[s]] = false
		srvList[s], srvList = srvList[len(srvList)-1], srvList[:len(srvList)-1]
	}

	res := chooser.Result{data.amountOfMetrics, 0}
	/*
		fmt.Printf("killedServers: %+v\n", killedServers)
		fmt.Printf("srvList: %+v\n", data.metricsMap[0])
	*/

	for _, srv := range killedServers {
		for _, backupNode := range data.metricsToNodeMap[srv] {
			if !serversStatus[backupNode] {
				res.MetricsLost++
			}
		}
	}

	res.MetricsLost /= 2

	return res
}
