package main

import (
	"strconv"
	"io/ioutil"
	"encoding/json"
	"log"
	"strings"

	"github.com/asdine/storm"
)

func scrape() {
	total_dump := []structmain__anime{}

	/* Open connection to boltdb db */
	db, err := storm.Open("animedom.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/* Scraping Anime List Page */
	t_anime_list := animeshow_animelist_page(url_anime_list)

	count := 0
	for t_anime_name, t_anime_link := range t_anime_list {
		count++
		if t_anime_name == "Psycho-Pass New Edit Version" {
			continue
		}
		/* Scrape individual anime episode listings page */
		meta_data := animeshow_episodelist_page(t_anime_link)
		/* XML fetch MAL API content for given anime */
		// Replacing spaces with '+' to query MAL
		mal_anime_xml_data, err := mal_fetch_data(t_anime_name)
		if err != nil {
			log.Fatal("30", err)
		}
		/* Temporary array based xml MAL fix */
		var entry_num int
		if t_anime_name == "Jinsei" {
			entry_num = 1
		} else if t_anime_name == "Beelzebub" {
			entry_num = 1
		} else if t_anime_name == "Million Doll" {
			entry_num = 1
		} else if t_anime_name == "Ajin" {
			entry_num = 25
		} else if t_anime_name == "Sengoku Musou" {
			entry_num = 1
		} else if t_anime_name == "Amnesia" {
			entry_num = 3
		} else if t_anime_name == "Hundred" {
			entry_num = 9
		} else if t_anime_name == "Itoshi no Muco" {
			entry_num = 3
		} else if t_anime_name == "Another" {
			entry_num = 20
		} else if t_anime_name == "Onigiri" {
			entry_num = 1
		} else if t_anime_name == "Charlotte" {
			entry_num = 1
		} else {
			entry_num = 0
		}

		/* Scrape MAL Trailer and Episode-names */
		if _DEBUG {
			log.Println(entry_num)
		}
		mal_anime_trailer := mal_fetch_trailer(mal_anime_xml_data.Entries[entry_num].ID)
		mal_anime_episode_names := mal_fetch_episodelist(mal_anime_xml_data.Entries[entry_num].ID, mal_anime_xml_data.Entries[entry_num].Title)

		anime := structmain__anime{}

		// Fetching score & ID & image from MAL
		anime.Score = mal_anime_xml_data.Entries[entry_num].Score
		anime.MAL_ID = mal_anime_xml_data.Entries[entry_num].ID
		anime.Image = strings.Replace(mal_anime_xml_data.Entries[entry_num].Image, ".jpg", "l.jpg", 1)
		anime.AnimeShow_Name = t_anime_name
		anime.SynonymNames = mal_anime_xml_data.Entries[entry_num].SynonymNames
		anime.AltName = meta_data.Altname
		anime.Genre = meta_data.Genre
		anime.AnimeShow_Description = meta_data.Description
		anime.Status = meta_data.Status
		anime.Type = meta_data.Type
		anime.Year = meta_data.Year
		anime.Trailer = mal_anime_trailer
		anime.MAL_English = mal_anime_xml_data.Entries[entry_num].EnglishName
		anime.MAL_Title = mal_anime_xml_data.Entries[entry_num].Title
		anime.MAL_Description = mal_anime_xml_data.Entries[entry_num].Synopsis

		t_episode_list := []structmain__episode{}
		for t_episode_number, t_episode_link := range meta_data.Episodes {
			/* Scrape per episode link in episode listing page */
			t_mirrorlist_to_extract := []structscraper__mirror_to_extract{}
			t_mirrorlist := []structmain__mirror{}
			temp_mirrorlist_links_to_scrape, boolean := animeshow_mirrorlist_links_to_scrape(t_episode_link)
			if !boolean {
				continue
			}
			for _, t_mirror := range temp_mirrorlist_links_to_scrape {
				/* Store pages to extract mirrors from */
				t_mirrorlist_to_extract = append(t_mirrorlist_to_extract, t_mirror)
			}
			for i, t_link_to_extract_iframe := range t_mirrorlist_to_extract {
				/* Iterate over stored mirror list links to fetch iframe */
				t_iframe := animeshow_mirrorlist_iframe(t_link_to_extract_iframe.Link)
				t_mirrorlist = append(t_mirrorlist, structmain__mirror{
					Name:t_mirrorlist_to_extract[i].Name,
					Iframe:t_iframe,
					SubDub: t_link_to_extract_iframe.SubDub,
				})
			}
			/* Assigning episode name */
			t_episode_list = append(t_episode_list, structmain__episode{
				Name: func() (string) {
					if t_episode_number >= len(mal_anime_episode_names) || t_episode_number < 0 {
						return anime.AnimeShow_Name + " Episode " + strconv.Itoa(t_episode_number + 1)
					}
					return mal_anime_episode_names[t_episode_number]
				}(), //anime.Name + " Episode " + strconv.Itoa(t_episode_number + 1),
				Mirrors:t_mirrorlist,
			})
		}
		/* Download Anime Image */
		mal_fetch_image(anime.Image, anime.MAL_ID)

		anime.EpisodeList = t_episode_list
		/* Append collected anime data to total_dump json object */
		total_dump = append(total_dump, anime)
		err = db.Save(anime)
		if err != nil {
			log.Fatal(err)
		}
	}
	jobj, err := json.Marshal(total_dump)
	if err != nil {
		log.Fatal(err)
	}

	/* Write JSON file to disk */
	err = ioutil.WriteFile("dump.json", jobj, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//func test_mal_names(){
//	t_anime_list := animeshow_animelist_page(url_anime_list)
//	var defected_anime_list string
//	var ok_anime_list string
//	for t_anime_name, _ := range t_anime_list {
//		/* XML fetch MAL API content for given anime */
//		// Replacing spaces with '+' to query MAL
//		mal_anime_xml_data, err := mal_fetch_data(t_anime_name)
//		if err != nil {
//			defected_anime_list += t_anime_name + "\n"
//			log.Println("[DEFECTIVE]", t_anime_name)
//			continue
//		}
//		ok_anime_list += "animeshow:\t" + t_anime_name + "\n" + "myanimelist:\t" + mal_anime_xml_data.Entries[0].Title + "\n\n"
//		log.Println("[PROCESSED]", t_anime_name, ":", mal_anime_xml_data.Entries[0].Title)
//	}
//	err := ioutil.WriteFile("defected_animes.txt", []byte(defected_anime_list), 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = ioutil.WriteFile("ok_animes.txt", []byte(ok_anime_list), 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	run_tests()
	scrape()
}