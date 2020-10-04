package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/hashstore/GoActorGo/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}

// NewLeafTagMatch creates leaf TagMatch
func NewLeafTagMatch(numOfTags int) *pb.TagMatch {
	matches := make([]*pb.TagMatch, numOfTags)
	for i := 0; i < numOfTags; i++ {
		matches[i] = &pb.TagMatch{
			Tag: randomStringFromSet("a", "b", "c", "d", "e", "f")}
	}
	return &pb.TagMatch{
		Matches:     matches,
		CombineAsOr: randomBool(),
		Negate:      randomBool(),
	}
}

// NewLeafTagMatch creates TagMatch tree
func NewTagMatch(depth int) *pb.TagMatch {
	if depth < 2 {
		return NewLeafTagMatch(randomInt(1, 3))
	}
	matches := make([]*pb.TagMatch, randomInt(1, 3))
	for i := 0; i < len(matches); i++ {
		matches[i] = NewTagMatch(depth - 1)
	}
	return &pb.TagMatch{
		Matches:     matches,
		CombineAsOr: randomBool(),
		Negate:      randomBool()}
}
