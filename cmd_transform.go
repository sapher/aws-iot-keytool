package main

type transformCommand struct {
	SrcKeystoreFile      string `long:"src-keystore-file" description:"Source pcks12 keystore file" required:"true"`
	SrcKeystorePassword  string `long:"src-keystore-password" description:"Source keystore password" required:"true"`
	DestKeystoreFile     string `long:"dest-keystore-file" description:"Destination jks keystore file" required:"true"`
	DestKeystorePassword string `long:"dest-keystore-password" description:"Destination keystore password" required:"true"`
	DestKeyPassword      string `long:"dest-key-password" description:"Destination key password" required:"true"`
	Alias                string `long:"alias" description:"Alias for the certificate and private key in the keystore" required:"true"`
}

var transformCmd transformCommand

func (cmd *transformCommand) Execute(args []string) error {
	transformKeystore(cmd.SrcKeystoreFile, cmd.SrcKeystorePassword, cmd.DestKeystoreFile, cmd.DestKeystorePassword, cmd.Alias, cmd.DestKeystorePassword)
	return nil
}

func init() {
	parser.AddCommand(
		"transform",
		"Transform a pkcs12 keystore into jks keystore compatible with AWS IoT device SDK",
		"Transform a pkcs12 keystore into jks keystore compatible with AWS IoT device SDK",
		&transformCmd,
	)
}
