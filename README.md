
# Ezd Daemon

## Running as service

The canonical way to run this program on *nix systems is `systemd` as a service.

Create a file `/etc/systemd/system/ezd-daemon.service`

```
[Unit]
Description=Ezd Daemon Service
After=network.target
StartLimitIntervalSec=0
[Service]
Type=simple
Restart=always
RestartSec=1
User=pi
ExecStart=/path/to/ezd-daemon

[Install]
WantedBy=multi-user.target
```

Where `ExecStart` points to the compiled binary.

Start the service:

```shell
systemctl start ezd-daemon
```

Stop the service: 

```shell
systemctl stop ezd-daemon
```

To automatically start on boot:

```shell
systemctl enable ezd-daemon
```

## Logs

Structured json logs are written to `./logs`.

To view stdout output from the running service:

```sh
journalctl -f -u ezd-daemon
```

### sources:

1. see: [Creating a Linux Service with systemd](https://medium.com/@benmorel/creating-a-linux-service-with-systemd-611b5c8b91d6) by [Benjamin Morel](https://medium.com/@benmorel)
2. [SUSE `systemctl` guide](https://documentation.suse.com/smart/systems-management/html/systemd-setting-up-service/index.html)


## pgsql DB

### Initialization

Build and run the docker container:

```sh
./scripts/pg-db.sh build
./scripts/pg-db.sh
```

The database should initialize on first run. If it does not, follow the steps below.

Create a shell in the running container and execute `init.sql`:

```sh
docker exec -it ezd-daemon-postgres bin/bash
```

Inside the container, replacting dbname/username with your values:

```sh
psql ezdd_db -Uezd_daemon < ~/init.sql
```
