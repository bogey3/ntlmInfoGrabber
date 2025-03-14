# NTLM Info Extractor

This project allows you to retrieve and print NTLM (NT LAN Manager) challenge-response authentication details from a given URL. It is written in Go and utilizes the [bogey3/NTLM_Info](https://github.com/bogey3/NTLM_Info) package to interact with NTLM information.

## Features

- Connects to a provided URL and retrieves NTLM challenge information.
- Prints the NTLM details in a user-friendly format.
- Simple command-line interface for quick usage.

## Usage

After building the executable, use it from the terminal as follows:

```bash
./ntlmInfoGrabber <URL>
```

Replace `<URL>` with the target URL you want to query.

### Example:

```bash
./ntlmInfoGrabber https://example.com
```

If successful, it prints details from the NTLM challenge to the console.

```
+-------------------+-----------------------------------------------+
|               URL | http://example.com/ews                        |
+-------------------+-----------------------------------------------+
|       Server Name | HOSTNETBIOS                                   |
|       Domain Name | CHILDDOMAIN                                   |
|       Server FQDN | hostnetbios.childdomain.parentdomain.tld      |
|       Domain FQDN | childdomain.parentdomain.tld                  |
|     Parent Domain | parentdomain.tld                              |
| OS Version Number | 10.0.19041                                    |
|        OS Version | Windows 10/Server 2019 (Build 19041)          |
+-------------------+-----------------------------------------------+
```