package main 

/*
  Creates encrypted config files for server credentials
  
  Intended to be used as a standalone app kept separate from
  any other app that uses the encrypted config file

  Because it uses an external private key and an embedded salt string,
  both of these values must be the same between this app as it is built
  and the app that uses the config

  To change the default salt string:
  go build/test/install -ldflags "-X secrets.salty my-new-salt-string"
*/

import (
    "flag"
    "fmt"
    "log"
    "io/ioutil"
    "os"
    "github.com/paulstuart/secrets"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [options] <encode | decode>\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(1)
}

func main() {
    var private, public, keyfile string
	flag.StringVar(&private, "u", private, "private config file with unencoded data")
	flag.StringVar(&public,  "e", public,  "public config file with encoded data")
	flag.StringVar(&keyfile, "k", keyfile, "file containing private key")
    flag.Usage = usage
    flag.Parse()
    args := flag.Args()
    if len(keyfile) == 0 {
        fmt.Println("\nNo key file specified\n")
        usage()
    }
    if len(private) == 0 {
        fmt.Println("\nNo private config file specified\n")
        usage()
    }
    if len(public) == 0 {
        fmt.Println("\nNo public config file specified\n")
        usage()
    }
    keydata, err := ioutil.ReadFile(keyfile) 
    if err != nil {
        log.Fatal(err.Error())
    }
    secrets.SetKey(string(keydata))
    switch {
    case len(args) == 0:
        usage()
    case "encode" == args[0]:
        if config, err := secrets.ConfigLoad(private); err != nil {
            log.Fatal(err.Error())
        } else {
            config.Private()
            config.Save(public)
        }
    case "decode" == args[0]:
        if config, err := secrets.ConfigLoad(public); err != nil {
            log.Fatal(err.Error())
        } else {
            config.Public()
            config.Save(private)
        }
    default:
        usage()
    }
}

