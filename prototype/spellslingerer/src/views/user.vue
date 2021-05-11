<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <h1>{{ user }}</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <deck-search :decks="decks" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import DeckSearch from '../components/decks/deck-search.vue';

const { mapGetters } = createNamespacedHelpers('decks');

export default {
  name: 'user',
  components: {
    DeckSearch,
  },
  data() {
    return {
      user: this.$route.params.id,
    };
  },
  computed: {
    ...mapGetters([
      'decksByUser',
    ]),
    decks() {
      return this.decksByUser(this.user);
    },
  },
};
</script>
