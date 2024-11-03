package spyglass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTeams(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewClient(WithUrl(server.URL))

	teams, _ := client.GetTeams()
	assert := assert.New(t)

	team1 := teams[0]
	assert.Equal("ATL", team1.ID)
	assert.Equal("ATL", team1.Tricode)
	assert.Equal("Atlanta Hawks", team1.Name)

	team2 := teams[15]
	assert.Equal("MIA", team2.ID)
	assert.Equal("MIA", team2.Tricode)
	assert.Equal("Miami Heat", team2.Name)
}

func TestGetPlayers(t *testing.T) {
	server := GetMockServer(t)
	defer server.Close()
	client := NewClient(WithUrl(server.URL))

	players, _ := client.GetPlayers(2023, "TOR")
	assert := assert.New(t)

	player1 := players[0]
	assert.Equal("barnesc01", player1.ID)
	assert.Equal("Scottie", player1.FirstName)
	assert.Equal("Barnes", player1.LastName)
	assert.Equal("4", player1.Number)
	assert.Equal(PowerForward, player1.Position)
	assert.Equal("TOR", player1.TeamID)
}
