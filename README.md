# TileCast

TileCast is a program that serves map tiles for a specified geographical region. It fetches the necessary map tiles in the **mbtiles** format, stores them on disk, and then uses the **martin** executable to serve them over HTTP.

## Features

- Specify a rectangular geographical region using top-left and bottom-right coordinates.
- Automatically fetches the corresponding mbtiles for the region.
- Stores the mbtiles on disk for future use.
- Uses the **martin** executable to serve the fetched tiles over HTTP.

## License

TileCast is licensed under the MIT License. See [LICENSE](LICENSE) for more information.
