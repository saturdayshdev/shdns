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

All commits MUST contain the suitable prefix. Acceptable prefixes:

- `ci`: Changes to our CI configuration files and scripts (e.g. Travis, CircleCI)
- `chore`: A tech task that's not actually a feature but still needs to be done (e.g. test integration with a 3rd-party service)
- `docs`: Documentation only changes
- `feat`: A new feature
- `fix`: A bug fix
- `perf`: A code change that improves performance
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `test`: Adding missing tests or correcting existing tests

### Team

- Oskar Wójcikiewicz (oskar@saturdaysheroes.dev)
