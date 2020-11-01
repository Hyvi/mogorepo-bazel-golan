# Monorepo Golang Application With Bazel

packages folder is a folder to store multiple project
    
## Auto Generate BUILD.bazel

bazel run //:gazelle

或

bazel run //:gazelle -- update-repos -from_file=go.mod

## Run Hello-World Application

bazel run  --verbose_failures  //packages/hello-world:hello-world