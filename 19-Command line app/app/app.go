package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{	
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name:  "servidores",
			Usage: "Busca o nome dos servidores na Internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}
