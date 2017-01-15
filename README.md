# Uptime

This program is similar to program [uptime](http://man.he.net/?topic=uptime&section=all).
But mine shows when computer powers up.
Of course I can use `uptime -s` to know this, but then it shows only when computer powers up without showing other parameters.

And I wrote simple program in Go.

## Usage

```
~$ go get github.com/janczer/uptime
~$
~$ uptime
~$ 21:33:40 up 0:58, power on 20:34:48, load average: 0.68, 1.52, 1.34
```