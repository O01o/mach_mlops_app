docker build --no-cache -t mach-mlops-swagger:latest .
docker run --rm -p 18180:80 mach-mlops-swagger:latest
