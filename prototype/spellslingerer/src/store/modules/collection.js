import allCards from '../data/cards.json';

export default {
  namespaced: true,
  state: () => ({
    collection: [],
  }),
  getters: {
    collection: ({ collection }) => collection.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
    quantities: ({ collection }) => Object.fromEntries(collection.map(({ id, quantity }) => [id, quantity])),
    cardMaterials: ({ collection }) => {
      const materials = {};
      collection.forEach(({ id, quantity }) => {
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
    updateQuantity({ collection }, { id, quantity }) {
      const cardIndex = collection.findIndex((card) => card.id === id);
      if (cardIndex >= 0) {
        if (quantity < 1) {
          collection.splice(cardIndex, 1);
        } else {
          collection[cardIndex].quantity = quantity;
        }
      } else {
        collection.push({ id, quantity });
      }
    },
  },
  actions: {
    updateQuantity: ({ commit }, { id, quantity }) => commit('updateQuantity', { id, quantity }),
  },
};
