<template>
  <v-card class="fill-height">
    <v-container>
      <slot name="label">
        <span class="pa-4 headline" v-if="label">{{ label }}</span>
      </slot>
      <slot name="header" />
      <v-row>
        <v-col
          cols="auto"
          v-for="card in cards"
          :key="card.id"
        >
          <card-grid-item :card="card" />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import CardGridItem from './card-grid-item.vue';

export default {
  name: 'card-grid',
  components: {
    CardGridItem,
  },
  props: {
    cards: {
      type: Array,
      required: true,
    },
    label: {
      type: String,
      default: '',
    },
  },
};
</script>
