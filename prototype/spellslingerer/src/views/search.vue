<template>
  <card-search no-quantity :cards="cards" />
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import CardSearch from '../components/cards/card-search.vue';

const { mapGetters } = createNamespacedHelpers('cards');

export default {
  name: 'search',
  components: {
    CardSearch,
  },
  computed: {
    ...mapGetters([
      'cards',
    ]),
  },
};
</script>
