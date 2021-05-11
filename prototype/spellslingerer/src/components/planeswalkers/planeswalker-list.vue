<template>
  <div>
    <planeswalker-list-item
      v-for="planeswalker in planeswalkers"
      v-bind="planeswalker"
      :key="planeswalker.name"
      @planeswalkerSelected="$emit('planeswalkerSelected', $event)"
    />
  </div>
</template>

<script>
import PlaneswalkerListItem from './planeswalker-list-item.vue';

export default {
  name: 'planeswalker-list',
  components: {
    PlaneswalkerListItem,
  },
  props: {
    planeswalkers: {
      type: Array,
      required: true,
    },
  },
};
</script>
