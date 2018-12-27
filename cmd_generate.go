package main

type generateCommand struct {
	DestKeystoreFile string `long:"dest-keystore-file" description:"Destination jks keystore file" required:"true"`
	CertificateFile  string `long:"certificate-file" description:"Certificate file issued by AWS IoT" required:"true"`
	PrivateKeyFile   string `long:"private-key-file" description:"Private key file issued by AWS IoT" required:"true"`
	Alias            string `long:"alias" description:"Alias for the certificate and private key in the keystore" required:"true"`
	Password         string `long:"password" description:"Password to encrypt the private key in keystore" required:"true"`
}

var generateCmd generateCommand

func (cmd *generateCommand) Execute(args []string) error {
	generateKeystore(cmd.DestKeystoreFile, cmd.CertificateFile, cmd.PrivateKeyFile, cmd.Alias, cmd.Password)
	return nil
}

func init() {
	parser.AddCommand(
		"generate",
		"Generate an usable jks keystore with AWS IoT device SDK",
		"Generate an usable jks keystore with AWS IoT device SDK",
		&generateCmd,
	)
}
