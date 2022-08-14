package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

//function that builds the query for the scrapper
func build_query(user string) string {
	return fmt.Sprintf("from:%s url:spotify", user)
}

//function that scraps the spotify urls from the user's tweets
func scrap_urls(user string, ctx context.Context) []string {
	scraper := twitterscraper.New()
	query := build_query(user)
	var urls []string
	// I'm sorry if you've posted more than 1000 spotify urls in tweets, everything will be better I'm sure.
	for tweet := range scraper.SearchTweets(ctx, query, 1000) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		urls = append(urls, tweet.URLs...)
	}
	return urls
}

//function that remove query from the url array
func remove_query(urls []string) []string {
	var no_query_urls []string
	for _, url := range urls {
		no_query_urls = append(no_query_urls, strings.Split(url, "?")[0])
	}
	return no_query_urls
}

//function that remove duplicates from the urls
func remove_duplicates(s []string) []string {
	result := make(map[string]bool)
	var unique_urls []string
	for _, str := range s {
		if _, ok := result[str]; !ok {
			result[str] = true
			unique_urls = append(unique_urls, str)
		}
	}
	return unique_urls
}

//function that segregate the urls into different categories tracks and playlists
func segregate_urls(urls []string) ([]string, []string, []string) {
	var tracks []string
	var playlists []string
	var other []string
	for _, url := range urls {
		if strings.Contains(url, "track") {
			tracks = append(tracks, url)
		} else if strings.Contains(url, "playlist") {
			playlists = append(playlists, url)
		} else {
			other = append(other, url)
		}
	}
	return tracks, playlists, other
}

//function that prints the urls in the terminal
func print_urls(urls []string, category string) {
	fmt.Println(category)
	for _, url := range urls {
		fmt.Println(url)
	}
}

func main() {
	user := os.Args[1]
	ctx := context.Background()
	urls := scrap_urls(user, ctx)
	urls = remove_query(urls)
	urls = remove_duplicates(urls)
	tracks, playlists, _ := segregate_urls(urls)
	print_urls(tracks, "Tracks")
	print_urls(playlists, "Playlists")
}
