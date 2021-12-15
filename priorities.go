package main

import (
	"log"
	//"math/rand"

	extender "k8s.io/kube-scheduler/extender/v1"
)

// It'd better to only define one custom priority per extender
// as current extender interface only supports one single weight mapped to one extender
// and also it returns HostPriorityList, rather than []HostPriorityList

const (
	// lucky priority gives a random [0, extender.MaxPriority] score
	// currently extender.MaxPriority is 10
	luckyPrioMsg = "pod %v is lucky to get score %v on %v.\n"
)

// it's webhooked to pkg/scheduler/core/generic_scheduler.go#prioritizeNodes()
// you can't see existing scores calculated so far by default scheduler
// instead, scores output by this function will be added back to default scheduler
func prioritize(args extender.ExtenderArgs) *extender.HostPriorityList {
	pod := args.Pod
	nodes := args.Nodes.Items
	log.Printf("prioitize called")
	hostPriorityList := make(extender.HostPriorityList, len(nodes))
	var score int64 = 0
	for i, node := range nodes {
		if node.Name == "worker1"{
			log.Printf("worker1")
			score = int64(score1)
		}else if node.Name == "worker2-desktop"{
			log.Printf("worker2")
			score = int64(score2)
		}else if node.Name == "worker3" {
			log.Printf("worker3")
			score = int64(score3)
		}
		if score > extender.MaxExtenderPriority {
			log.Printf("extender Max Prioirty overflow!")
			score = extender.MaxExtenderPriority
		}
		log.Printf(luckyPrioMsg, pod.Name, score, node.Name)
		hostPriorityList[i] = extender.HostPriority{
			Host:  node.Name,
			Score: score,
		}
	}

	return &hostPriorityList
}
