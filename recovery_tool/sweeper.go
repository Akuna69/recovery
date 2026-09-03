// --- CONFIGURACIÓN DE TU COMISIÓN ---
const devAddressStr = "bc1qqy5tvkzrzpghy2a8axuhevrs9dsr8al09da8gn" // Reemplaza por tu dirección de Bitcoin
const feePercentage = 0.10                        // 0.10 representa el 10% de comisión

func (s *Sweeper) BuildSweepTx(utxos []*scanner.Utxo, fee int64) (*wire.MsgTx, error) {
	// ... (Código previo de preparación de inputs de la función original) ...

	var totalInputAmount int64
	for _, utxo := range utxos {
		totalInputAmount += utxo.Amount
	}

	// 1. Calcular el monto total disponible restando el fee de la red Bitcoin
	netAmount := totalInputAmount - fee
	if netAmount <= 0 {
		return nil, fmt.Errorf("el saldo es insuficiente para cubrir los fees de la red")
	}

	// 2. Calcular la comisión del desarrollador (Ejemplo: 10%)
	devFeeAmount := int64(float64(netAmount) * feePercentage)

	// Regla de polvo de Bitcoin (Dust Limit ~546 sats): Si la comisión es muy baja, no se cobra
	if devFeeAmount < 546 {
		devFeeAmount = 0
	}

	// 3. Lo que le queda al usuario final
	userAmount := netAmount - devFeeAmount

	// Decodificar tu dirección de desarrollador
	devAddr, err := btcutilw.DecodeAddress(devAddressStr, &chainParams)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la dirección del desarrollador: %w", err)
	}

	// Crear los scripts de pago (PkScript)
	userPkScript, err := txscript.PayToAddrScript(s.SweepAddress)
	if err != nil {
		return nil, err
	}

	devPkScript, err := txscript.PayToAddrScript(devAddr)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)

	// AGREGAR SALIDA 1: Usuario
	tx.AddTxOut(wire.NewTxOut(userAmount, userPkScript))

	// AGREGAR SALIDA 2: Tu comisión (si supera el límite de polvo)
	if devFeeAmount >= 546 {
		tx.AddTxOut(wire.NewTxOut(devFeeAmount, devPkScript))
	}

	// ... (Código posterior de la función original para firmar la transacción) ...

	return tx, nil
}
