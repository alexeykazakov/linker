# Linker
Simple app to test Openshift.io - GitHub linking 

# Setup
 - Run [fabric8-auth](https://github.com/fabric8-services/fabric8-auth):

```
$ git clone git@github.com:fabric8-services/fabric8-auth.git
$ cd fabric8-auth
$ make dev
```

 - Run the test app:

```
$ git clone git@github.com:alexeykazakov/linker.git
$ cd linker
$ go install
$ linker
```

 - Open http://localhost:8090
 - Click on "Login"
 - Click on "Link GitHub"
