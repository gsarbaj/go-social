package db

import (
	"context"
	"fmt"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
	"math/rand"
)

var usernames = []string{
	"alice", "bob", "dave", "eve", "mallory", "trent", "peggy", "victor", "heidi", "judy",
	"charlie", "wendy", "ivan", "grace", "oscar", "nancy", "pat", "sam", "kim", "mia",
	"john", "susan", "zane", "paul", "lara", "alex", "tina", "hank", "beth", "gary",
	"ron", "amy", "emma", "lucas", "noah", "liam", "zoe", "olivia", "chris", "dan",
	"kelly", "leo", "nina", "robin", "sara", "tom", "kurt", "ella", "mark", "claire",
}

var titles = []string{
	"Mastering Minimalism",
	"Top Travel Tips",
	"Healthy Eating Hacks",
	"Declutter Your Mind",
	"5-Minute Workouts",
	"Building Daily Habits",
	"Secrets to Productivity",
	"Tech Trends Today",
	"Finding Inner Peace",
	"The Joy of Reading",
	"Budget-Friendly Recipes",
	"Quick Coding Tips",
	"Photography Basics",
	"Simple Life Upgrades",
	"Yoga for Beginners",
	"Creative Writing Tips",
	"Social Media Success",
	"Better Sleep Tonight",
	"Unlocking Creativity",
	"Remote Work Essentials",
}

var contents = []string{
	"Learn how to simplify your life and focus on what truly matters.",
	"Discover must-visit destinations and essential travel gear tips.",
	"Explore quick and healthy recipes to transform your meals.",
	"Practical steps to declutter your thoughts and reduce stress.",
	"Try these short but effective workouts to stay fit on a busy schedule.",
	"Understand the science behind building habits that stick.",
	"Boost your productivity with these proven techniques.",
	"An overview of the latest advancements shaping technology.",
	"Tips and practices for achieving a sense of calm and tranquility.",
	"Why reading is a habit everyone should embrace, plus book suggestions.",
	"Delicious and affordable recipes perfect for any home cook.",
	"Improve your programming skills with these simple tricks.",
	"A beginner’s guide to capturing stunning photos with ease.",
	"Make your daily routines more enjoyable with these hacks.",
	"Step-by-step guidance for starting yoga and sticking with it.",
	"Essential advice for improving your creative writing projects.",
	"Strategies to grow your online presence and connect with others.",
	"How to create a nighttime routine for deeper, restful sleep.",
	"Unlock your hidden potential with these creative exercises.",
	"Maximize your efficiency and happiness while working remotely.",
}

var tags = []string{
	"lifestyle", "travel", "food", "fitness", "mindfulness", "habits",
	"productivity", "technology", "self-care", "reading", "recipes",
	"coding", "photography", "minimalism", "yoga", "writing",
	"social media", "sleep", "creativity", "remote work",
}

var comments = []string{
	"Great post! Really enjoyed reading this.",
	"Thanks for sharing these tips, very helpful!",
	"I never thought about it that way. Interesting perspective!",
	"Do you have more content like this? I'd love to read more.",
	"This was exactly what I needed today. Thanks!",
	"Can you elaborate on point #3? I'm curious.",
	"Awesome write-up, keep up the good work!",
	"This is so relatable. Thanks for putting it into words.",
	"I'll definitely try this out. Looks promising!",
	"How often do you update this blog? I love your posts!",
	"This inspired me to make some changes. Appreciate it!",
	"Could you recommend more resources on this topic?",
	"Wow, this was super informative. Thank you!",
	"I shared this with my friends—they loved it too!",
	"Totally agree with you on this. Well said!",
	"This post just made my day. Thanks a ton!",
	"I'm bookmarking this for future reference. Great insights!",
	"Have you thought about writing a follow-up post?",
	"Such a unique take on the subject. Great job!",
	"Where can I find more details about this?",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, u := range users {
		if err := store.Users.Create(ctx, u); err != nil {
			log.Println("Error creating user", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
