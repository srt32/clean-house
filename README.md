# Clean House

A script to:

* Unfollow all the Twitter users you've followed.
* Unlike all the tweets you've liked.
* Unretweet all the tweets you've retweeted.

## Setup

[Download your entire Twitter history][history].

[history]: https://twitter.com/settings/account#tweet_export

[Create a Twitter app][app].

[app]: https://apps.twitter.com/app/new

Create a `.env` file:

```
TWITTER_CONSUMER_KEY=get-from-twitter-app
TWITTER_CONSUMER_SECRET=get-from-twitter-app
TWITTER_ACCESS_TOKEN=get-from-twitter-app
TWITTER_ACCESS_SECRET=get-from-twitter-app
```

Run the script:

```
go run main.go
```
