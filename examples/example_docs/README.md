From this directory, try the following commands:

### Install bood
`go get -u github.com/OlegVanyaGreatBand/kpi-lab-1/build/cmd/bood`

### Build the program

```
$ bood
INFO 2021/03/21 13:32:09 Adding doc actions for go module 'example_godoc'
INFO 2021/03/21 13:32:09 Ninja build file is generated at out/build.ninja
INFO 2021/03/21 13:32:09 Starting the build now
[1/1] Generating docs for example_godoc
```

Docs stored in file `out/docs/my-docs.txt`

When tests have succeed, nothing to do anymore:
```
$ bood
INFO 2021/03/21 13:32:30 Adding doc actions for go module 'example_godoc'
INFO 2021/03/21 13:32:30 Ninja build file is generated at out/build.ninja
INFO 2021/03/21 13:32:30 Starting the build now
ninja: no work to do.
```