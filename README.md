## Using Docker Container in Cloud Foundry
---
As we all know, we can push source code to CF directly, and CF will compile it and create a container to run our application. Life is so great with CF.

But sometimes, for some reasons, like our App need a special setup or we want to run  app on different platforms or infrastructures, we may already have a configured container that for our App. That won't block our way to CF at all. This post will show you how to push docker image to CF.


##### Enable docker feature for CF

we can turn on docker support by the following cf command

```
cf enable-feature-flag diego_docker
```
also we can turn it off by

```
cf disable-feature-flag diego_docker
```

##### Push docker image to CF

```
cf push cf-docker -o golang/alpine
```

Unlike the normal way, CF won't take our code and run our app inside the image we specified. CF would suppose that you already put everything into the image. We have to rebuild docker image every single time we push change to our repository.

We need to tell CF how to start our app inside the image by specifing the start command. We can either put it as an argument for `cf push` or put it into `manifest.yml` as below.

```
---
applications:
- name: cf-docker
  command: git clone https://github.com/kaleo211/cf-docker && cd cf-docker && mkdir -p app/tmp && go run main.go
```

In this example, we are using an official docker image from docker hub. In the start command, we clone our demo repo from Github, do something and run our code.

##### Update Diego with private docker registry
If you are in EMC network, you may not able to use Docker Hub due to certificate issues. In this case, you need to setup a private docker registry. The version of registry need to be V2 for now. Also, you have to redeploy your CF or Diego with the changes being showed below.
```
properties:
  garden:
    insecure_docker_registry_list:
    - 12.34.56.78:9000
  capi:
    stager:
      insecure_docker_registry_list:
      - 12.34.56.78:9000
```
Replace `12.34.56.78:9000` with your own docker registry ip and port.

Then, you need to create security group to reach your private docker registry. You can put definition of this security group into `docker.json` as

```
[
  {
    "destination": "12.34.56.78:9000",
    "protocol": "all"
  }
]
```
And run

```
cf create-security-group docker docker.json
cf bind-staging-security-group docker
```
Now you can push to CF again by
```
cf push -o 12.34.56.78:9000/your-image
```
