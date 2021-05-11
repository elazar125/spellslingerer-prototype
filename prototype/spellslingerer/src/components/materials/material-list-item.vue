<template>
  <v-card flat align="center">
    <v-img
      src="Materials/HellhoundFang.png"
      width="50px"
      contain
      class="pt-4"
    />
    <span class="subtitle">
      {{ name }}
    </span>
    <v-card-actions v-if="localQuantity !== null">
      <v-spacer />
      <v-text-field
        dense
        hide-details
        solo
        flat
        v-model="localQuantity"
        :readonly="readonly"
        style="max-width: 160px;"
      >
        <template v-slot:prepend v-if="!readonly">
          <v-btn x-small @click="if (localQuantity > 0) localQuantity--">
            <v-icon>mdi-minus</v-icon>
          </v-btn>
        </template>
        <template v-slot:append v-if="!readonly">
          <v-btn x-small @click="localQuantity++">
            <v-icon>mdi-plus</v-icon>
          </v-btn>
        </template>
      </v-text-field>
      <v-spacer />
    </v-card-actions>
  </v-card>
</template>

<script>
// TODO: Use correct image

export default {
  name: 'material-list-item',
  data() {
    return {
      localQuantity: this.quantity,
    };
  },
  props: {
    readonly: {
      type: Boolean,
      default: false,
    },
    id: {
      type: String,
      required: true,
    },
    image: {
      type: String,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    quantity: {
      type: Number,
      default: null,
    },
  },
  watch: {
    localQuantity(val) {
      this.$emit('updateQuantity', { materialId: this.id, quantity: val });
    },
  },
};
</script>
