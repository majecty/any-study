package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/jroimartin/gocui"
)

var repo *git.Repository
var commits []Commit

func main() {
	// Parse command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("Please provide a path to a Git repository")
	}
	repoPath := os.Args[1]

	// Open the Git repository
	var err error
	repo, err = git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Failed to get HEAD reference: %v", err)
	}

	fmt.Printf("Current HEAD: %s\n", ref.Hash())

	commits, err = parseCommits(repo, 10)
	if err != nil {
		log.Fatalf("Failed to parse commits: %v", err)
	}

	// Create CUI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalf("Failed to create GUI: %v", err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Fatalf("Failed to set keybindings: %v", err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("Failed to start main loop: %v", err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("commits", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Commits"
		v.Wrap = true
		v.Autoscroll = true

		for _, c := range commits {
			fmt.Fprintf(v, "%s - %s: %s\n", c.Hash[:7], c.Author, c.Message)
		}
	}
	return nil
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

type Commit struct {
	Hash    string
	Message string
	Author  string
	Date    time.Time
}

func parseCommits(repo *git.Repository, maxCount int) ([]Commit, error) {
	var commits []Commit

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD reference: %v", err)
	}

	iter, err := repo.Log(&git.LogOptions{
		From:  ref.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get commit iterator: %v", err)
	}

	// Iterate through the commits
	err = iter.ForEach(func(c *object.Commit) error {
		commits = append(commits, Commit{
			Hash:    c.Hash.String(),
			Message: c.Message,
			Author:  c.Author.Name,
			Date:    c.Author.When,
		})

		if len(commits) >= maxCount {
			return fmt.Errorf("reached max count")
		}

		return nil
	})

	if err != nil && err.Error() != "reached max count" {
		return nil, fmt.Errorf("failed to iterate commits: %v", err)
	}

	return commits, nil
}
