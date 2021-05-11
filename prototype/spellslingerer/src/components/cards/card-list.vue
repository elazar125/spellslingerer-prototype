<template>
  <div>
    <card-list-item
      v-for="card in cards"
      v-bind="card"
      :key="card.id"
      :no-quantity="noQuantity"
      :display-tile-image="displayTileImage"
      @updateQuantity="$emit('updateQuantity', $event)"
      @cardSelected="$emit('cardSelected', $event)"
    />
  </div>
</template>

<script>
import CardListItem from './card-list-item.vue';

export default {
  name: 'card-list',
  components: {
    CardListItem,
  },
  props: {
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
