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

### quran:ayah

Show an ayah (verse) of the Quran:

```
~ » deen quran:ayah 2:255
[Al-Baqara 2:255, Saheeh International] Allah - there is no deity except Him, the Ever-Living, the Sustainer of [all] existence...
```

The reference format is `surah:ayah` (e.g. `2:255` for Ayat al-Kursi). Ranges are also supported (e.g. `2:13-15`).

Multiple editions can be configured to show each verse in all editions.

## Configuration

Deen reads its configuration from `~/.config/deen/config.toml`:

```toml
[database]
path = "/home/user/.deendb"

[adhan]
method = 3
city = "London"
country = "United Kingdom"

[quran]
editions = ["en.asad"]
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

Prayer times are fetched from the [Aladhan API](https://aladhan.com) and cached locally in a BoltDB database. Quran verses are fetched from the [Al Quran Cloud API](https://alquran.cloud/api). Available editions can be found in the API documentation.

## License

Deen is available under the BSD 3-Clause License.
