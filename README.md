# Debugging containerized Go Apps with Delve + VSCode

Just an updated example following the steps from this [Medium post](https://medium.com/@kaperys/delve-into-docker-d6c92be2f823)

# Try it out

Steps:

1. Clone the repo locally
2. Open it with VSCode
3. Build the image and running it, like this

```
$ docker build -t delve-playground .
$ docker run --rm --publish 38081:38081 --publish 38080:38080 --security-opt=seccomp:unconfined --name delve-playground delve-playground
```

At this point the dummy app is running but not debugger process attached to it

4. In VSCode, set some breakpoints in main.go

For example: go to line 21 and press F9 or click in the left side of the line number). This toggle a breakpoint

5. Start debugging (F5)

Attach the debbuger from VSCode using the `Delve Docker` configuration.

Notice that F5 uses the last configuration, `Shift+Cmd+D` for opening `Run and Debug` where you will find the dropdown for configurations.

(Your debugging window will show something like ..."Starting Web server")

6. Send a request to it

For example:
```
$ curl http://127.0.0.1:38081/go
```

You will see how VSCode pauses the execution at that line... Now it's your turn, debug it!


## Restarting the debugger and known issues

To restart the debugger you have to:

* Stop the contanier (which uses Delvev for running your app)
* Stop the debugging session in VSCode (shift + F5)
* Start the container again (`docker run...`) and attach the debug session from VSCode (F5)

This restart is necessary in case you add or remove breakpoints or if something goes wrong (you might see errors like trying to read from a closed connection or stuff like that, etc...).

## Debugging locally

If you don't need Docker and you can just debug your code locally (this program is a good example), then just:

```
$ dlv debug ./main.go --listen=:38081 --headless=true --api-version=2 --log
```

on your terminal, and attach the debbuger from VSCode using the `Delve Local` configuration. Of course it requires `Delve` to be installed on your machine.

## Related useful links

[Delve](https://github.com/go-delve/delve)
[Go + VSCode](https://code.visualstudio.com/docs/languages/go)
[Docker](https://www.docker.com/)
