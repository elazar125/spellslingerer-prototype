<template>
  <v-card flat @click="$emit('planeswalkerSelected', id)">
    <v-row class="d-flex align-center">
      <v-col>
        <v-list-item three-line>
          <v-list-item-avatar tile size="100">
            <v-img :src="`Planeswalkers/${image}`" contain></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title class="headline">
              {{ name }}
              <v-icon class="ml-2" v-for="c in colour" :key="c">$vuetify.icons.{{ c }}</v-icon>
            </v-list-item-title>
            <v-list-item-subtitle v-for="abilityName in Object.keys(ability)" :key="abilityName">
              <span class="font-weight-bold">{{ abilityName }}:</span> {{ ability[abilityName] }}
            </v-list-item-subtitle>
            <div>
              Health: {{ health }}
            </div>
          </v-list-item-content>
        </v-list-item>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  name: 'planeswalker-list-item',
  props: {
    id: {
      type: String,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    image: {
      type: String,
      required: true,
    },
    colour: {
      type: Array,
      required: true,
    },
    ability: {
      type: Object,
      required: true,
    },
    health: {
      type: Number,
      required: true,
    },
    splashColours: {
      type: Array,
      required: true,
    },
    cardRules: {
      type: Array,
      required: true,
    },
    landRules: {
      type: Array,
      required: true,
    },
  },
};
</script>
