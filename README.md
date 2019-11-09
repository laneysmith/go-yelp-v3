# go-yelp-v3 (WIP)

This is a fork of [Justin Beckwith](https://github.com/JustinBeckwith/)'s [go-yelp](https://github.com/JustinBeckwith/go-yelp), adapted to use the [v3 Yelp Fusion API](https://www.yelp.com/developers/documentation/v3).

### TODO:

- [x] update AuthOptions to use v3 API keys
- [x] adapt DoSearch
- [ ] adapt DoSimpleSearch
- [ ] tests

---

go-yelp is a #golang wrapper for the Yelp REST API. It lets you do all kinds of interesting things like searching for businesses, getting user comments and ratings, and handling common errors. The library is written Go.

## Getting Started

To install go-yelp-v3, just use the `go get` command:

```sh
go get github.com/laneysmith/go-yelp-v3/yelp
```

When you're ready to start using the API, import the reference:

```go
import "github.com/laneysmith/go-yelp-v3/yelp"
```

### Authentication

All searches are performed through a client. To create a new client, you need provide a set of access keys necessary to use the v3 Yelp Fusion API. You can sign up for a Yelp developer account, and access your keys here:

[Yelp | Manage Keys](https://www.yelp.com/developers/v3/manage_app)

Keep these keys safe! There are a variety of ways to store them. I chose to store them in a config.json file which is not checked into the repository. To run the tests, you can create your own `config.json` file:

```json
{
  "YelpAPIKey": "YELP_API_KEY"
}
```

### The Search API

The simple search API enables searching for businesses with a term and a location (ex: coffee, Seattle). After you have your keys, create a client, and make a simple query:

```go
import "github.com/laneysmith/go-yelp-v3/yelp"

client := yelp.New(options, nil)
result, err := client.DoSimpleSearch("coffee", "seattle")
```

For more complex searches, the `DoSearch` method allows for searching based on a combination of general search criteria, and advanced location options:

```go
// Build an advanced set of search criteria that include
// general options, and location specific options.
options := SearchOptions{
	GeneralOptions: &GeneralOptions{
		Term: "food",
	},
	LocationOptions: &LocationOptions{
		"bellevue",
		&CoordinateOptions{
			Latitude:  null.FloatFrom(37.788022),
			Longitude: null.FloatFrom(-122.399797),
		},
	},
}

// Perform the search using the search options
result, err := client.DoSearch(options)
```

### The Business API

To directly search for a business by name, use the `client.GetBusiness(...)` method on the client:

```go
client := yelp.New(options, nil)
result, err := client.GetBusiness("yelp-san-francisco")
```

## License

This library is distributed under the [MIT License](http://opensource.org/licenses/MIT) found in the LICENSE file.

## Questions?

Feel free to submit an issue on the repository, or find me at [@JustinBeckwith](http://twitter.com/JustinBeckwith)
