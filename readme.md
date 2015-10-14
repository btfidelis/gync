# Gync

Tool for syncing files to dropbox your folder.
 
## Build 

*This project is under development and at the moment is just a working prototype*

Download the source

```sh
$ go get -d github.com/btfidelis/gync
```

Install the dependencies

```sh
$ go get github.com/codegangsta/cli
$ go get github.com/crackcomm/go-clitable
```

Now you should be able to build with
```sh
go install github.com/btfidelis/gync
```

## Usage
First, you'll need to configure the dropbox folder you wish to store your savegames do that by copying the config.json.dist file and renaming to config.json

```sh
$ cp config.json.dist config.json
```

Change the values of the config.json keys as needed (at the moment check interval does not work)

```json
{
	"BackupPath": "C:\\Users\\Bruno\\Dropbox\\.gync",
	"CheckInterval": 2
}
```

Now you can add the save game directory
```sh
$ gync add GameName "C:\Path\To\Save"
```

And start the daemon, witch will listen for changes in the save dir and replicate to your dropbox

```sh
$ gync start
```

## Third-Party Dependencies
* [CLI](https://github.com/codegangsta/cli) :  A small package for building command line apps in Go
* [CLI-TABLE](https://github.com/crackcomm/go-clitable) : Command line (ASCII) and Markdown table for Golang

## Todo
* Unit Tests
* Refactoring
* Add the functionality to restore saves
* Add the functionality to create profiles ("git like branches")
* Implementation with dropbox api

## Contribuiting
I am quite new to golang, so if you have any suggestions or you can help with some of the todo itens, feel free to open a feature issue. Please describe the new feature as well as implementation ideas.

Pull requests for bugs may be sent without creating any Feature issue. If you believe that you know of a solution for a bug that has been filed on Github, please leave a comment detailing your proposed fix

## License
See the file called LICENSE
