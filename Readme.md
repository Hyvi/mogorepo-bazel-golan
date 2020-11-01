# Monorepo Golang Application With Bazel

packages folder is a folder to store multiple project,

and share the same code (such as the Health handler in the dir "shared")
    
## Auto Generate BUILD.bazel

bazel run //:gazelle

或

bazel run //:gazelle -- update-repos -from_file=go.mod

## Run Hello-World Application

bazel run  --verbose_failures  //packages/hello-world:hello-world

## Create Dep graph

bazel query 'allpaths(packages/...,@com_github_gin_gonic_gin//:go_default_library)' --output graph | dot -Tpng > dep.png


## Run Hello-World's Router Unit-Test

bazel run  --verbose_failures  //packages/hello-world/router:router_test