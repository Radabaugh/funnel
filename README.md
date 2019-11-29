# Funnel

# Overview

This project provides a centralized way to run re-occuring data extraction + ingestion jobs. It is really a platform to schedule, run, and manage periodic ETL jobs from data sources to data destinations.

# Development

## Local Setup

After cloning the repo, you'll need to follow the `.env` example to create your own. The app can be run from the root directory with the following command:

```bash
go run cmd/main.go
```

### Development Database

For development, I'm using a Postgres database. Currently, this is the only supported database type. You can create the database with whatever you like, for example, I used `DBeaver` but you could also use a terminal tool like `psql`. Just be sure to note down the connection information in your `.env` file.

Below is the create table statement for the `jobs` table. This is currently the only table that Funnel uses.

```sql
CREATE TABLE jobs(
    id serial PRIMARY KEY,
    source TEXT NOT NULL,
    destination TEXT NOT NULL,
    interval TEXT NOT NULL,
    frequency INTEGER NOT NULL,
    job_type TEXT DEFAULT 'Database to Database',
    last_run TIMESTAMP,
    last_run_duration TEXT,
    number_of_runs INTEGER,
    next_run TIMESTAMP
);
```

## Production Setup

### Heroku CLI

The production enviroment is hosted on Heroku, you'll need to have the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed locally to push changes to the production server.


The Heroku environment settings can be accessed with `heroku config` once you have the Heroku CLI installed.

### Pushing Code to Heroku

Code can be pushed to Heroku with the following command:

```bash
git push heroku [branch_name]:master
```

### Accessing the Production Database

```bash
heroku psql
```

The `heroku psql` command will give you access to the production database where you can run any SQL command.

### Opening the Heroku App

The Heroku app can be accessed by the following command:

```bash
heroku open
```

This will open the app in a new tab in your browser.
