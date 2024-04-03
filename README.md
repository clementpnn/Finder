<div align="center">
  <img alt="Home Page" src="./assets/logo.png">
</div>

<h1 align="center">Finder</h1>
<p align="center">
  Personal Project
</p>

<p align="center">
  <a href="#introduction"><strong>Introduction</strong></a> ·
  <a href="#algorithm"><strong>Algorithm</strong></a> ·
  <a href="#setting-up-locally"><strong>Setting Up Locally</strong></a> ·
  <a href="#tech-stack"><strong>Tech Stack</strong></a> .
  <a href="#routes"><strong>Routes</strong></a>
</p>
<br/>

## Introduction

Uncovering the vast ocean of information on the web requires sophisticated tools that can navigate the complexities of content, context, and relevance. My project, which I've playfully nicknamed "Finder", is a powerful crawler that digs through websites and meticulously indexes every word on every page. It doesn’t stop there; it also calculates the frequency of each term and tracks the network of links between pages.

At the heart of "Finder" are the sophisticated TF-IDF (Term Frequency-Inverse Document Frequency) algorithm and PageRank (simplify). This dual approach allows “Finder” not only to index, but also to understand the importance of content, thus making search results more relevant and useful.

## Algorithm

TF = Total number of words in the page / Term frequency in the page
IDF = log(Number of pages containing the term / Total number of pages)
TF-IDF = TF×IDF
Total TF-IDF Score = ∑terms(TF-IDF of the term)

​PageRank = Number of incoming links

Final Score=α(Total TF-IDF Score)+β(PageRank)

## Setting Up Locally

To configure Finder locally, you will need to clone the repository and configure the following environment variables (in the .env file):

```
DB_HOST="database"
DB_PORT="5432"
DB_USER="root"
DB_PASSWORD="root"
DB_NAME="postgres"
```

To run the app locally, you can run the following commands:

```
docker-compose up -d
cd server
make db-up
```

## Tech Stack

Finder is built on the following stack:

**Back End:**

- [Go](https://go.dev/) - Programming Language
- [GoFiber](https://gofiber.io/) - Framework
- [Sqlc](https://sqlc.dev/) - ORM

**Front End:**

- [TypeScript](https://www.typescriptlang.org/) - Programming Language
- [React](https://fr.react.dev/) - JavaScript Library
- [Vite](https://vitejs.dev/) - Build Tool
- [Tanstack Router](https://tanstack.com/router/v1) – Routing
- [Tanstack Query](https://tanstack.com/query/latest) - Query Management
- [Ky](https://github.com/sindresorhus/ky) - fetching Library
- [TailwindCSS](https://tailwindcss.com/) – CSS Framework

**Database:**

- [PostgresSQL](https://www.postgresql.org/) - Relational Database

**Infrastructure & Deployment:**

- [Docker](https://www.docker.com/) - Containerize

## Routes

**Crawler Routes:**

`http://localhost:3000/swagger` Displays the crawler swagger

**Search Routes:**

`http://localhost:5173/` Displays the search page
