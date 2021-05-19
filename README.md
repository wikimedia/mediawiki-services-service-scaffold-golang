# service-template-go

### Docker Quickstart

Build a service variant with `blubber .pipeline/blubber.yaml <variant> > Dockerfile`,
and build using regular Docker tools.


For example, build and run a `development` variant of a service with:
```
blubber .pipeline/blubber.yaml development > Dockerfile
docker build -t service-scaffold-golang .
docker run -p 8000:8000  service-scaffold-golang
```

Connect to `http://localhost:8000/healthz`.
