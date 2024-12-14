package application

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"quake-logger/domain"
)

type LogParserService struct {
	Matches map[string]*domain.Match
}

func NewLogParserService() *LogParserService {
	return &LogParserService{Matches: make(map[string]*domain.Match)}
}

func (lp *LogParserService) ParseLogFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentMatch := ""
	matchCounter := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "InitGame") {
			matchCounter++
			currentMatch = "game_" + strconv.Itoa(matchCounter)
			lp.Matches[currentMatch] = &domain.Match{
				GameID:       currentMatch,
				Players:      make(map[string]struct{}),
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
		} else if strings.Contains(line, "Kill:") {
			lp.processKill(currentMatch, line)
		}
	}
	return scanner.Err()
}

func (lp *LogParserService) processKill(matchID, line string) {
	re := regexp.MustCompile(`\d+:\d+ Kill: \d+ \d+ \d+: (.+) killed (.+) by (\w+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 4 {
		killer := matches[1]
		killed := matches[2]
		means := matches[3]

		match := lp.Matches[matchID]

		if killer == "<world>" {
			match.Kills[killed]++
		} else {
			match.Players[killer] = struct{}{}
			match.Kills[killer]++
		}

		match.Players[killed] = struct{}{}
		match.TotalKills++
		match.KillsByMeans[means]++
	}
}

func (lp *LogParserService) GenerateReports() []*domain.Match {
	var keys []string
	for key := range lp.Matches {
		keys = append(keys, key)
	}

	// Order keys
	sort.Slice(keys, func(i, j int) bool {
		id1, _ := strconv.Atoi(strings.TrimPrefix(keys[i], "game_"))
		id2, _ := strconv.Atoi(strings.TrimPrefix(keys[j], "game_"))
		return id1 < id2
	})

	// Generate repost in a order
	reports := []*domain.Match{}
	for _, key := range keys {
		match := lp.Matches[key]
		match.GeneratePlayerList()
		reports = append(reports, match)
	}

	return reports
}
