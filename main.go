package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
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
}

func authenticate(email string, password string, host string) {
    var auth AuthSuccess

    user := &User{Email: email, Password: password}
    body, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }

    resty.R().
          SetHeader("Content-Type", "application/json").
          SetBody(body).
          SetResult(&auth).
          Post(fmt.Sprintf("https://%s/_/auth/authenticate", host))

    fmt.Println(auth.Data.Token)
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
            Value: "",
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
            Name: "label",
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
            Usage: "set encryption: cipher(DES, AES), or hasher(MD5, SHA1, SHA256), or more",
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
