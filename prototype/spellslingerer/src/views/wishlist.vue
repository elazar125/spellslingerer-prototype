<template>
  <v-row>
    <v-col cols="4">
      <card-grid :cards="wishlist" label="Wishlist" />
    </v-col>
    <v-col cols="8">
      <v-row class="mb-2">
        <v-sheet width="100%" min-height="215px">
          <material-list oneline readonly label="Required Materials" :materials="requiredMaterials" />
        </v-sheet>
      </v-row>
      <v-row>
        <card-search :cards="cardsWithQuantity" @updateQuantity="updateQuantity($event)" />
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import CardSearch from '../components/cards/card-search.vue';
import CardGrid from '../components/cards/card-grid.vue';
import MaterialList from '../components/materials/material-list.vue';

const { mapGetters: mapCardGetters } = createNamespacedHelpers('cards');
const { mapGetters: mapMaterialGetters } = createNamespacedHelpers('materials');
const { mapGetters: mapWishlistGetters, mapActions: mapWishlistActions } = createNamespacedHelpers('wishlist');

export default {
  name: 'wishlist',
  components: {
    CardSearch,
    CardGrid,
    MaterialList,
  },
  computed: {
    ...mapCardGetters([
      'cards',
    ]),
    ...mapMaterialGetters([
      'quantityStillRequired',
    ]),
    ...mapWishlistGetters([
      'wishlist',
      'quantities',
      'cardMaterials',
    ]),
    cardsWithQuantity() {
      return this.cards.map((card) => ({ ...card, quantity: this.quantities[card.id] || 0 }));
    },
    requiredMaterials() {
      return this.cardMaterials.map(this.quantityStillRequired).filter(({ quantity }) => quantity > 0);
    },
  },
  methods: {
    ...mapWishlistActions([
      'updateQuantity',
    ]),
  },
};
</script>
