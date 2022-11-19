package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Junkes887/translate/model"
	"github.com/go-redis/redis"
)

type Client struct {
	DB *redis.Client
}

func (client Client) Find(id string) model.Page {
	var page model.Page

	idRedis := client.DB.Get(id)

	b, _ := idRedis.Bytes()
	json.Unmarshal(b, &page)

	return page
}

func (client Client) Save(id string, page model.Page) {
	p, err := json.Marshal(page)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Dado salvo no redis ", id)
	client.DB.Set(id, p, 0)
}
