<template>
  <v-dialog v-model="open" max-width="1000px">
    <v-card>
      <card-list
        :cards="cards"
        :no-quantity="noQuantity"
        :display-tile-image="displayTileImage"
        @cardSelected="$emit('cardSelected', $event)"
      />
    </v-card>
  </v-dialog>
</template>

<script>
import CardList from './card-list.vue';

export default {
  name: 'card-list-modal',
  components: {
    CardList,
  },
  computed: {
    open: {
      get() {
        return this.value;
      },
      set() {
        this.$emit('input', false);
      },
    },
  },
  props: {
    value: {
      type: Boolean,
      required: true,
    },
    noQuantity: {
      type: Boolean,
      default: false,
    },
    displayTileImage: {
      type: Boolean,
      default: false,
    },
    cards: {
      type: Array,
      required: true,
    },
  },
};
</script>
