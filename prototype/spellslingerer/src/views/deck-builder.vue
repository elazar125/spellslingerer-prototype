<template>
  <v-row>
    <div v-if="currentDeck.isMine">
      <v-btn
        fab bottom right fixed
        color="success"
        @click="editMode = true"
        v-if="!editMode"
      >
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
      <v-speed-dial bottom right fixed v-else>
        <template #activator>
          <v-btn fab color="info">
            <v-icon>mdi-content-save-settings</v-icon>
          </v-btn>
        </template>
        <v-btn fab color="success" @click="saveDeck(); editMode = false">
          <v-icon>mdi-content-save</v-icon>
        </v-btn>
        <v-btn fab color="warning" @click="isDirty ? confirmResetModalOpen = true : editMode = false">
          <v-icon>mdi-undo</v-icon>
        </v-btn>
        <v-btn fab color="error" @click="confirmDeleteModalOpen = true">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-speed-dial>
      <confirm-dialog
        v-model="confirmDeleteModalOpen"
        @confirmed="deleteCurrentDeck()"
        title="Delete?"
        body="Really delete?"
        destructive-action
      />
      <confirm-dialog
        v-model="confirmResetModalOpen"
        @confirmed="resetDeck(); editMode = false"
        title="Reset?"
        body="Really reset?"
        destructive-action
      />
      <confirm-dialog
        v-model="confirmLeaveModalOpen"
        ref="leaveModal"
        title="Leave?"
        body="Really leave?"
        destructive-action
      />
    </div>
    <v-col cols="4">
      <card-grid :cards="currentDeck.cards">
        <template #label v-if="editMode">
          <v-row dense>
            <v-col cols="8">
              <v-text-field
                hide-details
                :value="currentDeck.name"
                @change="setName($event)"
              />
            </v-col>
            <v-col cols="4">
              <v-switch
                class="ml-4"
                dense
                label="Public"
                :input-value="currentDeck.isPublic"
                @change="setPublic($event)"
                :true-value="true"
                :false-value="false"
              />
            </v-col>
          </v-row>
        </template>
        <template #label v-else>
          <v-row dense>
            <v-col cols="11">
              <v-row class="mb-0">
                <v-col class="my-0 py-0">{{ currentDeck.name }}</v-col>
              </v-row>
              <v-row class="overline mt-0">
                <v-col class="my-0 py-0">
                  <router-link :to="`/user/${currentDeck.user}`">{{ currentDeck.user }}</router-link>
                </v-col>
              </v-row>
            </v-col>
            <v-col cols="1" v-if="currentDeck.isMine">
              <v-icon :color="currentDeck.isPublic ? 'primary' : ''">
                mdi-earth{{ currentDeck.isPublic ? '' : '-off' }}
              </v-icon>
            </v-col>
          </v-row>
        </template>
        <template #header>
          <v-row class="mt-4">
            <v-spacer />
            <v-col cols="auto" align="center">
              <div class="overline mb-2">Planeswalker</div>
              <v-img
                :src="`/Planeswalkers/${currentDeck.planeswalker.image}`"
                contain
                max-width="50px"
                @click="planeswalkerModalOpen = editMode"
              />
              <planeswalker-list-modal
                v-model="planeswalkerModalOpen"
                :planeswalkers="planeswalkers"
                @planeswalkerSelected="selectPlaneswalker($event)"
              />
            </v-col>
            <v-col cols="auto" v-if="currentDeck.planeswalker.splashColours.length">
              <v-menu offset-y>
                <template #activator="{ on }">
                  <div v-on="on" class="text-center">
                    <div class="overline mb-2">Splash</div>
                    <v-icon x-large>
                      $vuetify.icons.{{ currentDeck.splashColour || 'C' }}
                    </v-icon>
                  </div>
                </template>
                <v-list>
                  <v-list-item
                    v-for="colour in currentDeck.planeswalker.splashColours"
                    :key="colour"
                    @click="setSplash(colour)"
                  >
                    <v-icon
                      large
                      class="mr-2"
                    >
                      $vuetify.icons.{{ colour }}
                    </v-icon>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-col>
            <v-col cols="auto" align="center">
              <div class="overline mb-2">Land</div>
              <v-img
                :src="`/Cards/Icons/${currentDeck.land.image}`"
                contain
                max-height="50px"
                max-width="50px"
                @click="landModalOpen = editMode"
              />
              <card-list-modal
                v-model="landModalOpen"
                :cards="lands"
                @cardSelected="selectLand($event)"
                no-quantity
              />
            </v-col>
            <v-col cols="auto" align="center">
              <div class="overline mb-2">Tile Image</div>
              <v-img
                :src="`/Cards/Tile/${currentDeck.featureCard.tileImage}`"
                contain
                max-height="50px"
                max-width="80px"
                @click="featureModalOpen = editMode"
              />
              <card-list-modal
                v-model="featureModalOpen"
                :cards="currentDeck.cards"
                @cardSelected="selectFeatureCard($event)"
                no-quantity
                display-tile-image
              />
            </v-col>
            <v-spacer />
          </v-row>
          <v-divider />
        </template>
      </card-grid>
    </v-col>
    <v-col cols="8">
      <v-row class="mb-2">
        <v-card width="100%">
          <v-col>
            <v-textarea
              v-if="editMode"
              label="Description"
              :value="currentDeck.description"
              @change="setDescription($event)"
            />
            <div v-else>{{ currentDeck.description }}</div>
          </v-col>
        </v-card>
      </v-row>
      <v-row class="mb-2" v-if="editMode">
        <card-search
          :cards="cardsWithQuantity"
          :filters="currentDeck.planeswalker.cardRules"
          :parameters="{ splash: currentDeck.splashColour }"
          @updateQuantity="updateQuantity($event)"
        />
      </v-row>
      <v-row class="mb-2">
        <v-sheet width="100%" min-height="215px">
          <material-list oneline readonly label="Required Materials" :materials="requiredMaterials" />
        </v-sheet>
      </v-row>
      <v-row>
        <deck-charts :cards="currentDeck.cards" />
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import CardSearch from '../components/cards/card-search.vue';
import CardGrid from '../components/cards/card-grid.vue';
import PlaneswalkerListModal from '../components/planeswalkers/planeswalker-list-modal.vue';
import CardListModal from '../components/cards/card-list-modal.vue';
import DeckCharts from '../components/decks/deck-charts.vue';
import MaterialList from '../components/materials/material-list.vue';
import ConfirmDialog from '../components/confirm-dialog.vue';

const { mapGetters: mapDeckGetters, mapActions } = createNamespacedHelpers('decks');
const { mapGetters: mapCardGetters } = createNamespacedHelpers('cards');
const { mapGetters: mapPlaneswalkerGetters } = createNamespacedHelpers('planeswalkers');
const { mapGetters: mapMaterialGetters } = createNamespacedHelpers('materials');

export default {
  name: 'deck-builder',
  components: {
    CardSearch,
    CardGrid,
    PlaneswalkerListModal,
    CardListModal,
    DeckCharts,
    MaterialList,
    ConfirmDialog,
  },
  data() {
    return {
      editMode: this.$route.params.deckId === 'new',
      planeswalkerModalOpen: false,
      landModalOpen: false,
      featureModalOpen: false,
      confirmDeleteModalOpen: false,
      confirmResetModalOpen: false,
      confirmLeaveModalOpen: false,
    };
  },
  computed: {
    ...mapMaterialGetters([
      'quantityStillRequired',
    ]),
    ...mapDeckGetters([
      'currentDeck',
      'quantities',
      'cardMaterials',
      'isDirty',
    ]),
    ...mapCardGetters([
      'cards',
    ]),
    ...mapPlaneswalkerGetters([
      'planeswalkers',
    ]),
    lands() {
      return this.cards.filter((card) => this.currentDeck.planeswalker.landRules.every((func) => func({ card })));
    },
    cardsWithQuantity() {
      return this.cards.map((card) => ({ ...card, quantity: this.quantities[card.id] }));
    },
    requiredMaterials() {
      return this.cardMaterials.map(this.quantityStillRequired).filter(({ quantity }) => quantity > 0);
    },
  },
  methods: {
    ...mapActions([
      'updateQuantity',
      'setPlaneswalker',
      'setLand',
      'setFeatureCard',
      'setSplash',
      'setName',
      'setDescription',
      'setPublic',
      'loadDeck',
      'resetDeck',
      'saveDeck',
      'deleteDeck',
    ]),
    selectPlaneswalker(planeswalkerId) {
      this.planeswalkerModalOpen = false;
      this.setPlaneswalker(planeswalkerId);
      if (this.currentDeck.planeswalker.splashColours.length === 0) {
        this.setSplash(null);
      }
    },
    selectLand(landId) {
      this.landModalOpen = false;
      this.setLand(landId);
    },
    selectFeatureCard(featureCardId) {
      this.featureModalOpen = false;
      this.setFeatureCard(featureCardId);
    },
    deleteCurrentDeck() {
      this.deleteDeck();
      this.$router.push('/deck-builder');
    },
  },
  created() {
    this.loadDeck(this.$route.params.deckId);
  },
  beforeRouteLeave(to, from, next) {
    if (this.editMode && this.isDirty) {
      this.$refs.leaveModal.awaitResponse().then((confirmResult) => next(confirmResult));
    } else {
      next();
    }
  },
};
</script>
