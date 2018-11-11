package cmd

import (
	"github.com/urfave/cli"
	"github.com/xenolf/lego/log"
)

func createRevoke() cli.Command {
	return cli.Command{
		Name:   "revoke",
		Usage:  "Revoke a certificate",
		Action: revoke,
	}
}

func revoke(ctx *cli.Context) error {
	acc, client := setup(ctx, NewAccountsStorage(ctx))

	if acc.Registration == nil {
		log.Fatalf("Account %s is not registered. Use 'run' to register a new account.\n", acc.Email)
	}

	certsStorage := NewCertificatesStorage(ctx)
	certsStorage.CreateRootFolder()

	for _, domain := range ctx.GlobalStringSlice("domains") {
		log.Printf("Trying to revoke certificate for domain %s", domain)

		certBytes, err := certsStorage.ReadFile(domain, ".crt")
		if err != nil {
			log.Fatalf("Error while revoking the certificate for domain %s\n\t%v", domain, err)
		}

		err = client.Certificate.Revoke(certBytes)
		if err != nil {
			log.Fatalf("Error while revoking the certificate for domain %s\n\t%v", domain, err)
		}

		log.Println("Certificate was revoked.")
	}

	return nil
}
