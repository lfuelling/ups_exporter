# ups_exporter

Prometheus exporter for PowerWalker(?)/V7(?)/FSP(?) UPS SNMP Intelligent Slot addon cards.

You'll need SSH access to copy files to and execute commands on the SNMP card!

## Prerequisites

- Go >= 1.16
- SSH Access to the UPS SNMP card

## Usage

1. Clone this repository
2. Run `env GOOS=linux GOARCH=arm GOARM=5 go build -o ups_exporter_armv5`
3. Copy the file to `/usr/sbin/ups_exporter` on the SNMP card
4. Run `chmod +x /usr/sbin/ups_exporter` on the SNMP card
5. Copy `examples/etc_rc.d_init.d_ups_exporter` to `/etc/rc.d/init.d/ups_exporter` on the SNMP card
6. Reboot the device and go to `http://<ups-ip>:9122/metrics`

### BONUS: Disabling the webserver and/or snmp

Since you can't change the admin password because of bugs, it might be a good idea to disable the web interface completely.

This can be done by commenting in the `exit 0` at the beginning of `/etc/rc.d/init.d/webserver` and rebooting the device (same goes for `snmpd` in `/etc/rc.d/init.d/snmpserver`).
