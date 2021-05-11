import allPlaneswalkers from '../data/planeswalkers';

export default {
  namespaced: true,
  getters: {
    planeswalkers: () => Object.values(allPlaneswalkers).filter(({ id }) => id !== 'Default'),
    planeswalker: () => (planeswalkerId) => allPlaneswalkers[planeswalkerId],
    planeswalkerNames: () => Object.values(allPlaneswalkers).filter(({ id }) => id !== 'Default').map(({ name }) => name),
  },
};
