package miner

type InscribeMint struct {
	P        string `json:"p" mapstructure:"p"`               // Protocol: "bracle-20"
	Op       string `json:"op" mapstructure:"op"`             // Operation, "mint"
	Tick     string `json:"tick" mapstructure:"tick"`         // Tick Name
	Amt      string `json:"amt" mapstructure:"amt"`           // Amount, can be a string format number
	Solution string `json:"solution" mapstructure:"solution"` // Solution: <TICK>:<BTC_ADDRESS>:<DEPLOY_BLOCK_HEADER>:<NONCE>
}

type InscribeTransfer struct {
	P    string `json:"p" mapstructure:"p"`                       // Protocol: "bracle-20"
	Op   string `json:"op" mapstructure:"op"`                     // Operation, "transfer"
	Tick string `json:"tick" mapstructure:"tick"`                 // Tick Name
	Amt  string `json:"amt" mapstructure:"amt"`                   // Amount, can be a string format number
	To   string `json:"to,omitempty" mapstructure:"to,omitempty"` // To Address
}
