<template>
  <div>
    <v-img :src="`Cards/Icons/${card.image}`" width="50px" @click="modal = true">
      <v-img
        v-if="card.quantity > 1"
        src="Icons/two.png"
        style="position: absolute; bottom:0; right: 0;"
        width="25px"
      />
    </v-img>
    <card-detail-modal v-model="modal" :card="card" />
  </div>
</template>

<script>
import CardDetailModal from './card-details-modal.vue';

export default {
  name: 'card-grid-item',
  components: {
    CardDetailModal,
  },
  data() {
    return {
      modal: false,
    };
  },
  props: {
    card: {
      type: Object,
      required: true,
    },
  },
};
</script>
