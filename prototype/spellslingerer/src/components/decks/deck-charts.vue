<template>
  <v-card height="215" width="100%">
    <v-row>
      <v-col class="ml-6">
        <bar-chart
          :styles="styles"
          :options="costChartOptions"
          :chart-data="costChartData"
        />
      </v-col>
      <v-col>
        <doughnut-chart
          :styles="styles"
          :options="typeChartOptions"
          :chart-data="typeChartData"
        />
      </v-col>
      <v-col>
        <bar-chart
          :styles="styles"
          :options="rarityChartOptions"
          :chart-data="rarityChartData"
        />
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
import BarChart from '../charts/bar-chart.vue';
import DoughnutChart from '../charts/doughnut-chart.vue';
import CardSearchMixin from '../cards/card-search-mixin';

// TODO: Make colours less magic

export default {
  name: 'deck-charts',
  components: {
    BarChart,
    DoughnutChart,
  },
  mixins: [CardSearchMixin],
  props: {
    cards: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      styles: {
        width: '175px',
      },
      costChartOptions: {
        legend: { display: false },
        title: {
          display: true,
          text: 'Mana Cost',
        },
        scales: {
          yAxes: [{
            ticks: {
              stepSize: 1,
            },
          }],
        },
      },
      rarityChartOptions: {
        legend: { display: false },
        title: {
          display: true,
          text: 'Rarity',
        },
        scales: {
          yAxes: [{
            ticks: {
              stepSize: 1,
            },
          }],
        },
      },
      typeChartOptions: {
        title: {
          display: true,
          text: 'Type',
        },
        legend: { display: false },
      },
    };
  },
  computed: {
    costChartData() {
      const costCounts = Object.fromEntries(this.cardSearchOptions.manaCost.map((m) => [m, 0]));
      this.cards.forEach(({ cost, quantity }) => {
        costCounts[cost] += quantity;
      });

      return {
        labels: Object.keys(costCounts),
        datasets: [{
          backgroundColor: Object.keys(costCounts).map(() => '#BBBBBB'),
          data: Object.values(costCounts),
        }],
      };
    },
    typeChartData() {
      const typeCounts = Object.fromEntries(this.cardSearchOptions.type.map((c) => [c, 0]));
      this.cards.forEach(({ type, quantity }) => {
        typeCounts[type] += quantity;
      });

      return {
        labels: Object.keys(typeCounts),
        datasets: [{
          backgroundColor: [
            '#7E4C24',
            '#2A776C',
            '#7C4C44',
            '#D77A49',
            '#1e1e1e',
          ],
          borderColor: [
            '#7E4C24',
            '#2A776C',
            '#7C4C44',
            '#D77A49',
            '#1e1e1e',
          ],
          data: Object.values(typeCounts),
        }],
      };
    },
    rarityChartData() {
      const rarityCounts = Object.fromEntries(this.cardSearchOptions.rarity.map((r) => [r, 0]));
      this.cards.forEach(({ rarity, quantity }) => {
        rarityCounts[rarity] += quantity;
      });

      return {
        labels: Object.keys(rarityCounts),
        datasets: [{
          backgroundColor: [
            '#BBBBBB',
            '#37837D',
            '#965159',
            '#D77A49',
          ],
          data: Object.values(rarityCounts),
        }],
      };
    },
  },
};
</script>
