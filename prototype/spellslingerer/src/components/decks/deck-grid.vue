<template>
  <v-container>
    <v-row>
      <v-col
        cols="auto"
        v-for="deck in decks"
        :key="deck.name"
      >
        <deck-grid-item v-bind="deck" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import DeckGridItem from './deck-grid-item.vue';

export default {
  name: 'deck-grid',
  props: {
    decks: {
      type: Array,
      required: true,
    },
  },
  components: {
    DeckGridItem,
  },
};
</script>
