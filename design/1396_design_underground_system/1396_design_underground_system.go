package main

import "strings"

type UndergroundSystem struct {
	TotalTimeOfTrip map[string]int
	CountsOfTrip    map[string]int
	CheckInStation  map[int]string
	CheckInTime     map[int]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[string]int),
		make(map[string]int),
		make(map[int]string),
		make(map[int]int),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.CheckInStation[id] = stationName
	this.CheckInTime[id] = t
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	tripName := strings.Join([]string{this.CheckInStation[id], stationName}, "-")
	this.TotalTimeOfTrip[tripName] += t - this.CheckInTime[id]
	this.CountsOfTrip[tripName]++
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	tripName := strings.Join([]string{startStation, endStation}, "-")
	return float64(this.TotalTimeOfTrip[tripName]) / float64(this.CountsOfTrip[tripName])

}
