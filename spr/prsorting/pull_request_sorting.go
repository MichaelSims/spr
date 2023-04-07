package prsorting

import (
	"github.com/ejoffe/spr/git"
	"github.com/ejoffe/spr/github"
)

type LocalCommitsAndPrs struct {
	LocalCommits []git.Commit
	PullRequests []*github.PullRequest
}
type SortPrsByLocalCommits LocalCommitsAndPrs

func (pair SortPrsByLocalCommits) Len() int {
	return len(pair.PullRequests)
}

func (pair SortPrsByLocalCommits) Swap(i, j int) {
	pair.PullRequests[i], pair.PullRequests[j] = pair.PullRequests[j], pair.PullRequests[i]
}

func _pos(commits []git.Commit, commitID string) int {
	for i := 0; i < len(commits); i++ {
		if commits[i].CommitID == commitID {
			return i
		}
	}
	return -1
}

func (pair SortPrsByLocalCommits) Less(i, j int) bool {
	positionI := _pos(pair.LocalCommits, pair.PullRequests[i].Commit.CommitID)
	positionJ := _pos(pair.LocalCommits, pair.PullRequests[j].Commit.CommitID)
	return positionI < positionJ
}
