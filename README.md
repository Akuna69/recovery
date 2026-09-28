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
