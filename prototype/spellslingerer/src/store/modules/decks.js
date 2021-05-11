import isEqual from 'lodash-es/isEqual';
import allCards from '../data/cards.json';
import allPlaneswalkers from '../data/planeswalkers';

const emptyDeck = {
  id: '',
  name: '',
  description: '',
  planeswalkerId: 'Default',
  splashColour: null,
  landId: 'Default',
  featureCardId: 'Default',
  cards: [],
  isPublic: false,
};

export default {
  namespaced: true,
  state: () => ({
    decks: [
      {
        id: '1',
        name: 'Someone\'s Cool Deck',
        description: 'This is a pretty nifty deck!',
        planeswalkerId: 'Default',
        splashColour: null,
        landId: 'Default',
        featureCardId: 'Default',
        cards: [
          { id: 'AA', quantity: 1 },
          { id: 'AB', quantity: 2 },
        ],
        isPublic: true,
        user: 'not me',
      },
      {
        id: '2',
        name: 'My Cool Deck',
        description: 'This is a pretty nifty deck!',
        planeswalkerId: 'Default',
        splashColour: null,
        landId: 'Default',
        featureCardId: 'Default',
        cards: [
          { id: 'AA', quantity: 1 },
          { id: 'AB', quantity: 2 },
        ],
        isPublic: false,
        user: 'me',
      },
    ],
    currentDeck: { ...emptyDeck },
    currentCards: [],
  }),
  getters: {
    isDirty: ({ currentDeck, currentCards, decks }) => {
      const savedState = decks.find(({ id }) => id === currentDeck.id);
      const currentState = { ...currentDeck, cards: [...currentCards] };
      return !isEqual(savedState, currentState);
    },
    decks: ({ decks }, getters, { currentUser }) => decks.map((deck) => ({
      ...deck,
      colour: [...allPlaneswalkers[deck.planeswalkerId].colour, deck.splashColour].filter((colour) => colour),
      planeswalker: allPlaneswalkers[deck.planeswalkerId],
      land: allCards[deck.landId],
      featureCard: allCards[deck.featureCardId],
      cards: deck.cards.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
      isMine: deck.user === currentUser,
    })).filter(({ isPublic, isMine }) => isPublic || isMine),
    myDecks: ({ decks }, getters, { currentUser }) => decks.map((deck) => ({
      ...deck,
      colour: [...allPlaneswalkers[deck.planeswalkerId].colour, deck.splashColour].filter((colour) => colour),
      planeswalker: allPlaneswalkers[deck.planeswalkerId],
      land: allCards[deck.landId],
      featureCard: allCards[deck.featureCardId],
      cards: deck.cards.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
      isMine: deck.user === currentUser,
    })).filter(({ isMine }) => isMine),
    decksByUser: ({ decks }, getters, { currentUser }) => (selectedUser) => decks.map((deck) => ({
      ...deck,
      colour: [...allPlaneswalkers[deck.planeswalkerId].colour, deck.splashColour].filter((colour) => colour),
      planeswalker: allPlaneswalkers[deck.planeswalkerId],
      land: allCards[deck.landId],
      featureCard: allCards[deck.featureCardId],
      cards: deck.cards.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
      isMine: deck.user === currentUser,
    })).filter(({ isPublic, user }) => isPublic && user === selectedUser),
    quantities: ({ currentCards }) => Object.fromEntries(currentCards.map(({ id, quantity }) => [id, quantity])),
    currentDeck: ({ currentDeck, currentCards }, getters, { currentUser }) => ({
      ...currentDeck,
      planeswalker: allPlaneswalkers[currentDeck.planeswalkerId],
      land: allCards[currentDeck.landId],
      featureCard: allCards[currentDeck.featureCardId],
      cards: currentCards.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
      isMine: currentDeck.user === currentUser,
    }),
    cardMaterials: ({ currentCards }) => {
      const materials = {};
      currentCards.forEach(({ id, quantity }) => {
        Object.keys(allCards[id].materials).forEach((materialId) => {
          if (!materials[materialId]) {
            materials[materialId] = 0;
          }
          materials[materialId] += quantity * allCards[id].materials[materialId];
        });
      });
      return Object.keys(materials).map((materialId) => ({ materialId, quantity: materials[materialId] }));
    },
  },
  mutations: {
    updateQuantity({ currentCards }, { id, quantity }) {
      const cardIndex = currentCards.findIndex((card) => card.id === id);
      if (cardIndex >= 0) {
        if (quantity < 1) {
          currentCards.splice(cardIndex, 1);
        } else {
          currentCards[cardIndex].quantity = quantity;
        }
      } else {
        currentCards.push({ id, quantity });
      }
    },
    setPlaneswalker({ currentDeck }, planeswalkerId) {
      currentDeck.planeswalkerId = planeswalkerId;
    },
    setLand({ currentDeck }, landId) {
      currentDeck.landId = landId;
    },
    setFeatureCard({ currentDeck }, featureCardId) {
      currentDeck.featureCardId = featureCardId;
    },
    setSplash({ currentDeck }, colour) {
      currentDeck.splashColour = colour;
    },
    setName({ currentDeck }, name) {
      currentDeck.name = name;
    },
    setDescription({ currentDeck }, description) {
      currentDeck.description = description;
    },
    setPublic({ currentDeck }, isPublic) {
      currentDeck.isPublic = isPublic;
    },
    loadDeck(state, { deckId, currentUser }) {
      const generatedId = deckId === 'new' ? Math.random().toString() : deckId;
      const deck = state.decks.find(({ id }) => id === generatedId) || emptyDeck;
      state.currentDeck = { user: currentUser, ...deck, id: generatedId };
      state.currentCards = [...deck.cards];
    },
    resetDeck(state) {
      const deck = state.decks.find(({ id }) => id === state.currentDeck.id);
      state.currentDeck = { ...deck };
      state.currentCards = [...deck.cards];
    },
    saveDeck({ decks, currentDeck, currentCards }) {
      const deck = { ...currentDeck, cards: [...currentCards] };
      const deckIndex = decks.findIndex(({ id }) => id === deck.id);

      if (deckIndex >= 0) {
        decks.splice(deckIndex, 1, deck);
      } else {
        decks.push(deck);
      }
    },
    deleteDeck(state) {
      const deckIndex = state.decks.findIndex(({ id }) => id === state.currentDeck.id);
      if (deckIndex >= 0) {
        state.decks.splice(deckIndex, 1);
      }

      state.currentCards = [];
      state.currentDeck = { ...emptyDeck };
    },
  },
  actions: {
    updateQuantity: ({ commit }, { id, quantity }) => commit('updateQuantity', { id, quantity }),
    setPlaneswalker: ({ commit }, planeswalkerId) => commit('setPlaneswalker', planeswalkerId),
    setLand: ({ commit }, landId) => commit('setLand', landId),
    setFeatureCard: ({ commit }, featureCardId) => commit('setFeatureCard', featureCardId),
    setSplash: ({ commit }, colour) => commit('setSplash', colour),
    setName: ({ commit }, name) => commit('setName', name),
    setDescription: ({ commit }, description) => commit('setDescription', description),
    setPublic: ({ commit }, isPublic) => commit('setPublic', isPublic),
    loadDeck: ({ commit, rootState }, deckId) => commit('loadDeck', { deckId, currentUser: rootState.currentUser }),
    resetDeck: ({ commit }) => commit('resetDeck'),
    saveDeck: ({ commit }) => commit('saveDeck'),
    deleteDeck: ({ commit }) => commit('deleteDeck'),
  },
};
