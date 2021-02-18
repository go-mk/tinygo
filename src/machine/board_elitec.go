// +build arduino

package machine

// Return the current CPU frequency in hertz.
func CPUFrequency() uint32 {
	return 16000000
}

// Digital pins
const (
	D0  = PD2
	D1  = PD3
	D2  = PD1
	D3  = PD0
	D4  = PD4
	D6  = PD7
	D7  = PE6
	D8  = PB4
	D9  = PB5
	D10 = PB6
	D11 = PB7
	D12 = PC6
	D13 = PC7
	D14 = PB3
	D15 = PB1
	D16 = PB2
	D18 = PF7
	D19 = PF6
	D20 = PF5
	D21 = PF4
)

// LEDs on the EliteC
const (
  LED      = LED1
  LED1 Pin = PD5 // TX LED
  LED2 Pin = PB0 // RX LED
)

// I2C
const (
  SDA_PIN = PD1
  SCL_PIN = PD0
)

// UART pins
const (
	UART_TX_PIN Pin = PD1
	UART_RX_PIN Pin = PD0
)
