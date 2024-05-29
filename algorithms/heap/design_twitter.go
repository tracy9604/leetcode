package heap

import (
	"container/heap"
	"fmt"
	"time"
)

type TweetHeap []Tweet

func (h TweetHeap) Len() int {
	return len(h)
}

func (h TweetHeap) Less(i, j int) bool {
	return h[i].timestamp < h[j].timestamp
}

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type Tweet struct {
	tweetId   int
	timestamp int64
}

type Twitter struct {
	userTweets     map[int][]Tweet
	userFollowings map[int][]int
	tweets         map[int]Tweet
}

func Constructor() Twitter {
	return Twitter{
		userTweets:     make(map[int][]Tweet),
		userFollowings: make(map[int][]int),
		tweets:         make(map[int]Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	now := time.Now().Unix() + int64(len(this.tweets))
	tweet := Tweet{tweetId: tweetId, timestamp: now}
	this.tweets[tweetId] = tweet
	if _, found := this.userTweets[userId]; !found {
		this.userTweets[userId] = make([]Tweet, 0)
	}
	this.userTweets[userId] = append(this.userTweets[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := this.userTweets[userId]
	followees := this.userFollowings[userId]

	for _, followee := range followees {
		tweets = append(tweets, this.userTweets[followee]...)
	}

	maxTweet := 10

	minHeap := &TweetHeap{}
	heap.Init(minHeap)

	for _, tweet := range tweets {
		heap.Push(minHeap, tweet)
		if minHeap.Len() > maxTweet {
			heap.Pop(minHeap)
		}
	}

	ans := make([]int, 0)
	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Tweet)
		ans = append([]int{top.tweetId}, ans...)
	}

	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	followees, found := this.userFollowings[followerId]
	if !found {
		this.userFollowings[followerId] = make([]int, 0)
	}
	for _, userId := range followees {
		if userId == followeeId {
			return
		}
	}
	this.userFollowings[followerId] = append(this.userFollowings[followerId], followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followers, found := this.userFollowings[followerId]
	if !found {
		return
	}
	for idx, userId := range followers {
		if userId == followeeId {
			followers = append(followers[:idx], followers[idx+1:]...)
			this.userFollowings[followerId] = followers
		}
	}
}

func TestDesignTwitter() {
	obj := Constructor()
	obj.PostTweet(1, 5)
	obj.PostTweet(1, 3)
	obj.PostTweet(1, 101)
	obj.PostTweet(1, 13)
	obj.PostTweet(1, 10)
	obj.PostTweet(1, 2)
	obj.PostTweet(1, 94)
	obj.PostTweet(1, 505)
	obj.PostTweet(1, 333)
	fmt.Println(obj.GetNewsFeed(1))

}
