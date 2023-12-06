## fox publish

Builds, pushes, and deploys KubeFox Apps using the version of the currently checked out Git commit

```
fox publish [NAME] [flags]
```

### Options

```
  -t, --create-tag         create Git tag using the AppDeployment version
      --dry-run            submit server-side request without persisting the resource
      --force              force build even if component image exists
  -h, --help               help for publish
  -k, --kind string        if provided the built image will be loaded into the kind cluster
  -n, --namespace string   namespace of KubeFox Platform
      --no-cache           do not use cache when building image
  -p, --platform string    name of KubeFox Platform to utilize
      --skip-deploy        do not perform deployment after build
      --skip-push          do not push image after build
  -s, --version string     version to assign to the AppDeployment, making it immutable
      --wait duration      wait up the specified time for components to be ready
```

### Options inherited from parent commands

```
  -a, --app string      path to directory containing KubeFox App
  -i, --info            enable info output
  -o, --output string   output format, one of ["json", "yaml"] (default "yaml")
  -r, --repo string     path to directory containing Git repository
  -v, --verbose         enable verbose output
```

### SEE ALSO

* [fox](fox.md)	 - CLI for interacting with KubeFox

