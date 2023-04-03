<a name="readme-top"></a>

[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]




<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/mconnat/faceit-api-client">

[//]: # (    <img src="images/logo.png" alt="Logo" width="80" height="80">)
  </a>

<h3 align="center">Faceit API Client</h3>

  <p align="center">
    Wrapper for the FACEIT API in Golang
    <br />
    <strong></strong>
    <br />
    <br />
    <a href="https://github.com/mconnat/faceit-api-client/issues">Report Bug</a>
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
go get github.com/mconnat/faceit-api-client

....

import (
	"github.com/mconnat/faceit-api-client/pkg/client"
	"github.com/mconnat/faceit-api-client/pkg/models"
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
	params := map[string]interface{}{
		"nickname": "m0NESY",
	}
	player, err := faceit.GetPlayer(params)
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
	params := map[string]interface{}{
        "type":   "all",
        "offset": "0",
        "limit":  "20",
        }
	matches, err := faceit.GetMatchesByChampionshipID(championshipID, params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matches)
```

Use expanded params:

```go
	expandedParams := []string{"game"}
	params := map[string]interface{}{
		"expanded": expandedParams,
	}
	championship, err := faceit.GetChampionshipByID(championshipID, params)
	if err != nil {
        fmt.Println(err)
	}
        fmt.Println(matches)
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
[issues-url]:https://github.com/mconnat/faceit-api-client/issues
[license-shield]: https://img.shields.io/github/license/mconnat/faceit-api-client?style=flat-square
[license-url]: https://github.com/mconnat/faceit-api-client/blob/master/LICENSE.txt
