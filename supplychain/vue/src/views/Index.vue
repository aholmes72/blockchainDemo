<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="shipment" :fields="['ShippingNumber', 'ShipTo', 'Item', 'ShipDate', ]" module="supplychain" />
		<sp-type-form type="receipt" :fields="['ShippingNumber', 'ReceiptDate', ]" module="supplychain" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
