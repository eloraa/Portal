package utils

import (
	"math/rand"
	"regexp"
	"time"
)

var (
	// Create a local random generator
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	roomIDPattern = regexp.MustCompile(`^[a-zA-Z0-9]{8}$`)

	adjectives = []string{
		"Fluffy", "Adorable", "Bouncy", "Cheerful", "Dancing",
		"Elegant", "Friendly", "Gentle", "Happy", "Jolly",
		"Lively", "Magical", "Noble", "Peaceful", "Quirky",
		"Radiant", "Silly", "Tender", "Upbeat", "Vibrant",
		"Warm", "Zealous", "Bright", "Cozy", "Dreamy",
	}

	nouns = []string{
		"Cookie", "Panda", "Kitten", "Puppy", "Cloud",
		"Star", "Moon", "Sun", "Rainbow", "Flower",
		"Bird", "Butterfly", "Dragon", "Phoenix", "Unicorn",
		"Crystal", "Diamond", "Pearl", "River", "Ocean",
		"Mountain", "Forest", "Garden", "Meadow", "Valley",
	}
)

func ValidateRoomID(id string) bool {
	return roomIDPattern.MatchString(id)
}

func GenerateRoomName() string {
	adj := adjectives[rng.Intn(len(adjectives))]
	noun := nouns[rng.Intn(len(nouns))]
	return adj + noun
}

func GenerateShortID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rng.Intn(len(charset))]
	}
	return string(result)
}
