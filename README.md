# tufferina



```
docker build -t tufferina .

docker run --rm -p 8010:8010 tufferina
```


```
tufferina>docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/tufferina tufferina
```

The docker run command is used to run a container from an image,
The -it flag starts the container in an interactive mode (tie it to the current shell),
The --rm flag cleans out the container after it shuts down,
The --name mathapp-instance names the container mathapp-instance,
The -p 8010:8010 flag allows the container to be accessed at port 8010,
The -v $PWD/src:/go/src/mathapp is more involved. It maps the src/ directory from the machine to /go/src/mathapp in the container. This makes the development files available inside and outside the container, and
The mathapp part specifies the image name to use in the container.

