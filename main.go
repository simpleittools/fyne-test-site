package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

type BlogPost struct {
	title    string
	date     string
	content  string
	expanded bool
}

func main() {
	a := app.New()
	// This should not be needed anymore
	//a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Ryan's Blog")
	w.Resize(fyne.NewSize(800, 600))
	logo := widget.NewLabel("Ryan's Blog")
	logo.TextStyle = fyne.TextStyle{Bold: true}

	// TODO: Move entire navigation to its own file
	// Navigation buttons
	homeBtn := widget.NewButton("Home", func() {
		// TODO: Scroll to top
	})
	aboutBtn := widget.NewButton("About", func() {
		// TODO: scroll to about
	})
	blogBtn := widget.NewButton("Blog", func() {
		// TODO scroll to blog
	})
	contactBtn := widget.NewButton("Contact", func() {
		// TODO scroll to contact
	})

	// create the navbar
	navBar := container.NewHBox(
		layout.NewSpacer(),
		homeBtn,
		aboutBtn,
		blogBtn,
		contactBtn,
	)

	// create a responsive nav, containing the logo, the header, and the nav in a hamburger menu
	header := container.NewVBox(
		container.NewHBox(
			logo,
			layout.NewSpacer(),
			makeHamburgerMenu(navBar),
		), navBar,
	)

	// TODO: Move this to it's own file
	// About section
	aboutHeading := widget.NewLabel("About")
	aboutHeading.TextStyle = fyne.TextStyle{Bold: true}
	// TODO: Why use widget.NewLabel instead of canvas.NewText
	aboutText := widget.NewLabel("This is a blog about Ryan's life and projects")
	aboutText.Wrapping = fyne.TextWrapWord

	// TODO: Move this to it's own file
	// Hobbies Section
	hobbiesLabel := widget.NewLabel("My hobbies are:")
	// Add links
	programmingLink := widget.NewHyperlink("Programming", parseURL("https://youtube.com/@codingunknown"))
	musicLink := widget.NewHyperlink("Listening to Music", parseURL("https://youtube.com/@musicunknown_ak"))
	moviesLink := widget.NewHyperlink("Watching Movies", parseURL("https://youtube.com/@moviesunknown_ak"))

	booksLabel := widget.NewLabel("Books")
	booksLink := widget.NewHyperlink("We are Legion", parseURL("https://www.goodreads.com/book/show/32109569-we-are-legion-we-are-bob?ac=1&from_search=true&qid=v0pzblWBqS&rank=1"))
	booksContainer := container.NewHBox(booksLabel, widget.NewLabel(" - "), booksLink, widget.NewLabel("is one of my all time favorites"))

	scifiLabel := widget.NewLabel("Watching Science Fiction")

	teachingLabel := widget.NewLabel("Teaching People About")
	computersLink := widget.NewHyperlink("computers", parseURL("https://youtube.com/@windowsunknown"))
	teachingContainer := container.NewHBox(teachingLabel, widget.NewLabel(" "), computersLink)

	// create a container for all the hobbies
	hobbiesContainer := container.NewVBox(
		hobbiesLabel,
		container.NewHBox(widget.NewLabel(" - "), programmingLink),
		container.NewHBox(widget.NewLabel(" - "), musicLink),
		container.NewHBox(widget.NewLabel(" - "), moviesLink),
		container.NewHBox(widget.NewLabel(" - "), booksContainer),
		container.NewHBox(widget.NewLabel(" - "), scifiLabel),
		container.NewHBox(widget.NewLabel(" - "), teachingContainer),
	)

	// About container
	aboutContainer := container.NewVBox(
		aboutHeading,
		aboutText,
		hobbiesContainer,
	)

	// TODO: move this to it's own file
	blogHeading := widget.NewLabel("Blog")
	blogHeading.TextStyle = fyne.TextStyle{Bold: true}

	// list of blog posts is a slice of BlogPost from the struct
	blogPosts := []BlogPost{
		{
			title:    "Blog Post 1",
			date:     "09/29/2023",
			content:  "Here is my blog post. Here is a bunch of cool details.\n\nThis is after a br",
			expanded: false,
		},
		{
			title:    "Blog Post 2",
			date:     "09/24/2023",
			content:  "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Est rem inventore quo beatae magni! Odit quo saepe alias molestias fugiat doloribus dolorem consectetur incidunt, adipisci et perspiciatis enim at vero?",
			expanded: false,
		},
		{
			title:    "Blog Post 3",
			date:     "09/21/2023",
			content:  "Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestias, maxime quis corrupti sit natus ipsa ex in voluptatibus fugiat rerum neque corporis. Veniam sit aliquid cum sunt ad quis culpa.",
			expanded: false,
		},
	}

	// create the blog UI components
	var blogContainers []*fyne.Container
	for i := range blogPosts {
		post := &blogPosts[i]

		// create the title as a button
		titleBtn := widget.NewButton(post.title, func() {
			post.expanded = !post.expanded
			updateBlogPost(post, blogContainers[i])
		})

		// show the date and content
		dateLabel := widget.NewLabel(post.date)
		dateLabel.TextStyle = fyne.TextStyle{Italic: true}
		contentLabel := widget.NewLabel(post.content)
		contentLabel.Wrapping = fyne.TextWrapWord

		// detail container
		detailsContainer := container.NewVBox(
			dateLabel,
			contentLabel,
		)

		if !post.expanded {
			detailsContainer.Hide()
		}

		postContainer := container.NewVBox(
			titleBtn,
			detailsContainer,
		)

		blogContainers = append(blogContainers, postContainer)
	}

	// create a container for all blog posts
	blogPostsContainer := container.NewVBox()
	for _, c := range blogContainers {
		blogPostsContainer.Add(c)
	}

	blogContainer := container.NewVBox(
		blogHeading,
		blogPostsContainer,
	)

	// TODO: move this to it's own file
	contactHeading := widget.NewLabel("Contact Info")
	contactHeading.TextStyle = fyne.TextStyle{Bold: true}

	contactText := widget.NewLabel("Contact me through this form:")
	formLink := widget.NewHyperlink("Google form", parseURL("https://docs.google.com/forms/d/e/1FAIpQLScHE8eLyHMlfUP5KU6S7Tmiu2ZtrGkhPwY0m86Aaqalc6O6Gg/viewform?usp=sharing"))

	contactContainer := container.NewVBox(
		contactHeading,
		// TODO: Change this to a button
		container.NewHBox(contactText, widget.NewLabel(" "), formLink),
	)
	// Main content
	content := container.NewVBox(
		header,
		widget.NewSeparator(),
		widget.NewLabel("Home Page"),
		widget.NewSeparator(),
		aboutContainer,
		widget.NewSeparator(),
		blogContainer,
		widget.NewSeparator(),
		contactContainer,
	)

	// make the content scrollable
	scrollContainer := container.NewScroll(content)

	w.SetContent(scrollContainer)
	w.ShowAndRun()
}

// makeHamburgerMenu is a helper function to create the hamburger menu button that toggles visibility
func makeHamburgerMenu(navBar *fyne.Container) *widget.Button {
	isMenuVisible := true

	hamburgerBtn := widget.NewButtonWithIcon("", theme.MenuIcon(), func() {
		isMenuVisible = !isMenuVisible
		if isMenuVisible {
			navBar.Show()
		} else {
			navBar.Hide()
		}
	})

	return hamburgerBtn
}

// parseURL is a helper function to handle URL Parsing and error handling
func parseURL(urlStr string) *url.URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing the URL: ", err)
		return nil
	}
	return u
}

// updateBlogPost is a helper function to toggle blog post visibility
func updateBlogPost(post *BlogPost, container *fyne.Container) {
	details := container.Objects[1]
	if post.expanded {
		details.Show()
	} else {
		details.Hide()
	}
	container.Refresh()
}
