## shDNS

Service for automatic creation of Cloudflare DNS records.

### Usage

shDNS works based on Docker labels of containers and the container start event:

```yaml
labels:
  - shdns.type=A
  - shdns.name=test.domain.com
  - shdns.value=127.0.0.1
  - shdns.proxied=false
```

You can also run shDNS in a container:

```yaml
services:
  app:
    container_name: shdns
    image: ghcr.io/saturdayshdev/shdns:latest
    environment:
      - DOCKER_API_VERSION=1.42
      - CLOUDFLARE_API_KEY=${CLOUDFLARE_API_KEY}
      - CLOUDFLARE_EMAIL=${CLOUDFLARE_EMAIL}
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
```

### Environment

You can see a list of all environment variables below.

| Variable           | Description                    |
| ------------------ | ------------------------------ |
| CLOUDFLARE_API_KEY | Global API key for Cloudflare. |
| CLOUDFLARE_EMAIL   | Email for Cloudflare.          |

### Commits

All commits MUST contain a suitable prefix.

### Team

- Oskar Wójcikiewicz (oskar@saturdaysheroes.dev)
