<template>
  <v-row>
    <v-col cols="4">
      <card-grid :cards="collection" label="My Collection" />
    </v-col>
    <v-col cols="8">
      <card-search :cards="cardsWithQuantity" @updateQuantity="updateQuantity($event)" />
    </v-col>
  </v-row>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import CardSearch from '../components/cards/card-search.vue';
import CardGrid from '../components/cards/card-grid.vue';

const { mapGetters: mapCollectionGetters, mapActions } = createNamespacedHelpers('collection');
const { mapGetters: mapCardGetters } = createNamespacedHelpers('cards');

export default {
  name: 'collection',
  components: {
    CardSearch,
    CardGrid,
  },
  computed: {
    ...mapCollectionGetters([
      'collection',
      'quantities',
    ]),
    ...mapCardGetters([
      'cards',
    ]),
    cardsWithQuantity() {
      return this.cards.map((card) => ({ ...card, quantity: this.quantities[card.id] || 0 }));
    },
  },
  methods: {
    ...mapActions([
      'updateQuantity',
    ]),
  },
};
</script>
