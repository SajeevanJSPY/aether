# Pool Module

The **Pool Module** provides the foundational infrastructure for managing token liquidity pools in the Aether blockchain. It allows liquidity providers (LPs) to deposit collateral tokens into a shared pool and enables other modules (such as a perpetual trading module) to access pooled liquidity for trading or collateralization.

---

## ✨ Features

- **Liquidity Pool Creation**: Bootstrap a new liquidity pool with a specified collateral token.
- **Deposits**: Allow users to deposit tokens and receive pool shares.
- **Withdrawals**: Allow LPs to exit the pool by burning their shares and redeeming their collateral.
- **Accurate Accounting**: Maintain records of each depositor's share.
- **Modular Design**: Built for reuse across various financial modules like perpetuals or options.

---


## 🛠️ Messages

- `CreatePool(MsgCreatePool)`: Creates a new liquidity pool with specified token configuration.
- `RemovePool(MsgRemovePool)`: Removes an existing pool from the active set.
- `Deposit(MsgDeposit)`: Deposits tokens into the specified pool and mints LP shares.
- `Withdraw(MsgWithdraw)`: Burns LP shares and returns underlying tokens.
- `UpdateParams(MsgUpdateParams)`: Updates protocol parameters; callable only by the module authority.
  
Each message is authenticated and validated using Cosmos SDK's `MsgServer`.

---

## 🔍 Queries

- `GetPoolInfo(QueryGetPoolInfoRequest)`: Get details about a specific pool.
- `GetAllPoolInfos(QueryGetAllPoolInfosRequest)`: Retrieve a list of all active pools.
- `GetUserInfo(QueryGetUserInfoRequest)`: Get deposit info for a specific user in a pool.
- `Params(QueryParamsRequest)`: View the current protocol parameters.

---

