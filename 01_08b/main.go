package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "friends.json"

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

var hasHeardTheGossip = make(map[string]bool)

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	hasHeardTheGossip[f.ID] = true
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections
type Friends struct {
	fmap map[string]Friend
}

// getFriend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.fmap[id]
}

// getRandomFriend returns an random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.fmap)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

func getFriendIdsHasNotHeard(friends []string) []string {
	hasNotHead := []string{}
	for _, friendId := range friends {
		if !hasHeardTheGossip[friendId] {
			hasNotHead = append(hasNotHead, friendId)
		}
	}

	return hasNotHead
}

// spreadGossip ensures that all the friends in the map have heard the news
func spreadGossip(root Friend, friends Friends) {
	if !hasHeardTheGossip[root.ID] {
		root.hearGossip()
	}
	for _, friendId := range getFriendIdsHasNotHeard(root.Friends) {
		friend := friends.getFriend(friendId)
		spreadGossip(friend, friends)
	}
}

func main() {
	friends := importData()
	root := friends.getRandomFriend()
	root.hearGossip()
	spreadGossip(root, friends)
}

// importData reads the input data from file and
// creates the friends map.
func importData() Friends {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []Friend
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	fm := make(map[string]Friend, len(data))
	for _, d := range data {
		fm[d.ID] = d
	}

	return Friends{
		fmap: fm,
	}
}
