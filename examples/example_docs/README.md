From this directory, try the following commands:

### Install bood
`go get -u github.com/OlegVanyaGreatBand/kpi-lab-1/build/cmd/bood`

### Build the program

```
$ bood
INFO 2021/03/21 11:24:53 Adding build & test actions for go binary module 'example_godoc'
INFO 2021/03/21 11:24:53 Ninja build file is generated at out/build.ninja
INFO 2021/03/21 11:24:53 Starting the build now
[1/1] Build example_godoc as Go binary
```
Docs stored in file `out/docs/my-docs.txt`

When tests have succeed, nothing to do anymore:
```
$ bood
INFO 2021/03/21 11:35:24 Adding build & test actions for go binary module 'example_test'
INFO 2021/03/21 11:35:24 Ninja build file is generated at out/build.ninja
INFO 2021/03/21 11:35:24 Starting the build now
ninja: no work to do.
```