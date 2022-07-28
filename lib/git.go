package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type GitRepo struct {
	repo *git.Repository
}

type CommitLog struct {
	Hash    string
	Message string
	Name    string
	Email   string
	Time    time.Time
}

type TagInfo struct {
	Hash string
	Name string
}

func indent(t string) string {
	var output []string
	for _, line := range strings.Split(t, "\n") {
		if len(line) != 0 {
			line = "    " + line
		}

		output = append(output, line)
	}

	return strings.Join(output, "\n")
}

func (c CommitLog) String() string {
	return fmt.Sprintf(
		"commit %s\nAuthor: %s <%s>\nDate:   %s\n\n%s\n",
		c.Hash, c.Name, c.Email,
		c.Time, indent(c.Message),
	)
}

func Open() GitRepo {
	repo, err := git.PlainOpen("./.git")
	IfError(err)

	return GitRepo{repo: repo}
}

func (g *GitRepo) GetHead() string {
	head, err := g.repo.Head()
	IfError(err)

	return head.Hash().String()
}

func (g *GitRepo) GetCommitBetweenHead(from string, to string) []CommitLog {
	iter, err := g.repo.Log(&git.LogOptions{From: plumbing.NewHash(from)})
	IfError(err)

	var messages []CommitLog
	iter.ForEach(func(c *object.Commit) error {
		messages = append(messages,
			CommitLog{
				Hash:    c.Hash.String(),
				Message: c.Message,
				Name:    c.Author.Name,
				Email:   c.Author.Email,
				Time:    c.Author.When,
			},
		)

		if to != "" && c.Hash.String() == to {
			return storer.ErrStop
		}
		return nil
	})

	return messages
}

func (g *GitRepo) GetTags() []TagInfo {
	iter, err := g.repo.Tags()
	IfError(err)

	var tags []TagInfo
	iter.ForEach(func(t *plumbing.Reference) error {
		tags = append(tags,
			TagInfo{
				Hash: t.Hash().String(),
				Name: t.Name().Short(),
			},
		)

		return nil
	})

	return tags
}
