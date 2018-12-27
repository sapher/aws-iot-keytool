package main

type createCommand struct {
	KeystoreFilename string `long:"keystore-filename" description:"Filename for the generated JKS keystore" required:"true"`
	CertificateFile  string `long:"certificate-file" description:"Certificate file issued by AWS IoT" required:"true"`
	PrivateKeyFile   string `long:"private-key-file" description:"Private key file issued by AWS IoT" required:"true"`
	Alias            string `long:"alias" description:"Alias for the certificate and private key in keystore" required:"true"`
	Password         string `long:"password" description:"Password to encrypt private key in keystore" required:"true"`
}

var createCmd createCommand

func (cmd *createCommand) Execute(args []string) error {
	createKeystoreWithCertificates(cmd.KeystoreFilename, cmd.CertificateFile, cmd.PrivateKeyFile, cmd.Alias, cmd.Password)
	return nil
}

func init() {
	parser.AddCommand(
		"create",
		"Create a pkcs12 keystore with certificate and private key file issued by AWS IoT",
		"Create a pkcs12 keystore with certificate and private key file issued by AWS IoT",
		&createCmd,
	)
}
