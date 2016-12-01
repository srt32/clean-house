package main

import (
	"log"
	"os"

	"github.com/croaky/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func newClient() *twitter.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

func getTweets(client *twitter.Client, maxID *int64) ([]twitter.Tweet, error) {
	params := &twitter.UserTimelineParams{
		Count:           200,
		IncludeRetweets: twitter.Bool(true),
	}
	if maxID != nil {
		params.MaxID = *maxID
	}

	tweets, _, err := client.Timelines.UserTimeline(params)
	return tweets, err
}

func deleteRetweets(client *twitter.Client, maxID *int64) {
	tweets, _ := getTweets(client, maxID)

	if len(tweets) > 0 {
		var newMaxID int64

		for _, tweet := range tweets {
			if tweet.Retweeted {
				client.Statuses.Destroy(tweet.ID, nil)
			}

			newMaxID = tweet.ID
		}

		deleteRetweets(client, &newMaxID)
	}
}

func getFavorites(client *twitter.Client, maxID *int64) ([]twitter.Tweet, error) {
	params := &twitter.FavoriteListParams{
		Count: 200,
	}
	if maxID != nil {
		params.MaxID = *maxID
	}

	favorites, _, err := client.Favorites.List(params)
	return favorites, err
}

func deleteFavorites(client *twitter.Client, maxID *int64) {
	favorites, _ := getFavorites(client, maxID)

	if len(favorites) > 0 {
		var newMaxID int64

		for _, favorite := range favorites {
			client.Favorites.Destroy(&twitter.FavoriteDestroyParams{ID: favorite.ID})
			newMaxID = favorite.ID
		}

		deleteFavorites(client, &newMaxID)
	}
}

func deleteFriendships(client *twitter.Client) {
	ids, _, err := client.Friends.IDs(&twitter.FriendIDParams{Count: 5000})

	if err != nil {
		for _, id := range ids.IDs {
			client.Friendships.Destroy(&twitter.FriendshipDestroyParams{UserID: id})
		}
	}
}

func main() {
	client := newClient()
	deleteRetweets(client, nil)
	deleteFavorites(client, nil)
	deleteFriendships(client)
}
