# ucnbrew
a boilerplate for rest api in golang
gin, gorm, logrus (log, elastic, log to db), env


## CLI ucnbrew

Install cli

```sh
go install github.com/saucon/ucnbrew@latest
```

Don't forget to setup your path to go/bin.

A tool to generate boilerplate.
For example: <br />
packagename -> github.com/saucon <br />
appname -> helloapp <br />

```sh
ucnbrew brew [packagename]/[appname] [appname]
ucnbrew brew github.com/saucon/helloapp helloapp
```
