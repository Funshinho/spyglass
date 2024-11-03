package spyglass

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

const NBA_REFERENCE_ENDPOINT = "https://basketball-reference.com"

type Client struct {
	Url string
}

type Option func(client *Client)

func WithUrl(url string) Option {
	return func(client *Client) {
		client.Url = url
	}
}

func NewClient(opts ...Option) *Client {
	if opts == nil {
		return &Client{Url: NBA_REFERENCE_ENDPOINT}
	}
	client := &Client{}
	for _, applyOpt := range opts {
		applyOpt(client)
	}
	return client
}

// GetTeams returns the list of teams for the current season
func (c Client) GetTeams() ([]Team, error) {
	collector := colly.NewCollector()
	teams := make([]Team, 0)

	collector.OnRequest(getVisitingPage)
	collector.OnResponse(getResponseStatus)

	collector.OnXML("//div[@id=\"all_teams_active\"]//th[@data-stat=\"franch_name\"]/a", func(e *colly.XMLElement) {
		regex := regexp.MustCompile(`\/teams\/([^}]*)\/`)
		match := regex.FindStringSubmatch(e.Attr("href"))
		tricode := getModifiedTricode(match[1])
		team := Team{
			ID:      tricode,
			Name:    e.Text,
			Tricode: tricode,
		}
		teams = append(teams, team)
	})

	err := collector.Visit(c.Url + "/teams")
	if err != nil {
		fmt.Println(err)
		return []Team{}, err
	}
	return teams, nil
}

func getModifiedTricode(tricode string) string {
	switch tricode {
	case "NJN":
		return "BRK"
	case "NOH":
		return "NOP"
	case "CHA":
		return "CHO"
	}
	return tricode
}

// GetPlayers returns the list of teams for the given season by list of teams, or all teams if no teams specified
func (c Client) GetPlayers(year int, teams ...string) ([]Player, error) {
	collector := colly.NewCollector()
	extensions.RandomUserAgent(collector)
	players := make([]Player, 0)

	collector.OnRequest(getVisitingPage)
	collector.OnResponse(getResponseStatus)

	collector.OnHTML("#roster > tbody", func(e *colly.HTMLElement) {
		regex := regexp.MustCompile(`\/teams\/([^}]*)\/` + strconv.Itoa(year) + `.html`)
		match := regex.FindStringSubmatch(e.Request.URL.Path)
		teamID := match[1]

		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			regex = regexp.MustCompile(`\/players\/.\/([^}]*)\.html`)
			match = regex.FindStringSubmatch(h.ChildAttr("td[data-stat=\"player\"] > a", "href"))
			playerID := match[1]

			fullName := h.ChildText("td[data-stat=\"player\"] > a")
			position := h.ChildText("td[data-stat=\"pos\"]")
			name := strings.Split(fullName, " ")
			firstName := name[0]
			lastName := strings.ReplaceAll(fullName, name[0]+" ", "")

			player := Player{
				ID:        playerID,
				FirstName: firstName,
				LastName:  lastName,
				Position:  Position(position),
				TeamID:    teamID,
			}
			players = append(players, player)
		})
	})
	if teams == nil {
		refTeams, err := c.GetTeams()
		if err != nil {
			fmt.Println(err)
			return []Player{}, err
		}
		teams = make([]string, 0)
		for _, refTeam := range refTeams {
			teams = append(teams, refTeam.ID)
		}
	}
	for _, team := range teams {
		err := collector.Visit(c.Url + "/teams/" + team + "/" + strconv.Itoa(year) + ".html")
		if err != nil {
			fmt.Println(err)
			return []Player{}, err
		}
	}
	return players, nil
}

func getVisitingPage(r *colly.Request) {
	fmt.Printf("Visiting page %s\n", r.URL.String())
}

func getResponseStatus(r *colly.Response) {
	fmt.Println("Status:", r.StatusCode)
}
