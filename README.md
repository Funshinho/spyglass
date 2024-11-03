# Spyglass

Go client to fetch NBA players and teams info, as well as game statistics from basketball reference

## Usage

### Installation

```shell
go get github.com/Funshinho/spyglass
```

### Importing

```go
 import "github.com/Funshinho/spyglass
```

### Usage

```go
players := spyglass.GetPlayers(2024)               // Returns the roster for all teams
players := spyglass.GetPlayers(2024, "MIA", "LAL") // Returns the roster for the given teams
teams := spyglass.GetTeams()                       // Returns the list of teams
```
