# TUI CV

## Mandatory screenshot ;)

![Screenshot](https://raw.githubusercontent.com/maitesin/tui-cv/master/screenshot.png)

## Building the Docker Image

```
docker build -t resume .
docker tag resume maitesin/resume
docker login
docker push maitesin/resume
```
