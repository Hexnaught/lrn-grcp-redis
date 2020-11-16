## Sources for resolving some issues when setting up

There are 3 very similarly named tools, one of them is newer and had a PR open as of April 1st 2020, has since been merged, but a lot of guides/tutorials/'get started' information is seemingly outdated.

https://github.com/golang/protobuf/issues/1070

### Issue with methods in auto-generated files not found

I resolved this by using the 'dev' version specified in this reddit thread

https://www.reddit.com/r/grpc/comments/ih5qn5/undefined_grpcsupportpackageisversion7/


### Issue with missing `go_package` 'option'

https://stackoverflow.com/questions/61666805/correct-format-of-protoc-go-package
