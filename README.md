# script.shutdown-or-reboot
This is a reimplementation of the [first script I wrote](https://github.com/JagdCake/bash.scripts/blob/master/scripts/shut_down.sh) while learning BASH. Now I'm learning Go.

It does the same things the original did, but better:
- Choose between shutdown, reboot and cancel
- Close all applications specified in [config.go](./config/config.go)
- Log...
    - system startup time and date
    - system shutdown / reboot time and date
    - system uptime
    
    ...to a log file specified in [config.go](./config/config.go)
- Power system off OR reboot after a specified amount of time

### Usage:
- update the config file
- run `make` to test and build the script
- execute `./shutdown-or-reboot`
