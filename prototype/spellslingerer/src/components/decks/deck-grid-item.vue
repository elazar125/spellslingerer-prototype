<template>
  <v-card max-width="300px" :to="`/decks/${id}`">
    <v-img
      :src="`Cards/Tile/${featureCard.tileImage}`"
      gradient="to bottom, transparent, rgba(25,25,25,.7)"
      class="align-end"
    >
      <v-card-title style="position: absolute; top: 0; width: 100%; align-items: flex-start;" class="mt-0">
        <v-col cols="auto" class="pt-0">
          <v-row class="mb-0">{{ name }}</v-row>
          <v-row class="overline mt-0">
            <router-link :to="`/user/${user}`">{{ user }}</router-link>
          </v-row>
        </v-col>
        <v-spacer />
        <v-icon v-if="isMine" :color="isPublic ? 'primary' : ''" class="mt-1">
          mdi-earth{{ isPublic ? '' : '-off' }}
        </v-icon>
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-icon
          v-for="c in colour"
          :key="c"
          large
          class="mr-2"
        >
          $vuetify.icons.{{ c }}
        </v-icon>
        <v-img
          :src="`Planeswalkers/${planeswalker.image}`"
          max-width="50px"
          class="mr-2"
        />
        <v-img
          :src="`Cards/Icons/${land.image}`"
          max-width="40px"
        />
      </v-card-actions>
    </v-img>
  </v-card>
</template>

<script>
export default {
  name: 'deck-grid-item',
  props: {
    id: {
      type: String,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    planeswalker: {
      type: Object,
      required: true,
    },
    colour: {
      type: Array,
      required: true,
    },
    land: {
      type: Object,
      required: true,
    },
    featureCard: {
      type: Object,
      required: true,
    },
    isPublic: {
      type: Boolean,
      required: true,
    },
    user: {
      type: String,
      required: true,
    },
    isMine: {
      type: Boolean,
      required: true,
    },
  },
};
</script>
