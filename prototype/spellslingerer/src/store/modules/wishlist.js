import allCards from '../data/cards.json';

export default {
  namespaced: true,
  state: () => ({
    wishlist: [],
  }),
  getters: {
    wishlist: ({ wishlist }) => wishlist.map(({ id, quantity }) => ({ ...allCards[id], quantity })),
    quantities: ({ wishlist }) => Object.fromEntries(wishlist.map(({ id, quantity }) => [id, quantity])),
    cardMaterials: ({ wishlist }) => {
      const materials = {};
      wishlist.forEach(({ id, quantity }) => {
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
    updateQuantity({ wishlist }, { id, quantity }) {
      const cardIndex = wishlist.findIndex((card) => card.id === id);
      if (cardIndex >= 0) {
        if (quantity < 1) {
          wishlist.splice(cardIndex, 1);
        } else {
          wishlist[cardIndex].quantity = quantity;
        }
      } else {
        wishlist.push({ id, quantity });
      }
    },
  },
  actions: {
    updateQuantity: ({ commit }, { id, quantity }) => commit('updateQuantity', { id, quantity }),
  },
};
