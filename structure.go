package main

import "encoding/xml"

/* Scraper Structs */
type structscraper__episodelist_page struct {
	Type        string
	Year        string
	Status      string
	Genre       []string
	Altname     string
	Description string
	Episodes    []string
}

type structscraper__mirror_to_extract struct {
	Name string
	Link string
	SubDub string
}

/* myanimelist.net structs */
type structmal_api_anime struct {
	Anime   xml.Name `xml:"anime"`
	Entries []structmal_api_entry `xml:"entry"`
}

type structmal_api_entry struct {
	ID            string `xml:"id"`
	Title         string `xml:"title"`
	EnglishName   string `xml:"english"`
	SynonymNames  string `xml:"synonyms"` // Semi-colon+space seperated
	EpisodesCount string `xml:"episodes"`
	Score         string `xml:"score"`
	Type          string `xml:"type"`
	Status        string `xml:"status"`
	StartDate     string `xml:"start_date"`
	EndDate       string `xml:"end_date"`
	Synopsis      string `xml:"synopsis"`
	Image         string `xml:"image"`
}

/* MAIN STRUCTURE */
type structmain__anime struct {
	MAL_ID                string `json:"MAL ID" storm:"id"`
	AnimeShow_Name        string `json:"AnimeShow Name" storm:"index"`
	MAL_Title             string `json:"MAL Title" storm:"index"`
	MAL_English           string `json:"MAL English Title" storm:"index"`
	AltName               string `json:"Alternate Name, omitempty" storm:"index"`
	Genre                 []string `json:"Genre" storm:"index"`
	AnimeShow_Description string `json:"AnimeShow Description" storm:"index"`
	MAL_Description       string `json:"MAL Description" storm:"index"`
	Score                 string `json:"Score" storm:"index"`
	Studio                string `json:"Studio Name" storm:"index"`
	Status                string `json:"Status" storm:"index"`
	Type                  string `json:"Type" storm:"index"`
	SynonymNames          string `json:"Synonym Names" storm:"index"`
	Year                  string `json:"Year" storm:"index"`
	Image                 string `json:"Image Url" storm:"index"`
	Trailer               string `json:"Trailer" storm:"index"`
	EpisodeList           []structmain__episode `json:"Episode List" storm:"index"`
}

type structmain__episode struct {
	Name    string     `json:"Episode Name" storm:"index"`
	Mirrors []structmain__mirror `json:"Episode Mirrors" storm:"index"`
}

type structmain__mirror struct {
	Name   string `json:"Mirror Name" storm:"index"`
	Iframe string `json:"Embed Code" storm:"index"`
	SubDub string `json:"SubDub" storm:"index"`
}