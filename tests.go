package main

import (
	"log"
	"os"
	"fmt"
)

//
// MAL specific testing
//

/* Test grabbing xml response from MAL */
func test__mal_fetch_data() (bool) {
	test_data := []string{"Trinity Seven OVA", "Aldnoah Zero", "Brynhildr in the Darkness", "Chaos Head"}
	test_answers := []string{"28285", "22729", "21431", "4975"}
	for i, v := range test_data {
		obj, err := mal_fetch_data(v)
		if err != nil {
			log.Fatal("test__mal_fetch_data failed", err)
			return false
		}
		if obj.Entries[0].ID != test_answers[i] {
			return false
		}
	}
	return true
}

/* Test scraping episodes list from MAL */
//func test__mal_fetch_episodelist() (bool) {
//	test_answers := []string{"Boot Up", "", ""}
//	test_data := map[string]string{"4975":"ChäoS;HEAd", "21431":"Gokukoku no Brynhildr",
//		"27775":"Plastic Memories", "21639":"Yu☆Gi☆Oh! Arc-V"}
//	for id, title := range test_data{
//		obj := mal_fetch_episodelist(id, title)
//		if obj =
//	}
//
//	return true
//}

/* Test scraping trailer from MAL */
func test__mal_fetch_trailer() (bool) {
	obj := mal_fetch_trailer("31804")
	if obj != "http://www.youtube.com/embed/CiqZUdyrwBk?enablejsapi=1&wmode=opaque&autoplay=1" {
		return false
	}
	return true
}

//
// animeshow.tv specific testing
//

/* Testing scraping entire anime list from animeshow.tv */
func test__animeshow_animelist_page() (bool) {
	obj := animeshow_animelist_page(url_anime_list)
	count := len(obj)
	if count <= 500 {
		return false
	}
	return true
}

/* Testing scraping from episodes listing page in animeshow.tv */
func test__animeshow_episodelist_page() (bool) {
	test_data := []string{"http://animeshow.tv/Hundred/", "http://animeshow.tv/Aldnoah-Zero/",
		"http://animeshow.tv/Btooom/", "http://animeshow.tv/Gate-Jieitai-Kanochi-nite-Kaku-Tatakaeri/"}
	test_answers := []string{"TV Series", "TV Series, 12 Episodes", "TV Series, 12 Episodes", "TV Series, 12 Episodes"}
	for i, v := range test_data {
		obj := animeshow_episodelist_page(v)
		if obj.Type != test_answers[i] {
			return false
		}
	}
	return true
}

/* Testing scraping mirrors from animeshow.tv */
func test__animeshow_mirrorlist_iframe() (bool) {
	test_data := []string{"http://animeshow.tv/Gate-Jieitai-Kanochi-nite-Kaku-Tatakaeri-episode-6/",
		"http://animeshow.tv/Hundred-episode-5/",
		"http://animeshow.tv/Kuma-Miko-episode-5/",
		"http://animeshow.tv/Sakamoto-desu-ga-episode-3/"}
	test_answers := []string{`<iframe width="100%" height="100%" id="video_embed" scrolling="no" src="http://www.mp4upload.com/embed-nbpct8u6syi4.html" allowfullscreen=""></iframe>`,
		`<iframe width="100%" height="100%" id="video_embed" scrolling="no" src="http://www.mp4upload.com/embed-d1bzlyzjv2bq.html" allowfullscreen=""></iframe>`,
		`<iframe width="100%" height="100%" id="video_embed" scrolling="no" src="http://www.mp4upload.com/embed-dqffyin361rk.html" allowfullscreen=""></iframe>`,
		`<iframe width="100%" height="100%" id="video_embed" scrolling="no" src="http://www.mp4upload.com/embed-yaxx1w1uuxvu.html" allowfullscreen=""></iframe>`}
	for i, v := range test_data {
		obj := animeshow_mirrorlist_iframe(v)
		if obj != test_answers[i] {
			return false
		}
	}
	return true
}

/* Testing scraping hinter episode links for mirror iframes on animeshow.tv */
func test__animeshow_mirrorlist_links_to_scrape() (bool) {
	obj, exists := animeshow_mirrorlist_links_to_scrape("http://animeshow.tv/Sakamoto-desu-ga-episode-1/")
	if !exists {
		log.Fatal("Doesn't exist")
	}
	if obj[0].Name != "MP4UPLOAD" {
		return false
	}
	return true
}

//
// Run the tests
//

func run_tests() {
	fmt.Print("Test grabbing xml response from MAL... ")
	if !test__mal_fetch_data() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}

	fmt.Print("Test scraping trailer from MAL... ")
	if !test__mal_fetch_trailer() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}

	fmt.Print("Testing scraping entire anime list from animeshow.tv... ")
	if !test__animeshow_animelist_page() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}

	fmt.Print("Testing scraping from episodes listing page in animeshow.tv... ")
	if !test__animeshow_episodelist_page() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}

	fmt.Print("Testing scraping mirrors from animeshow.tv... ")
	if !test__animeshow_mirrorlist_iframe() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}

	fmt.Print("Testing scraping hinter episode links for mirror iframes on animeshow.tv... ")
	if !test__animeshow_mirrorlist_links_to_scrape() {
		fmt.Println("Failed")
		os.Exit(1)
	} else {
		fmt.Println("Passed")
	}
}