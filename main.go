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

func authenticate(email string, password string) {
    user := &User{Email: email, Password: password}
    body, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }

    resp, err := resty.R().
          SetHeader("Content-Type", "application/json").
          SetBody(body).
          Post("https://2s.re.kr/_/auth/authenticate")

    fmt.Println(resp, err)
}

func main() {
    app := cli.NewApp()
    app.Name = "Catswords Community CLI"
    app.Usage = "Message broadcaster and receiver"
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
    app.Copyright = "(C) 2019 Catswords Research."

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "server",
            Value: "https://2s.re.kr/",
            Usage: "set server url",
        },
        cli.StringFlag{
            Name: "lang",
            Value: "english",
            Usage: "set language",
        },
        cli.StringFlag{
            Name: "format",
            Value: "text",
            Usage: "set message type: text, json, xml, or more",
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
    }

    app.Action = func(c *cli.Context) error {
        if c.String("token") == "" {
            authenticate(c.String("email"), c.String("password"))
            return nil
        }
        
        return nil
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
