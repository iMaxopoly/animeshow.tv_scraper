package main

import (
	"fmt"
	"strings"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func animeshow_mirrorlist_links_to_scrape(url string) ([]structscraper__mirror_to_extract, bool) {
	if _DEBUG {
		log.Println("scrape__mirrorlist_links_to_scrape", url)
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		if strings.Contains(err.Error(), "no Host in request URL") {
			file, err := os.OpenFile("noaccess_mirrorlist_toscrape.txt", os.O_APPEND | os.O_WRONLY, 0600)
			if err != nil {
				log.Fatal("scrape__mirrorlist_links_to_scrape write_file_err1", err)
			}
			defer file.Close()

			if _, err = file.WriteString(url + "\n"); err != nil {
				log.Fatal("scrape__mirrorlist_links_to_scrape write_file_err2", err)
			}
			return []structscraper__mirror_to_extract{}, false
		} else {
			log.Fatal("scrape__mirrorlist_links_to_scrape 1", err)
		}

	}

	mirrorlist_doc := doc.Find(".container.main").
		Find(".row").
		Find(".col-lg-12.col-md-12.col-sm-12.col-xs-12").
		Find("#main").
		Find("#episode").
		Find("#episode_mirrors").
		Find(".row").
		Find(".col-lg-6.col-md-6.col-sm-12.col-xs-12")

	mirror_to_extract_struct := []structscraper__mirror_to_extract{}
	default_mirror := mirrorlist_doc.Find(".episode_mirrors_wraper.episode_mirrors_wraper_focus").Find(".episode_mirrors_name").Text()
	default_subdub := mirrorlist_doc.Find(".episode_mirrors_wraper.episode_mirrors_wraper_focus").Find(".episode_mirrors_type_sub").Text()
	mirror_to_extract_struct = append(mirror_to_extract_struct, structscraper__mirror_to_extract{Name: default_mirror, Link: url, SubDub:default_subdub})

	for mirror := range mirrorlist_doc.Nodes {
		link_path := mirrorlist_doc.Eq(mirror).Find("a")
		link, exists := link_path.Attr("href")
		if exists {
			mirror_name := link_path.Find(".episode_mirrors_wraper").Find(".episode_mirrors_name").Text()
			if _DEBUG {
				fmt.Println(link, mirror_name)
			}
			subdub := link_path.Find(".episode_mirrors_wraper").Find(".episode_mirrors_type_sub").Text()
			mirror_to_extract_struct = append(mirror_to_extract_struct, structscraper__mirror_to_extract{Name: mirror_name, Link: link, SubDub:subdub})
		}
	}
	return mirror_to_extract_struct, true
}

func animeshow_mirrorlist_iframe(url string) (string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("scrape__mirrorlist_iframe 1", err)
		return ""
	}

	iframe_object, err := doc.Find(".container.main").
		Find(".row").
		Find(".col-lg-12.col-md-12.col-sm-12.col-xs-12").
		Find("#main").
		Find("#episode").
		Find(".embed_wraper").
		Find(".embed.embeded").
		First().Html()

	if err != nil {
		log.Fatal("scrape__mirrorlist_iframe 2", err)
	}

	iframe_object = strings.TrimPrefix(iframe_object, " <nil>")
	if _DEBUG {
		fmt.Println(iframe_object)
	}
	return iframe_object

}

func animeshow_episodelist_page(url string) (structscraper__episodelist_page) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal("scrape__episodelist_page 1", err)
	}

	base_doc := doc.Find(".container.main").Find(".row").Find(".col-lg-9.col-md-9.col-sm-8.col-xs-12").Find("#main").
		Find("#anime")

	// Fetch Type
	meta_anime_doc := base_doc.Find(".row").Find(".col-lg-6.col-md-6.col-sm-12.col-xs-12.anime_info").
		Find(".col-lg-9.col-md-9.col-sm-9.col-xs-9")
	meta_anime_type := meta_anime_doc.Eq(0).Text()
	if _DEBUG {
		fmt.Println(url, "Type: ", meta_anime_type)
	}

	// Fetch Year
	meta_anime_year := meta_anime_doc.Eq(1).Text()
	if _DEBUG {
		fmt.Println(url, "Year: ", meta_anime_year)
	}

	// Fetch Status
	meta_anime_status := meta_anime_doc.Eq(2).Text()
	if _DEBUG {
		fmt.Println(url, "Status: ", meta_anime_status)
	}

	// Fetch Genre
	meta_anime_genre := strings.Split(meta_anime_doc.Eq(3).Text(), ", ")
	if _DEBUG {
		fmt.Println(url, "Genre: ", meta_anime_genre)
		fmt.Println(url, "Genre length: ", len(meta_anime_genre))
	}

	// Fetch Alternative Title
	meta_anime_altname := base_doc.Find(".row").Find(".col-lg-6.col-md-6.col-sm-12.col-xs-12.anime_info").
		Find(".alternative_titles").Find("ul").Find("li").Text()
	if _DEBUG {
		fmt.Println(url, "Alternate name: ", meta_anime_altname)
	}

	// Fetching description
	meta_anime_description, err := base_doc.Find(".anime_discription").Html()
	if err != nil {
		log.Fatal(err)
	}
	if _DEBUG {
		fmt.Println(url, "Description: ", meta_anime_description)
	}

	// Fetching episode links
	var meta_anime_episodes []string
	episodelist_doc := base_doc.Closest("#main").Find("#episodes_list").
		Find(".episodes_list_result")

	for episode := len(episodelist_doc.Nodes) - 1; episode >= 0; episode-- {
		link, _ := episodelist_doc.Eq(episode).Find("a").Attr("href")
		if _DEBUG {
			fmt.Println(link)
		}
		meta_anime_episodes = append(meta_anime_episodes, link)
	}

	data := structscraper__episodelist_page{
		Type      : meta_anime_type,
		Year       : meta_anime_year,
		Status     : meta_anime_status,
		Genre      : meta_anime_genre,
		Altname    : meta_anime_altname,
		Description: meta_anime_description,
		Episodes   : meta_anime_episodes,
	}

	return data
}

func animeshow_animelist_page(url string) (map[string]string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal("scrape__animelist_page 1", err)
	}

	animelist_map := make(map[string]string)

	animelist_doc := doc.Find(".container.main").
		Find(".row").
		Find(".col-lg-9.col-md-9.col-sm-8.col-xs-12").
		Find("#main").
		Find("#anime_list").
		Find(".anime_list_result")

	if _DEBUG {
		fmt.Println("anime_list_result len", len(animelist_doc.Nodes))
	}

	for node := range animelist_doc.Nodes {
		container_doc := animelist_doc.Eq(node).Find("ul").Find("li")

		if _DEBUG {
			fmt.Println("li len", len(container_doc.Nodes))
		}
		for subnode := range container_doc.Nodes {
			link, _ := container_doc.Eq(subnode).Find("a").Attr("href")
			name := container_doc.Eq(subnode).Find("a").Text()
			if _DEBUG {
				fmt.Println(name, link)
			}
			animelist_map[name] = link
		}
	}
	return animelist_map
}
