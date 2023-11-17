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

## Prerequisites

1. Add all necesssary environment variables. Check [Environment](#environment) for more info.
2. Consume the available aliases by running `source dev.sh`.
3. Run the app using the `start` alias.

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
