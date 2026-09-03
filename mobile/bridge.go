package mobile

import (
	"github.com/Akuna69/recovery/recovery_tool"
)

// StartRecovery ejecuta el barrido pasándole las credenciales necesarias
func StartRecovery(emergencyKitPDFPath string, targetAddress string) string {
	err := recovery_tool.Run(emergencyKitPDFPath, targetAddress)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "Transacción de recuperación enviada con éxito"
}
