// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendSellOrder } from "./types/ibcdex/tx";
import { MsgCancelSellOrder } from "./types/ibcdex/tx";
import { MsgSendCreatePair } from "./types/ibcdex/tx";
import { MsgSendBuyOrder } from "./types/ibcdex/tx";
import { MsgCancelBuyOrder } from "./types/ibcdex/tx";


const types = [
  ["/code0xff.interchange.ibcdex.MsgSendSellOrder", MsgSendSellOrder],
  ["/code0xff.interchange.ibcdex.MsgCancelSellOrder", MsgCancelSellOrder],
  ["/code0xff.interchange.ibcdex.MsgSendCreatePair", MsgSendCreatePair],
  ["/code0xff.interchange.ibcdex.MsgSendBuyOrder", MsgSendBuyOrder],
  ["/code0xff.interchange.ibcdex.MsgCancelBuyOrder", MsgCancelBuyOrder],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgSendSellOrder: (data: MsgSendSellOrder): EncodeObject => ({ typeUrl: "/code0xff.interchange.ibcdex.MsgSendSellOrder", value: data }),
    msgCancelSellOrder: (data: MsgCancelSellOrder): EncodeObject => ({ typeUrl: "/code0xff.interchange.ibcdex.MsgCancelSellOrder", value: data }),
    msgSendCreatePair: (data: MsgSendCreatePair): EncodeObject => ({ typeUrl: "/code0xff.interchange.ibcdex.MsgSendCreatePair", value: data }),
    msgSendBuyOrder: (data: MsgSendBuyOrder): EncodeObject => ({ typeUrl: "/code0xff.interchange.ibcdex.MsgSendBuyOrder", value: data }),
    msgCancelBuyOrder: (data: MsgCancelBuyOrder): EncodeObject => ({ typeUrl: "/code0xff.interchange.ibcdex.MsgCancelBuyOrder", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
