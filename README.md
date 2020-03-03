# histidine
CLI histogram tool

## usage

```
Usage of ./histidine:
  -f string
    	Input format or units [d=golang duration h=hours m=minutes s=seconds i=milliseconds u=microseconds n=nanoseconds (default "d")
  -version
    	Show version
```

See [time.ParseDuration](https://pkg.go.dev/time?tab=doc#ParseDuration) for golang duration format details.

## Example Usage

```
cat samples | ./histidine -f s
Distribution in seconds:
Total: 50
70.125 	 ................................
98 	 ........
121.125	 ................................
160.5 	 ........
177.5 	 ............
214.25 	 ................
248.75 	 ................
297.17 	 ............
336.5 	 ....
368.5 	 ........
401.5 	 ....
469 	 ........
569.5 	 ....
607.5 	 ....
665.5 	 ....
774.83 	 ............
869.5 	 ....
1029.5 	 ....
1113.5 	 ....
1356.5 	 ....

p50 seconds: 214.25
p90 seconds: 774.8333333333334
p95 seconds: 1029.5
p99 seconds: 1356.5
```
