package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zexk/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage :%s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	err = s.cfg.SetUser(name)
	fmt.Printf("user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]

	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Printf("user has been registered:")
	printUser(usr)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table: %w", err)
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching all users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	//TODO hardcoded for now
	url := "https://www.wagslane.dev/index.xml"
	//url := cmd.Args[0]
	feed, err := fetchFeed(context.Background(), url)
	unescapeHtml(feed)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v <url>", err)
	}

	printFeed(feed)
	return nil
}

func printUser(user database.User) {
	fmt.Printf("* id:	%v\n", user.ID)
	fmt.Printf("* name:	%v\n", user.Name)
}

func printFeed(feed *RSSFeed) {
	fmt.Printf("%v\n", feed.Channel.Title)
	fmt.Printf("%v\n", feed.Channel.Link)
	fmt.Printf("%v\n", feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Printf("%v\n", item.Title)
		fmt.Printf("%v\n", item.Link)
		fmt.Printf("%v\n", item.Description)
	}
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>\n", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching current user: %w\n", err)
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
			Url:       url,
			UserID:    current_user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating feed: %w\n", err)
	}

	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v\n", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		username, err := s.db.GetUsername(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error fetching feed user id: %w", err)
		}
		fmt.Printf("%v\n", username)
	}
	return nil
}
