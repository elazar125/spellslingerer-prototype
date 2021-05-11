<template>
  <v-dialog v-model="open" max-width="1000px">
    <v-card>
      <card-details v-bind="card" />
    </v-card>
  </v-dialog>
</template>

<script>
import CardDetails from './card-details.vue';

export default {
  name: 'card-detail-modal',
  components: {
    CardDetails,
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
    card: {
      type: Object,
      required: true,
    },
  },
};
</script>
