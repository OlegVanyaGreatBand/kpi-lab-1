From this directory, try the following commands:

### Install bood
`go get -u github.com/OlegVanyaGreatBand/kpi-lab-1/build/cmd/bood`

### Build the program
```
$ bood
INFO 2021/03/08 10:27:43 Adding build & test actions for go binary module 'example'
INFO 2021/03/08 10:27:43 Ninja build file is generated at out/build.ninja
INFO 2021/03/08 10:27:43 Starting the build now
[2/2] Test example as Go binary
```

Every second build will fail

```
$ bood
INFO 2021/03/08 10:21:13 Adding build & test actions for go binary module 'example'
INFO 2021/03/08 10:21:13 Ninja build file is generated at out/build.ninja
INFO 2021/03/08 10:21:13 Starting the build now
[1/1] Test example as Go binary
FAILED: out/tests/test.txt 
cd . && go test ./... > out/tests/test.txt
ninja: build stopped: subcommand failed.
INFO 2021/03/08 10:21:13 Error invoking ninja build. See logs above.
```

Test results stored in file `out/tests/test.txt`

Built binary stored in `out/bin/example`

When tests have succeed, nothing to do anymore:

```
$ bood
INFO 2021/03/08 10:25:50 Adding build & test actions for go binary module 'example'
INFO 2021/03/08 10:25:50 Ninja build file is generated at out/build.ninja
INFO 2021/03/08 10:25:50 Starting the build now
ninja: no work to do.
```

Tests will be run again if source or test files have been changed.