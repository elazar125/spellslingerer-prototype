<template>
  <v-container>
    <v-row>
      <v-col cols="3">
        <v-img :src="`Cards/FullArt/${image}`" />
      </v-col>
      <v-col cols="9">
        <v-row>
          <router-link :to="`/cards/${id}`" class="white--text display-2">{{ name }}</router-link>
        </v-row>
        <v-row v-if="cost !== null">
          <span>Mana Cost: <v-icon>$vuetify.icons.{{ cost }}</v-icon></span>
        </v-row>
        <v-row v-if="power !== null">
          <span>Power: <v-icon>$vuetify.icons.{{ power }}</v-icon></span>
        </v-row>
        <v-row v-if="health !== null">
          <span>Health: <v-icon>$vuetify.icons.{{ health }}</v-icon></span>
        </v-row>
        <v-row>
          <span>Ability: {{ ability }}</span>
        </v-row>
        <v-row>
          <span>Rarity: {{ rarity }}</span>
        </v-row>
        <v-row>
          <span>Type: {{ type }}</span>
        </v-row>
        <v-row>
          <span>Subtype: {{ subtype.join(' ') }}</span>
        </v-row>
        <v-row>
          <span>Artist: {{ artist }}</span>
        </v-row>
        <v-row>
          <span>Colour: <v-icon v-for="c in colour" :key="c">$vuetify.icons.{{ c }}</v-icon></span>
        </v-row>
        <v-row>
          <span>Set: {{ set }}</span>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'card-details',
  props: {
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
    name: {
      type: String,
      required: true,
    },
    cost: {
      type: Number,
      required: true,
    },
    power: {
      type: Number,
      required: true,
    },
    health: {
      type: Number,
      required: true,
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
  },
};
</script>
