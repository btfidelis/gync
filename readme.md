# Gync

Cli and daemon for syncing game saves to dropbox.
 
## Build 

*This project is under development and at the moment is just a working prototype*

Clone the repository in your go workspace

```sh
$ git clone https://github.com/btfidelis/gync.git
```

Install the dependencies

```sh
$ go get github.com/codegangsta/cli
$ go get github.com/crackcomm/clitable
```

Now you should be able to build with
```sh
go build github.com/btfidelis/gync/gync-daemon/gync.go
```

## Usage
First, you'll need to configure the dropbox folder you wish to store your savegames do that by coping the config.json.dist file and renaming to config.json

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

And start the deamon, witch will listen for changes in the save dir and replicate to your dropbox

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

## Contribuiting
I am quite new to golang, so if you have any suggestions or you can help with some of the todo itens, feel free to open a feature issue. Please describe the new feature as well as implementation ideas.

Pull requests for bugs may be sent without creating any Feature issue. If you believe that you know of a solution for a bug that has been filed on Github, please leave a comment detailing your proposed fix

## License
See the file called LICENSE
