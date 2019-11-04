# Funnel

# Overview

This project provides a centralized way to run re-occuring data extraction + ingestion jobs. It is really a platform to schedule, run, and manage periodic ETL jobs from data sources to data destinations.

# Development

## Local Setup

### Build and Run

After cloning the repo, you'll need to build the latest version of the code before it can be run:

```bash
cd cmd
go build -o bin/funnel -v .
```

To view the web app run this command:

```bash
heroku local web
```

In a browser, navigate to `localhost:5000`.

**NOTE:** To run the app locally, you'll need to be sure that the paths below in `cmd/main.go` look like this:

```go
router.LoadHTMLGlob("templates/*.tmpl.html")
router.Static("/static", "static")
```

## Heroku Setup

### Pushing Code to Heroku

Code can be pushed to Heroku with the following command:

```bash
git push heroku [branch_name]:master
```
**NOTE:** Heroku requires the paths to be a bit different than how they need to be for local development. You'll need to be sure that the following two paths in `cmd/main.go` look like this:

```go
router.LoadHTMLGlob("cmd/templates/*.tmpl.html")
router.Static("/static", "cmd/static")
```

### Opening the Heroku App

The Heroku app can be accessed by the following command:

```bash
heroku open
```

This will open the app in a new tab in your browser.
