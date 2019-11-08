# Funnel

# Overview

This project provides a centralized way to run re-occuring data extraction + ingestion jobs. It is really a platform to schedule, run, and manage periodic ETL jobs from data sources to data destinations.

# Development

## Local Setup

After cloning the repo, you'll need to follow the `.env` examples to create your own. The app can be run from the root directory with the following command:

```bash
go run cmd/main.go
```

## Production Setup

### Heroku CLI

The production enviroment is hosted on Heroku, you'll need to have the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed locally to push changes to the production server.

### Pushing Code to Heroku

Be sure that you have followed the production `.env` example file to create your own. Code can be pushed to Heroku with the following command:

```bash
git push heroku [branch_name]:master
```

### Opening the Heroku App

The Heroku app can be accessed by the following command:

```bash
heroku open
```

This will open the app in a new tab in your browser.
