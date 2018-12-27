# AWS IoT SDK - Keystore tool

Simple tool to generate Keystore usable with AWS IoT Device SDK for Java.
AWS IoT Device SDK is compatible only with JKS keystore.
We have to create a PKCS12 keystore first with a private key and certificate issued by [AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/device-certs-create.html), then we convert that keystore into a JKS keystore. 