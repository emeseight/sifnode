import { ActionContext } from "..";
import { Asset, AssetAmount } from "../../entities";
import notify from "../../api/utils/Notifications";
import JSBI from "jsbi";

function isOriginallySifchainNativeToken(asset: Asset) {
  return ["erowan", "rowan"].includes(asset.symbol);
}
// listen for 50 confirmations
// Eventually this should be set on ebrelayer
// to centralize the business logic
const ETH_CONFIRMATIONS = 50;

export default ({
  api,
  store,
}: ActionContext<
  "SifService" | "EthbridgeService" | "EthereumService",
  "wallet"
>) => {
  const actions = {
    getSifTokens() {
      return api.SifService.getSupportedTokens();
    },
    getEthTokens() {
      return api.EthereumService.getSupportedTokens();
    },
    async unpeg(assetAmount: AssetAmount) {
      const [lockOrBurnFn, feeNumber] = isOriginallySifchainNativeToken(
        assetAmount.asset
      )
        ? [api.EthbridgeService.lockToEthereum, "18332015000000000"]
        : [api.EthbridgeService.burnToEthereum, "16164980000000000"];

      const tx = await lockOrBurnFn({
        assetAmount,
        ethereumRecipient: store.wallet.eth.address,
        fromAddress: store.wallet.sif.address,
        feeAmount: AssetAmount(Asset.get("ceth"), JSBI.BigInt(feeNumber)),
      });

      return await api.SifService.signAndBroadcast(tx.value.msg);
    },
    async peg(assetAmount: AssetAmount) {
      const lockOrBurnFn = isOriginallySifchainNativeToken(assetAmount.asset)
        ? api.EthbridgeService.burnToSifchain
        : api.EthbridgeService.lockToSifchain;

      return await new Promise<any>(done => {
        lockOrBurnFn(store.wallet.sif.address, assetAmount, ETH_CONFIRMATIONS)
          .onTxHash(done)
          .onError(err => {
            const payload: any = err.payload;
            notify({ type: "error", message: payload.message ?? err });
          })
          .onComplete(({ txHash }) => {
            notify({
              type: "success",
              message: `Transfer ${txHash} has succeded.`,
            });
          });
      });
    },
  };

  return actions;
};