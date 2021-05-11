<template>
  <v-form>
    <v-container>
      <v-row>
        <v-col>
          <v-text-field label="Name" v-model="cardSearchParameters.name" />
        </v-col>
        <v-col>
          <v-text-field label="Abilities" v-model="cardSearchParameters.ability" />
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="1">
          <v-label>Cost</v-label>
        </v-col>
        <v-col cols="5">
          <v-btn-toggle
            v-model="cardSearchParameters.manaCost"
            multiple
            borderless
            tile
          >
            <v-btn
              v-for="i in cardSearchOptions.manaCost"
              :key="i"
              :value="i"
              icon
              small
            >
              <v-icon>$vuetify.icons.{{ i }}</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>
        <v-col cols="1">
          <v-label>Type</v-label>
        </v-col>
        <v-col cols="5">
          <v-btn-toggle
            v-model="cardSearchParameters.type"
            multiple
            tile
            borderless
            dense
          >
            <v-btn
              v-for="i in cardSearchOptions.type"
              :key="i"
              :value="i"
              text
            >{{i}}</v-btn>
          </v-btn-toggle>
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="1">
          <v-label>Power</v-label>
        </v-col>
        <v-col cols="5">
          <v-btn-toggle
            v-model="cardSearchParameters.power"
            multiple
            borderless
            tile
          >
            <v-btn
              v-for="i in cardSearchOptions.power"
              :key="i"
              :value="i"
              icon
              small
            >
              <v-icon>$vuetify.icons.{{ i }}</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>
        <v-col cols="4">
          <v-autocomplete
            label="Subtype"
            :items="cardSearchOptions.subtype"
            clearable
            multiple
            v-model="cardSearchParameters.subtype"
          />
        </v-col>
        <v-col cols="2">
          <v-switch
            v-model="cardSearchParameters.exactSubtype"
            label="Use Exact"
            dense
          ></v-switch>
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="1">
          <v-label>Health</v-label>
        </v-col>
        <v-col cols="5">
          <v-btn-toggle
            v-model="cardSearchParameters.health"
            multiple
            borderless
            tile
          >
            <v-btn
              v-for="i in cardSearchOptions.health"
              :key="i"
              :value="i"
              icon
              small
            >
              <v-icon>$vuetify.icons.{{ i }}</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>
        <v-col cols="6">
          <v-autocomplete
            label="Set / Planeswalker"
            :items="cardSearchOptions.set"
            item-text="name"
            item-value="key"
            clearable
            multiple
            v-model="cardSearchParameters.set"
          />
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="1">
          <v-label>Colour</v-label>
        </v-col>
        <v-col cols="3">
          <v-btn-toggle
            v-model="cardSearchParameters.colour"
            multiple
            borderless
            tile
          >
            <v-btn
              v-for="i in cardSearchOptions.colour"
              :key="i"
              :value="i"
              icon
              small
            >
              <v-icon>$vuetify.icons.{{ i }}</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>
        <v-col cols="2">
          <v-switch
            v-model="cardSearchParameters.exactColour"
            label="Use Exact"
            dense
          ></v-switch>
        </v-col>
        <v-col cols="6">
          <v-autocomplete
            label="Artist"
            :items="cardSearchOptions.artist"
            clearable
            multiple
            v-model="cardSearchParameters.artist"
          />
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="1">
          <v-label>Rarity</v-label>
        </v-col>
        <v-col cols="5">
          <v-btn-toggle
            v-model="cardSearchParameters.rarity"
            multiple
            tile
            borderless
            dense
          >
            <v-btn
              v-for="i in cardSearchOptions.rarity"
              :key="i"
              :value="i"
              text
            >{{i}}</v-btn>
          </v-btn-toggle>
        </v-col>
        <v-col cols="6">
          <v-radio-group
            row
            label="Legendary"
            v-model="cardSearchParameters.legendary"
            dense
          >
            <v-radio
              v-for="i in cardSearchOptions.legendary"
              :key="i"
              :label="i.toString() | capitalize"
              :value="i"
            />
          </v-radio-group>
        </v-col>
      </v-row>
      <v-row class="d-flex align-center">
        <v-col cols="6">
          <v-radio-group
            row
            label="Materials Available"
            v-model="cardSearchParameters.materialsAvailable"
            dense
          >
            <v-radio
              v-for="i in cardSearchOptions.materialsAvailable"
              :key="i"
              :label="i.toString() | capitalize"
              :value="i"
            />
          </v-radio-group>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn @click="$emit('search', { ...cardSearchParameters })">Search</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import capitalize from 'lodash-es/capitalize';
import CardSearchParametersMixin from './card-search-mixin';

// TODO: Build a Reset feature

export default {
  name: 'card-filter',
  filters: {
    capitalize,
  },
  mixins: [CardSearchParametersMixin],
};
</script>
