## Google images
This project grabs images from google search (primarily kitten images), resizes them, 
then Secures them with a very sophisticated key (cuz kitten images are most valuable) and finally stores them in pg

### Setup project
Make a `config.yaml` in `conf` folder (or use CONFIG env var to indicate where your config file is)
you can use example-config as a starting point

grab a google SearchEngineId from google's [Custom Search Engine](https://programmablesearchengine.google.com/controlpanel/all)

And Api Key from [here](https://console.cloud.google.com/apis/credentials)

simply run the project using `go run cmd/main.go` (and address the nags until it doesn't nag no more)

You can also use the provided `Dockerfile` to build project and make an image from it

### Endpoints

Now if you want to download images use link below:

GET http://localhost:9000/download/start?query=bad%20kittens&count=10

keep in mind that due to Google generating result pages with 10 items most/least `count` ceils to upper 10 like 0-10 = 10, 11-20 = 20 ...

You can download the grabbed images using link below:

GET http://localhost:9000/image/1/view (will download image file)

