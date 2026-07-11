package main

import "snippetbox.shivanshu.in/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
