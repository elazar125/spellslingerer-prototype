<template>
  <v-card flat @click="$emit('cardSelected', id)">
    <v-row class="d-flex align-center">
      <v-col cols="9">
        <v-list-item three-line>
          <v-list-item-avatar tile size="100">
            <v-img v-if="displayTileImage" :src="`Cards/Tile/${tileImage}`" contain></v-img>
            <v-img v-else src="/Cards/Icons/Default.png" :srcset="`Cards/FullArt/${image}`" contain></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <div class="overline">
              <v-icon small v-if="cost !== null">$vuetify.icons.{{ cost }}</v-icon>
              {{ type }}{{ subtype.length ? ` (${subtype.join(' ')})` : '' }}
              <v-icon small v-for="c in colour" :key="c">$vuetify.icons.{{ c }}</v-icon>
            </div>
            <v-list-item-title class="headline">
              {{ name }}
            </v-list-item-title>
            <v-list-item-subtitle>
              <v-icon small v-if="power !== null">$vuetify.icons.{{ power }}</v-icon>
              <span v-if="power !== null && health !== null">/</span>
              <v-icon small v-if="health !== null">$vuetify.icons.{{ health }}</v-icon>
              {{ ability }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-col>
      <v-col cols="3">
        <v-text-field
          dense
          solo
          flat
          v-model="localQuantity"
          v-if="!noQuantity"
        >
          <template v-slot:prepend>
            <v-btn x-small @click="if (localQuantity > 0) localQuantity--">
              <v-icon>mdi-minus</v-icon>
            </v-btn>
          </template>
          <template v-slot:append>
            <v-btn x-small @click="if (localQuantity < maxQuantity) localQuantity++">
              <v-icon>mdi-plus</v-icon>
            </v-btn>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  name: 'card-list-item',
  data() {
    return {
      localQuantity: this.quantity,
    };
  },
  computed: {
    maxQuantity() {
      return this.legendary || this.type === 'Land' ? 1 : 2;
    },
  },
  props: {
    noQuantity: {
      type: Boolean,
      default: false,
    },
    displayTileImage: {
      type: Boolean,
      default: false,
    },
    id: {
      type: String,
      required: true,
    },
    set: {
      type: String,
      required: true,
    },
    image: {
      type: String,
      required: true,
    },
    tileImage: {
      type: String,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    cost: {
      type: Number,
      default: null,
    },
    power: {
      type: Number,
      default: null,
    },
    health: {
      type: Number,
      default: null,
    },
    ability: {
      type: String,
      required: true,
    },
    rarity: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      required: true,
    },
    subtype: {
      type: Array,
      required: true,
    },
    artist: {
      type: String,
      required: true,
    },
    colour: {
      type: Array,
      required: true,
    },
    legendary: {
      type: Boolean,
      required: true,
    },
    quantity: {
      type: Number,
      default: 0,
    },
  },
  watch: {
    localQuantity(val) {
      this.$emit('updateQuantity', { id: this.id, quantity: val });
    },
  },
};
</script>
