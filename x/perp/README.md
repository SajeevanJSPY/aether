# Perpetual Module

The **Perpetual Module** implements a decentralized perpetual futures market on a Cosmos SDK-based blockchain. It allows users to open leveraged long or short positions, tracks funding payments to maintain price parity with an underlying index, and handles liquidation logic to ensure solvency.

---

## 🧩 Features

- Open and close long/short positions with leverage
- Track and apply funding rate mechanics
- Query liquidity and funding information
- Identify liquidation candidates based on margin and leverage
- On-chain funding rate management and settlement

---

## 🛠️ Transactions

| RPC | Description |
|-----|-------------|
| `OpenPosition` | Opens a new long or short position with margin and leverage. |
| `ClosePosition` | Closes an open position, realizing profit/loss and releasing margin. |
| `Liquidate` | Liquidates undercollateralized positions to maintain system solvency. |
| `UpdateFundingRate` | Updates the global funding rate and index (typically by oracle/admin). |
| `ApplyFundingPayment` | Applies accumulated funding to active positions (adjusts their funding liability). |
| `SettleFundingPayment` | Allows traders to settle their accrued funding payments voluntarily. |

---

## 🔍 Queries

| RPC | Description |
|-----|-------------|
| `GetPosition` | Returns detailed information about a specific open position. |
| `GetLiquidity` | Returns available liquidity in the perpetual pool. |
| `GetLiquidationCandidates` | Returns a list of positions eligible for liquidation. |
| `GetFundingRate` | Returns the current funding rate and cumulative index. |
| `GetPositionFundingInfo` | Returns funding-related information for a specific position. |
| `GetAccruedFunding` | Returns the total accrued funding payment for a position. |

---

## 🧮 Core Concepts

### 🏦 Margin & Leverage
Users open positions by depositing margin and selecting a leverage factor. The position notional = margin × leverage.

### ↔️ Long and Short
- Longs profit when the token price rises.
- Shorts profit when the token price falls.

### 💸 Funding Rate
- Periodic payments exchanged between longs and shorts to keep the perpetual price close to the spot price.
- Funding payments are tracked per position and can be settled or applied over time.

### 💀 Liquidation
Positions that fall below the maintenance margin ratio are flagged for liquidation to protect the protocol from bad debt.

---
