# Raspberry PI Fan Control

## Build from source
```sh
make
```

## Usage
```
Usage of fancontrol:
    -off int
        (degrees Celsius) Fan shuts off at this temperature. (default 55)
    -on int
        (degrees Celsius) Fan kicks on at this temperature. (default 65)
    -pin int
        Which GPIO pin you're using to control the fan. (default 17)
    -sleep int
        (seconds) How often we check the core temperature. (default 5)
```

## systemd service
```ini
[Unit]
Description=raspberry pi fan control

[Service]
ExecStart=/usr/local/bin/fancontrol -pin 18
ExecStop=pkill -f /usr/local/bin/fancontrol

[Install]
WantedBy=default.target
```