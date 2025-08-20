// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Music
type Myfavor_music struct {
	ID_music string `json:"id"`
	Name     string `json:"name"`
	Singer   string `json:"singer"`
}

var myfavor_musics = []Myfavor_music{
	{ID_music: "152059", Name: "Propose[プロポーズ]", Singer: "9Lena(Origin) และ Supeacha(Thai Cover)"},
	{ID_music: "434633", Name: "The Flame of Love[愛が灯る]", Singer: "Rokudenashi [ロクデナシ]"},
	{ID_music: "642429", Name: "Hated by Life Itself[命に嫌われている](Inochi ni Kirawareteiru)", Singer: "Kanzaki Iori(Origin) และ BEARISS (Thai cover)"},
	{ID_music: "658158", Name: "Nightglow", Singer: "Tanya Chua จากเกม Honkai impact"},
}

func getMyfavor_music(c *gin.Context) {
	ID_filled := c.Query("id")
	if ID_filled != "" {
		var filtered2 []Myfavor_music
		for _, myfavor_music := range myfavor_musics {
			if fmt.Sprint(myfavor_music.ID_music) == ID_filled {
				filtered2 = append(filtered2, myfavor_music)
			}
		}
		c.JSON(http.StatusOK, filtered2)
		return
	}
	//return
	c.JSON(http.StatusOK, myfavor_musics)
}

// struct Yuribook Bruhhh
type Yuri_Book struct {
	ID_Novel string  `json:"id"`
	Name     string  `json:"name"`
	Shipping string  `json:"ship"`
	Price    float64 `json:"price"`
}

var yuris = []Yuri_Book{
	{ID_Novel: "517253", Name: "[Fic GuP]บทรักแห่งรถถัง(Yuri)DxK", Shipping: "Darjeeling x Kay", Price: 0.00},
	{ID_Novel: "507355", Name: "It's Not 'Ran' but 'Master'", Shipping: "Moca x Ran", Price: 1500.00},
	{ID_Novel: "136606", Name: "Kiss You a Lot", Shipping: "Arisa x Kasumiy", Price: 1900.00},
	{ID_Novel: "870200", Name: "Heya de Kimi, Kakushinhan", Shipping: "Kokoro x Misaki", Price: 2000.00},
	{ID_Novel: "558807", Name: "Citrus", Shipping: "Yuzu x Mei", Price: 999999.00},
	{ID_Novel: "040225", Name: "Honeymoon Baby", Shipping: "Maki x Nico", Price: 50000000000000.00},
}

func getYuri_book(c *gin.Context) {
	ID_filled := c.Query("id")
	if ID_filled != "" {
		var filtered2 []Yuri_Book
		for _, yuri := range yuris {
			if fmt.Sprint(yuri.ID_Novel) == ID_filled {
				filtered2 = append(filtered2, yuri)
			}
		}
		c.JSON(http.StatusOK, filtered2)
		return
	}
	//return
	c.JSON(http.StatusOK, yuris)
}
func main() {
	r := gin.Default()
	r.GET("/yuri", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "yuri book"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/myfavor", getMyfavor_music)
		api.GET("/yuri", getYuri_book)
	}

	r.Run(":8080")
}
