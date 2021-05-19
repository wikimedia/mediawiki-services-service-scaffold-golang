# service-template-go

### Docker Quickstart
```
blubber .pipeline/blubber.yaml development > Dockerfile
docker build -t service-scaffold-golang .
docker run -p 8000:8000  service-scaffold-golang
```

Connect to `http://localhost:8000/healthz`.
