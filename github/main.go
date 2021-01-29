package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

var (
	nowTime = time.Now()

	yesterday = nowTime.AddDate(0, 0, -1)

	lastWeek = nowTime.AddDate(0, 0, -7)

	lastMonth = nowTime.AddDate(0, -1, 0)

	lastYear = nowTime.AddDate(-1, 0, 0)

	json = jsoniter.ConfigFastest
)

type OriginProject struct {
	Repo   string   `json:"repo"`
	Org    string   `json:"org"`
	Tags    []string `json:"tags"`
	Name   string   `json:"name"`
	Branch string   `json:"branch"`
	Topics  []string `json:"topics"`
}

// LoadFile -
func LoadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//  UnmarshalListProjects - 
func UnmarshalListProjects(data []byte) ([]OriginProject, error) {
	var OriginProjects []OriginProject

	err := json.Unmarshal(data, &OriginProjects)
	if err != nil {
		return nil, err
	}

	return OriginProjects, nil
}

type Tag struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type JSONFile struct {
	Date     string    `json:"date"`
	Tags     []Tag     `json:"tags"`
	Projects []Project `json:"projects"`
}

// Project -
type Project struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Trends      struct {
		Daily   int `json:"daily"`
		Weekly  int `json:"weekly"`
		Monthly int `json:"monthly"`
		Yearly  int `json:"yearly"`
	} `json:"trends"`
	Tags             []string `json:"tags"`
	ContributorCount int      `json:"contributor_count"`
	PushedAt         string   `json:"pushed_at"`
	OwnerID          int64    `json:"owner_id"`
	CreatedAt        string   `json:"created_at"`

	// below special to npm
	Npm string `json:"npm"`
	// origin is npm, now show be go.dev
	Downloads int `json:"downloads"`

	Icon   string `json:"icon"`
	Branch string `json:"branch"`
	URL    string `json:"url"`
}

// format

// branch: "main"
// contributor_count: 5
// created_at: "2021-01-14"
// description: "An ultra-light UI runtime"
// downloads: 1609
// full_name: "forgojs/forgo"
// name: "Forgo"
// npm: "forgo"
// owner_id: 77432949
// pushed_at: "2021-01-28"
// stars: 224
// tags: ["framework"]
// trends: {daily: 1}
// url: "https://forgojs.org/"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "e0b6e1cc3314887d976b5c5d9c97003f327cb76b"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	raw, err := LoadFile("./list-projects.json")
	if err != nil {
		log.Println(err)
		return
	}

	listPro, err := UnmarshalListProjects(raw)
	if err != nil {
		log.Println(err)
		return 
	}
	log.Println(listPro)

	var jsonfile = &JSONFile{}
	// date
	jsonfile.Date = time.Now().Format(time.RFC3339)

	// tags
	var Tags2Code = make(map[string]string)
	for _, v := range listPro {
		for _, t := range v.Tags {
			if Tags2Code[t] != "" {
				continue
			}
			Tags2Code[t] = t
			jsonfile.Tags = append(jsonfile.Tags, Tag{Name: t, Code: t})
		}
	}

	// get projects down
	var org = "abserari"
	var repo = "abserari"
	project, err := getProject(context.Background(), client, org, repo)

	jsonfile.Projects = append(jsonfile.Projects, *project)

	data, err := json.Marshal(jsonfile)
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func getProject(ctx context.Context, client *github.Client, org, repo string) (*Project, error) {
	startTime := time.Now()
	log.Println("start now at ", startTime.Sub(startTime))
	var project = &Project{}
	stats, _, err := client.Repositories.Get(ctx, org, repo)
	if err != nil {
		return nil, err
	}
	err = getStarsTrending(ctx, client, project, org, repo)
	if err != nil {
		return nil, err
	}
	log.Println("end network at ", startTime.Sub(time.Now()))

	project.Name = repo
	project.Stars = stats.GetStargazersCount()
	project.FullName = stats.GetFullName()
	project.Description = stats.GetDescription()
	project.OwnerID = stats.GetOwner().GetID()
	project.PushedAt = stats.GetCreatedAt().Format("2006-01-02")
	project.CreatedAt = stats.GetCreatedAt().Format("2006-01-02")
	project.URL = stats.GetURL()
	opt := &github.ListContributorsOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	contributors, _, err := client.Repositories.ListContributors(context.Background(), org, repo, opt)
	project.ContributorCount = len(contributors)

	project.Tags = []string{"framework"}

	// get all pages of results
	return project, nil
}

func getStarsTrending(ctx context.Context, client *github.Client, project *Project, org, repo string) error {
	var (
		daily   int
		weekly  int
		monthly int
		yearly  int
	)

	var stars []*github.Stargazer
	var page int = 0
	for {
		stargazers, resp, err := client.Activity.ListStargazers(context.Background(), org, repo, &github.ListOptions{Page: page})
		if err != nil {
			log.Println(err)
			return err
		}
		stars = append(stars, stargazers...)
		if resp.NextPage == 0 {
			break
		}
		page = resp.NextPage
	}

	for _, v := range stars {
		if v.StarredAt.Time.After(yesterday) {
			daily++
			weekly++
			monthly++
			yearly++
		} else if v.StarredAt.Time.After(lastWeek) {
			weekly++
			monthly++
			yearly++
		} else if v.StarredAt.Time.After(lastMonth) {
			monthly++
			yearly++
		} else if v.StarredAt.Time.After(lastYear) {
			monthly++
			yearly++
		}
	}
	project.Trends.Daily, project.Trends.Weekly, project.Trends.Monthly, project.Trends.Yearly = daily, weekly, monthly, yearly
	return nil
}
