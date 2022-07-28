package main

import (
	"fmt"

	"github.com/JisungPark0319/versioning/lib"
)

func main() {
	git := lib.Open()

	head := git.GetHead()

	tags := git.GetTags()
	tag := tags[len(tags)-1]

	for _, commit := range git.GetCommitBetweenHead(head, tag.Hash) {
		fmt.Print(commit.Message)
	}

}
