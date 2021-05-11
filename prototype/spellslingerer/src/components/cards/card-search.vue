<template>
  <v-expansion-panels multiple focusable accordion v-model="panels">
    <v-expansion-panel>
      <v-expansion-panel-header>Filter</v-expansion-panel-header>
      <v-expansion-panel-content>
        <card-filter @search="cardSearchParameters = $event" />
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>Sort</v-expansion-panel-header>
      <v-expansion-panel-content>
        <card-sorter />
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>Results</v-expansion-panel-header>
      <v-expansion-panel-content>
        <card-list
          :cards="results"
          @updateQuantity="$emit('updateQuantity', $event)"
          @cardSelected="$emit('cardSelected', $event)"
          :no-quantity="noQuantity"
        />
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import {
  isEmptyOrIncludes,
  isEmptyOrMatches,
  isEmptyOrBaz,
  isAnyOrValue,
} from '../../libraries/filters';
import CardFilter from './card-filter.vue';
import CardSorter from './card-sorter.vue';
import CardList from './card-list.vue';
import CardSearchParametersMixin from './card-search-mixin';

const { mapGetters } = createNamespacedHelpers('materials');

function isAnyOrFoo(include, materialsRequired, mox, materialsAvailable) {
  const moxesAvailable = materialsAvailable.find((material) => material.id === mox).quantity;
  const moxesRequired = 10 * Object.keys(materialsRequired).reduce((materialsMissing, materialId) => {
    const quantityAvailable = materialsAvailable.find((material) => material.id === materialId).quantity;
    const quantityrequired = materialsRequired[materialId];
    return quantityAvailable >= quantityrequired ? materialsMissing : materialsMissing + quantityrequired - quantityAvailable;
  }, 0);
  const hasEnoughMaterials = moxesAvailable >= moxesRequired;
  return ['Any', hasEnoughMaterials].includes(include);
}

const defaultFilters = [
  ({ card, parameters }) => isEmptyOrBaz(parameters.name, card.name),
  ({ card, parameters }) => isEmptyOrBaz(parameters.ability, card.ability),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.manaCost, card.cost),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.power, card.power),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.health, card.health),
  ({ card, parameters }) => isEmptyOrMatches(parameters.colour, card.colour, parameters.exactColour),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.rarity, card.rarity),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.type, card.type),
  ({ card, parameters }) => isEmptyOrMatches(parameters.subtype, card.subtype, parameters.exactSubtype),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.set, card.set),
  ({ card, parameters }) => isEmptyOrIncludes(parameters.artist, card.artist),
  ({ card, parameters }) => isAnyOrValue(parameters.legendary, card.legendary),
  ({ card, parameters }) => isAnyOrValue(parameters.collectible, card.collectible),
  ({ card, parameters, materials }) => isAnyOrFoo(parameters.materialsAvailable, card.materials, card.mox, materials),
];

export default {
  name: 'card-search',
  components: {
    CardFilter,
    CardSorter,
    CardList,
  },
  data() {
    return {
      panels: [2],
    };
  },
  mixins: [CardSearchParametersMixin],
  props: {
    noQuantity: {
      type: Boolean,
      default: false,
    },
    cards: {
      type: Array,
      required: true,
    },
    filters: {
      type: Array,
      default() {
        return [];
      },
    },
    parameters: {
      type: Object,
      default() {
        return {};
      },
    },
  },
  computed: {
    ...mapGetters([
      'materials',
    ]),
    results() {
      const filters = [...defaultFilters, ...this.filters];
      // TODO: Pass in deck splash colour
      return this.cards.filter((card) => filters.every((func) => func({
        card,
        parameters: this.cardSearchParameters,
        materials: this.materials,
        ...this.parameters,
      })));
    },
  },
};
</script>
