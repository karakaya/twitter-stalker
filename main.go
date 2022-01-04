package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/karakaya/twitter-stalker/pkg/twitter"
)

var interval time.Duration
var userId int64

func init() {

	flag.DurationVar(&interval, "interval", time.Minute*10, "Check Interval")
	flag.Int64Var(&userId, "userId", 0, "Twitter User ID")
	flag.Parse()
}

var old, current twitter.UserFollower

func main() {

	for c := time.Tick(interval); ; <-c {

		current = check(old)
		old = current

	}

}
func diff(old, current twitter.UserFollower) {

	oldId := pushIdsToArr(old)
	currentId := pushIdsToArr(current)
	fmt.Println(oldId)
	fmt.Println(currentId)
	//TODO: Calc diff

}
func check(old twitter.UserFollower) twitter.UserFollower {
	twitter := twitter.Tw{ID: userId}
	updated := twitter.Followers()
	diff(old, updated)

	return updated
}

func pushIdsToArr(follower twitter.UserFollower) []int64 {
	arr := make([]int64, 0)
	for _, data := range follower.Data {
		i, _ := strconv.Atoi(data.ID)
		arr = append(arr, int64(i))
	}
	return arr
}
