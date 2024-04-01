<p align="center">
<img height="150" width="150" src="https://cdn.simpleicons.org/premierleague/gray" alt="premier-league-logo"/>
</p>

<div align="center">
    <h1>Premier League CLI</h1>
    <img src="https://img.shields.io/github/v/release/digitalghost-dev/premier-league-cli?style=flat-square&logo=git&logoColor=38003C&label=Release%20Version&labelColor=EEE&color=38003C" alt="version-label">
    <img src="https://img.shields.io/docker/image-size/digitalghostdev/pl-cli/v0.4.2?arch=arm64&style=flat-square&logo=docker&logoColor=38003C&labelColor=EEE&color=38003C" alt="docker-image-size">
</div>

<div align="center">
    <img src="https://img.shields.io/github/actions/workflow/status/digitalghost-dev/premier-league-cli/go_tests.yml?style=flat-square&logo=go&logoColor=00ADD8&label=Tests&labelColor=EEE&color=00ADD8" alt="tests-label">
    <img src="https://img.shields.io/github/go-mod/go-version/digitalghost-dev/premier-league-cli?style=flat-square&logo=Go&labelColor=EEE&color=00ADD8" alt="go-version">
</div>

## Overview
A tool for viewing data related to the Premier League but through a command line!
This project is an addition to a data engineering project that I've been working on.

## Links
* [Data Engineering Project](https://github.com/digitalghost-dev/premier-league)

## Infrastructure
Using Cloud Scheduler, a Cloud Function is ran on a schedule that exports data from BigQuery to Cloud Storage daily at 6AM PST.
![data-pipeline-flowchart](https://storage.googleapis.com/premier_league_bucket/premier_league_cli/cloud_functions_pipeline.png)

## Usage
**Note:** An internet connection is required to fetch the data from the Cloud Bucket.
```
Welcome! This tool displays statistics and data for English Premier League in the terminal.
Data is currently updated every day at 6:00 PST.

Usage:
  pl-cli [subcommand] [flag]

Available Commands:
         standings:      Renders the current standings of the English Premier League.
```

## Install

### Docker
```bash
docker run --rm -it digitalghostdev/pl-cli:v0.4.2
```

## Examples
<details><summary><b>Example</b> - Full table output</summary>

![example-standings-output](https://storage.googleapis.com/premier_league_bucket/premier_league_cli/standings_table_output.png)

</details>

<details><summary><b>Example</b> - Champions League output</summary>

![example-standings-c-flag-output](https://storage.googleapis.com/premier_league_bucket/premier_league_cli/standings_table_c_flag_output.png)

</details>


## Security