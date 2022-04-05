package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeckTestSuite struct {
	suite.Suite
	DefaultDeckSize int
}

func (suite *DeckTestSuite) SetupTest() {
	suite.DefaultDeckSize = 16
	os.Remove("_decktesting")
}

func TestNewDeckDefaultGo(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to be 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[len(d)-1])
	}
}

func (suite *DeckTestSuite) TestNewDeck() {
	d := newDeck()
	assert.Equal(suite.T(), suite.DefaultDeckSize, len(d), fmt.Sprintf("newDeck() should return deck with length of %v", suite.DefaultDeckSize))
}

func (suite *DeckTestSuite) TestSaveToDeckAndNewFromFile() {
	savedDeck := newDeck()
	savedDeck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	suite.Equal(len(savedDeck), len(loadedDeck), "newDeckFromFile() should match the length of savedDeck")

	os.Remove("_decktesting")
}

func TestDeckTestSuite(t *testing.T) {
	suite.Run(t, new(DeckTestSuite))
}
