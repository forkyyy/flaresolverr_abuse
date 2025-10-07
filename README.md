<h2>Golang Flaresolverr Abuser</h2>

<h3>Coded by forky</h3>

<h4>This is made to abuse open flaresolverr servers, never patched, still waiting, if you are getting abused, open an issue on their github repo or make ACL rules on your network.</h4>


<h1>Installation:</h1>

```sh
apt install snap snapd -y
snap install go --classic
curl -sL https://deb.nodesource.com/setup_16.x | sudo bash -
sudo apt -y install nodejs
npm i axios
```

<h1>Setup:</h1>

```sh
go build -o main main.go
```

<h3>Scraping for Flaresolverr servers (change the API Key on scraper.js)</h3><br>

```sh
node scraper.js > list.txt
```

<h3>Usage:</h3><br>

```sh
./main https://website.com
```


## Hope this gets fixed soon, as its allowing people to get free browser sessions
