// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
)

// Contract is the model entity for the Contract schema.
type Contract struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// ChainType holds the value of the "chain_type" field.
	ChainType string `json:"chain_type,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID int32 `json:"chain_id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Symbol holds the value of the "symbol" field.
	Symbol string `json:"symbol,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// BlockNum holds the value of the "block_num" field.
	BlockNum uint64 `json:"block_num,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// TxTime holds the value of the "tx_time" field.
	TxTime uint32 `json:"tx_time,omitempty"`
	// ProfileURL holds the value of the "profile_url" field.
	ProfileURL string `json:"profile_url,omitempty"`
	// BaseURL holds the value of the "base_url" field.
	BaseURL string `json:"base_url,omitempty"`
	// BannerURL holds the value of the "banner_url" field.
	BannerURL string `json:"banner_url,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contract) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contract.FieldCreatedAt, contract.FieldUpdatedAt, contract.FieldDeletedAt, contract.FieldChainID, contract.FieldBlockNum, contract.FieldTxTime:
			values[i] = new(sql.NullInt64)
		case contract.FieldChainType, contract.FieldAddress, contract.FieldName, contract.FieldSymbol, contract.FieldCreator, contract.FieldTxHash, contract.FieldProfileURL, contract.FieldBaseURL, contract.FieldBannerURL, contract.FieldDescription, contract.FieldRemark:
			values[i] = new(sql.NullString)
		case contract.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contract", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contract fields.
func (c *Contract) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case contract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = uint32(value.Int64)
			}
		case contract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = uint32(value.Int64)
			}
		case contract.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = uint32(value.Int64)
			}
		case contract.FieldChainType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_type", values[i])
			} else if value.Valid {
				c.ChainType = value.String
			}
		case contract.FieldChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				c.ChainID = int32(value.Int64)
			}
		case contract.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case contract.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case contract.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				c.Symbol = value.String
			}
		case contract.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				c.Creator = value.String
			}
		case contract.FieldBlockNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_num", values[i])
			} else if value.Valid {
				c.BlockNum = uint64(value.Int64)
			}
		case contract.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				c.TxHash = value.String
			}
		case contract.FieldTxTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tx_time", values[i])
			} else if value.Valid {
				c.TxTime = uint32(value.Int64)
			}
		case contract.FieldProfileURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_url", values[i])
			} else if value.Valid {
				c.ProfileURL = value.String
			}
		case contract.FieldBaseURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field base_url", values[i])
			} else if value.Valid {
				c.BaseURL = value.String
			}
		case contract.FieldBannerURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_url", values[i])
			} else if value.Valid {
				c.BannerURL = value.String
			}
		case contract.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case contract.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				c.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Contract.
// Note that you need to call Contract.Unwrap() before calling this method if this Contract
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contract) Update() *ContractUpdateOne {
	return (&ContractClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contract) Unwrap() *Contract {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contract is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contract) String() string {
	var builder strings.Builder
	builder.WriteString("Contract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", c.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("chain_type=")
	builder.WriteString(c.ChainType)
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ChainID))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("symbol=")
	builder.WriteString(c.Symbol)
	builder.WriteString(", ")
	builder.WriteString("creator=")
	builder.WriteString(c.Creator)
	builder.WriteString(", ")
	builder.WriteString("block_num=")
	builder.WriteString(fmt.Sprintf("%v", c.BlockNum))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(c.TxHash)
	builder.WriteString(", ")
	builder.WriteString("tx_time=")
	builder.WriteString(fmt.Sprintf("%v", c.TxTime))
	builder.WriteString(", ")
	builder.WriteString("profile_url=")
	builder.WriteString(c.ProfileURL)
	builder.WriteString(", ")
	builder.WriteString("base_url=")
	builder.WriteString(c.BaseURL)
	builder.WriteString(", ")
	builder.WriteString("banner_url=")
	builder.WriteString(c.BannerURL)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(c.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Contracts is a parsable slice of Contract.
type Contracts []*Contract

func (c Contracts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
