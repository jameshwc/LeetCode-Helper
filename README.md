# LeetCode-Helper
Get the LeetCode information of a user via Golang.

## Set-up
You should put your cookie of leetcode session and csrftoken in a file named 'config.toml'.
It would be something like this:
 ```toml
 [connection]
session = "eyJ0eXAiOi..."
csrftoken = "krReU9..."
```

## Usage
Clone the repo and ```go run helper/main.go``` and you would get a README.md with all the problems you solved and get the record of the date on which you run the helper.
You can visit [my another repo](https://github.com/jameshwc/LeetCode-ans) to check how I use it.
