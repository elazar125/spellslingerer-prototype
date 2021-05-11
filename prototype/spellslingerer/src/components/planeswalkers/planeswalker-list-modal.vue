<template>
  <v-dialog v-model="open" max-width="1000px">
    <v-card>
      <planeswalker-list
        :planeswalkers="planeswalkers"
        @planeswalkerSelected="$emit('planeswalkerSelected', $event)"
      />
    </v-card>
  </v-dialog>
</template>

<script>
import PlaneswalkerList from './planeswalker-list.vue';

export default {
  name: 'planeswalker-list-modal',
  components: {
    PlaneswalkerList,
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
    planeswalkers: {
      type: Array,
      required: true,
    },
  },
};
</script>
