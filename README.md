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
boxscores := gobert.GetBoxscores("20220101")
players := gobert.GetPlayers(2023, "MIA")
teams := gobert.GetTeams()
```
