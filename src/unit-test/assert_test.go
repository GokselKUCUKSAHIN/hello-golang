package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	CampaignId int `json:"campaignId"`
	Listings   []struct {
		Id     string   `json:"id"`
		Tags   []string `json:"tags"`
		Stamps []struct {
			Type        string  `json:"type"`
			ImageUrl    string  `json:"imageUrl"`
			Position    string  `json:"position"`
			AspectRatio float64 `json:"aspectRatio"`
			Priority    int     `json:"priority"`
		} `json:"stamps"`
		SearchableTags []interface{} `json:"searchableTags"`
	} `json:"listings"`
}

type Model struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Friends *[]*Model `json:"friends"`
}

const resp = `[{"campaignId":600066,"listings":[{"id":"08719d8ac615d1ba369592121557ada9","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]}]},{"campaignId":600077,"listings":[{"id":"d5c84c54826be8004cecf5f81fc8b1a5","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]},{"id":"5176d2023369fec51f970607dcc9a485","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]},{"id":"1d77aba6d8c67f6a64a7b240feb29a6d","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]},{"id":"08719d8ac615d1ba369592121557ada9","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]},{"id":"dfccd07478f107e3f2c2d59e84b09d1f","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]},{"id":"32481994f403f0fbbb21c338e8aad25e","tags":["kategori_encoksatanlar","bestseller"],"stamps":[{"type":"TypeA","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/marketing/datascience/automation/2020/12/9/EnCokSatan_202012091129.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeA","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800},{"type":"TypeB","imageUrl":"/stamp/en-bestselling.png","position":"upper-left","aspectRatio":0.25,"priority":800}],"searchableTags":[]}]}]`

func deepEqual(x, y *[]*Model) bool {
	if len(*x) != len(*y) {
		return false
	}
	for i := range *x {
		if !deepEqualPerson((*x)[i], (*y)[i]) {
			return false
		}
	}
	return true
}

// deepEqualPerson compares two Person structs for deep equality.
func deepEqualPerson(x, y *Model) bool {
	if x.Name != y.Name || x.Surname != y.Surname {
		return false
	}
	return deepEqual(x.Friends, y.Friends)
}

func TestJsonUnmarshal(t *testing.T) {
	var expected []TestModel
	expectedErr := json.Unmarshal([]byte(resp), &expected)
	assert.Nil(t, expectedErr)
	fmt.Println("~~~~~~~")
	const myJson1 = `[{"name":"John","surname":"Doe","friends":[{"name":"Jane","surname":"Doe","friends":[]},{"name":"Bob","surname":"Smith","friends":[{"name":"Alice","surname":"Smith","friends":[]}]}]},{"name":"Jane","surname":"Doe","friends":[{"name":"John","surname":"Doe","friends":[]},{"name":"Alice","surname":"Smith","friends":[]}]},{"name":"Bob","surname":"Smith","friends":[{"name":"Alice","surname":"Smith","friends":[]},{"name":"John","surname":"Doe","friends":[{"name":"Jane","surname":"Doe","friends":[]}]}]},{"name":"Alice","surname":"Smith","friends":[{"name":"Bob","surname":"Smith","friends":[]},{"name":"Jane","surname":"Doe","friends":[]}]}]`
	// const myJson2 = `[{"name":"Alice","surname":"Smith","friends":[{"name":"Jane","surname":"Doe","friends":[]},{"name":"Bob","surname":"Smith","friends":[]}]},{"name":"John","surname":"Doe","friends":[{"name":"Jane","surname":"Doe","friends":[]},{"name":"Bob","surname":"Smith","friends":[{"name":"Alice","surname":"Smith","friends":[]}]}]},{"name":"Jane","surname":"Doe","friends":[{"name":"John","surname":"Doe","friends":[]},{"name":"Alice","surname":"Smith","friends":[]}]},{"name":"Bob","surname":"Smith","friends":[{"name":"Alice","surname":"Smith","friends":[]},{"name":"John","surname":"Doe","friends":[{"name":"Jane","surname":"Doe","friends":[]}]}]}]`
	var modelA, modelB *[]*Model
	modelErrA := json.Unmarshal([]byte(myJson1), &modelA)
	modelErrB := json.Unmarshal([]byte(myJson1), &modelB)
	// modelErrB := json.Unmarshal([]byte(myJson2), &modelB)
	assert.Nil(t, modelErrA)
	assert.Nil(t, modelErrB)
	fmt.Println(deepEqual(modelA, modelB))
}
