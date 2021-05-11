import allMaterials from '../data/materials.json';

export default {
  namespaced: true,
  state: () => ({
    materialQuantities: Object.fromEntries(Object.keys(allMaterials).map((materialId) => [materialId, 0])),
  }),
  getters: {
    materials: ({ materialQuantities }) => Object.keys(materialQuantities).map((materialId) => ({ ...allMaterials[materialId], quantity: materialQuantities[materialId] })),
    material: ({ materialQuantities }) => (materialId) => ({ ...allMaterials[materialId], quantity: materialQuantities[materialId] }),
    currentQuantity: ({ materialQuantities }) => (materialId) => materialQuantities[materialId],
    quantityStillRequired: ({ materialQuantities }) => ({ materialId, quantity }) => ({
      ...allMaterials[materialId],
      quantity: quantity - materialQuantities[materialId],
    }),
  },
  mutations: {
    updateQuantity({ materialQuantities }, { materialId, quantity }) {
      materialQuantities[materialId] = quantity;
    },
  },
  actions: {
    updateQuantity: ({ commit }, { materialId, quantity }) => commit('updateQuantity', { materialId, quantity }),
  },
};
