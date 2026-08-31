# Decypharr (TwistedRat fork)

![ui](docs/src/assets/images/index.png)

> **This is a personal fork of [sirrobot01/decypharr](https://github.com/sirrobot01/decypharr)** with fixes and features not yet merged upstream. The `main` branch is the active deploy target. See [Fork Changes](#fork-changes) below.

**Decypharr** is a **Media Gateway** for Debrid services and Usenet written in Go.

## What is Decypharr?

Decypharr provides a unified interface for Sonarr, Radarr, and other *Arr applications to access Debrid providers and
Usenet streaming.

## Features

- Mock Qbittorent and Sabnzbd API that supports the Arrs (Sonarr, Radarr, Lidarr etc)
- Multiple Debrid and usenet providers support with a single interface
- Direct Usenet streaming via NNTP (no separate download client required)
- Read-only NFSv4 and SMB servers for the same libraries, including custom virtual folders

## Supported Debrid Providers

- [Real Debrid](https://real-debrid.com)
- [Torbox](https://torbox.app)
- [Debrid Link](https://debrid-link.com)
- [All Debrid](https://alldebrid.com)
- [Premiumize](https://www.premiumize.me)

## Quick Start

### Docker (Recommended)

```yaml
services:
  decypharr:
    image: cy01/blackhole:latest
    container_name: decypharr
    ports:
      - "8282:8282"
      # Optional: NFSv4 (when NFS is enabled in Settings)
      # - "2049:20490/tcp"
      # Optional: SMB — Windows clients require host port 445 (when SMB is enabled in Settings)
      # - "445:1445/tcp"
    volumes:
      - /mnt/:/mnt:rshared
      - ./configs/:/app # config.json must be in this directory
    restart: unless-stopped
    devices:
      - /dev/fuse:/dev/fuse:rwm
    cap_add:
      - SYS_ADMIN
    security_opt:
      - apparmor:unconfined
```

> Prefer not to self-host? A managed Decypharr instance is available
> via [ElfHosted](https://store.elfhosted.com/product/decypharr/?utm_source=github&utm_medium=readme&utm_campaign=decypharr-readme),
> preconfigured alongside Sonarr/Radarr to route requests to your debrid provider (7-day trial).

## Documentation

For complete documentation, please visit our [Documentation](https://docs.decypharr.com).

## Fork Changes

Active fixes on this fork's `main` branch (based on upstream v2.5):

| Commit | Fix |
|--------|-----|
| `2d84e49` | **NNTP 430 → repair sweep**: article-not-found errors during streaming now mark the entry dirty so the repair sweep triggers an arr re-search instead of retrying indefinitely |
| `6570c42` | **Multi-episode import**: after Sonarr imports a grabbed episode, a folder-based re-import fires 30s later so secondary episodes in multi-episode files (e.g. `S09E23E24`) are also linked — fixes season packs where two episodes share one file |
| `7a17a66` | **Blu-ray main feature selection**: when a torrent contains multiple `.m2ts` files, only the largest (the main feature) is exposed; menus, trailers, and bonus streams are hidden to prevent Plex from flooding the debrid CDN with probe requests |

Config options added by this fork: `bd_main_file_only`, `download_uid`, `download_gid` — see [Configuration Reference](docs/src/content/docs/guides/configuration.md#fork-specific-features-twistedratdecypharr).

Old fork history is preserved in the `backup/main-2026-08-30` branch.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
