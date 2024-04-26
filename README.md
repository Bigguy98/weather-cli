## Introduce

This project is a CLI (command line interface) to get weather of cities around the world, written in Golang.

Some details about it:
- Using [cobra](https://github.com/spf13/cobra) lib to create CLI in easy way
- Using [termui](https://github.com/gizak/termui) lib to show charts on console
- Calling [Free Weather APIs](https://open-meteo.com/) to get city location & weather data 

## Demo

![Peek 2024-04-26 00-40](https://github.com/Bigguy98/weather-cli/assets/27953500/0d5a2c55-7cbd-413f-bf4d-e192f90e5d38)

## Instruction
To build cli app, run this command:
> go build -o bin/weather-cli main.go

To run builded app:
> bin/weather-cli [command]
