<template>
  <deck-search :decks="myDecks" />
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import DeckSearch from '../components/decks/deck-search.vue';

const { mapGetters } = createNamespacedHelpers('decks');

export default {
  name: 'my-decks',
  components: {
    DeckSearch,
  },
  computed: {
    ...mapGetters([
      'myDecks',
    ]),
  },
};
</script>
