package main

import (
	"fmt"
	"sort"
)

type schedule struct {
	start  float64
	end    float64
	weight int
}

var schedules = []schedule{
	{6, 8, 2},
	{6.5, 12, 1},
	{6.5, 7, 2},
	{7, 8, 2},
	{7.5, 10, 3},
	{8, 9, 2},
	{8, 10, 1},
	{9, 12, 2},
	{9.5, 10, 4},
	{10, 11, 2},
	{10, 12, 3},
	{11, 12, 7},
}

type timeWithAction struct {
	time   float64
	weight int
	join   bool
}

type byTime []timeWithAction

func (b byTime) Len() int {
	return len(b)
}

func (b byTime) Less(i, j int) bool {
	if b[i].time == b[j].time {
		if b[i].join {
			return false
		} else {
			return true
		}
	}
	return b[i].time < b[j].time
}

func (b byTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func bestTimeToParty(schedules []schedule, me schedule) (float64, int) {
	times := make([]timeWithAction, 0, 2*len(schedules))
	for _, s := range schedules {
		times = append(times,
			timeWithAction{time: s.start, weight: s.weight, join: true},
			timeWithAction{time: s.end, weight: s.weight, join: false},
		)
	}
	sort.Sort(byTime(times))

	current := 0
	bestTime := float64(0)
	bestNumPeople := 0
	for _, t := range times {
		if t.time >= me.end {
			break
		}
		if t.join {
			current += t.weight
			if t.time >= me.start && current > bestNumPeople {
				bestTime = t.time
				bestNumPeople = current
			}
		} else {
			current -= t.weight
		}
	}

	return bestTime, bestNumPeople
}

func main() {
	fmt.Println(bestTimeToParty(schedules, schedule{start: 6, end: 12}))
	fmt.Println(bestTimeToParty(schedules, schedule{start: 6, end: 8}))
	fmt.Println(bestTimeToParty(schedules, schedule{start: 10, end: 11}))
}
