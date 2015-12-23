package main

import (
	"errors"
	"strings"

	"github.com/cloudnativego/gogo-engine"
)

type inMemoryMatchRepository struct {
	matches []gogo.Match
}

// NewRepository creates a new in-memory match repository
func newInMemoryRepository() *inMemoryMatchRepository {
	repo := &inMemoryMatchRepository{}
	repo.matches = []gogo.Match{}
	return repo
}

func (repo *inMemoryMatchRepository) addMatch(match gogo.Match) (err error) {
	repo.matches = append(repo.matches, match)
	return err
}

func (repo *inMemoryMatchRepository) getMatches() []gogo.Match {
	return repo.matches
}

func (repo *inMemoryMatchRepository) getMatch(id string) (match gogo.Match, err error) {
	found := false
	for _, target := range repo.matches {
		if strings.Compare(target.ID, id) == 0 {
			match = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find match in repository")
	}
	return match, err
}
