# ethermine_exporter_go (WIP)

This is a prometheus exporter for Ethermine.org written in GO

## Usage

### Build

```
docker build -t andreasottosson/ethermine_exporter_go .
```

### Run

```
docker run -e ADDRESS='<ethereum-address>' -p 9118:9118 andreasottosson/ethermine_exporter_go
```

Then access the metrics at http://localhost:9118


## License

MIT

[Inspiration](https://github.com/jcrowthe/ethermine_exporter) Written in Python 3
