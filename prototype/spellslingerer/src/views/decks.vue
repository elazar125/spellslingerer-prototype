<template>
  <deck-search :decks="decks" />
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import DeckSearch from '../components/decks/deck-search.vue';

const { mapGetters } = createNamespacedHelpers('decks');

export default {
  name: 'decks',
  components: {
    DeckSearch,
  },
  computed: {
    ...mapGetters([
      'decks',
    ]),
  },
};
</script>
