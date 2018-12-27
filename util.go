package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// exitCommandNotFound Check if a command exist in PATH, exit program if not found
func exitCommandNotFound(commandName string) {
	_, err := exec.LookPath(commandName)
	if err != nil {
		fmt.Printf("error: `%s` command not found in PATH\n", commandName)
		os.Exit(0)
	}
}

// executeCmdAndWait Execute a command, wait for completion
func executeCmdAndWait(cmd *exec.Cmd) {
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		printBytes(stderr)
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		printBytes(stderr)
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		printBytes(stderr)
		log.Fatal(err)
	}
}

// printBytes Print array of byte to console
func printBytes(bytes []byte) {
	fmt.Printf("%s\n", string(bytes))
}

// createKeystoreWithCertificates Create a keystore with certificates generate by AWS IoT
func createKeystoreWithCertificates(keystoreFileName string, certificateFile string, privateKeyFile string, alias string, keyPassword string) {
	cmd := exec.Command("openssl", "pkcs12", "-export", "-in", certificateFile, "-inkey", privateKeyFile, "-out", keystoreFileName, "-name", alias, "-password", "pass:"+keyPassword)
	executeCmdAndWait(cmd)
}

// transformKeystore Transform a PCKS12 keystore into a JKS keystore usable with AWS IoT Device SDK
func transformKeystore(srcKeystoreFile string, srcKeystorePassword string, destKeystoreFile string, destKeystorePassword string, alias string, destKeyPassword string) {
	cmd := exec.Command("keytool", "-importkeystore", "-srckeystore", srcKeystoreFile, "-srcstoretype", "PKCS12", "-srcstorepass", srcKeystorePassword, "-alias", alias, "-deststorepass", destKeystorePassword, "-destkeypass", destKeyPassword, "-destkeystore", destKeystoreFile)
	executeCmdAndWait(cmd)
}

// verifyAliasKeystore Verify that a keystore contain a specific alias secured by a password
func verifyAliasKeystore(keystoreFile string, keystorePassword string, alias string) {
	cmd := exec.Command("keytool", "-list", "-keystore", keystoreFile, "-storepass", keystorePassword, "-alias", alias)
	executeCmdAndWait(cmd)
}

// generateKeystore Generate a keystore. Create the keystore, transform into a keystore usable with AWS IoT Device SDK and verify that is has been successfully generated
func generateKeystore(destKeystoreFile string, certificateFile string, privateKeyFile string, alias string, keyPassword string) {
	temporaryKeystoreFile := "file"
	createKeystoreWithCertificates(temporaryKeystoreFile, certificateFile, privateKeyFile, alias, keyPassword)
	transformKeystore(temporaryKeystoreFile, keyPassword, destKeystoreFile, keyPassword, alias, keyPassword)
	verifyAliasKeystore(destKeystoreFile, keyPassword, alias)
	os.Remove(temporaryKeystoreFile)
}
