
# Simple Revel Web APP with MongoDB and Redis

A high-productivity web framework for the [Go language](http://www.golang.org/).

# Install packages

[Glide](https://github.com/Masterminds/glide).

    glide install

To get packages use:

    glide get github.com/Masterminds/cookoo

### Start the web server localhost:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## [Heroku](https://revelapi.herokuapp.com/) 

Deploy to heroku: `git push heroku master`

## [Postman Collection](public/Comments.postman_collection.json)    

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

