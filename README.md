# Spotify links twitter scrapper (SLTS)

<p align="center">
  <img src="https://github.com/victormagalhaess/spotify-links-twitter-scrapper/blob/main/SLTS.png?raw=true" width="200" alt="SLTS Logo">
</p>

Have you ever thought "How many spotify links I've shared on twitter in the past?. Well, not many people think about it, but I kinda feel like checking anyway!

SLTS is a really simple implementation over [@n0madic](github.com/n0madic)'s [twitter scrapper](github.com/n0madic/twitter-scraper)! It aims to find all tracks and playlists that you've shared in the past.

Due to no reason other than a fear of beeing blocked from twitter's api, the search is limited to match 1000 tweets with spotify links. However, it can be extended to as much as anyone please, just changing the limit on line 23 at main.go.

To run, you must pass an user @ as parameter to the program:

```sh
go run main.go TwitterBrasil
```

Make sure to install all dependencies with ```go get``` before running! 

# Credits

The SLTS logo was made using Freepik flaticon free icons:

<a href="https://www.flaticon.com/free-icons/spotify" title="spotify icons">Spotify icons created by Freepik - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/twitter" title="twitter icons">Twitter icons created by Pixel perfect - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/chain" title="chain icons">Chain icons created by Smashicons - Flaticon</a>

And also makes use of Go Gopher:

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)