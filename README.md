
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

To automatically start on boot:

```shell
systemctl enable ezd-daemon
```

### sources:

1. see: [Creating a Linux Service with systemd](https://medium.com/@benmorel/creating-a-linux-service-with-systemd-611b5c8b91d6) by [Benjamin Morel](https://medium.com/@benmorel)
2. [SUSE `systemctl` guide](https://documentation.suse.com/smart/systems-management/html/systemd-setting-up-service/index.html)

