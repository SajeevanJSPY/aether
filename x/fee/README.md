# x/fee Module

The x/fee module is a Cosmos SDK module responsible for handling protocol fee configuration, tracking collected fees, and providing controlled mechanisms for parameter updates via governance or an authority account.


## ✨ Features

- 🔧 Configurable Fee Parameters - Allows setting protocol-level fee percentage via governance or designated authority.
- 💰 Fee Calculation - Modules can use this module to calculate fees based on transaction values.
- 📊 Tracking Total Collected Fees - Keeps a running total of all protocol fees collected.


## ⚙️ CLI Commands

### 🔍 Query

```bash
    fee query params
```
Displays the current protocol fee parameters.
```bash
    fee query total-fee-collected
```
Returns the total fee collected by the module.


### ✏️ Transactions

```bash
    fee tx update-params [flags]
```
Updates the fee parameters (e.g., fee percentage). Only executable by the module’s authority (typically the governance module or a specified admin address).


## ✅ Integration

Other modules can import and use the Keeper.CalculateFee(sdk.Coin) function to compute fees based on the current protocol percentage.
