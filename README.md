# Deen

Command-line companion for prayer timings, Quran verses, and more.

## Install

```sh
go install github.com/hjr265/deen@latest
```

## Usage

### adhan:next

Show the next prayer time:

```
~ » deen adhan:next
Fajr 04:05
```

## Configuration

Deen reads its configuration from `~/.config/deen/config.toml`:

```toml
[Database]
Path = "/home/user/.deendb"

[Adhan]
Method = 3
City = "London"
Country = "United Kingdom"
```

### Prayer Calculation Methods

| Method | Description |
|--------|-------------|
| 0 | Shia Ithna Ansari |
| 1 | University of Islamic Sciences, Karachi |
| 2 | Islamic Society of North America |
| 3 | Muslim World League |
| 4 | Umm Al-Qura University, Makkah |
| 5 | Egyptian General Authority of Survey |
| 7 | Institute of Geophysics, University of Tehran |
| 8 | Gulf Region |
| 9 | Kuwait |
| 10 | Qatar |
| 11 | Majlis Ugama Islam Singapura, Singapore |
| 12 | Union Organization Islamic De France |
| 13 | Diyanet Isleri Baskanligi, Turkey |
| 14 | Spiritual Administration of Muslims of Russia |

Prayer times are fetched from the [Aladhan API](https://aladhan.com) and cached locally in a BoltDB database.

## License

Deen is available under the BSD 3-Clause License.
