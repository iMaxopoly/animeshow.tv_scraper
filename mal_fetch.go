package main

import (
	"log"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func mal_fetch_image(path string, name string){
	// don't worry about errors
	response, e := http.Get(path)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	// open a file for writing
	file, err := os.Create(fmt.Sprintf("images/%s.jpg", name))
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func mal_fetch_data(anime_name string) (structmal_api_anime, error) {
	client := &http.Client{}
	if anime_name == "Chaos Head" {
		anime_name = "Chaos;Head"
	} else if anime_name == "JoJo's Bizarre Adventure: Stardust Crusaders Season 2" {
		anime_name = "JoJo no Kimyou na Bouken: Stardust Crusaders 2nd Season"
	} else if anime_name == "Cardfight!! Vanguard G Stride Gate-hen" {
		anime_name = "Cardfight!! Vanguard Third Season"
	} else if anime_name == "Uta no Prince-sama Maji Love 2000%" {
		anime_name = "Uta no Prince Sama 2"
	} else if anime_name == "Osomatsu-san Year-End Special" {
		anime_name = "Osomatsu-san Special"
	} else if anime_name == "Kaitou Joker Season 2" {
		anime_name = "Kaitou Joker 2nd Season"
	} else if anime_name == "Saki: Nationals" {
		anime_name = "Saki: The Nationals"
	} else if anime_name == "Silver Spoon Season 2" {
		anime_name = "Silver Spoon 2nd Season"
	} else if anime_name == "Aldnoah Zero" {
		anime_name = "Aldnoah.Zero"
	} else if anime_name == "Aldnoah Zero Season 2" {
		anime_name = "Aldnoah.Zero 2nd Season"
	} else if anime_name == "Fate Zero Second Season" {
		anime_name = "Fate/Zero 2nd Season"
	} else if anime_name == "Re Hamatora" {
		anime_name = "Re: Hamatora: Season 2"
	} else if anime_name == "Futsuu no Joshikousei ga Locodol Yatte Mita Special" {
		anime_name = "Futsuu no Joshikousei ga [Locodol] Yattemita.: Nagarekawa, Annai Shitemita."
	} else if anime_name == "Kaitou Joker Season 3" {
		anime_name = "Kaitou Joker 3rd Season"
	} else if anime_name == "Selector Infected WIXOSS Specials" {
		anime_name = "Selector Infected WIXOSS: Midoriko-san to Piruruku-tan"
	} else if anime_name == "Gintama 2015" {
		anime_name = "Gintama' (2015)"
	} else if anime_name == "Ore, Twintails ni Narimasu." {
		anime_name = "Gonna be the Twin-Tail!!"
	} else if anime_name == "Fate/stay night: Unlimited Blade Works (TV) Season 2" {
		anime_name = "Fate/stay night: Unlimited Blade Works 2nd Season"
	} else if anime_name == "Futsuu no Joshikousei ga Locodol Yatte Mita" {
		anime_name = "Futsuu no Joshikousei ga [Locodol] Yattemita."
	} else if anime_name == "Active Raid: Kidou Kyoushuushitsu Dai Hakkei" {
		anime_name = "Active Raid: Kidou Kyoushuushitsu Dai Hachi Gakari"
	} else if anime_name == "Fate/kaleid liner Prisma Illya 2wei!" {
		anime_name = "Fate/kaleid liner Prisma☆Illya 2wei!"
	} else if anime_name == "Sailor Moon: Crystal Season III" {
		anime_name = "Bishoujo Senshi Sailor Moon Crystal Season III"
	} else if anime_name == "Ace of Diamond Season 2" {
		anime_name = "Diamond no Ace: Second Season"
	} else if anime_name == "Norn9: Norn+Nonet" {
		anime_name = "Norn9"
	} else if anime_name == "Ai Tenchi Muyo!" {
		anime_name = "Ai Tenchi Muyou!"
	} else if anime_name == "Shingeki no Bahamut - Genesis" {
		anime_name = "Shingeki no Bahamut: Genesis"
	} else if anime_name == "Futsuu no Joshikousei ga Locodol Yatte Mita OVA" {
		anime_name = "Futsuu no Joshikousei ga [Locodol] Yattemita. OVA"
	} else if anime_name == "Fate Zero" {
		anime_name = "Fate/Zero"
	} else if anime_name == "Hunter X Hunter 2011" {
		anime_name = "Hunter x Hunter (2011)"
	} else if anime_name == "Haikyu!! Second Season" {
		anime_name = "Haikyuu!! Second Season"
	} else if anime_name == "Sailor Moon: Crystal" {
		anime_name = "Bishoujo Senshi Sailor Moon Crystal"
	} else if anime_name == "Fate/stay night: Unlimited Blade Works (TV)" {
		anime_name = "Fate/stay night: Unlimited Blade Works"
	} else if anime_name == "Sabagebu! - Survival Game Club" {
		anime_name = "Sabagebu!"
	} else if anime_name == "Nisekoi Season 2" {
		anime_name = "Nisekoi: False Love"
	} else if anime_name == "Uta no Prince-sama: Maji Love Revolutions" {
		anime_name = "Uta no Prince Sama Revolutions"
	} else if anime_name == "Fairy Tail 2014" {
		anime_name = "Fairy Tail (2014)"
	} else if anime_name == "Oregairu Season 2" {
		anime_name = "My Teen Romantic Comedy SNAFU TOO!"
	} else if anime_name == "Rozen Maiden 2013" {
		anime_name = "Rozen Maiden (2013)"
	} else if anime_name == "Sengoku Basara Judge End" {
		anime_name = "Sengoku Basara: Judge End"
	} else if anime_name == "Baby Steps Season 2" {
		anime_name = "Baby Steps 2nd Season"
	} else if anime_name == "Tokyo Ghoul Season 2" {
		anime_name = "Tokyo Ghoul 2nd Season"
	} else if anime_name == "Luck & Logic" {
		anime_name = "Luck and Logic"
	} else if anime_name == "Kore wa Zombie Desu ka? of the Dead" {
		anime_name = "Is this A Zombie? of the Dead"
	} else if anime_name == "Toriko" {
		anime_name = "Toriko (2011)"
	} else if anime_name == "OreImo Season 2" {
		anime_name = "Oreimo 2"
	} else if anime_name == "Shin Strange+" {
		anime_name = "Strange Plus Second Season"
	} else if anime_name == "Ai Mai Mi" {
		anime_name = "Choboraunyopomi Gekijou Ai Mai Mii"
	} else if anime_name == "Himegoto" {
		anime_name = "Secret Princess Himegoto"
	} else if anime_name == "Maken-Ki! Two" {
		anime_name = "Maken-Ki! Second Season"
	} else if anime_name == "Highschool DxD New" {
		anime_name = "High School DxD New"
	} else if anime_name == "Strange+" {
		anime_name = "Strange Plus"
	} else if anime_name == "Shoujo-tachi wa Kouya wo Mezasu" {
		anime_name = "Girls Beyond the Wasteland"
	} else if anime_name == "Shounen Maid" {
		anime_name = "Boy Maid"
	}
	anime_name = strings.Replace(anime_name, " ", "+", -1)
	if _DEBUG {
		log.Println(anime_name)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://myanimelist.net/api/anime/search.json?q=%s", anime_name), nil)
	req.SetBasicAuth(mal_username, mal_password)
	resp, err := client.Do(req)
	if err != nil {
		return structmal_api_anime{}, err
		//log.Fatal("mal_fetch_data 1", anime_name, err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return structmal_api_anime{}, err
		//log.Fatal("mal_fetch_data 2", anime_name, err)
	}

	xobj := structmal_api_anime{}
	err = xml.Unmarshal(bodyText, &xobj)
	if err != nil {
		return structmal_api_anime{}, err
		//log.Fatal("mal_fetch_data 3", anime_name, err)
	}
	return xobj, nil
}

func mal_fetch_episodelist(mal_id string, mal_title string) ([]string) {
	//*[@id="content"]/table/tbody/tr/td[2]/div[1]/div[2]/table/tbody/tr/td/table[1]/tbody/tr[2]/td[3]/a
	//*[@id="content"]/table/tbody/tr/td[2]/div[1]/div[2]/table/tbody/tr/td/table[1]/tbody/tr[3]/td[3]/a
	mal_title = strings.Replace(mal_title, "%", "", -1)
	doc, err := goquery.NewDocument(fmt.Sprintf("http://myanimelist.net/anime/%s/%s/episode", mal_id, mal_title))
	if err != nil {
		log.Fatal("mal_fetch_episodelist 1", err)
	}

	var result []string
	episodes_list :=
		doc.Find("#content").
			Find("table").
			Find("tbody").
			Find("tr").
			Find("td").Next().
			Find("div").
			Find("div").Next().
			Find("table").
			Find("tbody").
			Find("tr").Find("td").Find(".mt8.episode_list.js-watch-episode-list.ascend").Find("tbody").Find("tr.episode-list-data")

	if _DEBUG {
		log.Println(len(episodes_list.Nodes))
	}

	for node := range episodes_list.Nodes {
		container_doc := episodes_list.Eq(node).Find("td").Find(".fl-l.fw-b ")

		if _DEBUG {
			log.Println(container_doc.Text())
		}
		result = append(result, container_doc.Text())
	}
	return result
}

func mal_fetch_trailer(mal_id string) (string) {
	// //*[@id="content"]/table/tbody/tr/td[2]/div[1]/table/tbody/tr[1]/td/div[1]/div[1]/div[2]/div[1]/a
	doc, err := goquery.NewDocument(fmt.Sprintf("http://myanimelist.net/anime/%s/", mal_id))
	if err != nil {
		log.Fatal("mal_fetch_trailer 1", err)
	}

	trailer, exists :=
		doc.Find("#content").
			Find("table").
			Find("tbody").
			Find("tr").
			Find("td").Next().
			Find("div").
			Find("table").
			Find("tbody").
			Find("tr").
			Find("td").
			Find("div").
			Find("div").
			Find("div").Next().
			Find("div").
			Find("a").Attr("href")

	if !exists {
		return "Nil"
	}
	return trailer
}