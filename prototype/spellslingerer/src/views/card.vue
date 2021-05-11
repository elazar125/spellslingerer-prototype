<template>
  <card-details v-bind="getCardFromRoute()" />
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import CardDetails from '../components/cards/card-details.vue';

const { mapGetters } = createNamespacedHelpers('cards');

export default {
  name: 'card',
  components: {
    CardDetails,
  },
  computed: {
    ...mapGetters([
      'card',
    ]),
  },
  methods: {
    getCardFromRoute() {
      return this.card(this.$route.params.id);
    },
  },
};
</script>
