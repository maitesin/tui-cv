# TUI CV

## Building the Docker Image

```
docker build -t resume .
docker tag resume maitesin/resume
docker login
docker push maitesin/resume
```
