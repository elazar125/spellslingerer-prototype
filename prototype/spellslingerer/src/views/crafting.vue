<template>
  <v-row>
    <material-list @updateQuantity="updateQuantity($event)" :materials="materials" />
  </v-row>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import MaterialList from '../components/materials/material-list.vue';

const { mapGetters, mapActions } = createNamespacedHelpers('materials');

export default {
  name: 'crafting',
  components: {
    MaterialList,
  },
  computed: {
    ...mapGetters([
      'materials',
    ]),
  },
  methods: {
    ...mapActions([
      'updateQuantity',
    ]),
  },
};
</script>
