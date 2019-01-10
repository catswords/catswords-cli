package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
    "io/ioutil"
    "github.com/urfave/cli"
    "gopkg.in/resty.v1"
)

type User struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type AuthToken struct {
    Token string `json:"token"`
}

type AuthSuccess struct {
    Data AuthToken `json:"data"`
    Public bool `json:"public"`
}

type MessageContext struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Agent string `json:"agent"`
    Format string `json:"format"`
    Delimiter string `json:"delimiter"`
    Encoding string `json:"encoding"`
    Mime string `json:"mime"`
    Label string `json:"label"`
    Encryption string `json:"encryption"`
    EncryptionKey string `json:"encryption_key"`
    EncryptionIv string `json:"encryption_iv"`
    PrivateKey string `json:"private_key"`
    PublicKey string `json:"public_key"`
    HashFunction string `json:"hash_function"`
    HashValue string `json:"hash_value"`
    Mnemonic string `json:"mnemonic"`
    IntNetwork string `json:"int_network"`
    IntAddress string `json:"int_address"`
    ExtNetwork string `json:"ext_network"`
    ExtAddress string `json:"ext_address"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func authenticate(email string, password string, host string) {
    var auth AuthSuccess

    user := &User{Email: email, Password: password}
    body, err := json.Marshal(user)
    check(err)

    resty.R().
          SetHeader("Content-Type", "application/json").
          SetBody(body).
          SetResult(&auth).
          Post(fmt.Sprintf("https://%s/_/auth/authenticate", host))

    tokenBytes := []byte(auth.Data.Token)
    err2 := ioutil.WriteFile("token.dat", tokenBytes, 0644)
    check(err2)
}

func main() {
    app := cli.NewApp()
    app.Name = "Catswords Community CLI"
    app.Usage = "Message broadcaster"
    app.Version = "0.1"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        cli.Author{
            Name:  "Go Namhyeon",
            Email: "gnh1201@gmail.com",
        },
        cli.Author{
            Name:  "Catswords Community",
            Email: "support@exts.kr",
        },
        cli.Author{
            Name:  "Catswords Research",
            Email: "support@2s.re.kr",
        },
    }
    app.Copyright = "(c) 2019 Catswords Research."

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "host",
            Value: "2s.re.kr",
            Usage: "set server hostname",
        },
        cli.StringFlag{
            Name: "lang",
            Value: "english",
            Usage: "set language",
        },
        cli.StringFlag{
            Name: "email",
            Value: "",
            Usage: "set user email",
        },
        cli.StringFlag{
            Name: "password",
            Value: "",
            Usage: "set user password",
        },
        cli.StringFlag{
            Name: "token",
            Value: "",
            Usage: "set access token",
            FilePath: "token.dat",
        },
        cli.StringFlag{
            Name: "message",
            Value: "",
            Usage: "set message it will send to server",
        },
        cli.StringFlag{
            Name: "format",
            Value: "text",
            Usage: "set message type: text, json, xml, rfc5424(syslog), or more",
        },
        cli.StringFlag{
            Name: "delimiter",
            Value: ",",
            Usage: "set delimiter: comma, or pipeline, or more",
        },
        cli.StringFlag{
            Name: "encoding",
            Value: "utf-8",
            Usage: "set encoding: character set, or encapsulation, or more",
        },
        cli.StringFlag{
            Name: "mine",
            Value: "text/plain",
            Usage: "set media type: text/plain, or application/json, or more",
        },
        cli.StringFlag{
            Name: "label,labels",
            Value: "",
            Usage: "set custom label(s) to each columns",
        },
        cli.StringFlag{
            Name: "agent",
            Value: "",
            Usage: "set custom agent name",
        },
        cli.StringFlag{
            Name: "encryption,enc",
            Value: "",
            Usage: "set encryption algorithm: des, or aes, or more",
        },
        cli.StringFlag{
            Name: "encryption-key,ekey",
            Value: "",
            Usage: "set encryption key",
        },
        cli.StringFlag{
            Name: "encryption-iv,eiv",
            Value: "",
            Usage: "set encryption IV",
        },
        cli.StringFlag{
            Name: "private-key,privkey",
            Value: "",
            Usage: "set private key",
        },
        cli.StringFlag{
            Name: "public-key,pubkey",
            Value: "",
            Usage: "set public key",
        },
        cli.StringFlag{
            Name: "hash-function,hasher",
            Value: "",
            Usage: "set hash function(s) with delimiter: md5, or sha1, or sha256, or more",
        },
        cli.StringFlag{
            Name: "hash-value,hash",
            Value: "",
            Usage: "set hash value(s) with delimiter",
        },
        cli.StringFlag{
            Name: "mnemonic",
            Value: "",
            Usage: "set mnemonic",
        },
        cli.StringFlag{
            Name: "int-network,innet",
            Value: "",
            Usage: "set internal network name",
        },
        cli.StringFlag{
            Name: "int-address,inaddr",
            Value: "",
            Usage: "set address of specified internal network",
        },
        cli.StringFlag{
            Name: "ext-network,exnet",
            Value: "",
            Usage: "set external network name",
        },
        cli.StringFlag{
            Name: "ext-address,exaddr",
            Value: "",
            Usage: "set address of specified external network",
        },
    }

    app.Action = func(c *cli.Context) error {
        if c.String("token") == "" {
            authenticate(c.String("email"), c.String("password"), c.String("host"))
            return nil
        }
        
        
        
        return nil
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
