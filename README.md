<p align="center">

# gomark

</p>

<p align="center">

## About

</p>

This will be an App to do benchmark test for API. The goal for version 1.0 will be to have a cli version.

This cli version will configure the tests through parameters and JSON or YAML files.

The results must be delivered in JSON, PDF or HTML files.

For future versions, we want to be able to be used programmatically with other languages.

<p align="center">

## Getting Involved

</p>

To help this project, you can read some Issues of our repository, if you want to help us with one, feel free to help.

In our Wiki we have a page with more details about our app.

We ask that when you use conventional commits when contributing. You can use a guide by running:

```
npm run commit
```

<p align="center">

## Starting to code

</p>

After cloning the repositor, you must install npm modules that are used for lint commit

```
npm install
```

After that, when you commit to the repository, please do "npm run commit" to guide you through a conventional commit. Commits out of the pattern will not be accepted.

The project is made in go, you can use go locally by running

```
go run ./cmd/app
```

You can develop locally without go installed, but with docker installed, by running:

```
docker compose up --build -d && docker compose logs -f
```

In this way, you will have a live reload environment with go running in a docker container. This docker composer file also have an API configured to be used during tests.

---

<p align="center">
This software is released under the AGPL-3.0 license
</p>
