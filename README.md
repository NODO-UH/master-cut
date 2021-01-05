# Maste Cut

Master Cut is a tool for manage CUT and UNCUT the user quota. Each Quota Scraper service notify to Master Cut if an user over quota.

## Table of Contents

- [Maste Cut](#maste-cut)
  - [Table of Contents](#table-of-contents)
  - [Configuration](#configuration)

## Configuration

Master Cut recive requests from diferent Quota Scraper services, and each one is part of a group. For example, group `Trabajadores` have two Quota Scraper services, one per server, and shares the CUT file. So we need to tell to Master Cut what groups exists, and CUT file and script. For that, we use JSON configuration like that:

```json
{
    "groups": [
        {
            "name": "trabajadores",
            "file": "trabajadores.cut",
            "script": "trabajadores.sh"
        }
    ]
}
```

Configuration path is set with argument `--conf`.
