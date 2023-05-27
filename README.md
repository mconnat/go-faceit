<a name="readme-top"></a>

[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]




<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/mconnat/go-faceit">

[//]: # (    <img src="images/logo.png" alt="Logo" width="80" height="80">)
  </a>

<h3 align="center">Go Faceit</h3>

  <p align="center">
    Wrapper for the FACEIT API in Golang
    <br />
    <strong></strong>
    <br />
    <br />
    <a href="https://github.com/mconnat/go-faceit/issues">Report Bug</a>
  </p>
</div>



<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
  </ol>
</details>



## About The Project

This project is meant to wrap the Face IT Data API and Chat API in golang.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

You need a Valid API token from FaceIT
```
go get github.com/mconnat/go-faceit

....

import (
	"github.com/mconnat/go-faceit/pkg/client"
	"github.com/mconnat/go-faceit/pkg/models"
)

```

### Prerequisites

* golang >= 1.19

<!-- USAGE EXAMPLES -->
## Usage

Use classic params:
```go
	faceit, err := client.New("API-TOKEN")
	if err != nil {
		log.Println(err)
	}
	params :=  client.GetPlayerParams{
		Nickname: "m0NESY",
	}

	player, err := faceit.GetPlayer(&params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(player.PlayerId)
```

Use multiple params:
```go
	faceit, err := client.New("API-TOKEN")
	if err != nil {
		log.Println(err)
	}
	from := time.Date(2023, 4, 01, 01, 00, 00, 00, time.Local).Unix()
	to := time.Now().Unix()
	params := client.GetPlayerHistoryParams{
		Game:   "csgo",
		From:   from,
		To:     to,
		Offset: 0,
		Limit:  10,
	}
	matches, err := faceit.GetPlayerHistory(player.PlayerId, &params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matches)
```

Use expanded params:

```go
	expandedParams := []string{"game"}
	params := client.GetChampionshipByIDParams{
		Expanded: expandedParams,
	}
	championship, err := faceit.GetChampionshipByID("f6117c36-cbe0-4f58-99b1-9e5e91ec4741", &params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(championship)
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>




## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



[issues-shield]: https://img.shields.io/github/issues/mconnat/faceit-api-client?style=flat-square
[issues-url]:https://github.com/mconnat/go-faceit/issues
[license-shield]: https://img.shields.io/github/license/mconnat/faceit-api-client?style=flat-square
[license-url]: https://github.com/mconnat/go-faceit/blob/master/LICENSE.txt
