package celestia

import (
	"encoding/base64"
	"time"

	"github.com/shopspring/decimal"
)

// HeaderResponse -
type HeaderResponse struct {
	Header       Header       `json:"header"`
	Commit       Commit       `json:"commit"`
	ValidatorSet ValidatorSet `json:"validator_set"`
	Dah          Dah          `json:"dah"`
}

// ValidatorSet -
type ValidatorSet struct {
	Validators []Validator `json:"validators"`
	Proposer   Validator   `json:"proposer"`
}

// Header -
type Header struct {
	Version            Version   `json:"version"`
	ChainID            string    `json:"chain_id"`
	Height             uint64    `json:"height"`
	Time               time.Time `json:"time"`
	LastBlockID        BlockID   `json:"last_block_id"`
	LastCommitHash     string    `json:"last_commit_hash"`
	DataHash           string    `json:"data_hash"`
	ValidatorsHash     string    `json:"validators_hash"`
	NextValidatorsHash string    `json:"next_validators_hash"`
	ConsensusHash      string    `json:"consensus_hash"`
	AppHash            string    `json:"app_hash"`
	LastResultsHash    string    `json:"last_results_hash"`
	EvidenceHash       string    `json:"evidence_hash"`
	ProposerAddress    string    `json:"proposer_address"`
}

// Version -
type Version struct {
	Block int `json:"block"`
}

// BlockID -
type BlockID struct {
	Hash  string       `json:"hash"`
	Parts BlockIDParts `json:"parts"`
}

// BlockIDParts -
type BlockIDParts struct {
	Total int    `json:"total"`
	Hash  string `json:"hash"`
}

// Commit -
type Commit struct {
	Height     uint64      `json:"height"`
	Round      int         `json:"round"`
	BlockID    BlockID     `json:"block_id"`
	Signatures []Signature `json:"signatures"`
}

// Signature -
type Signature struct {
	BlockIDFlag      uint64    `json:"block_id_flag"`
	ValidatorAddress string    `json:"validator_address"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `json:"signature"`
}

// Validator -
type Validator struct {
	Address          string `json:"address"`
	PubKey           string `json:"pub_key"`
	VotingPower      int    `json:"voting_power"`
	ProposerPriority int    `json:"proposer_priority"`
}

// Dah -
type Dah struct {
	RowRoots    []string `json:"row_roots"`
	ColumnRoots []string `json:"column_roots"`
}

// NamespaceData -
type NamespaceData struct {
	Data   []string `json:"data,omitempty"`
	Height uint64   `json:"height"`
}

// GetBytes - returns array of bytes array from `Data` field of `NamespaceData` response
func (nd NamespaceData) GetBytes() ([][]byte, error) {
	result := make([][]byte, len(nd.Data))
	for i := range nd.Data {
		decoded, err := base64.StdEncoding.DecodeString(nd.Data[i])
		if err != nil {
			return nil, err
		}
		result[i] = decoded
	}

	return result, nil
}

// NamespaceShares -
type NamespaceShares struct {
	Shares []string `json:"shares,omitempty"`
	Height uint64   `json:"height"`
}

// DataAvailableResponse -
type DataAvailableResponse struct {
	Available                 bool   `json:"available"`
	ProbabilityOfAvailability string `json:"probability_of_availability"`
}

// Balance -
type Balance struct {
	Denom  string          `json:"denom"`
	Amount decimal.Decimal `json:"amount"`
}

// SubmitTx -
type SubmitTx struct {
	Tx string `json:"tx"`
}

// SubmittedTx -
type SubmittedTx struct {
	Txhash    string `json:"txhash"`
	Codespace string `json:"codespace"`
	Code      int64  `json:"code"`
	RawLog    string `json:"raw_log"`
	Logs      []Log  `json:"logs,omitempty"`
	GasWanted uint64 `json:"gas_wanted"`
}

// SubmitPfd -
type SubmitPfd struct {
	NamespaceID string `json:"namespace_id"`
	Data        string `json:"data"`
	GasLimit    uint64 `json:"gas_limit"`
}

// SubmittedPfd -
type SubmittedPfd struct {
	Height int     `json:"height"`
	Txhash string  `json:"txhash"`
	Data   string  `json:"data"`
	RawLog string  `json:"raw_log"`
	Logs   []Log   `json:"logs"`
	Events []Event `json:"events"`
}

// Log -
type Log struct {
	MsgIndex int64   `json:"msg_index"`
	Events   []Event `json:"events"`
}

// Event -
type Event struct {
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

// Attribute -
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index *bool  `json:"index"`
}

// IsIndex -
func (attr Attribute) IsIndex() bool {
	if attr.Index == nil {
		return false
	}
	return *attr.Index
}
