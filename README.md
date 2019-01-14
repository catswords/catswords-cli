# catswords-cli
catswords-cli

## tutorial
```
1. tutorial video
https://youtu.be/VJhGeX5J1wA

2. autienticate
$ ./catswords-cli --email [your email] --password [your password]

3. send message
$ ./catswords-cli --message "hello world" --network-id [your network id]
```

## show help
```
$ ./catswords-cli help
NAME:
   Catswords Community CLI - Message broadcaster

USAGE:
   catswords-cli [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHORS:
   Go Namhyeon <gnh1201@gmail.com>
   Catswords Community <support@exts.kr>
   Catswords Research <support@2s.re.kr>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value                           set server hostname (default: "2s.re.kr")
   --lang value                           set language (default: "english")
   --email value                          set user email
   --password value                       set user password
   --token value                          set access token [token.dat]
   --protocol value                       set protocol: https, or http, or more (default: "https")
   --action value                         set action: send, or recv, or refresh (default: "send")
   --message value                        set message it will send to server
   --file value                           set file path instead of message
   --format value                         set message type: text, json, xml, rfc5424(syslog), or more (default: "text")
   --delimiter value                      set delimiter: comma, or pipeline, or more (default: "comma")
   --encoding value                       set encoding: character set, or encapsulation, or more (default: "utf-8")
   --mine value                           set media type: text/plain, or application/json, or more (default: "text/plain")
   --label value, --labels value          set label(s) with delimiter
   --agent value                          set custom agent name
   --encryption value, --enc value        set encryption algorithm: des, or aes, or more
   --encryption-key value, --ekey value   set encryption key
   --encryption-iv value, --eiv value     set encryption IV
   --private-key value, --privkey value   set private key
   --public-key value, --pubkey value     set public key
   --hash-function value, --hasher value  set hash function(s) with delimiter: md5, or sha1, or sha256, or more
   --hash-value value, --hash value       set hash value(s) with delimiter
   --mnemonic value                       set mnemonic
   --int-network value, --innet value     set internal network name
   --int-address value, --inaddr value    set address of specified internal network
   --ext-network value, --exnet value     set external network name
   --ext-address value, --exaddr value    set address of specified external network
   --network-id value, --netid value      set network ID
   --access-key value, --akey value       set access key
   --access-secret value, --asec value    set access secret
   --help, -h                             show help
   --version, -v                          print the version

COPYRIGHT:
   (c) 2019 Catswords Research.

```
