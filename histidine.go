package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/VividCortex/gohistogram"
)

var (
	Version = "dev"
)

func main() {
	var (
		doVersion bool
		format    = "d"
	)

	flag.BoolVar(&doVersion, "version", false, "Show version")
	flag.StringVar(&format, "f", format, "Input format or units [d=golang duration h=hours m=minutes s=seconds i=milliseconds u=microseconds n=nanoseconds")
	flag.Parse()

	if doVersion {
		fmt.Println("histidine", Version)
		os.Exit(0)
	}

	var conv ConvFunc
	conv, err := format2Conv(format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		flag.Usage()
	}

	nh := gohistogram.NewHistogram(20)
	scanner := bufio.NewScanner(os.Stdin)

	var lineNum uint
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		sec, err := conv(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error line %d: %s\n", lineNum, err)
			os.Exit(1)
		}

		nh.Add(sec)
	}
	fmt.Println("Distribution in seconds:")
	fmt.Println(nh.String())
	fmt.Println("p50    sec:", nh.Quantile(0.5))
	fmt.Println("p90    sec:", nh.Quantile(0.9))
	fmt.Println("p95    sec:", nh.Quantile(0.95))
	fmt.Println("p99    sec:", nh.Quantile(0.99))
	fmt.Println("p99.9  sec:", nh.Quantile(0.999))
	fmt.Println("p99.99 sec:", nh.Quantile(0.9999))

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func format2Conv(format string) (ConvFunc, error) {
	var conv ConvFunc
	switch format {
	case "d":
		conv = convDuration
	case "h":
		conv = convF(1.0 / (60.0 * 60.0))
	case "m":
		conv = convF(1.0 / 60.0)
	case "s":
		conv = convF(1)
	case "i":
		conv = convF(1000)
	case "u":
		conv = convF(1000000)
	case "n":
		conv = convF(1000000000)
	default:
		return nil, fmt.Errorf("invalid format %q", format)
	}
	return conv, nil
}

type ConvFunc func(string) (float64, error)

func convDuration(s string) (float64, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return d.Seconds(), nil
}

func convF(d float64) ConvFunc {
	return func(s string) (float64, error) {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return f / d, nil
	}
}
