# IpList 0.1.0

Iplist package that it's compatible with firehol lists

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

`brew install golang`


### Installing

`go get github.com/0x4139/iplist`

```
response, err := http.Get("https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset")
	if err != nil {
		t.Fatalf("error while obtaining the list %s", err.Error())
	}
	defer response.Body.Close()

	iplist, err := NewFromReader(response.Body)
	if err != nil {
		t.Fatalf("Error while building the iplist %s", err.Error())
	}
	ip := "2.50.11.28"

	if iplist.Contains(ip) == false {
		t.Fatalf("List should contain the ip %s", ip)
	}
```


End with an example of getting some data out of the system or using it for a little demo

## Running the tests
`go test`

## Contributing

Please read [CONTRIBUTING.md]() for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Vali Malinoiu** - *Initial work* -

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
