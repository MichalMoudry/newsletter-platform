package main

import (
	"fmt"
	"log"

	"newsletter-platform/config"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service"

	"github.com/google/uuid"
	//netchi "newsletter-platform/transport"
	//"newsletter-platform/transport/model"
	//"github.com/go-chi/jwtauth/v5"
)

func main() {
	cfg, cfgErr := config.ReadConfigFromFile("C:/Users/tomas/Desktop/newsletter-platform-main/config.json")
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	if err := database.OpenDb(cfg.ConnectionString); err != nil {
		log.Fatal(err)
	}
	/*
		tokenAuth := jwtauth.New("HS512", []byte(cfg.SecurityString), nil)
		handler := netchi.Initialize(
			cfg.Port,
			tokenAuth,
			model.NewServiceCollection(tokenAuth),
		)

		serverErr := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
		if serverErr != nil {
			log.Fatal(serverErr)
		}
	*/

	// Initialize the newsletter service
	newsletterService, err := service.NewNewsletterService(cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Create new newsletter
	authorID, err := uuid.Parse("9a09cf1c-99ee-4772-8086-c5a71706b816")
	if err != nil {
		log.Fatal(err)
	}

	newsletter := model.NewNewsletter("My Newsletter", "Newsletter description", authorID)
	err = newsletterService.CreateNewsletter(newsletter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Newsletter created!")

	// Create new post
	post := model.NewPost("My Post", "Post content", "John Doe", newsletter.Nltr_ID)
	err = newsletterService.CreatePost(post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post created!")

	// Delete newsletter
	err = newsletterService.DeleteNewsletter(newsletter.Nltr_ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Newsletter deleted!")

	// Delete post
	err = newsletterService.DeletePost(post.Post_ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post deleted!")
}
