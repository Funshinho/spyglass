# BoxscGobertore

Go client to fetch NBA players and teams info, as well as game statistics from basketball reference

## Usage

### Installation

```shell
go get github.com/Funshinho/gobert
```

### Importing

```go
 import "github.com/Funshinho/gobert
```

### Usage

```go
players := gobert.GetPlayers(2023)               // Returns the roster for all teams
players := gobert.GetPlayers(2023, "MIA", "LAL") // Returns the roster for the given teams
teams := gobert.GetTeams()                       // Returns the list of teams
```
