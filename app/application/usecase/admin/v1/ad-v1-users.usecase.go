package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	redisService "github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/redis/services"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/olivere/elastic/v7"
)

// UsersUseCase interface
type UsersUseCase interface {
	Create(ctx *fiber.Ctx, Body *models.User) (err error)
	Find(ctx *fiber.Ctx, Users []models.User) (data interface{}, count uint, err error)
	FindOne(ctx *fiber.Ctx, Users *models.User) (data interface{}, err error)
}

// NewUsersUseCase Instantiate the UseCase
func NewUsersUseCase() UsersUseCase {
	return &usersUseCase{
		service:           services.NewUsersService(),
		serviceRole:       services.NewRoleService(),
		serviceRedisUsers: redisService.NewUsersServiceRedis(),
	}
}

type usersUseCase struct {
	service           services.UsersService
	serviceRole       services.RoleService
	serviceRedisUsers redisService.UsersServiceRedis
}

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	}
}`

func ELK(ctx *fiber.Ctx) {
	// Starting with elastic.v5, you must pass a context to execute each service

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient(elastic.SetURL("http://es01:9200"), elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://es01:9200").Do(ctx.Context())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://es01:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx.Context())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}

	// // Index a tweet (using JSON serialization)
	// tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	// put1, err := client.Index().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	Id("1").
	// 	BodyJson(tweet1).
	// 	Do(ctx.Context())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// // Index a second tweet (by string)
	// tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	// put2, err := client.Index().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	Id("2").
	// 	BodyString(tweet2).
	// 	Do(ctx.Context())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// // Get tweet with specified ID
	// get1, err := client.Get().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	Id("1").
	// 	Do(ctx.Context())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// if get1.Found {
	// 	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	// }

	// // Flush to make sure the documents got written.
	// _, err = client.Flush().Index("twitter").Do(ctx.Context())
	// if err != nil {
	// 	panic(err)
	// }

	// // Search with a term query
	// termQuery := elastic.NewTermQuery("user", "olivere")
	// searchResult, err := client.Search().
	// 	Index("twitter").   // search in index "twitter"
	// 	Query(termQuery).   // specify the query
	// 	Sort("user", true). // sort by "user" field, ascending
	// 	From(0).Size(10).   // take documents 0-9
	// 	Pretty(true).       // pretty print request and response JSON
	// 	Do(ctx.Context())             // execute
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }

	// // searchResult is of type SearchResult and returns hits, suggestions,
	// // and all kinds of other information from Elasticsearch.
	// fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// // Each is a convenience function that iterates over hits in a search result.
	// // It makes sure you don't need to check for nil values in the response.
	// // However, it ignores errors in serialization. If you want full control
	// // over iterating the hits, see below.
	// var ttyp Tweet
	// for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	// 	if t, ok := item.(Tweet); ok {
	// 		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	// 	}
	// }
	// // TotalHits is another convenience function that works even when something goes wrong.
	// fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// // Here's how you iterate through results with full control over each step.
	// if searchResult.Hits.TotalHits > 0 {
	// 	fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

	// 	// Iterate through results
	// 	for _, hit := range searchResult.Hits.Hits {
	// 		// hit.Index contains the name of the index

	// 		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	// 		var t Tweet
	// 		err := json.Unmarshal(*hit.Source, &t)
	// 		if err != nil {
	// 			// Deserialization failed
	// 		}

	// 		// Work with tweet
	// 		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	// 	}
	// } else {
	// 	// No hits
	// 	fmt.Print("Found no tweets\n")
	// }

	// // Update a tweet by the update API of Elasticsearch.
	// // We just increment the number of retweets.
	// update, err := client.Update().Index("twitter").Type("tweet").Id("1").
	// 	Script(elastic.NewScriptInline("ctx.Context()._source.retweets += params.num").Lang("painless").Param("num", 1)).
	// 	Upsert(map[string]interface{}{"retweets": 0}).
	// 	Do(ctx.Context())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)

	// // ...

	// // Delete an index.
	// deleteIndex, err := client.DeleteIndex("twitter").Do(ctx.Context())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// if !deleteIndex.Acknowledged {
	// 	// Not acknowledged
	// }
}

// Create usecase Users
func (fn *usersUseCase) Create(ctx *fiber.Ctx, Body *models.User) (err error) {
	Body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.Version = 1
	if response := fn.service.Create(Body); response.Error != nil {
		return response.Error
	}

	// Match role to user
	if Body.RoleID != 0 {
		Role := new(models.Role)
		res := fn.serviceRole.FindOne(Role, Body.RoleID)
		if res.Error != nil {
			return res.Error
		}

		Body.Role, err = json.Marshal(Role)
		if err != nil {
			return err
		}
	}

	return err
}

// Find usecase Users
func (fn *usersUseCase) Find(ctx *fiber.Ctx, Users []models.User) (data interface{}, count uint, err error) {
	version, count, err := fn.service.GetVersionCount()
	if err != nil {
		return
	}

	// get list with redis cache
	strVersion := strconv.Itoa(int(version))
	cache := helper.GetCache(ctx, strVersion)
	if cache.Err() != nil {
		if cache.Err().Error() == "redis: nil" {
			responseUsers, err2 := fn.service.Find(ctx, Users)
			if err2 != nil {
				return nil, 0, err2
			}

			err2 = helper.SetCache(ctx, strVersion, responseUsers)
			if err2 != nil {
				return nil, 0, err2
			}

			log.Println("dari postgres loo")
			return responseUsers, count, nil
		}
		return
	}

	byte, err := cache.Bytes()
	err = jsoniter.Unmarshal(byte, &Users)
	if err != nil {
		return nil, 0, err
	}

	log.Println("dari redis loo")
	return Users, count, nil
}

// Find usecase Users
func (fn *usersUseCase) FindOne(ctx *fiber.Ctx, Users *models.User) (data interface{}, err error) {
	// ELK(ctx)
	id := ctx.Params("id")
	if response := fn.service.FindOne(Users, id); response.Error != nil {
		err = response.Error
		return
	}

	// Match role to user
	if Users.RoleID != 0 {
		Role := new(models.Role)
		res := fn.serviceRole.FindOne(Role, Users.RoleID)
		if res.Error != nil {
			return nil, res.Error
		}

		Users.Role, err = json.Marshal(Role)
		if err != nil {
			return nil, err
		}
	}

	return Users, nil
}
