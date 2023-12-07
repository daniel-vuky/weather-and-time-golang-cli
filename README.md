# weather-and-time-golang-cli
Get the weather and time data

# How to use
Step 1: register an account from **https://www.weatherapi.com/**
Step 2: register API key from website dashboard
Step 3: config uri and api key for go: **go run main.go config -u "http://api.weatherapi.com/v1/current.json" -a "{API_KEY}"**
Step 4: Now we can try with some city: **go run main.go weather-and-time London** // Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name
