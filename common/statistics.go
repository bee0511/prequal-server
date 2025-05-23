package common

import (
	"fmt"
	"log"
	"sync"
	"sort"
)

type ReplicaStatisticsParameters struct {
	SuccessfulRequests int
	FailedRequests     int
}

var ReplicaStatistics map[string]ReplicaStatisticsParameters
var lock = sync.RWMutex{}

func InitializeStatistics(replicas []string) {
	lock.Lock()
	defer lock.Unlock()
	ReplicaStatistics = make(map[string]ReplicaStatisticsParameters)
	for _, replica := range replicas {
		AppendReplicaKey(replica)
	}
}

func AppendReplicaKey(replica string) {
	log.Print("Appending new replica to stat: ", replica)
	ReplicaStatistics[replica] = ReplicaStatisticsParameters{
		SuccessfulRequests: 0,
		FailedRequests:     0,
	}
}

func RemoveReplicaKey(replica string) {
	lock.Lock()
	defer lock.Unlock()
	delete(ReplicaStatistics, replica)
}

func IncrementSuccessfulRequests(replica string) {
	lock.Lock()
	defer lock.Unlock()
	if stats, exists := ReplicaStatistics[replica]; exists {
		ReplicaStatistics[replica] = ReplicaStatisticsParameters{
			SuccessfulRequests: stats.SuccessfulRequests + 1,
			FailedRequests:     stats.FailedRequests,
		}
	} else {
		ReplicaStatistics[replica] = ReplicaStatisticsParameters{
			SuccessfulRequests: 1,
			FailedRequests:     0,
		}
	}
}

func IncrementFailedRequests(replica string) {
	lock.Lock()
	defer lock.Unlock()
	if stats, exists := ReplicaStatistics[replica]; exists {
		ReplicaStatistics[replica] = ReplicaStatisticsParameters{
			SuccessfulRequests: stats.SuccessfulRequests,
			FailedRequests:     stats.FailedRequests + 1,
		}
	} else {
		ReplicaStatistics[replica] = ReplicaStatisticsParameters{
			SuccessfulRequests: 0,
			FailedRequests:     1,
		}
	}

}

type statDataArr struct {
	ReplicaName string                      `json:"replica_name"`
	Statistics  ReplicaStatisticsParameters `json:"statistics"`
}

func TransformMapToJson() []statDataArr {
	var jsonArray []statDataArr

	lock.RLock()
	defer lock.RUnlock()
	var keys []string
	for name := range ReplicaStatistics {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for _, name := range keys {
		stats := ReplicaStatistics[name]
		jsonArray = append(jsonArray, statDataArr{
			ReplicaName: name,
			Statistics:  stats,
		})
	}

	// fmt.Printf("Final JSON Array: %+v\n", jsonArray)

	return jsonArray
}

func PrintStatistics() {
	lock.RLock()
	defer lock.RUnlock()

	if len(ReplicaStatistics) == 0 {
		log.Println("No statistics data available.")
		return
	}

	fmt.Println("Replica Statistics:")
	for url, stats := range ReplicaStatistics {
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("  Successful Requests: %d\n", stats.SuccessfulRequests)
		fmt.Printf("  Failed Requests: %d\n", stats.FailedRequests)
		fmt.Println("---")
	}
}
