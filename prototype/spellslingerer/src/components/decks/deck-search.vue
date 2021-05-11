<template>
  <div>
    <v-btn fab bottom right fixed color="success" to="/decks/new">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-card>
      <v-container>
        <v-row dense>
          <v-col>
            <v-text-field
              label="Deck name"
              v-model="deckName"
              hide-details
            />
          </v-col>
          <v-col>
            <v-autocomplete
              label="Planeswalker"
              :items="planeswalkerNames"
              v-model="selectedPlaneswalkerNames"
              hide-details
              multiple
            />
          </v-col>
          <v-col>
            <v-autocomplete
              label="Land"
              :items="landNames"
              v-model="selectedLandNames"
              hide-details
              multiple
            />
          </v-col>
          <v-col>
            <v-autocomplete
              label="Card name"
              :items="cardNames"
              v-model="selectedCardNames"
              hide-details
              multiple
            />
          </v-col>
          <v-col cols="auto">
            <v-btn fab color="primary">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
    <v-row>
      <deck-grid :decks="results" />
    </v-row>
  </div>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import {
  isEmptyOrIncludes,
  isEmptyOrBaz,
  isEmptyOrIncludesAll,
} from '../../libraries/filters';
import DeckGrid from './deck-grid.vue';

const { mapGetters: mapCardGetters } = createNamespacedHelpers('cards');
const { mapGetters: mapPlaneswalkerGetters } = createNamespacedHelpers('planeswalkers');

const defaultFilters = [
  ({ name }, { deckName }) => isEmptyOrBaz(deckName, name),
  ({ planeswalker }, { selectedPlaneswalkerNames }) => isEmptyOrIncludes(selectedPlaneswalkerNames, planeswalker.name),
  ({ land }, { selectedLandNames }) => isEmptyOrIncludes(selectedLandNames, land.name),
  ({ cards }, { selectedCardNames }) => isEmptyOrIncludesAll(selectedCardNames, cards.map(({ name }) => name)),
];

export default {
  name: 'deck-search',
  components: {
    DeckGrid,
  },
  data() {
    return {
      deckName: '',
      selectedPlaneswalkerNames: [],
      selectedLandNames: [],
      selectedCardNames: [],
    };
  },
  props: {
    decks: {
      type: Array,
      required: true,
    },
    filters: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  computed: {
    ...mapCardGetters([
      'cardNames',
      'landNames',
    ]),
    ...mapPlaneswalkerGetters([
      'planeswalkerNames',
    ]),
    results() {
      const filters = [...defaultFilters, ...this.filters];
      return this.decks.filter((deck) => filters.every((func) => func(deck, {
        deckName: this.deckName,
        selectedPlaneswalkerNames: this.selectedPlaneswalkerNames,
        selectedLandNames: this.selectedLandNames,
        selectedCardNames: this.selectedCardNames,
      })));
    },
  },
};
</script>
