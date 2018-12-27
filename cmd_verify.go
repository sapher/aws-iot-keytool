package main

type verifyCommand struct {
	KeystoreFile     string `long:"keystore-file" description:"JKS keystore file" required:"true"`
	KeystorePassword string `long:"keystore-password" description:"Keystore password" required:"true"`
	Alias            string `long:"alias" description:"Alias in keystore" required:"true"`
}

var verifCmd verifyCommand

func (cmd *verifyCommand) Execute(args []string) error {
	verifyAliasKeystore(cmd.KeystoreFile, cmd.KeystorePassword, cmd.Alias)
	return nil
}

func init() {
	parser.AddCommand(
		"verify",
		"Verify that a generated jks keystore contain a specific alias",
		"Verify that a generated jks keystore contain a specific alias",
		&verifCmd,
	)
}
